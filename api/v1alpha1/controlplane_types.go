package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ControlPlaneSpec define la especificación del control plane multi-cluster
type ControlPlaneSpec struct {
	// Clusters contiene la configuración de todos los clusters gestionados
	// +kubebuilder:validation:MinItems=1
	Clusters []ClusterSpec `json:"clusters"`

	// ArgoCD define la configuración de ArgoCD para GitOps
	// +optional
	ArgoCD *ArgoCDSpec `json:"argocd,omitempty"`

	// Crossplane define la configuración de Crossplane para infraestructura
	// +optional
	Crossplane *CrossplaneSpec `json:"crossplane,omitempty"`

	// Vault define la configuración de Vault para secrets
	// +optional
	Vault *VaultSpec `json:"vault,omitempty"`

	// AWS define la configuración de AWS para salud de regiones
	// +optional
	AWS *AWSSpec `json:"aws,omitempty"`

	// Replication define la estrategia de replicación de estado
	// +optional
	Replication *ReplicationSpec `json:"replication,omitempty"`

	// Failover define la configuración de disaster recovery
	// +optional
	Failover *FailoverSpec `json:"failover,omitempty"`
}

// ClusterSpec define la especificación de un cluster individual
type ClusterSpec struct {
	// Name es el nombre identificador del cluster
	// +kubebuilder:validation:Pattern=^[a-z0-9-]+$
	Name string `json:"name"`

	// Region es la región AWS donde reside el cluster
	// +kubebuilder:validation:Pattern=^[a-z]{2}-[a-z]+-[0-9]+$
	Region string `json:"region"`

	// Endpoint es la URL del API server del cluster
	// +kubebuilder:validation:Format=uri
	Endpoint string `json:"endpoint"`

	// IsMain indica si este cluster es el principal (primary)
	// +optional
	// +kubebuilder:default=false
	IsMain bool `json:"isMain,omitempty"`

	// IsDrTarget indica si este cluster es objetivo de disaster recovery
	// +optional
	// +kubebuilder:default=false
	IsDrTarget bool `json:"isDrTarget,omitempty"`

	// Priority define la prioridad del cluster para failover
	// +optional
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	Priority int `json:"priority,omitempty"`

	// Labels contiene etiquetas adicionales para el cluster
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
}

// ArgoCDSpec define la configuración de ArgoCD
type ArgoCDSpec struct {
	// Enabled indica si ArgoCD está habilitado
	// +optional
	// +kubebuilder:default=true
	Enabled bool `json:"enabled,omitempty"`

	// Version es la versión de ArgoCD a desplegar
	// +optional
	// +kubebuilder:default="v2.9.0"
	Version string `json:"version,omitempty"`

	// Namespace es el namespace donde ArgoCD será instalado
	// +optional
	// +kubebuilder:default="argocd"
	Namespace string `json:"namespace,omitempty"`

	// RepoURL es la URL del repositorio Git con las definiciones de aplicaciones
	// +kubebuilder:validation:Format=uri
	RepoURL string `json:"repoURL"`

	// RepoBranch es la rama del repositorio a usar
	// +optional
	// +kubebuilder:default="main"
	RepoBranch string `json:"repoBranch,omitempty"`

	// SyncPolicy define la política de sincronización automática
	// +optional
	SyncPolicy *ArgoCDSyncPolicy `json:"syncPolicy,omitempty"`

	// Clusters es la lista de clusters registrados en ArgoCD
	// +optional
	Clusters []ArgoCDClusterRef `json:"clusters,omitempty"`
}

// ArgoCDSyncPolicy define la política de sincronización de ArgoCD
type ArgoCDSyncPolicy struct {
	// AutoSync habilita la sincronización automática
	// +optional
	// +kubebuilder:default=true
	AutoSync bool `json:"autoSync,omitempty"`

	// RetryConfiguration contiene la configuración de reintentos
	// +optional
	RetryConfiguration *ArgoCDRetryConfig `json:"retryConfiguration,omitempty"`

	// PrunePropagationPolicy define cómo se propagan las eliminaciones
	// +optional
	// +kubebuilder:validation:Enum=foreground;background;orphan
	// +kubebuilder:default="foreground"
	PrunePropagationPolicy string `json:"prunePropagationPolicy,omitempty"`

	// SelfHeal habilita la auto-recuperación
	// +optional
	// +kubebuilder:default=true
	SelfHeal bool `json:"selfHeal,omitempty"`
}

