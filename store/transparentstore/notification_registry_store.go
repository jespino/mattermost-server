package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentNotificationRegistryStore struct {
	baseStore store.NotificationRegistryStore
}

func (s TransparentNotificationRegistryStore) Save(notification *model.NotificationRegistry) (*model.NotificationRegistry, *model.AppError) {
	return s.baseStore.Save(notification)
}

func (s TransparentNotificationRegistryStore) MarkAsReceived(ackId string, time int64) *model.AppError {
	return s.baseStore.MarkAsReceived(ackId, time)
}

func (s TransparentNotificationRegistryStore) UpdateSendStatus(ackId string, status string) *model.AppError {
	return s.baseStore.UpdateSendStatus(ackId, status)
}
