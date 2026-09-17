package failover

import (
	"context"
	"fmt"
	"sync"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/example/controlplane-operator/pkg/healthcheck"
)

const (
	DefaultFailureThreshold    = 3
	DefaultFailureWindow       = 3 * time.Minute
	PromotionTimeout           = 30 * time.Second
	HealthCheckInterval        = 30 * time.Second
)

type ClusterRole string

const (
	RolePrimary   ClusterRole = "primary"
	RoleStandby   ClusterRole = "standby"
	RolePromoting ClusterRole = "promoting"
)

type FailureRecord struct {
	ClusterID    string
	FailureCount int
	FirstFailure time.Time
	LastFailure  time.Time
}

type Mediator struct {
	client         client.Client
	clusterStates  map[string]ClusterState
	failureRecords map[string]*FailureRecord
	mutex          sync.RWMutex
	config         *MediatorConfig
}

type MediatorConfig struct {
	FailureThreshold int
	FailureWindow    time.Duration
	QuorumSize       int
	EnableSplitBrain bool
}

type ClusterState struct {
	ClusterID     string
	Role          ClusterRole
	Region        string
	IsHealthy     bool
	LastHealth    time.Time
	PromotedAt    *time.Time
	FailoverCount int
}

func NewMediator(cli client.Client, cfg *MediatorConfig) *Mediator {
	if cfg == nil {
		cfg = &MediatorConfig{
			FailureThreshold: DefaultFailureThreshold,
			FailureWindow:    DefaultFailureWindow,
			QuorumSize:       2,
			EnableSplitBrain: false,
		}
	}
	return &Mediator{
		client:         cli,
		clusterStates:  make(map[string]ClusterState),
		failureRecords: make(map[string]*FailureRecord),
		config:         cfg,
	}
}

func (m *Mediator) RegisterCluster(ctx context.Context, clusterID, region string, initialRole ClusterRole) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.clusterStates[clusterID] = ClusterState{
		ClusterID: clusterID,
		Role:      initialRole,
		Region:    region,
		IsHealthy: true,
		LastHealth: time.Now(),
	}

	m.failureRecords[clusterID] = &FailureRecord{
		ClusterID:    clusterID,
		FailureCount: 0,
	}

	log.FromContext(ctx).Info("Cluster registered in mediator",
		"cluster", clusterID,
		"region", region,
		"role", initialRole)

	return nil
}

func (m *Mediator) HandleHealthCheckResult(ctx context.Context, clusterID string, healthy bool) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	logger := log.FromContext(ctx)
	state, exists := m.clusterStates[clusterID]
	if !exists {
		return fmt.Errorf("cluster %s not registered", clusterID)
	}

	record := m.failureRecords[clusterID]

	if healthy {
		state.IsHealthy = true
		state.LastHealth = time.Now()
		m.clusterStates[clusterID] = state

		if record.FailureCount > 0 {
			logger.Info("Cluster recovered", "cluster", clusterID, "previous_failures", record.FailureCount)
			record.FailureCount = 0
			record.FirstFailure = time.Time{}
		}
		return nil
	}

	state.IsHealthy = false
	m.clusterStates[clusterID] = state

	now := time.Now()
	if record.FailureCount == 0 {
		record.FirstFailure = now
	}
	record.LastFailure = now
	record.FailureCount++

	logger.Info("Health check failed",
		"cluster", clusterID,
		"failure_count", record.FailureCount,
		"threshold", m.config.FailureThreshold)

	if m.shouldTriggerFailover(ctx, clusterID, record) {
		return m.initiateFailover(ctx, clusterID)
	}

	return nil
}

func (m *Mediator) shouldTriggerFailover(ctx context.Context, clusterID string, record *FailureRecord) bool {
	if record.FailureCount < m.config.FailureThreshold {
		return false
	}

	windowExpired := time.Since(record.FirstFailure) > m.config.FailureWindow
	if !windowExpired {
		return false
	}

	logger := log.FromContext(ctx)
	logger.Info("Failure threshold reached",
		"cluster", clusterID,
		"failures", record.FailureCount,
		"window", m.config.FailureWindow)

	return true
}

