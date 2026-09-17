package crossplane

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/iam"
	crossplanev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	ProvisionTimeout = 15 * time.Minute
	PollInterval     = 30 * time.Second
)

type ProvisionedResource struct {
	ID        string
	Type      string
	Endpoint  string
	Status    string
	CreatedAt time.Time
}

type RDSConfig struct {
	InstanceClass    string
	Engine           string
	EngineVersion    string
	AllocatedStorage int64
	MultiAZ          bool
	StorageEncrypted bool
	BackupRetention  int64
}

type S3Config struct {
	BucketName    string
	Region        string
	Versioning    bool
	Encryption    string
	PublicAccess  bool
	LifecyclePolicy *LifecyclePolicy
}

type LifecyclePolicy struct {
	Transitions []LifecycleTransition
}

type LifecycleTransition struct {
	Days          int
	StorageClass  string
}

type IAMRoleConfig struct {
	RoleName        string
	PolicyARNs      []string
	AssumeRolePolicy string
	Tags            map[string]string
}

type CrossplaneProvider struct {
	client       client.Client
	awsConfig    *AWSConfig
	resourceDefs map[string]resource.Managed
}

type AWSConfig struct {
	Region    string
	AccountID string
}

func NewCrossplaneProvider(k8sClient client.Client, awsCfg *AWSConfig) *CrossplaneProvider {
	return &CrossplaneProvider{
		client:    k8sClient,
		awsConfig: awsCfg,
		resourceDefs: make(map[string]resource.Managed),
	}
}

func (cp *CrossplaneProvider) ProvisionRDS(ctx context.Context, name string, cfg RDSConfig) (*ProvisionedResource, error) {
	log.Log.Info("Provisioning RDS instance", "name", name, "class", cfg.InstanceClass)

	svc := rds.New(cp.newAWSSession())

	input := &rds.CreateDBInstanceInput{
		DBInstanceIdentifier: aws.String(name),
		DBInstanceClass:      aws.String(cfg.InstanceClass),
		Engine:               aws.String(cfg.Engine),
		EngineVersion:        aws.String(cfg.EngineVersion),
		AllocatedStorage:     aws.Int64(cfg.AllocatedStorage),
		MultiAZ:              aws.Bool(cfg.MultiAZ),
		StorageEncrypted:     aws.Bool(cfg.StorageEncrypted),
		BackupRetentionPeriod: aws.Int64(cfg.BackupRetention),
		MasterUsername:       aws.String("admin"),
		MasterUserPassword:   aws.String(generatePassword()),
		PubliclyAccessible:   aws.Bool(false),
		tags: []*rds.Tag{
			{Key: aws.String("managed-by"), Value: aws.String("crossplane")},
			{Key: aws.String("environment"), Value: aws.String("production")},
		},
	}

	result, err := svc.CreateDBInstance(input)
	if err != nil {
		return nil, fmt.Errorf("failed to create RDS instance: %w", err)
	}

	pr := &ProvisionedResource{
		ID:        *result.DBInstance.DBInstanceArn,
		Type:      "RDS",
		Endpoint:  *result.DBInstance.Endpoint.Address,
		Status:    *result.DBInstance.DBInstanceStatus,
		CreatedAt: time.Now(),
	}

	log.Log.Info("RDS instance provisioned", "arn", pr.ID, "endpoint", pr.Endpoint)
	return pr, nil
}

