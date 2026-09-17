package vault

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	vaultapi "github.com/hashicorp/vault/api"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	DefaultVaultPath = "secret/data/controlplane"
	MaxRetries       = 3
	RetryInterval    = 2 * time.Second
)

type SecretRotation struct {
	Enabled        bool
	Interval       time.Duration
	GracePeriod    time.Duration
	RotationPolicy string
}

type SecretReplication struct {
	SourceRegion    string
	TargetRegions   []string
	ReplicationMode string
	FailoverEnabled bool
}

type VaultSecret struct {
	Key       string
	Value     string
	Version   int
	CreatedAt time.Time
	ExpiresAt *time.Time
}

type ExternalSecretSpec struct {
	SecretStore    string
	RemoteKey      string
	RefreshInterval time.Duration
	Data           []RemoteRef
}

type RemoteRef struct {
	RemoteKey string
	Property  string
}

type VaultClient struct {
	client      *vaultapi.Client
	config      *VaultConfig
	secretCache map[string]*VaultSecret
}

type VaultConfig struct {
	Address       string
	Token         string
	Namespace     string
	MountPath     string
	KVVersion     string
	TLSConfig     TLSConfig
	RetryConfig   RetryConfig
}

type TLSConfig struct {
	CACert     string
	CAPem      string
	ClientCert string
	ClientKey string
	SkipVerify bool
}

type RetryConfig struct {
	MaxRetries int
	Interval   time.Duration
}

func NewVaultClient(config VaultConfig) (*VaultClient, error) {
	if config.Address == "" {
		return nil, fmt.Errorf("Vault address is required")
	}

	vaultConfig := &vaultapi.Config{
		Address:    config.Address,
		MaxRetries: config.RetryConfig.MaxRetries,
	}

	if config.TLSConfig.CACert != "" {
		vaultConfig.CAFile = config.TLSConfig.CACert
	}

	if config.TLSConfig.SkipVerify {
		vaultConfig.TLSConfig.Insecure = true
	}

	client, err := vaultapi.NewClient(vaultConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create Vault client: %w", err)
	}

	if config.Token != "" {
		client.SetToken(config.Token)
	}

	if config.Namespace != "" {
		client.SetNamespace(config.Namespace)
	}

	return &VaultClient{
		client:      client,
		config:      &config,
		secretCache: make(map[string]*VaultSecret),
	}, nil
}

func (v *VaultClient) WriteSecret(ctx context.Context, path string, data map[string]interface{}) error {
	fullPath := v.buildPath(path)

	secret, err := v.client.KVv2(v.config.MountPath).Put(ctx, fullPath, data)
	if err != nil {
		return fmt.Errorf("failed to write secret to Vault: %w", err)
	}

	log.Log.Info("Secret written to Vault", "path", fullPath, "version", secret.Version)
	return nil
}

func (v *VaultClient) ReadSecret(ctx context.Context, path string) (map[string]interface{}, error) {
	fullPath := v.buildPath(path)

	secret, err := v.client.KVv2(v.config.MountPath).Get(ctx, fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read secret from Vault: %w", err)
	}

	data := make(map[string]interface{})
	for k, val := range secret.Data {
		data[k] = val
	}

	log.Log.Info("Secret read from Vault", "path", fullPath)
	return data, nil
}

func (v *VaultClient) DeleteSecret(ctx context.Context, path string) error {
	fullPath := v.buildPath(path)

	err := v.client.KVv2(v.config.MountPath).Delete(ctx, fullPath)
	if err != nil {
		return fmt.Errorf("failed to delete secret from Vault: %w", err)
	}

	log.Log.Info("Secret deleted from Vault", "path", fullPath)
	return nil
}

func (v *VaultClient) ListSecrets(ctx context.Context, path string) ([]string, error) {
	fullPath := v.buildPath(path)

	secrets, err := v.client.KVv2(v.config.MountPath).List(ctx, fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to list secrets from Vault: %w", err)
	}

	return secrets.Keys, nil
}

