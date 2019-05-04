package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaPluginStore struct {
	baseStore store.PluginStore
	root      *KafkaStore
}

func (s KafkaPluginStore) SaveOrUpdate(keyVal *model.PluginKeyValue) store.StoreChannel {
	return s.baseStore.SaveOrUpdate(keyVal)
}

func (s KafkaPluginStore) CompareAndSet(keyVal *model.PluginKeyValue, oldValue []byte) (bool, *model.AppError) {
	return s.baseStore.CompareAndSet(keyVal, oldValue)
}

func (s KafkaPluginStore) Get(pluginId string, key string) store.StoreChannel {
	return s.baseStore.Get(pluginId, key)
}

func (s KafkaPluginStore) Delete(pluginId string, key string) store.StoreChannel {
	return s.baseStore.Delete(pluginId, key)
}

func (s KafkaPluginStore) DeleteAllForPlugin(PluginId string) store.StoreChannel {
	return s.baseStore.DeleteAllForPlugin(PluginId)
}

func (s KafkaPluginStore) DeleteAllExpired() store.StoreChannel {
	return s.baseStore.DeleteAllExpired()
}

func (s KafkaPluginStore) List(pluginId string, page int, perPage int) store.StoreChannel {
	return s.baseStore.List(pluginId, page, perPage)
}
