package rbac

import (
	"context"
	"fmt"
	"sync"
	"time"

	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	DefaultSyncInterval  = 5 * time.Minute
	MaxRetryAttempts    = 3
	PropagationTimeout  = 30 * time.Second
)

type SyncDirection string

const (
	SyncDirectionPush SyncDirection = "push"
	SyncDirectionPull SyncDirection = "pull"
	SyncDirectionBoth SyncDirection = "bidirectional"
)

type Permission struct {
	Subjects    []rbacv1.Subject
	RoleRef     rbacv1.RoleRef
	Resources   []string
	Verbs       []string
	ClusterWide bool
}

type SyncStatus struct {
	ClusterID    string
	Timestamp    time.Time
	Success      bool
	Permissions  int
	ErrorMessage string
}

type FederatorConfig struct {
	SyncInterval     time.Duration
	Direction        SyncDirection
	EnableAutoSync   bool
	ConflictStrategy ConflictResolution
}

type ConflictResolution string

const (
	ConflictOverwrite ConflictResolution = "overwrite"
	ConflictMerge     ConflictResolution = "merge"
	ConflictFail      ConflictResolution = "fail"
)

type RBACFederator struct {
	client     client.Client
	config     *FederatorConfig
	policies   map[string]*Permission
	policyMux  sync.RWMutex
	syncStatus map[string]*SyncStatus
	statusMux  sync.RWMutex
	clusters   map[string]ClusterEndpoint
}

type ClusterEndpoint struct {
	ClusterID string
	Region    string
	Endpoint  string
	APIClient client.Client
}

func NewRBACFederator(cli client.Client, cfg *FederatorConfig) *RBACFederator {
	if cfg == nil {
		cfg = &FederatorConfig{
			SyncInterval:     DefaultSyncInterval,
			Direction:        SyncDirectionPush,
			EnableAutoSync:   true,
			ConflictStrategy: ConflictMerge,
		}
	}

	return &RBACFederator{
		client:     cli,
		config:     cfg,
		policies:   make(map[string]*Permission),
		syncStatus: make(map[string]*SyncStatus),
		clusters:   make(map[string]ClusterEndpoint),
	}
}

func (f *RBACFederator) RegisterCluster(ctx context.Context, endpoint ClusterEndpoint) error {
	f.statusMux.Lock()
	defer f.statusMuncUnlock()

	f.clusters[endpoint.ClusterID] = endpoint
	log.FromContext(ctx).Info("Cluster registered for RBAC federation",
		"cluster", endpoint.ClusterID,
		"region", endpoint.Region)

	return nil
}

func (f *RBACFederator) DefineGlobalPolicy(ctx context.Context, policyID string, perm Permission) error {
	f.policyMux.Lock()
	defer f.policyMux.Unlock()

	if len(perm.Verbs) == 0 {
		return fmt.Errorf("permission %s must have at least one verb", policyID)
	}

	if len(perm.Resources) == 0 {
		return fmt.Errorf("permission %s must have at least one resource", policyID)
	}

	f.policies[policyID] = &perm

	log.FromContext(ctx).Info("Global policy defined",
		"policy", policyID,
		"resources", perm.Resources,
		"verbs", perm.Verbs,
		"clusterWide", perm.ClusterWide)

	return nil
}

func (f *RBACFederator) SyncToCluster(ctx context.Context, clusterID string) error {
	f.statusMux.Lock()
	endpoint, exists := f.clusters[clusterID]
	f.statusMux.Unlock()

	if !exists {
		return fmt.Errorf("cluster %s not registered", clusterID)
	}

	logger := log.FromContext(ctx).WithValues("target_cluster", clusterID)

	f.policyMux.RLock()
	policies := make([]*Permission, 0, len(f.policies))
	for _, p := range f.policies {
		policies = append(policies, p)
	}
	f.policyMux.RUnlock()

	successCount := 0
	var lastErr error

	for _, policy := range policies {
		err := f.applyPolicyToCluster(ctx, endpoint, policy)
		if err != nil {
			logger.Error(err, "Failed to apply policy", "policy", policy)
			lastErr = err
			continue
		}
		successCount++
	}

	status := &SyncStatus{
		ClusterID:   clusterID,
		Timestamp:   time.Now(),
		Success:     lastErr == nil,
		Permissions: successCount,
	}
	if lastErr != nil {
		status.ErrorMessage = lastErr.Error()
	}

	f.statusMux.Lock()
	f.syncStatus[clusterID] = status
	f.statusMux.Unlock()

	logger.Info("Sync completed",
		"success", status.Success,
		"permissions", successCount)

	return lastErr
}