func (m *Mediator) initiateFailover(ctx context.Context, failedClusterID string) error {
	logger := log.FromContext(ctx)

	failedState := m.clusterStates[failedClusterID]
	if failedState.Role != RolePrimary {
		logger.Info("Failover skipped - cluster is not primary", "cluster", failedClusterID, "role", failedState.Role)
		return nil
	}

	if !m.config.EnableSplitBrain {
		if !m.hasQuorum(ctx, failedState.Region) {
			return fmt.Errorf("cannot failover: insufficient quorum in region %s", failedState.Region)
		}
	}

	standbyClusters := m.getStandbyClustersInRegion(failedState.Region, failedClusterID)
	if len(standbyClusters) == 0 {
		return fmt.Errorf("no standby clusters available in region %s", failedState.Region)
	}

	promotedCluster := standbyClusters[0]
	promotedState := m.clusterStates[promotedCluster]

	logger.Info("Initiating failover",
		"from", failedClusterID,
		"to", promotedCluster,
		"region", failedState.Region)

	promotedState.Role = RolePromoting
	m.clusterStates[promotedCluster] = promotedState

	if err := m.promoteCluster(ctx, promotedCluster); err != nil {
		logger.Error(err, "Failed to promote cluster", "cluster", promotedCluster)
		promotedState.Role = RoleStandby
		m.clusterStates[promotedCluster] = promotedState
		return err
	}

	now := time.Now()
	promotedState.Role = RolePrimary
	promotedState.PromotedAt = &now
	promotedState.FailoverCount++
	m.clusterStates[promotedCluster] = promotedState

	failedState.Role = RoleStandby
	m.clusterStates[failedClusterID] = failedState

	m.failureRecords[failedClusterID].FailureCount = 0

	logger.Info("Failover completed",
		"new_primary", promotedCluster,
		"old_primary", failedClusterID,
		"failover_count", promotedState.FailoverCount)

	return nil
}

func (m *Mediator) promoteCluster(ctx context.Context, clusterID string) error {
	logger := log.FromContext(ctx)

	logger.Info("Executing promotion sequence", "cluster", clusterID)

	time.Sleep(100 * time.Millisecond)

	logger.Info("Promotion completed", "cluster", clusterID)
	return nil
}

func (m *Mediator) getStandbyClustersInRegion(region, excludeID string) []string {
	var standbys []string
	for id, state := range m.clusterStates {
		if id != excludeID && state.Region == region && state.Role == RoleStandby && state.IsHealthy {
			standbys = append(standbys, id)
		}
	}
	return standbys
}

func (m *Mediator) hasQuorum(ctx context.Context, region string) bool {
	healthyCount := 0
	for _, state := range m.clusterStates {
		if state.Region == region && state.IsHealthy {
			healthyCount++
		}
	}
	return healthyCount >= m.config.QuorumSize
}

func (m *Mediator) GetClusterState(clusterID string) (ClusterState, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	state, ok := m.clusterStates[clusterID]
	return state, ok
}

func (m *Mediator) GetAllClusterStates() map[string]ClusterState {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	copyStates := make(map[string]ClusterState, len(m.clusterStates))
	for k, v := range m.clusterStates {
		copyStates[k] = v
	}
	return copyStates
}

func (m *Mediator) GetPrimaryCluster(region string) (string, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	for id, state := range m.clusterStates {
		if state.Region == region && state.Role == RolePrimary && state.IsHealthy {
			return id, true
		}
	}
	return "", false
}

func (m *Mediator) RunHealthCheckLoop(ctx context.Context, checker healthcheck.ClusterHealthChecker) error {
	ticker := time.NewTicker(HealthCheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			m.mutex.RLock()
			clusterIDs := make([]string, 0, len(m.clusterStates))
			for id := range m.clusterStates {
				clusterIDs = append(clusterIDs, id)
			}
			m.mutex.RUnlock()

			for _, clusterID := range clusterIDs {
				healthy, err := checker.CheckClusterHealth(ctx, clusterID)
				if err != nil {
					log.FromContext(ctx).Error(err, "Health check error", "cluster", clusterID)
					continue
				}
				if err := m.HandleHealthCheckResult(ctx, clusterID, healthy); err != nil {
					log.FromContext(ctx).Error(err, "Failed to handle health result", "cluster", clusterID)
				}
			}
		}
	}
}

func (m *Mediator) ListClustersByRole(role ClusterRole) []ClusterState {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	var result []ClusterState
	for _, state := range m.clusterStates {
		if state.Role == role {
			result = append(result, state)
		}
	}
	return result
}

func (m *Mediator) ListClustersByRegion(region string) []ClusterState {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	var result []ClusterState
	for _, state := range m.clusterStates {
		if state.Region == region {
			result = append(result, state)
		}
	}
	return result
}

var _ healthcheck.ClusterHealthChecker = (*Mediator)(nil)