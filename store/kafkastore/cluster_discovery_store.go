package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaClusterDiscoveryStore struct {
	baseStore store.ClusterDiscoveryStore
	root      *KafkaStore
}

func (s KafkaClusterDiscoveryStore) Save(discovery *model.ClusterDiscovery) *model.AppError {
	return s.baseStore.Save(discovery)
}

func (s KafkaClusterDiscoveryStore) Delete(discovery *model.ClusterDiscovery) (bool, *model.AppError) {
	return s.baseStore.Delete(discovery)
}

func (s KafkaClusterDiscoveryStore) Exists(discovery *model.ClusterDiscovery) (bool, *model.AppError) {
	return s.baseStore.Exists(discovery)
}

func (s KafkaClusterDiscoveryStore) GetAll(discoveryType string, clusterName string) ([]*model.ClusterDiscovery, *model.AppError) {
	return s.baseStore.GetAll(discoveryType, clusterName)
}

func (s KafkaClusterDiscoveryStore) SetLastPingAt(discovery *model.ClusterDiscovery) *model.AppError {
	return s.baseStore.SetLastPingAt(discovery)
}

func (s KafkaClusterDiscoveryStore) Cleanup() *model.AppError {
	return s.baseStore.Cleanup()
}
