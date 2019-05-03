package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentPluginStore struct {
	baseStore store.PluginStore
}

func (s TransparentPluginStore) SaveOrUpdate(keyVal *model.PluginKeyValue) store.StoreChannel {
	return s.baseStore.SaveOrUpdate(keyVal)
}

func (s TransparentPluginStore) CompareAndSet(keyVal *model.PluginKeyValue, oldValue []byte) (bool, *model.AppError) {
	return s.baseStore.CompareAndSet(keyVal, oldValue)
}

func (s TransparentPluginStore) Get(pluginId string, key string) store.StoreChannel {
	return s.baseStore.Get(pluginId, key)
}

func (s TransparentPluginStore) Delete(pluginId string, key string) store.StoreChannel {
	return s.baseStore.Delete(pluginId, key)
}

func (s TransparentPluginStore) DeleteAllForPlugin(PluginId string) store.StoreChannel {
	return s.baseStore.DeleteAllForPlugin(PluginId)
}

func (s TransparentPluginStore) DeleteAllExpired() store.StoreChannel {
	return s.baseStore.DeleteAllExpired()
}

func (s TransparentPluginStore) List(pluginId string, page int, perPage int) store.StoreChannel {
	return s.baseStore.List(pluginId, page, perPage)
}
