package main

import (
	"flag"
	"os"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	controlplanev1alpha1 "github.com/example/controlplane-operator/api/v1alpha1"
	"github.com/example/controlplane-operator/controllers"
	"github.com/example/controlplane-operator/pkg/argocd"
	"github.com/example/controlplane-operator/pkg/aws"
	"github.com/example/controlplane-operator/pkg/crossplane"
	"github.com/example/controlplane-operator/pkg/failover"
	"github.com/example/controlplane-operator/pkg/healthcheck"
	"github.com/example/controlplane-operator/pkg/rbac"
	"github.com/example/controlplane-operator/pkg/replication"
	"github.com/example/controlplane-operator/pkg/vault"
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	_ = controlplanev1alpha1.AddToScheme(scheme)
}

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	var probeAddr string
	var clusterConfigPath string

	flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "leader-elect", false,
		"Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.")
	flag.StringVar(&clusterConfigPath, "cluster-config", "/etc/controlplane/clusters.yaml",
		"Path to the cluster topology configuration file.")

	opts := zap.Options{
		Development: true,
	}
	opts.BindFlags(flag.CommandLine)
	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     metricsAddr,
		Port:                   9443,
		HealthProbeBindAddress: probeAddr,
		LeaderElection:         enableLeaderElection,
		LeaderElectionID:       "controlplane-operator-lock",
	})
	if err != nil {
		setupLog.Error(err, "Unable to start manager")
		os.Exit(1)
	}

	clusterConfig, err := loadClusterConfiguration(clusterConfigPath)
	if err != nil {
		setupLog.Error(err, "Failed to load cluster configuration")
		os.Exit(1)
	}

	argocdClient, err := argocd.NewClient(clusterConfig.ArgoCD)
	if err != nil {
		setupLog.Error(err, "Unable to initialize ArgoCD client")
		os.Exit(1)
	}

	crossplaneProvider, err := crossplane.NewProvider(clusterConfig.Crossplane)
	if err != nil {
		setupLog.Error(err, "Unable to initialize Crossplane provider")
		os.Exit(1)
	}

	vaultClient, err := vault.NewClient(clusterConfig.Vault)
	if err != nil {
		setupLog.Error(err, "Unable to initialize Vault client")
		os.Exit(1)
	}

	awsConfig, err := aws.NewConfig(clusterConfig.AWS)
	if err != nil {
		setupLog.Error(err, "Unable to initialize AWS configuration")
		os.Exit(1)
	}

	topology, err := replication.NewTopology(clusterConfig.Clusters)
	if err != nil {
		setupLog.Error(err, "Unable to initialize replication topology")
		os.Exit(1)
	}

	failoverMediator := failover.NewMediator(topology, argocdClient, vaultClient)

	healthChecker := healthcheck.NewChecker(argocdClient, awsConfig)

	rbacFederator, err := rbac.NewFederator(mgr.GetClient(), clusterConfig.Clusters)
	if err != nil {
		setupLog.Error(err, "Unable to initialize RBAC federator")
		os.Exit(1)
	}

	if err = (&controllers.ControlPlaneReconciler{
		Client:            mgr.GetClient(),
		Scheme:           mgr.GetScheme(),
		ArgoCD:            argocdClient,
		Crossplane:       crossplaneProvider,
		Vault:             vaultClient,
		AWSConfig:         awsConfig,
		Topology:          topology,
		FailoverMediator:  failoverMediator,
		HealthChecker:     healthChecker,
		RbacFederator:     rbacFederator,
		ClusterConfig:     clusterConfig,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "Unable to create controller")
		os.Exit(1)
	}

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "Unable to set up health check")
		os.Exit(1)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "Unable to set up ready check")
		os.Exit(1)
	}

	setupLog.Info("Starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "Problem running manager")
		os.Exit(1)
	}
}

type ClusterConfiguration struct {
	Clusters   []ClusterInfo `json:"clusters"`
	ArgoCD     ArgoCDConfig  `json:"argocd"`
	Crossplane CrossplaneConfig `json:"crossplane"`
	Vault      VaultConfig   `json:"vault"`
	AWS        AWSConfig     `json:"aws"`
}

type ClusterInfo struct {
	Name     string `json:"name"`
	Region   string `json:"region"`
	Endpoint string `json:"endpoint"`
	Priority int    `json:"priority"`
}

type ArgoCDConfig struct {
	Server   string `json:"server"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CrossplaneConfig struct {
	Provider string `json:"provider"`
	Region   string `json:"region"`
}

type VaultConfig struct {
	Address  string `json:"address"`
	Token    string `json:"token"`
	Path     string `json:"path"`
}

type AWSConfig struct {
	Region    string `json:"region"`
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

func loadClusterConfiguration(configPath string) (*ClusterConfiguration, error) {
	return &ClusterConfiguration{
		Clusters: []ClusterInfo{
			{Name: "primary-us-east-1", Region: "us-east-1", Endpoint: "https://kube.us-east-1.example.com", Priority: 1},
			{Name: "primary-us-west-2", Region: "us-west-2", Endpoint: "https://kube.us-west-2.example.com", Priority: 2},
			{Name: "primary-eu-west-1", Region: "eu-west-1", Endpoint: "https://kube.eu-west-1.example.com", Priority: 3},
			{Name: "primary-ap-southeast-1", Region: "ap-southeast-1", Endpoint: "https://kube.ap-southeast-1.example.com", Priority: 4},
		},
		ArgoCD: ArgoCDConfig{
			Server:   "argocd.example.com",
			Username: "admin",
			Password: "changeme",
		},
		Crossplane: CrossplaneConfig{
			Provider: "aws",
			Region:   "us-east-1",
		},
		Vault: VaultConfig{
			Address: "https://vault.example.com",
			Token:   "root",
			Path:    "secret/data/controlplane",
		},
		AWS: AWSConfig{
			Region:    "us-east-1",
			AccessKey: "",
			SecretKey: "",
		},
	}, nil
}