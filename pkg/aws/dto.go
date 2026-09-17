package aws

import (
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/route53"
)

type HealthCheckMessage struct {
	ClusterID       string            `json:"cluster_id"`
	Region          string            `json:"region"`
	Timestamp       time.Time         `json:"timestamp"`
	Status          HealthStatus      `json:"status"`
	Components      []ComponentStatus `json:"components"`
	Metadata        map[string]string `json:"metadata,omitempty"`
}

type HealthStatus string

const (
	HealthStatusHealthy   HealthStatus = "healthy"
	HealthStatusDegraded  HealthStatus = "degraded"
	HealthStatusUnhealthy HealthStatus = "unhealthy"
	HealthStatusUnknown   HealthStatus = "unknown"
)

type ComponentStatus struct {
	Name               string                 `json:"name"`
	Type               ComponentType          `json:"type"`
	Status             HealthStatus           `json:"status"`
	Message            string                 `json:"message,omitempty"`
	LastCheck          time.Time              `json:"last_check"`
	AvailabilityZones  []string               `json:"availability_zones,omitempty"`
	TargetGroupHealth  *TargetGroupHealth     `json:"target_group_health,omitempty"`
	Route53HealthCheck *Route53HealthCheckDTO `json:"route53_health,omitempty"`
}

type ComponentType string

const (
	ComponentTypeAPIServer    ComponentType = "api_server"
	ComponentTypeETCD         ComponentType = "etcd"
	ComponentTypeWorkerNode   ComponentType = "worker_node"
	ComponentTypeLoadBalancer ComponentType = "load_balancer"
	ComponentTypeRoute53      ComponentType = "route53"
	ComponentTypeCrossplane   ComponentType = "crossplane"
	ComponentTypeArgoCD       ComponentType = "argocd"
)

type TargetGroupHealth struct {
	TargetGroupARN  string              `json:"target_group_arn"`
	TargetGroupName string              `json:"target_group_name"`
	LoadBalancerARN string              `json:"load_balancer_arn"`
	HealthyTargets  int64               `json:"healthy_targets"`
	UnhealthyTargets int64              `json:"unhealthy_targets"`
	TotalTargets    int64               `json:"total_targets"`
	TargetStatuses  []TargetStatusDetail `json:"target_statuses"`
}

type TargetStatusDetail struct {
	TargetID  string         `json:"target_id"`
	Port      int64          `json:"port"`
	Health    string         `json:"health"`
	Reason    string         `json:"reason,omitempty"`
}

type Route53HealthCheckDTO struct {
	HealthCheckID        string   `json:"health_check_id"`
	HealthCheckStatus    string   `json:"health_check_status"`
	HealthCheckType      string   `json:"health_check_type"`
	FailureReason        string   `json:"failure_reason,omitempty"`
	Regions              []string `json:"regions"`
	IPAddress            string   `json:"ip_address,omitempty"`
	FullyQualifiedDomain string   `json:"fqdn,omitempty"`
}

type FailoverEvent struct {
	EventID            string               `json:"event_id"`
	EventType          FailoverEventType    `json:"event_type"`
	Timestamp          time.Time            `json:"timestamp"`
	SourceCluster      string               `json:"source_cluster"`
	SourceRegion       string               `json:"source_region"`
	TargetCluster      string               `json:"target_cluster,omitempty"`
	TargetRegion       string               `json:"target_region,omitempty"`
	Trigger            FailoverTrigger      `json:"trigger"`
	Severity           FailoverSeverity     `json:"severity"`
	State              FailoverState        `json:"state"`
	AffectedComponents []string             `json:"affected_components"`
	Actions            []FailoverAction     `json:"actions"`
	RollbackAvailable  bool                 `json:"rollback_available"`
	Metadata           map[string]string    `json:"metadata,omitempty"`
}

type FailoverEventType string

const (
	FailoverEventTypePlanned   FailoverEventType = "planned"
	FailoverEventTypeEmergency FailoverEventType = "emergency"
	FailoverEventTypeDrill     FailoverEventType = "drill"
)

type FailoverTrigger string

const (
	FailoverTriggerManual          FailoverTrigger = "manual"
	FailoverTriggerHealthCheck     FailoverTrigger = "health_check"
	FailoverTriggerLatency         FailoverTrigger = "latency_threshold"
	FailoverTriggerErrorRate       FailoverTrigger = "error_rate_threshold"
	FailoverTriggerRegionFailure   FailoverTrigger = "region_failure"
)

type FailoverSeverity string

const (
	FailoverSeverityInfo     FailoverSeverity = "info"
	FailoverSeverityWarning  FailoverSeverity = "warning"
	FailoverSeverityCritical FailoverSeverity = "critical"
)

type FailoverState string

