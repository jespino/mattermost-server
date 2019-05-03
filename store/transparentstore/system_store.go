package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentSystemStore struct {
	baseStore store.SystemStore
}

func (s TransparentSystemStore) Save(system *model.System) store.StoreChannel {
	return s.baseStore.Save(system)
}

func (s TransparentSystemStore) SaveOrUpdate(system *model.System) store.StoreChannel {
	return s.baseStore.SaveOrUpdate(system)
}

func (s TransparentSystemStore) Update(system *model.System) store.StoreChannel {
	return s.baseStore.Update(system)
}

func (s TransparentSystemStore) Get() store.StoreChannel {
	return s.baseStore.Get()
}

func (s TransparentSystemStore) GetByName(name string) store.StoreChannel {
	return s.baseStore.GetByName(name)
}

func (s TransparentSystemStore) PermanentDeleteByName(name string) store.StoreChannel {
	return s.baseStore.PermanentDeleteByName(name)
}
