package healthcheck

import (
	"context"
	"fmt"
	"sync"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	DefaultTimeout       = 30 * time.Second
	DefaultRetries       = 3
	DefaultRetryInterval = 5 * time.Second
	ComponentAPIServer   = "api-server"
	ComponentEtcd        = "etcd"
	ComponentScheduler   = "scheduler"
	ComponentController  = "controller-manager"
	ComponentProxy       = "kube-proxy"
)

type CheckResult struct {
	ClusterID    string
	Timestamp    time.Time
	Healthy      bool
	Components   map[string]ComponentStatus
	LatencyMs    int64
	ErrorMessage string
}

type ComponentStatus struct {
	Name      string
	Healthy   bool
	LatencyMs int64
	Message   string
}

type ClusterHealthChecker interface {
	CheckClusterHealth(ctx context.Context, clusterID string) (bool, error)
	GetClusterStatus(ctx context.Context, clusterID string) (*CheckResult, error)
	CheckComponent(ctx context.Context, clusterID, component string) (*ComponentStatus, error)
}

type HealthCheckConfig struct {
	Timeout           time.Duration
	Retries           int
	RetryInterval     time.Duration
	EnableDeepCheck   bool
	CheckComponents   []string
}

type DistributedChecker struct {
	client       client.Client
	config       *HealthCheckConfig
	resultCache  map[string]*CheckResult
	cacheMutex   sync.RWMutex
	regionGroups map[string][]string
}

func NewDistributedChecker(cli client.Client, cfg *HealthCheckConfig) *DistributedChecker {
	if cfg == nil {
		cfg = &HealthCheckConfig{
			Timeout:         DefaultTimeout,
			Retries:         DefaultRetries,
			RetryInterval:   DefaultRetryInterval,
			EnableDeepCheck: true,
			CheckComponents: []string{ComponentAPIServer, ComponentEtcd, ComponentScheduler, ComponentController},
		}
	}

	return &DistributedChecker{
		client:       cli,
		config:       cfg,
		resultCache:  make(map[string]*CheckResult),
		regionGroups: make(map[string][]string),
	}
}

func (d *DistributedChecker) RegisterClusterToRegion(ctx context.Context, clusterID, region string) error {
	d.regionGroups[region] = append(d.regionGroups[region], clusterID)
	log.FromContext(ctx).Info("Cluster registered for health checks", "cluster", clusterID, "region", region)
	return nil
}

func (d *DistributedChecker) CheckClusterHealth(ctx context.Context, clusterID string) (bool, error) {
	result, err := d.GetClusterStatus(ctx, clusterID)
	if err != nil {
		return false, err
	}
	return result.Healthy, nil
}

func (d *DistributedChecker) GetClusterStatus(ctx context.Context, clusterID string) (*CheckResult, error) {
	d.cacheMutex.RLock()
	if cached, ok := d.resultCache[clusterID]; ok && time.Since(cached.Timestamp) < d.config.Timeout {
		d.cacheMutex.RUnlock()
		return cached, nil
	}
	d.cacheMutex.RUnlock()

	logger := log.FromContext(ctx)
	startTime := time.Now()

	result := &CheckResult{
		ClusterID:  clusterID,
		Timestamp:  startTime,
		Components: make(map[string]ComponentStatus),
	}

	if d.config.EnableDeepCheck {
		for _, component := range d.config.CheckComponents {
			compStatus, err := d.CheckComponent(ctx, clusterID, component)
			if err != nil {
				logger.Error(err, "Component check failed", "cluster", clusterID, "component", component)
				result.Components[component] = ComponentStatus{
					Name:    component,
					Healthy: false,
					Message: err.Error(),
				}
				continue
			}
			result.Components[component] = *compStatus
		}
	} else {
		apiHealthy, err := d.checkAPIServer(ctx, clusterID)
		if err != nil {
			result.Healthy = false
			result.ErrorMessage = err.Error()
			d.updateCache(clusterID, result)
			return result, nil
		}
		result.Components[ComponentAPIServer] = ComponentStatus{
			Name:    ComponentAPIServer,
			Healthy: apiHealthy,
		}
	}

	result.Healthy = d.isClusterHealthy(result)
	result.LatencyMs = time.Since(startTime).Milliseconds()

	d.updateCache(clusterID, result)

	logger.Info("Health check completed",
		"cluster", clusterID,
		"healthy", result.Healthy,
		"latency_ms", result.LatencyMs)

	return result, nil
}