func (v *VaultClient) RotateSecret(ctx context.Context, path string, rotation SecretRotation) error {
	log.Log.Info("Initiating secret rotation", "path", path, "interval", rotation.Interval)

	data, err := v.ReadSecret(ctx, path)
	if err != nil {
		return fmt.Errorf("failed to read secret for rotation: %w", err)
	}

	data["_rotation_timestamp"] = time.Now().Unix()
	data["_rotation_version"] = data["_rotation_version"].(int) + 1

	err = v.WriteSecret(ctx, path, data)
	if err != nil {
		return fmt.Errorf("failed to write rotated secret: %w", err)
	}

	log.Log.Info("Secret rotated successfully", "path", path)
	return nil
}

func (v *VaultClient) ReplicateSecret(ctx context.Context, path string, replication SecretReplication) error {
	sourcePath := v.buildPathInRegion(path, replication.SourceRegion)

	sourceData, err := v.client.KVv2(v.config.MountPath).Get(ctx, sourcePath)
	if err != nil {
		return fmt.Errorf("failed to read source secret: %w", err)
	}

	for _, targetRegion := range replication.TargetRegions {
		targetPath := v.buildPathInRegion(path, targetRegion)

		err = v.client.KVv2(v.config.MountPath).Put(ctx, targetPath, sourceData.Data)
		if err != nil {
			log.Log.Error(err, "Failed to replicate secret to region", "region", targetRegion)
			if !replication.FailoverEnabled {
				return fmt.Errorf("failed to replicate secret to %s: %w", targetRegion, err)
			}
		}

		log.Log.Info("Secret replicated", "from", replication.SourceRegion, "to", targetRegion)
	}

	return nil
}

func (v *VaultClient) buildPath(path string) string {
	if strings.HasPrefix(path, "/") {
		return path
	}
	return filepath.Join(DefaultVaultPath, path)
}

func (v *VaultClient) buildPathInRegion(path string, region string) string {
	basePath := v.buildPath(path)
	parts := strings.Split(basePath, "/")
	if len(parts) > 0 {
		parts = append([]string{region}, parts...)
	}
	return strings.Join(parts, "/")
}

func (v *VaultClient) GetSecretVersion(ctx context.Context, path string, version int) (map[string]interface{}, error) {
	fullPath := v.buildPath(path)

	secret, err := v.client.KVv2(v.config.MountPath).GetVersion(ctx, fullPath, version)
	if err != nil {
		return nil, fmt.Errorf("failed to read secret version: %w", err)
	}

	data := make(map[string]interface{})
	for k, val := range secret.Data {
		data[k] = val
	}

	return data, nil
}

func (v *VaultClient) UnwrapToken(ctx context.Context, token string) (map[string]interface{}, error) {
	secret, err := v.client.Logical().Unwrap(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("failed to unwrap token: %w", err)
	}

	return secret.Data, nil
}

func (v *VaultClient) HealthCheck(ctx context.Context) error {
	_, err := v.client.Sys().Health()
	return err
}

type ExternalSecretReconciler struct {
	client     *VaultClient
	k8sClient  client.Client
	namespace  string
}

func NewExternalSecretReconciler(vaultClient *VaultClient, k8sClient client.Client, namespace string) *ExternalSecretReconciler {
	return &ExternalSecretReconciler{
		client:    vaultClient,
		k8sClient: k8sClient,
		namespace: namespace,
	}
}

func (r *ExternalSecretReconciler) Reconcile(ctx context.Context, req types.NamespacedName) error {
	log.Log.Info("Reconciling external secret", "name", req.Name, "namespace", req.Namespace)

	data, err := r.client.ReadSecret(ctx, req.Name)
	if err != nil {
		if errors.IsNotFound(err) {
			log.Log.Info("Secret not found in Vault, skipping", "name", req.Name)
			return nil
		}
		return fmt.Errorf("failed to read secret: %w", err)
	}

	log.Log.Info("External secret reconciled", "name", req.Name, "keys", len(data))
	return nil
}

var _ client.Reconciler = &ExternalSecretReconciler{}

type VaultConfig struct {
	Address     string
	Token       string
	Namespace   string
	MountPath   string
}