func (f *RBACFederator) applyPolicyToCluster(ctx context.Context, endpoint ClusterEndpoint, policy *Permission) error {
	roleName := "federated-role"

	role := &rbacv1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Name:      roleName,
			Namespace: "default",
			Labels: map[string]string{
				"federated":   "true",
				"managed-by": "controlplane-operator",
			},
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups:     []string{"*"},
				Resources:     policy.Resources,
				Verbs:         policy.Verbs,
				ResourceNames: nil,
			},
		},
	}

	err := endpoint.APIClient.Create(ctx, role)
	if err != nil && !errors.IsAlreadyExists(err) {
		return fmt.Errorf("failed to create role: %w", err)
	}

	if errors.IsAlreadyExists(err) {
		existing := &rbacv1.Role{}
		if err := endpoint.APIClient.Get(ctx, client.ObjectKey{Name: roleName, Namespace: "default"}, existing); err != nil {
			return fmt.Errorf("failed to get existing role: %w", err)
		}

		if f.config.ConflictStrategy == ConflictOverwrite {
			existing.Rules = role.Rules
			return endpoint.APIClient.Update(ctx, existing)
		} else if f.config.ConflictStrategy == ConflictMerge {
			existing.Rules = mergeRules(existing.Rules, role.Rules)
			return endpoint.APIClient.Update(ctx, existing)
		}
	}

	roleBinding := &rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "federated-role-binding",
			Namespace: "default",
			Labels: map[string]string{
				"federated":   "true",
				"managed-by": "controlplane-operator",
			},
		},
		Subjects: policy.Subjects,
		RoleRef: rbacv1.RoleRef{
			APIGroup: rbacv1.GroupName,
			Kind:     "Role",
			Name:     roleName,
		},
	}

	err = endpoint.APIClient.Create(ctx, roleBinding)
	if err != nil && !errors.IsAlreadyExists(err) {
		return fmt.Errorf("failed to create role binding: %w", err)
	}

	return nil
}

func mergeRules(existing, new []rbacv1.PolicyRule) []rbacv1.PolicyRule {
	ruleMap := make(map[string]rbacv1.PolicyRule)

	for _, r := range existing {
		key := fmt.Sprintf("%v:%v", r.APIGroups, r.Resources)
		ruleMap[key] = r
	}

	for _, r := range new {
		key := fmt.Sprintf("%v:%v", r.APIGroups, r.Resources)
		if existingRule, ok := ruleMap[key]; ok {
			mergedVerbs := uniqueMerge(existingRule.Verbs, r.Verbs)
			existingRule.Verbs = mergedVerbs
			ruleMap[key] = existingRule
		} else {
			ruleMap[key] = r
		}
	}

	result := make([]rbacv1.PolicyRule, 0, len(ruleMap))
	for _, r := range ruleMap {
		result = append(result, r)
	}

	return result
}

func uniqueMerge(a, b []string) []string {
	existing := make(map[string]bool)
	result := append([]string{}, a...)

	for _, v := range b {
		if !existing[v] {
			existing[v] = true
			result = append(result, v)
		}
	}

	return result
}

func (f *RBACFederator) SyncAllClusters(ctx context.Context) error {
	f.statusMux.RLock()
	clusterIDs := make([]string, 0, len(f.clusters))
	for id := range f.clusters {
		clusterIDs = append(clusterIDs, id)
	}
	f.statusMux.RUnlock()

	var wg sync.WaitGroup
	errChan := make(chan error, len(clusterIDs))

	for _, clusterID := range clusterIDs {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			if err := f.SyncToCluster(ctx, id); err != nil {
				errChan <- fmt.Errorf("cluster %s: %w", id, err)
			}
		}(clusterID)
	}

	wg.Wait()
	close(errChan)

	var errors []error
	for err := range errChan {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return fmt.Errorf("sync failures: %v", errors)
	}

	return nil
}

func (f *RBACFederator) GetSyncStatus(clusterID string) (*SyncStatus, bool) {
	f.statusMux.RLock()
	defer f.statusMux.RUnlock()
	status, ok := f.syncStatus[clusterID]
	return status, ok
}

func (f *RBACFederator) GetAllSyncStatuses() map[string]*SyncStatus {
	f.statusMux.RLock()
	defer f.statusMux.RUnlock()

	copyStatus := make(map[string]*SyncStatus, len(f.syncStatus))
	for k, v := range f.syncStatus {
		copyStatus[k] = v
	}

	return copyStatus
}

func (f *RBACFederator) GetPolicy(policyID string) (*Permission, bool) {
	f.policyMux.RLock()
	defer f.policyMux.RUnlock()
	policy, ok := f.policies[policyID]
	return policy, ok
}

func (f *RBACFederator) ListPolicies() []*Permission {
	f.policyMux.RLock()
	defer f.policyMux.RUnlock()

	policies := make([]*Permission, 0, len(f.policies))
	for _, p := range f.policies {
		policies = append(policies, p)
	}

	return policies
}

func (f *RBACFederator) SyncRegion(ctx context.Context, region string) error {
	f.statusMux.RLock()
	var regionClusters []string
	for id, endpoint := range f.clusters {
		if endpoint.Region == region {
			regionClusters = append(regionClusters, id)
		}
	}
	f.statusMux.RUnlock()

	if len(regionClusters) == 0 {
		return fmt.Errorf("no clusters found in region %s", region)
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(regionClusters))

	for _, clusterID := range regionClusters {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			if err := f.SyncToCluster(ctx, id); err != nil {
				errChan <- err
			}
		}(clusterID)
	}

	wg.Wait()
	close(errChan)

	var errors []error
	for err := range errChan {
		if err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("region sync failures: %v", errors)
	}

	return nil
}

func (f *RBACFederator) RunAutoSyncLoop(ctx context.Context) error {
	if !f.config.EnableAutoSync {
		return nil
	}

	ticker := time.NewTicker(f.config.SyncInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			log.FromContext(ctx).Info("Starting auto-sync cycle")
			if err := f.SyncAllClusters(ctx); err != nil {
				log.FromContext(ctx).Error(err, "Auto-sync failed")
			}
		}
	}
}