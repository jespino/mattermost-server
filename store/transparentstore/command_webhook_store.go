package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentCommandWebhookStore struct {
	baseStore store.CommandWebhookStore
}

func (s TransparentCommandWebhookStore) Save(webhook *model.CommandWebhook) store.StoreChannel {
	return s.baseStore.Save(webhook)
}

func (s TransparentCommandWebhookStore) Get(id string) store.StoreChannel {
	return s.baseStore.Get(id)
}

func (s TransparentCommandWebhookStore) TryUse(id string, limit int) store.StoreChannel {
	return s.baseStore.TryUse(id, limit)
}

func (s TransparentCommandWebhookStore) Cleanup() {
	s.baseStore.Cleanup()
}