const (
	FailoverStatePending    FailoverState = "pending"
	FailoverStateInProgress FailoverState = "in_progress"
	FailoverStateCompleted  FailoverState = "completed"
	FailoverStateFailed     FailoverState = "failed"
	FailoverStateRolledBack FailoverState = "rolled_back"
)

type FailoverAction struct {
	ActionType FailoverActionType `json:"action_type"`
	Target     string             `json:"target"`
	Parameters map[string]string  `json:"parameters,omitempty"`
	Status     string             `json:"status"`
	Error      string             `json:"error,omitempty"`
}

type FailoverActionType string

const (
	FailoverActionTypeUpdateDNS           FailoverActionType = "update_dns"
	FailoverActionTypeSwitchLoadBalancer  FailoverActionType = "switch_load_balancer"
	FailoverActionTypePromoteReplica      FailoverActionType = "promote_replica"
	FailoverActionTypeSyncState           FailoverActionType = "sync_state"
	FailoverActionTypeNotify              FailoverActionType = "notify"
)

type ClusterHealthSummary struct {
	ClusterID         string            `json:"cluster_id"`
	Region            string            `json:"region"`
	OverallStatus     HealthStatus      `json:"overall_status"`
	LastUpdated       time.Time         `json:"last_updated"`
	APIServerStatus   HealthStatus      `json:"api_server_status"`
	WorkerNodesStatus HealthStatus      `json:"worker_nodes_status"`
	StorageStatus     HealthStatus      `json:"storage_status"`
	NetworkStatus     HealthStatus      `json:"network_status"`
	IntegrationsStatus map[string]HealthStatus `json:"integrations_status"`
	AlertCount        int                `json:"alert_count"`
}

type RegionHealthSummary struct {
	Region            string                  `json:"region"`
	OverallStatus     HealthStatus            `json:"overall_status"`
	Clusters          []ClusterHealthSummary  `json:"clusters"`
	CrossRegionLatency time.Duration          `json:"cross_region_latency,omitempty"`
	FailoverReadiness FailoverReadiness       `json:"failover_readiness"`
}

type FailoverReadiness struct {
	ReadyForFailover bool      `json:"ready_for_failover"`
	LastCheck        time.Time `json:"last_check"`
	StateSyncStatus  string    `json:"state_sync_status"`
	SecretsReady     bool      `json:"secrets_ready"`
	RbacReady        bool      `json:"rbac_ready"`
	BlockingIssues   []string  `json:"blocking_issues,omitempty"`
}

type DNSRecordUpdate struct {
	HostedZoneID     string            `json:"hosted_zone_id"`
	RecordName       string            `json:"record_name"`
	RecordType       string            `json:"record_type"`
	Action           string            `json:"action"`
	TargetValue      string            `json:"target_value"`
	TargetType       string            `json:"target_type"`
	TTL              int64             `json:"ttl"`
	HealthCheckID    string            `json:"health_check_id,omitempty"`
	FailoverRecord   *Route53FailoverRecord `json:"failover_record,omitempty"`
}

type Route53FailoverRecord struct {
	FailoverRecordType string `json:"failover_record_type"`
	SetID             string `json:"set_id"`
	HealthCheckID     string `json:"health_check_id"`
}

type LoadBalancerTarget struct {
	LoadBalancerARN  string `json:"load_balancer_arn"`
	LoadBalancerName string `json:"load_balancer_name"`
	LoadBalancerType string `json:"load_balancer_type"`
	DNSName          string `json:"dns_name"`
	Scheme           string `json:"scheme"`
	VPCID            string `json:"vpc_id"`
	Regions          []string `json:"regions"`
	TargetGroups     []string `json:"target_groups"`
}

type StateSyncMessage struct {
	SyncID          string            `json:"sync_id"`
	SourceCluster   string            `json:"source_cluster"`
	TargetCluster   string            `json:"target_cluster"`
	Timestamp       time.Time         `json:"timestamp"`
	StateVersion    string            `json:"state_version"`
	Resources       []StateResource   `json:"resources"`
	SyncStatus      StateSyncStatus   `json:"sync_status"`
	ConflictResolution string         `json:"conflict_resolution,omitempty"`
}

type StateResource struct {
	ResourceType string `json:"resource_type"`
	ResourceName string `json:"resource_name"`
	Namespace    string `json:namespace,omitempty"`
	Version      string `json:"version"`
	Checksum     string `json:"checksum"`
}

type StateSyncStatus string

const (
	StateSyncStatusPending   StateSyncStatus = "pending"
	StateSyncStatusInFlight StateSyncStatus = "in_flight"
	StateSyncStatusSynced   StateSyncStatus = "synced"
	StateSyncStatusFailed   StateSyncStatus = "failed"
	StateSyncStatusConflict StateSyncStatus = "conflict"
)