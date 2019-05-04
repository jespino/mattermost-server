package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaSchemeStore struct {
	baseStore store.SchemeStore
	root      *KafkaStore
}

func (s KafkaSchemeStore) Save(scheme *model.Scheme) store.StoreChannel {
	return s.baseStore.Save(scheme)
}

func (s KafkaSchemeStore) Get(schemeId string) store.StoreChannel {
	return s.baseStore.Get(schemeId)
}

func (s KafkaSchemeStore) GetByName(schemeName string) store.StoreChannel {
	return s.baseStore.GetByName(schemeName)
}

func (s KafkaSchemeStore) GetAllPage(scope string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetAllPage(scope, offset, limit)
}

func (s KafkaSchemeStore) Delete(schemeId string) store.StoreChannel {
	return s.baseStore.Delete(schemeId)
}

func (s KafkaSchemeStore) PermanentDeleteAll() store.StoreChannel {
	return s.baseStore.PermanentDeleteAll()
}
