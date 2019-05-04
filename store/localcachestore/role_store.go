package localcachestore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type LocalCacheRoleStore struct {
	baseStore store.RoleStore
	rootStore *LocalCacheStore
}

func (s LocalCacheRoleStore) Save(role *model.Role) store.StoreChannel {
	if len(role.Name) != 0 {
		defer s.rootStore.doInvalidateCacheCluster(s.rootStore.roleCache, role.Name)
	}
	return s.baseStore.Save(role)
}

func (s LocalCacheRoleStore) Get(roleId string) store.StoreChannel {
	// Roles are cached by name, as that is most commonly how they are looked up.
	// This means that no caching is supported on roles being looked up by ID.
	return s.baseStore.Get(roleId)
}

func (s LocalCacheRoleStore) GetAll() store.StoreChannel {
	// Roles are cached by name, as that is most commonly how they are looked up.
	// This means that no caching is supported on roles being listed.
	return s.baseStore.GetAll()
}

func (s LocalCacheRoleStore) GetByName(name string) store.StoreChannel {
	return store.Do(func(r *store.StoreResult) {
		if role := s.rootStore.doStandardReadCache(s.rootStore.roleCache, name); role != nil {
			r.Data = role
			return
		}

		resultGet := <-s.baseStore.GetByName(name)
		if resultGet.Err != nil {
			r.Err = resultGet.Err
			return
		}
		s.rootStore.doStandardAddToCache(s.rootStore.roleCache, name, resultGet.Data)
		r.Data = resultGet.Data
	})
}

func (s LocalCacheRoleStore) GetByNames(names []string) store.StoreChannel {
	return store.Do(func(result *store.StoreResult) {
		var foundRoles []*model.Role
		var rolesToQuery []string

		for _, roleName := range names {
			if role := s.rootStore.doStandardReadCache(s.rootStore.roleCache, roleName); role != nil {
				foundRoles = append(foundRoles, role.(*model.Role))
			} else {
				rolesToQuery = append(rolesToQuery, roleName)
			}
		}

		resultGet := <-s.baseStore.GetByNames(rolesToQuery)

		if resultGet.Data != nil {
			rolesFound := resultGet.Data.([]*model.Role)
			for _, role := range rolesFound {
				s.rootStore.doStandardAddToCache(s.rootStore.roleCache, role.Name, role)
			}
			result.Data = append(foundRoles, rolesFound...)
		}
	})
}

func (s LocalCacheRoleStore) Delete(roleId string) store.StoreChannel {
	return store.Do(func(result *store.StoreResult) {
		resultDelete := <-s.baseStore.Delete(roleId)

		if resultDelete.Err == nil {
			role := resultDelete.Data.(*model.Role)
			s.rootStore.doInvalidateCacheCluster(s.rootStore.roleCache, role.Name)
		}
		result.Data = resultDelete.Data
		result.Err = resultDelete.Err
	})
}

func (s LocalCacheRoleStore) PermanentDeleteAll() store.StoreChannel {
	defer s.rootStore.roleCache.Purge()
	defer s.rootStore.doClearCacheCluster(s.rootStore.roleCache)

	return s.baseStore.PermanentDeleteAll()
}
