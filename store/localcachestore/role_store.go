package localcachestore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type LocalCacheRoleStore struct {
	store.RoleStore
	rootStore *LocalCacheStore
}

func (s LocalCacheRoleStore) Save(role *model.Role) store.StoreChannel {
	if len(role.Name) != 0 {
		defer s.rootStore.doInvalidateCacheCluster(s.rootStore.roleCache, role.Name)
	}
	return s.RoleStore.Save(role)
}

func (s LocalCacheRoleStore) GetByName(name string) store.StoreChannel {
	return store.Do(func(r *store.StoreResult) {
		if role := s.rootStore.doStandardReadCache(s.rootStore.roleCache, name); role != nil {
			r.Data = role
			return
		}

		resultGet := <-s.RoleStore.GetByName(name)
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

		resultGet := <-s.RoleStore.GetByNames(rolesToQuery)

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
		resultDelete := <-s.RoleStore.Delete(roleId)

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

	return s.RoleStore.PermanentDeleteAll()
}
