package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentRoleStore struct {
	baseStore store.RoleStore
}

func (s TransparentRoleStore) Save(role *model.Role) store.StoreChannel {
	return s.baseStore.Save(role)
}

func (s TransparentRoleStore) Get(roleId string) store.StoreChannel {
	return s.baseStore.Get(roleId)
}

func (s TransparentRoleStore) GetAll() store.StoreChannel {
	return s.baseStore.GetAll()
}

func (s TransparentRoleStore) GetByName(name string) store.StoreChannel {
	return s.baseStore.GetByName(name)
}

func (s TransparentRoleStore) GetByNames(names []string) store.StoreChannel {
	return s.baseStore.GetByNames(names)
}

func (s TransparentRoleStore) Delete(roldId string) store.StoreChannel {
	return s.baseStore.Delete(roldId)
}

func (s TransparentRoleStore) PermanentDeleteAll() store.StoreChannel {
	return s.baseStore.PermanentDeleteAll()
}