// ArgoCDRetryConfig define la configuración de reintentos
type ArgoCDRetryConfig struct {
	// Limit es el número máximo de reintentos
	// +optional
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:default=3
	Limit int `json:"limit,omitempty"`

	// Backoff define la estrategia de backoff
	// +optional
	Backoff *ArgoCDRetryBackoff `json:"backoff,omitempty"`
}

// ArgoCDRetryBackoff define el backoff para reintentos
type ArgoCDRetryBackoff struct {
	// Duration es la duración inicial del backoff
	// +optional
	// +kubebuilder:default="5s"
	Duration string `json:"duration,omitempty"`

	// Factor es el factor multiplicador del backoff
	// +optional
	// +kubebuilder:validation:Minimum=1.0
	// +kubebuilder:default=2.0
	Factor float64 `json:"factor,omitempty"`

	// MaxDuration es la duración máxima del backoff
	// +optional
	// +kubebuilder:default="3m"
	MaxDuration string `json:"maxDuration,omitempty"`
}

// ArgoCDClusterRef referencia un cluster en ArgoCD
type ArgoCDClusterRef struct {
	// Name es el nombre del cluster en ArgoCD
	Name string `json:"name"`

	// ConfigServer es la URL del servidor de configuración
	// +optional
	ConfigServer string `json:"configServer,omitempty"`

	// Labels contiene etiquetas para el cluster en ArgoCD
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
}

// CrossplaneSpec define la configuración de Crossplane
type CrossplaneSpec struct {
	// Enabled indica si Crossplane está habilitado
	// +optional
	// +kubebuilder:default=true
	Enabled bool `json:"enabled,omitempty"`

	// Version es la versión de Crossplane a desplegar
	// +optional
	// +kubebuilder:default="v1.13.0"
	Version string `json:"version,omitempty"`

	// ProviderConfigs contiene las configuraciones de proveedores
	// +optional
	ProviderConfigs []CrossplaneProviderConfig `json:"providerConfigs,omitempty"`

	// CompositionRevisions contiene las revisiones de composiciones
	// +optional
	CompositionRevisions []CompositionRevisionSpec `json:"compositionRevisions,omitempty"`
}

// CrossplaneProviderConfig define la configuración de un proveedor
type CrossplaneProviderConfig struct {
	// Name es el nombre del provider config
	Name string `json:"name"`

	// Region es la región del proveedor
	Region string `json:"region"`

	// CredentialsSecretRef referencia el secret con credenciales
	CredentialsSecretRef SecretRef `json:"credentialsSecretRef"`
}

// CompositionRevisionSpec define una revisión de composición
type CompositionRevisionSpec struct {
	// Name es el nombre de la composición
	Name string `json:"name"`

	// Revision es el número de revisión
	// +kubebuilder:validation:Minimum=1
	Revision int `json:"revision"`
}

// VaultSpec define la configuración de Vault
type VaultSpec struct {
	// Enabled indica si Vault está habilitado
	// +optional
	// +kubebuilder:default=true
	Enabled bool `json:"enabled,omitempty"`

	// Address es la dirección del servidor Vault
	// +kubebuilder:validation:Format=uri
	Address string `json:"address"`

	// Namespace es el namespace de Vault (para Vault Enterprise)
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// AuthMethod es el método de autenticación
	// +optional
	// +kubebuilder:validation:Enum=kubernetes;aws;azure;gcp;token
	// +kubebuilder:default="kubernetes"
	AuthMethod string `json:"authMethod,omitempty"`

	// SecretsEngine es el motor de secrets a usar
	// +optional
	// +kubebuilder:validation:Enum=kv-v2;kv-v1;dynamic
	// +kubebuilder:default="kv-v2"
	SecretsEngine string `json:"secretsEngine,omitempty"`

	// Paths contiene las rutas de secrets a replicar
	// +optional
	Paths []string `json:"paths,omitempty"`

	// ReplicationInterval es el intervalo de replicación de secrets
	// +optional
	// +kubebuilder:default="5m"
	ReplicationInterval string `json:"replicationInterval,omitempty"`
}

