package argocd

import (
	"context"
	"fmt"
	"time"

	argoclient "github.com/argoproj/argo-cd/v2/pkg/apiclient"
	appclient "github.com/argoproj/argo-cd/v2/pkg/apiclient/application"
	clusterclient "github.com/argoproj/argo-cd/v2/pkg/apiclient/cluster"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	DefaultArgoCDTimeout = 30 * time.Second
	SyncRetryLimit       = 3
	SyncRetryInterval    = 5 * time.Second
)

type SyncPolicy struct {
	AutoSync          bool     `json:"autoSync"`
	SyncOptions       []string `json:"syncOptions"`
	SelfHeal          bool     `json:"selfHeal"`
	PrunePropagation  string   `json:"prunePropagation"`
	DryRun            bool     `json:"dryRun"`
}

type ClusterFederationConfig struct {
	SourceCluster    string
	TargetClusters   []string
	SyncPolicy       SyncPolicy
	ProjectName      string
	NamespaceMapping map[string]string
}

type ArgoCDClient struct {
	client    argoclient.Client
	opts      argoclient.ClientOptions
	namespace string
}

func NewArgoCDClient(config ArgoCDConfig, namespace string) (*ArgoCDClient, error) {
	if config.Server == "" {
		return nil, fmt.Errorf("ArgoCD server address is required")
	}

	opts := argoclient.ClientOptions{
		ServerAddr: config.Server,
		AuthToken:  config.Token,
		Insecure:   config.Insecure,
		TLSClientConfig: argoclient.TLSClientConfig{
			Insecure: config.Insecure,
		},
	}

	return &ArgoCDClient{
		opts:      opts,
		namespace: namespace,
	}, nil
}

func (a *ArgoCDClient) Connect(ctx context.Context) error {
	client, err := argoclient.NewClient(&a.opts)
	if err != nil {
		return fmt.Errorf("failed to create ArgoCD client: %w", err)
	}
	a.client = client
	return nil
}

func (a *ArgoCDClient) SyncApplication(ctx context.Context, appName string, syncPolicy SyncPolicy) error {
	if a.client == nil {
		return fmt.Errorf("ArgoCD client not connected")
	}

	appClient, err := a.client.NewApplicationClient()
	if err != nil {
		return fmt.Errorf("failed to create application client: %w", err)
	}

	req := &appclient.ApplicationSyncRequest{
		Name:         &appName,
		DryRun:       syncPolicy.DryRun,
		SyncOptions:  v1alpha1.SyncOptions(syncPolicy.SyncOptions),
		Prune:        true,
		RetryStrategy: &v1alpha1.RetryStrategy{
			Limit: int64(SyncRetryLimit),
		},
	}

	syncResp, err := appClient.Sync(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to sync application %s: %w", appName, err)
	}

	log.Log.Info("Application sync initiated", "app", appName, "syncResult", syncResp.Result)
	return nil
}

func (a *ArgoCDClient) GetApplicationState(ctx context.Context, appName string) (*v1alpha1.ApplicationState, error) {
	if a.client == nil {
		return nil, fmt.Errorf("ArgoCD client not connected")
	}

	appClient, err := a.client.NewApplicationClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create application client: %w", err)
	}

	app, err := appClient.Get(ctx, &appclient.ApplicationQuery{
		Name:    &appName,
		Project: a.namespace,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get application %s: %w", appName, err)
	}

	return &app.Status, nil
}

func (a *ArgoCDClient) RegisterCluster(ctx context.Context, cluster *v1alpha1.Cluster) error {
	if a.client == nil {
		return fmt.Errorf("ArgoCD client not connected")
	}

	clusterClient, err := a.client.NewClusterClient()
	if err != nil {
		return fmt.Errorf("failed to create cluster client: %w", err)
	}

	_, err = clusterClient.Create(ctx, &clusterclient.ClusterCreateRequest{
		Cluster: cluster,
	})
	if err != nil {
		return fmt.Errorf("failed to register cluster %s: %w", cluster.Server, err)
	}

	log.Log.Info("Cluster registered in ArgoCD", "server", cluster.Server, "name", cluster.Name)
	return nil
}

func (a *ArgoCDClient) FederateClusters(ctx context.Context, config ClusterFederationConfig) error {
	if a.client == nil {
		return fmt.Errorf("ArgoCD client not connected")
	}

	appClient, err := a.client.NewApplicationClient()
	if err != nil {
		return fmt.Errorf("failed to create application client: %w", err)
	}

	for _, targetCluster := range config.TargetClusters {
		app := &v1alpha1.Application{
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("%s-%s", config.ProjectName, targetCluster),
				Namespace: a.namespace,
			},
			Spec: v1alpha1.ApplicationSpec{
				Project: config.ProjectName,
				Source: v1alpha1.ApplicationSource{
					RepoURL:        "https://github.com/example/controlplane-repo",
					Path:           config.NamespaceMapping[targetCluster],
					targetRevision: "main",
				},
				Destination: v1alpha1.ApplicationDestination{
					Server:    targetCluster,
					Namespace: "default",
				},
				SyncPolicy: &v1alpha1.SyncPolicy{
					Automated: &v1alpha1.SyncPolicyAutomated{
						SelfHeal: config.SyncPolicy.SelfHeal,
						Prune:    true,
					},
				},
			},
		}

		_, err = appClient.Create(ctx, &appclient.ApplicationCreateRequest{
			Application: app,
		})
		if err != nil {
			if !errors.IsAlreadyExists(err) {
				return fmt.Errorf("failed to create federated application for cluster %s: %w", targetCluster, err)
			}
			log.Log.Info("Application already exists, updating", "cluster", targetCluster)
		}

		log.Log.Info("Cluster federated in ArgoCD", "source", config.SourceCluster, "target", targetCluster)
	}

	return nil
}

func (a *ArgoCDClient) SetAutoSync(ctx context.Context, appName string, policy SyncPolicy) error {
	if a.client == nil {
		return fmt.Errorf("ArgoCD client not connected")
	}

	return nil
}

func (a *ArgoCDClient) GetClusterList(ctx context.Context) (*v1alpha1.ClusterList, error) {
	if a.client == nil {
		return nil, fmt.Errorf("ArgoCD client not connected")
	}

	clusterClient, err := a.client.NewClusterClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create cluster client: %w", err)
	}

	return clusterClient.List(ctx, &clusterclient.ClusterQuery{})
}

func (a *ArgoCDClient) DeleteCluster(ctx context.Context, server string) error {
	if a.client == nil {
		return fmt.Errorf("ArgoCD client not connected")
	}

	clusterClient, err := a.client.NewClusterClient()
	if err != nil {
		return fmt.Errorf("failed to create cluster client: %w", err)
	}

	err = clusterClient.Delete(ctx, &clusterclient.ClusterQuery{Server: server})
	if err != nil {
		return fmt.Errorf("failed to delete cluster %s: %w", server, err)
	}

	log.Log.Info("Cluster deleted from ArgoCD", "server", server)
	return nil
}

type ArgoCDConfig struct {
	Server   string
	Token    string
	Insecure bool
}

func (a *ArgoCDClient) Reconcile(ctx context.Context, req types.NamespacedName) error {
	log.Log.Info("Reconciling ArgoCD resources", "name", req.Name)

	return nil
}

var _ client.Reconciler = &ArgoCDClient{}