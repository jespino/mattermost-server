package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentSchemeStore struct {
	baseStore store.SchemeStore
}

func (s TransparentSchemeStore) Save(scheme *model.Scheme) store.StoreChannel {
	return s.baseStore.Save(scheme)
}

func (s TransparentSchemeStore) Get(schemeId string) store.StoreChannel {
	return s.baseStore.Get(schemeId)
}

func (s TransparentSchemeStore) GetByName(schemeName string) store.StoreChannel {
	return s.baseStore.GetByName(schemeName)
}

func (s TransparentSchemeStore) GetAllPage(scope string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetAllPage(scope, offset, limit)
}

func (s TransparentSchemeStore) Delete(schemeId string) store.StoreChannel {
	return s.baseStore.Delete(schemeId)
}

func (s TransparentSchemeStore) PermanentDeleteAll() store.StoreChannel {
	return s.baseStore.PermanentDeleteAll()
}