func (d *DistributedChecker) CheckComponent(ctx context.Context, clusterID, component string) (*ComponentStatus, error) {
	startTime := time.Now()

	switch component {
	case ComponentAPIServer:
		healthy, err := d.checkAPIServer(ctx, clusterID)
		return &ComponentStatus{
			Name:      component,
			Healthy:   healthy,
			LatencyMs: time.Since(startTime).Milliseconds(),
			Message:   "API server responding",
		}, err
	case ComponentEtcd:
		healthy, err := d.checkEtcd(ctx, clusterID)
		return &ComponentStatus{
			Name:      component,
			Healthy:   healthy,
			LatencyMs: time.Since(startTime).Milliseconds(),
			Message:   "etcd cluster healthy",
		}, err
	case ComponentScheduler:
		healthy, err := d.checkScheduler(ctx, clusterID)
		return &ComponentStatus{
			Name:      component,
			Healthy:   healthy,
			LatencyMs: time.Since(startTime).Milliseconds(),
			Message:   "scheduler operational",
		}, err
	case ComponentController:
		healthy, err := d.checkControllerManager(ctx, clusterID)
		return &ComponentStatus{
			Name:      component,
			Healthy:   healthy,
			LatencyMs: time.Since(startTime).Milliseconds(),
			Message:   "controller manager running",
		}, err
	default:
		return nil, fmt.Errorf("unknown component: %s", component)
	}
}

func (d *DistributedChecker) checkAPIServer(ctx context.Context, clusterID string) (bool, error) {
	time.Sleep(10 * time.Millisecond)
	return true, nil
}

func (d *DistributedChecker) checkEtcd(ctx context.Context, clusterID string) (bool, error) {
	time.Sleep(20 * time.Millisecond)
	return true, nil
}

func (d *DistributedChecker) checkScheduler(ctx context.Context, clusterID string) (bool, error) {
	time.Sleep(15 * time.Millisecond)
	return true, nil
}

func (d *DistributedChecker) checkControllerManager(ctx context.Context, clusterID string) (bool, error) {
	time.Sleep(15 * time.Millisecond)
	return true, nil
}

func (d *DistributedChecker) isClusterHealthy(result *CheckResult) bool {
	if len(result.Components) == 0 {
		return false
	}
	for _, comp := range result.Components {
		if !comp.Healthy {
			return false
		}
	}
	return true
}

func (d *DistributedChecker) updateCache(clusterID string, result *CheckResult) {
	d.cacheMutex.Lock()
	defer d.cacheMutex.Unlock()
	d.resultCache[clusterID] = result
}

func (d *DistributedChecker) GetCachedResult(clusterID string) (*CheckResult, bool) {
	d.cacheMutex.RLock()
	defer d.cacheMutex.RUnlock()
	result, ok := d.resultCache[clusterID]
	return result, ok
}

func (d *DistributedChecker) InvalidateCache(clusterID string) {
	d.cacheMutex.Lock()
	defer d.cacheMutex.Unlock()
	delete(d.resultCache, clusterID)
}

func (d *DistributedChecker) GetRegionStatus(ctx context.Context, region string) (*CheckResult, error) {
	d.cacheMutex.RLock()
	clusterIDs := d.regionGroups[region]
	d.cacheMutex.RUnlock()

	if len(clusterIDs) == 0 {
		return nil, fmt.Errorf("no clusters registered in region %s", region)
	}

	regionResult := &CheckResult{
		ClusterID:  region,
		Timestamp:  time.Now(),
		Components: make(map[string]ComponentStatus),
		Healthy:    true,
	}

	healthyClusters := 0
	for _, clusterID := range clusterIDs {
		result, err := d.GetClusterStatus(ctx, clusterID)
		if err != nil {
			regionResult.Healthy = false
			continue
		}
		if result.Healthy {
			healthyClusters++
		}
		regionResult.Components[clusterID] = ComponentStatus{
			Name:    clusterID,
			Healthy: result.Healthy,
			Message: fmt.Sprintf("%d/%d components healthy",
				len(result.Components), len(result.Components)),
		}
	}

	if healthyClusters == 0 {
		regionResult.Healthy = false
		regionResult.ErrorMessage = "no healthy clusters in region"
	}

	return regionResult, nil
}

func (d *DistributedChecker) GetAllRegionsStatus(ctx context.Context) (map[string]*CheckResult, error) {
	results := make(map[string]*CheckResult)

	d.cacheMutex.RLock()
	regions := make([]string, 0, len(d.regionGroups))
	for region := range d.regionGroups {
		regions = append(regions, region)
	}
	d.cacheMutex.RUnlock()

	for _, region := range regions {
		result, err := d.GetRegionStatus(ctx, region)
		if err != nil {
			log.FromContext(ctx).Error(err, "Failed to get region status", "region", region)
			continue
		}
		results[region] = result
	}

	return results, nil
}

func (d *DistributedChecker) CheckClusterWithRetry(ctx context.Context, clusterID string) (bool, error) {
	var lastErr error
	for i := 0; i < d.config.Retries; i++ {
		healthy, err := d.CheckClusterHealth(ctx, clusterID)
		if err == nil {
			return healthy, nil
		}
		lastErr = err
		log.FromContext(ctx).Info("Health check retry",
			"cluster", clusterID,
			"attempt", i+1,
			"max", d.config.Retries,
			"error", err)

		if i < d.config.Retries-1 {
			time.Sleep(d.config.RetryInterval)
		}
	}
	return false, lastErr
}