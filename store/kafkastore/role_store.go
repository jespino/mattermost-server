package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaRoleStore struct {
	baseStore store.RoleStore
	root      *KafkaStore
}

func (s KafkaRoleStore) Save(role *model.Role) store.StoreChannel {
	return s.baseStore.Save(role)
}

func (s KafkaRoleStore) Get(roleId string) store.StoreChannel {
	return s.baseStore.Get(roleId)
}

func (s KafkaRoleStore) GetAll() store.StoreChannel {
	return s.baseStore.GetAll()
}

func (s KafkaRoleStore) GetByName(name string) store.StoreChannel {
	return s.baseStore.GetByName(name)
}

func (s KafkaRoleStore) GetByNames(names []string) store.StoreChannel {
	return s.baseStore.GetByNames(names)
}

func (s KafkaRoleStore) Delete(roldId string) store.StoreChannel {
	return s.baseStore.Delete(roldId)
}

func (s KafkaRoleStore) PermanentDeleteAll() store.StoreChannel {
	return s.baseStore.PermanentDeleteAll()
}
