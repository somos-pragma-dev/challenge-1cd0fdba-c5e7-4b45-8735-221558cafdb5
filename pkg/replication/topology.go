package replication

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.etcd.io/etcd/client/v3"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	controlplanev1alpha1 "github.com/example/controlplane-operator/api/v1alpha1"
)

type TopologyConfig struct {
	ClusterID       string
	Region          string
	EtcdEndpoints   []string
	ConsulEndpoints []string
	QuorumSize      int
	Timeout         time.Duration
}

type ClusterNode struct {
	ID        string
	Region    string
	Endpoint  string
	IsLeader  bool
	LastSeen  time.Time
	Health    bool
}

type ReplicationTopology struct {
	config        TopologyConfig
	nodes         map[string]*ClusterNode
	etcdClient    *clientv3.Client
	mu            sync.RWMutex
	watcherCancel context.CancelFunc
}

func NewReplicationTopology(cfg TopologyConfig) (*ReplicationTopology, error) {
	if cfg.Timeout == 0 {
		cfg.Timeout = 10 * time.Second
	}
	if cfg.QuorumSize == 0 {
		cfg.QuorumSize = 3
	}

	rt := &ReplicationTopology{
		config: cfg,
		nodes:  make(map[string]*ClusterNode),
	}

	if len(cfg.EtcdEndpoints) > 0 {
		cli, err := clientv3.New(clientv3.Config{
			Endpoints:   cfg.EtcdEndpoints,
			DialTimeout: cfg.Timeout,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create etcd client: %w", err)
		}
		rt.etcdClient = cli
	}

	return rt, nil
}

func (rt *ReplicationTopology) Start(ctx context.Context) error {
	logger := log.FromContext(ctx)

	if rt.etcdClient == nil {
		return fmt.Errorf("etcd client not initialized")
	}

	rt.watchClusterState(ctx)

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if err := rt.syncClusterState(ctx); err != nil {
				logger.Error(err, "failed to sync cluster state")
			}
		}
	}
}

func (rt *ReplicationTopology) watchClusterState(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	rt.watcherCancel = cancel

	rchan := rt.etcdClient.Watch(ctx, "/clusters/", clientv3.WithPrefix())

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case wresp := <-rchan:
				if wresp.Err() != nil {
					log.Log.Error(wresp.Err(), "watch error on cluster state")
					continue
				}
				for _, ev := range wresp.Events {
					rt.handleClusterEvent(ev)
				}
			}
		}
	}()
}

func (rt *ReplicationTopology) handleClusterEvent(ev *clientv3.Event) {
	rt.mu.Lock()
	defer rt.mu.Unlock()

	switch ev.Type {
	case clientv3.EventTypePut:
		nodeID := string(ev.Kv.Key)
		rt.nodes[nodeID] = &ClusterNode{
			ID:       nodeID,
			LastSeen: time.Now(),
			Health:   true,
		}
	case clientv3.EventTypeDelete:
		nodeID := string(ev.Kv.Key)
		delete(rt.nodes, nodeID)
	}
}

func (rt *ReplicationTopology) syncClusterState(ctx context.Context) error {
	logger := log.FromContext(ctx)

	resp, err := rt.etcdClient.Get(ctx, "/clusters/", clientv3.WithPrefix())
	if err != nil {
		return fmt.Errorf("failed to get cluster state: %w", err)
	}

	rt.mu.Lock()
	now := time.Now()
	for _, kv := range resp.Kvs {
		nodeID := string(kv.Key)
		if _, ok := rt.nodes[nodeID]; !ok {
			rt.nodes[nodeID] = &ClusterNode{ID: nodeID}
		}
		rt.nodes[nodeID].LastSeen = now
		rt.nodes[nodeID].Health = true
	}
	rt.mu.Unlock()

	logger.Info("cluster state synced", "nodeCount", len(rt.nodes))
	return nil
}

func (rt *ReplicationTopology) GetHealthyNodes(ctx context.Context) []*ClusterNode {
	rt.mu.RLock()
	defer rt.mu.RUnlock()

	var healthy []*ClusterNode
	for _, node := range rt.nodes {
		if node.Health && time.Since(node.LastSeen) < 2*time.Minute {
			healthy = append(healthy, node)
		}
	}
	return healthy
}

func (rt *ReplicationTopology) CheckQuorum(ctx context.Context) (bool, error) {
	healthy := rt.GetHealthyNodes(ctx)
	return len(healthy) >= rt.config.QuorumSize, nil
}

