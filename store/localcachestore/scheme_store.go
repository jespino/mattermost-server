package localcachestore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type LocalCacheSchemeStore struct {
	store.SchemeStore
	rootStore *LocalCacheStore
}

func (s LocalCacheSchemeStore) Save(scheme *model.Scheme) store.StoreChannel {
	if len(scheme.Id) != 0 {
		defer s.rootStore.doInvalidateCacheCluster(s.rootStore.schemeCache, scheme.Id)
	}
	return s.SchemeStore.Save(scheme)
}

func (s LocalCacheSchemeStore) Get(schemeId string) store.StoreChannel {
	return store.Do(func(r *store.StoreResult) {
		if scheme := s.rootStore.doStandardReadCache(s.rootStore.schemeCache, schemeId); scheme != nil {
			r.Data = scheme
			return
		}

		result := <-s.SchemeStore.Get(schemeId)
		if result.Err != nil {
			r.Err = result.Err
			return
		}

		s.rootStore.doStandardAddToCache(s.rootStore.schemeCache, schemeId, result.Data)

		r.Data = result.Data
	})
}

func (s LocalCacheSchemeStore) Delete(schemeId string) store.StoreChannel {
	defer s.rootStore.doInvalidateCacheCluster(s.rootStore.schemeCache, schemeId)
	defer s.rootStore.doClearCacheCluster(s.rootStore.roleCache)

	return s.SchemeStore.Delete(schemeId)
}

func (s LocalCacheSchemeStore) PermanentDeleteAll() store.StoreChannel {
	defer s.rootStore.doClearCacheCluster(s.rootStore.schemeCache)
	defer s.rootStore.doClearCacheCluster(s.rootStore.roleCache)

	return s.SchemeStore.PermanentDeleteAll()
}