// AWSSpec define la configuración de AWS
type AWSSpec struct {
	// Regions contiene la configuración por región
	Regions []AWSRegionConfig `json:"regions"`

	// HealthCheckInterval es el intervalo de health checks
	// +optional
	// +kubebuilder:default="30s"
	HealthCheckInterval string `json:"healthCheckInterval,omitempty"`

	// FailoverThreshold define el umbral para triggering de failover
	// +optional
	FailoverThreshold *AWSFailoverThreshold `json:"failoverThreshold,omitempty"`
}

// AWSRegionConfig define la configuración de una región AWS
type AWSRegionConfig struct {
	// Name es el nombre de la región (ej. us-east-1)
	Name string `json:"name"`

	// IsPrimary indica si es la región primaria
	// +optional
	IsPrimary bool `json:"isPrimary,omitempty"`

	// HealthCheckConfig contiene la configuración de health check
	// +optional
	HealthCheckConfig *AWSHealthCheckConfig `json:"healthCheckConfig,omitempty"`
}

// AWSHealthCheckConfig define la configuración de health check
type AWSHealthCheckConfig struct {
	// Endpoint es el endpoint a verificar
	// +optional
	Endpoint string `json:"endpoint,omitempty"`

	// Timeout es el timeout del health check
	// +optional
	// +kubebuilder:default="5s"
	Timeout string `json:"timeout,omitempty"`

	// Interval es el intervalo entre health checks
	// +optional
	// +kubebuilder:default="10s"
	Interval string `json:"interval,omitempty"`
}

// AWSFailoverThreshold define los umbrales para failover
type AWSFailoverThreshold struct {
	// MinHealthyClusters es el número mínimo de clusters saludables para no hacer failover
	// +kubebuilder:validation:Minimum=1
	MinHealthyClusters int `json:"minHealthyClusters"`

	// MaxLatencyMs es la latencia máxima aceptable en milisegundos
	// +kubebuilder:validation:Minimum=0
	MaxLatencyMs int `json:"maxLatencyMs"`

	// ErrorRateThreshold es el umbral de tasa de errores (0.0 - 1.0)
	// +kubebuilder:validation:Minimum=0.0
	// +kubebuilder:validation:Maximum=1.0
	ErrorRateThreshold float64 `json:"errorRateThreshold"`
}

// ReplicationSpec define la configuración de replicación de estado
type ReplicationSpec struct {
	// Backend es el backend de replicación a usar
	// +optional
	// +kubebuilder:validation:Enum=etcd;consul;dynamodb
	// +kubebuilder:default="etcd"
	Backend string `json:"backend,omitempty"`

	// Endpoints son los endpoints de los clusters de replicación
	// +kubebuilder:validation:MinItems=3
	Endpoints []string `json:"endpoints"`

	// QuorumSize es el tamaño del quórum para consistencia
	// +optional
	// +kubebuilder:validation:Minimum=1
	QuorumSize int `json:"quorumSize,omitempty"`

	// SyncInterval es el intervalo de sincronización
	// +optional
	// +kubebuilder:default="1s"
	SyncInterval string `json:"syncInterval,omitempty"`

	// ConsistencyMode define el modo de consistencia
	// +optional
	// +kubebuilder:validation:Enum=eventual;strong
	// +kubebuilder:default="eventual"
	ConsistencyMode string `json:"consistencyMode,omitempty"`
}