func (rt *ReplicationTopology) ElectLeader(ctx context.Context) (string, error) {
	healthy := rt.GetHealthyNodes(ctx)
	if len(healthy) == 0 {
		return "", fmt.Errorf("no healthy nodes available for leader election")
	}

	sort.Slice(healthy, func(i, j int) bool {
		return healthy[i].LastSeen.Before(healthy[j].LastSeen)
	})

	leader := healthy[0]
	rt.mu.Lock()
	rt.nodes[leader.ID].IsLeader = true
	rt.mu.Unlock()

	return leader.ID, nil
}

func (rt *ReplicationTopology) ReplicateState(ctx context.Context, key string, value []byte) error {
	hasQuorum, err := rt.CheckQuorum(ctx)
	if err != nil {
		return err
	}
	if !hasQuorum {
		return fmt.Errorf("cannot replicate: no quorum (need %d, have %d)", 
			rt.config.QuorumSize, len(rt.GetHealthyNodes(ctx)))
	}

	ctx, cancel := context.WithTimeout(ctx, rt.config.Timeout)
	defer cancel()

	_, err = rt.etcdClient.Put(ctx, key, string(value))
	return err
}

func (rt *ReplicationTopology) GetReplicatedState(ctx context.Context, key string) ([]byte, error) {
	resp, err := rt.etcdClient.Get(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("failed to get replicated state: %w", err)
	}
	if len(resp.Kvs) == 0 {
		return nil, nil
	}
	return resp.Kvs[0].Value, nil
}

func (rt *ReplicationTopology) Stop() {
	if rt.watcherCancel != nil {
		rt.watcherCancel()
	}
	if rt.etcdClient != nil {
		rt.etcdClient.Close()
	}
}

// Replicator define el contrato para implementar estrategias de replicación
type Replicator interface {
	Replicate(ctx context.Context, state controlplanev1alpha1.ControlPlane) error
	GetReplicatedState(ctx context.Context, key string) ([]byte, error)
	CheckQuorum(ctx context.Context) (bool, error)
}

// EtcdReplicator implementa la replicación usando etcd como store distribuido
type EtcdReplicator struct {
	topology *ReplicationTopology
}

func NewEtcdReplicator(topology *ReplicationTopology) *EtcdReplicator {
	return &EtcdReplicator{topology: topology}
}

func (r *EtcdReplicator) Replicate(ctx context.Context, cp controlplanev1alpha1.ControlPlane) error {
	key := fmt.Sprintf("/clusters/%s/controlplane", cp.Name)
	value := []byte(cp.Status.State)
	return r.topology.ReplicateState(ctx, key, value)
}

func (r *EtcdReplicator) GetReplicatedState(ctx context.Context, key string) ([]byte, error) {
	return r.topology.GetReplicatedState(ctx, key)
}

func (r *EtcdReplicator) CheckQuorum(ctx context.Context) (bool, error) {
	return r.topology.CheckQuorum(ctx)
}

// GetClusterInfo obtiene información del cluster desde el CRD
func GetClusterInfo(ctx context.Context, c client.Client, nn types.NamespacedName) (*ClusterInfo, error) {
	cp := &controlplanev1alpha1.ControlPlane{}
	err := c.Get(ctx, nn, cp)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	return &ClusterInfo{
		Name:      cp.Name,
		Namespace: cp.Namespace,
		Region:    cp.Spec.Region,
		State:     string(cp.Status.State),
	}, nil
}

// ClusterInfo representa la información básica de un cluster
type ClusterInfo struct {
	Name      string
	Namespace string
	Region    string
	State     string
}

// SyncClusterResources sincroniza recursos entre clusters usando la topología
func SyncClusterResources(ctx context.Context, topology *ReplicationTopology, resources []client.Object) error {
	hasQuorum, err := topology.CheckQuorum(ctx)
	if err != nil {
		return err
	}
	if !hasQuorum {
		return fmt.Errorf("cannot sync resources: quorum not reached")
	}

	for _, res := range resources {
		key := fmt.Sprintf("/clusters/%s/%s/%s", res.GetName(), res.GetObjectKind().GroupVersionKind().Kind, res.GetNamespace())
		value, err := client.MarshallJSON(res)
		if err != nil {
			return fmt.Errorf("failed to marshal resource: %w", err)
		}
		if err := topology.ReplicateState(ctx, key, value); err != nil {
			return fmt.Errorf("failed to replicate resource %s: %w", res.GetName(), err)
		}
	}
	return nil
}