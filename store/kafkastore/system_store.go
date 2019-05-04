package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaSystemStore struct {
	baseStore store.SystemStore
	root      *KafkaStore
}

func (s KafkaSystemStore) Save(system *model.System) store.StoreChannel {
	return s.baseStore.Save(system)
}

func (s KafkaSystemStore) SaveOrUpdate(system *model.System) store.StoreChannel {
	return s.baseStore.SaveOrUpdate(system)
}

func (s KafkaSystemStore) Update(system *model.System) store.StoreChannel {
	return s.baseStore.Update(system)
}

func (s KafkaSystemStore) Get() store.StoreChannel {
	return s.baseStore.Get()
}

func (s KafkaSystemStore) GetByName(name string) store.StoreChannel {
	return s.baseStore.GetByName(name)
}

func (s KafkaSystemStore) PermanentDeleteByName(name string) store.StoreChannel {
	return s.baseStore.PermanentDeleteByName(name)
}
