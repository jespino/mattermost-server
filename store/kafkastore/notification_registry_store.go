package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaNotificationRegistryStore struct {
	baseStore store.NotificationRegistryStore
	root      *KafkaStore
}

func (s KafkaNotificationRegistryStore) Save(notification *model.NotificationRegistry) (*model.NotificationRegistry, *model.AppError) {
	return s.baseStore.Save(notification)
}

func (s KafkaNotificationRegistryStore) MarkAsReceived(ackId string, time int64) *model.AppError {
	return s.baseStore.MarkAsReceived(ackId, time)
}

func (s KafkaNotificationRegistryStore) UpdateSendStatus(ackId string, status string) *model.AppError {
	return s.baseStore.UpdateSendStatus(ackId, status)
}