func (cp *CrossplaneProvider) ProvisionS3(ctx context.Context, name string, cfg S3Config) (*ProvisionedResource, error) {
	log.Log.Info("Provisioning S3 bucket", "name", name, "region", cfg.Region)

	svc := s3.New(cp.newAWSSession())

	input := &s3.CreateBucketInput{
		Bucket: aws.String(name),
		CreateBucketConfiguration: &s3.CreateBucketConfiguration{
			LocationConstraint: aws.String(cfg.Region),
		},
	}

	_, err := svc.CreateBucket(input)
	if err != nil {
		return nil, fmt.Errorf("failed to create S3 bucket: %w", err)
	}

	if cfg.Versioning {
		err = svc.PutBucketVersioning(&s3.PutBucketVersioningInput{
			Bucket: aws.String(name),
			VersioningConfiguration: &s3.VersioningConfiguration{
				Status: aws.String(s3.BucketVersioningStatusEnabled),
			},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to enable versioning: %w", err)
		}
	}

	if cfg.Encryption != "" {
		err = svc.PutBucketEncryption(&s3.PutBucketEncryptionInput{
			Bucket: aws.String(name),
			ServerSideEncryptionConfiguration: &s3.ServerSideEncryptionConfiguration{
				Rules: []*s3.ServerSideEncryptionRule{
					{
						ApplyServerSideEncryptionByDefault: &s3.ServerSideEncryptionByDefault{
							SSEAlgorithm: aws.String(cfg.Encryption),
						},
					},
				},
			},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to enable encryption: %w", err)
		}
	}

	pr := &ProvisionedResource{
		ID:        name,
		Type:      "S3",
		Endpoint:  fmt.Sprintf("s3://%s", name),
		Status:    "Available",
		CreatedAt: time.Now(),
	}

	log.Log.Info("S3 bucket provisioned", "bucket", name)
	return pr, nil
}

func (cp *CrossplaneProvider) ProvisionIAMRole(ctx context.Context, name string, cfg IAMRoleConfig) (*ProvisionedResource, error) {
	log.Log.Info("Provisioning IAM role", "name", name)

	svc := iam.New(cp.newAWSSession())

	input := &iam.CreateRoleInput{
		RoleName:                 aws.String(name),
		AssumeRolePolicyDocument: aws.String(cfg.AssumeRolePolicy),
		tags: []*iam.Tag{
			{Key: aws.String("managed-by"), Value: aws.String("crossplane")},
		},
	}

	role, err := svc.CreateRole(input)
	if err != nil {
		return nil, fmt.Errorf("failed to create IAM role: %w", err)
	}

	for _, policyArn := range cfg.PolicyARNs {
		_, err = svc.AttachRolePolicy(&iam.AttachRolePolicyInput{
			RoleName:  aws.String(name),
			PolicyArn: aws.String(policyArn),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to attach policy %s: %w", policyArn, err)
		}
	}

	pr := &ProvisionedResource{
		ID:        *role.Role.Arn,
		Type:      "IAM",
		Endpoint:  *role.Role.Arn,
		Status:    "Active",
		CreatedAt: time.Now(),
	}

	log.Log.Info("IAM role provisioned", "arn", pr.ID)
	return pr, nil
}

func (cp *CrossplaneProvider) DeleteResource(ctx context.Context, resourceType, identifier string) error {
	log.Log.Info("Deleting resource", "type", resourceType, "id", identifier)

	switch resourceType {
	case "RDS":
		return cp.deleteRDS(ctx, identifier)
	case "S3":
		return cp.deleteS3(ctx, identifier)
	case "IAM":
		return cp.deleteIAM(ctx, identifier)
	default:
		return fmt.Errorf("unknown resource type: %s", resourceType)
	}
}

func (cp *CrossplaneProvider) deleteRDS(ctx context.Context, identifier string) error {
	svc := rds.New(cp.newAWSSession())
	_, err := svc.DeleteDBInstance(&rds.DeleteDBInstanceInput{
		DBInstanceIdentifier: aws.String(identifier),
		SkipFinalSnapshot:     aws.Bool(true),
		DeleteAutomatedBackups: aws.Bool(true),
	})
	return err
}

func (cp *CrossplaneProvider) deleteS3(ctx context.Context, identifier string) error {
	svc := s3.New(cp.newAWSSession())
	_, err := svc.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: aws.String(identifier),
	})
	return err
}

func (cp *CrossplaneProvider) deleteIAM(ctx context.Context, identifier string) error {
	svc := iam.New(cp.newAWSSession())
	_, err := svc.DeleteRole(&iam.DeleteRoleInput{
		RoleName: aws.String(identifier),
	})
	return err
}

func (cp *CrossplaneProvider) newAWSSession() *aws.Config {
	return aws.NewConfig().WithRegion(cp.awsConfig.Region)
}

func generatePassword() string {
	return "temporary-password-change-me"
}

type CrossplaneConfig struct {
	ProviderName string
	Region       string
}

func (cp *CrossplaneProvider) Reconcile(ctx context.Context) error {
	log.Log.Info("Reconciling Crossplane resources")
	return nil
}

var _ client.Reconciler = &CrossplaneProvider{}