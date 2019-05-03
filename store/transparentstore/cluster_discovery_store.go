package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentClusterDiscoveryStore struct {
	baseStore store.ClusterDiscoveryStore
}

func (s TransparentClusterDiscoveryStore) Save(discovery *model.ClusterDiscovery) *model.AppError {
	return s.baseStore.Save(discovery)
}

func (s TransparentClusterDiscoveryStore) Delete(discovery *model.ClusterDiscovery) (bool, *model.AppError) {
	return s.baseStore.Delete(discovery)
}

func (s TransparentClusterDiscoveryStore) Exists(discovery *model.ClusterDiscovery) (bool, *model.AppError) {
	return s.baseStore.Exists(discovery)
}

func (s TransparentClusterDiscoveryStore) GetAll(discoveryType string, clusterName string) ([]*model.ClusterDiscovery, *model.AppError) {
	return s.baseStore.GetAll(discoveryType, clusterName)
}

func (s TransparentClusterDiscoveryStore) SetLastPingAt(discovery *model.ClusterDiscovery) *model.AppError {
	return s.baseStore.SetLastPingAt(discovery)
}

func (s TransparentClusterDiscoveryStore) Cleanup() *model.AppError {
	return s.baseStore.Cleanup()
}