// FailoverSpec define la configuración de disaster recovery
type FailoverSpec struct {
	// Enabled indica si el failover está habilitado
	// +optional
	// +kubebuilder:default=true
	Enabled bool `json:"enabled,omitempty"`

	// Strategy define la estrategia de failover
	// +optional
	// +kubebuilder:validation:Enum=automatic;manual;semi-automatic
	// +kubebuilder:default="automatic"
	Strategy string `json:"strategy,omitempty"`

	// DetectionTimeout es el timeout para detección de fallos
	// +optional
	// +kubebuilder:default="30s"
	DetectionTimeout string `json:"detectionTimeout,omitempty"`

	// PromotionTimeout es el timeout para promoción de nuevo primary
	// +optional
	// +kubebuilder:default="5m"
	PromotionTimeout string `json:"promotionTimeout,omitempty"`

	// HealthCheckRetries es el número de reintentos para health checks
	// +optional
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:default=3
	HealthCheckRetries int `json:"healthCheckRetries,omitempty"`

	// PrePromotionHooks son hooks a ejecutar antes de la promoción
	// +optional
	PrePromotionHooks []string `json:"prePromotionHooks,omitempty"`

	// PostPromotionHooks son hooks a ejecutar después de la promoción
	// +optional
	PostPromotionHooks []string `json:"postPromotionHooks,omitempty"`
}

// SecretRef referencia un Secret de Kubernetes
type SecretRef struct {
	// Name es el nombre del secret
	Name string `json:"name"`

	// Namespace es el namespace del secret
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

// ControlPlaneStatus define el estado observado del control plane
type ControlPlaneStatus struct {
	// ObservedGeneration es la generación observada
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Clusters contiene el estado de cada cluster
	// +optional
	Clusters []ClusterStatus `json:"clusters,omitempty"`

	// ArgoCD contiene el estado de ArgoCD
	// +optional
	ArgoCD *ArgoCDStatus `json:"argocd,omitempty"`

	// Crossplane contiene el estado de Crossplane
	// +optional
	Crossplane *CrossplaneStatus `json:"crossplane,omitempty"`

	// Vault contiene el estado de Vault
	// +optional
	Vault *VaultStatus `json:"vault,omitempty"`

	// Conditions contiene las condiciones actuales
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// FailoverStatus contiene el estado del failover
	// +optional
	FailoverStatus *FailoverStatus `json:"failoverStatus,omitempty"`
}

// ClusterStatus contiene el estado de un cluster
type ClusterStatus struct {
	// Name es el nombre del cluster
	Name string `json:"name"`

	// Region es la región del cluster
	Region string `json:"region"`

	// Ready indica si el cluster está listo
	// +optional
	Ready bool `json:"ready,omitempty"`

	// Health indica el estado de salud del cluster
	// +optional
	// +kubebuilder:validation:Enum=healthy;degraded;unhealthy;unknown
	Health string `json:"health,omitempty"`

	// LastHeartbeat es la última vez que se recibió heartbeat
	// +optional
	LastHeartbeat *metav1.Time `json:"lastHeartbeat,omitempty"`

	// Message contiene mensajes adicionales
	// +optional
	Message string `json:"message,omitempty"`
}

// ArgoCDStatus contiene el estado de ArgoCD
type ArgoCDStatus struct {
	// Ready indica si ArgoCD está listo
	// +optional
	Ready bool `json:"ready,omitempty"`

	// Version es la versión desplegada
	// +optional
	Version string `json:"version,omitempty"`

	// SyncStatus es el estado de sincronización
	// +optional
	SyncStatus string `json:"syncStatus,omitempty"`

	// RegisteredClusters es el número de clusters registrados
	// +optional
	RegisteredClusters int `json:"registeredClusters,omitempty"`

	// Message contiene mensajes adicionales
	// +optional
	Message string `json:"message,omitempty"`
}

// CrossplaneStatus contiene el estado de Crossplane
type CrossplaneStatus struct {
	// Ready indica si Crossplane está listo
	// +optional
	Ready bool `json:"ready,omitempty"`

	// Version es la versión desplegada
	// +optional
	Version string `json:"version,omitempty"`

	// ProviderConfigsReady es el número de provider configs listos
	// +optional
	ProviderConfigsReady int `json:"providerConfigsReady,omitempty"`

	// CompositionsReady es el número de composiciones listas
	// +optional
	CompositionsReady int `json:"compositionsReady,omitempty"`

	// Message contiene mensajes adicionales
	// +optional
	Message string `json:"message,omitempty"`
}

// VaultStatus contiene el estado de Vault
type VaultStatus struct {
	// Ready indica si Vault está listo
	// +optional
	Ready bool `json:"ready,omitempty"`

	// Address es la dirección de Vault
	// +optional
	Address string `json:"address,omitempty"`

	// Sealed indica si Vault está sellado
	// +optional
	Sealed bool `json:"sealed,omitempty"`

	// ReplicationStatus es el estado de replicación de secrets
	// +optional
	ReplicationStatus string `json:"replicationStatus,omitempty"`

	// Message contiene mensajes adicionales
	// +optional
	Message string `json:"message,omitempty"`
}

// FailoverStatus contiene el estado del failover
type FailoverStatus struct {
	// Active indica si hay un failover activo
	// +optional
	Active bool `json:"active,omitempty"`

	// SourceCluster es el cluster de origen (el que falló)
	// +optional
	SourceCluster string `json:"sourceCluster,omitempty"`

	// TargetCluster es el cluster objetivo (el nuevo primary)
	// +optional
	TargetCluster string `json:"targetCluster,omitempty"`

	// StartedAt es cuando empezó el failover
	// +optional
	StartedAt *metav1.Time `json:"startedAt,omitempty"`

	// CompletedAt es cuando completó el failover
	// +optional
	CompletedAt *metav1.Time `json:"completedAt,omitempty"`

	// State es el estado actual del failover
	// +optional
	// +kubebuilder:validation:Enum=detecting;promoting;replicating;completed;failed
	State string `json:"state,omitempty"`

	// Message contiene mensajes adicionales
	// +optional
	Message string `json:"message,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:path=controlplanes,scope=Cluster,shortName=cp
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status"
// +kubebuilder:printcolumn:name="Clusters",type="integer",JSONPath=".status.clusters\.size()"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// ControlPlane es el recurso que representa la configuración del control plane multi-cluster
type ControlPlane struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec es la especificación del control plane
	Spec ControlPlaneSpec `json:"spec,omitempty"`

	// Status es el estado observado del control plane
	// +optional
	Status ControlPlaneStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ControlPlaneList contiene una lista de ControlPlane
type ControlPlaneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ControlPlane `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ControlPlane{}, &ControlPlaneList{})
}

// Helper functions para condiciones
const (
	ConditionArgoCDReady     = "ArgoCDReady"
	ConditionCrossplaneReady = "CrossplaneReady"
	ConditionVaultReady      = "VaultReady"
	ConditionClustersReady   = "ClustersReady"
	ConditionFailoverReady   = "FailoverReady"
	ConditionReplicating     = "Replicating"
)

// IsReady verifica si el control plane está listo
func (cp *ControlPlane) IsReady() bool {
	for _, cond := range cp.Status.Conditions {
		if cond.Type == "Ready" && cond.Status == metav1.ConditionTrue {
			return true
		}
	}
	return false
}

// GetActiveFailover retorna el failover activo si existe
func (cp *ControlPlane) GetActiveFailover() *FailoverStatus {
	if cp.Status.FailoverStatus != nil && cp.Status.FailoverStatus.Active {
		return cp.Status.FailoverStatus
	}
	return nil
}

// GetPrimaryCluster retorna el cluster primario
func (cp *ControlPlane) GetPrimaryCluster() *ClusterSpec {
	for i := range cp.Spec.Clusters {
		if cp.Spec.Clusters[i].IsMain {
			return &cp.Spec.Clusters[i]
		}
	}
	return nil
}

// GetClustersByRegion retorna los clusters de una región específica
func (cp *ControlPlane) GetClustersByRegion(region string) []ClusterSpec {
	var result []ClusterSpec
	for _, c := range cp.Spec.Clusters {
		if c.Region == region {
			result = append(result, c)
		}
	}
	return result
}