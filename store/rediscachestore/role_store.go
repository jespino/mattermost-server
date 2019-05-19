package rediscachestore

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/mattermost/mattermost-server/mlog"
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type RedisCacheRoleStore struct {
	store.RoleStore
	client    *redis.Client
	rootStore *RedisCacheStore
}

func (s RedisCacheRoleStore) Save(role *model.Role) store.StoreChannel {
	key := buildRedisKeyForRoleName(role.Name)

	defer func() {
		if err := s.client.Del(key).Err(); err != nil {
			mlog.Error("Redis failed to remove key " + key + " Error: " + err.Error())
		}
	}()

	return s.RoleStore.Save(role)
}

func (s RedisCacheRoleStore) GetByName(name string) store.StoreChannel {
	return store.Do(func(r *store.StoreResult) {
		key := buildRedisKeyForRoleName(name)

		var role *model.Role
		found, err := s.rootStore.load(key, &role)
		if err != nil {
			mlog.Error("Redis encountered an error on read: " + err.Error())
		} else if found {
			r.Data = role
			return
		}

		resultGet := <-s.RoleStore.GetByName(name)
		if resultGet.Err == nil {
			if err := s.rootStore.save(key, resultGet.Data, REDIS_EXPIRY_TIME); err != nil {
				mlog.Error("Redis encountered and error on write: " + err.Error())
			}
		}
		r.Data = resultGet.Data
	})
}

func (s RedisCacheRoleStore) GetByNames(names []string) store.StoreChannel {
	return store.Do(func(result *store.StoreResult) {
		var foundRoles []*model.Role
		var rolesToQuery []string

		for _, roleName := range names {
			var role *model.Role
			found, err := s.rootStore.load(buildRedisKeyForRoleName(roleName), &role)
			if err == nil && found {
				foundRoles = append(foundRoles, role)
			} else {
				rolesToQuery = append(rolesToQuery, roleName)
				if err != nil {
					mlog.Error("Redis encountered an error on read: " + err.Error())
				}
			}
		}

		resultGet := <-s.RoleStore.GetByNames(rolesToQuery)

		if resultGet.Err == nil {
			rolesFound := resultGet.Data.([]*model.Role)
			for _, role := range rolesFound {
				if err := s.rootStore.save(buildRedisKeyForRoleName(role.Name), role, REDIS_EXPIRY_TIME); err != nil {
					mlog.Error("Redis encountered and error on write: " + err.Error())
				}
			}
			resultGet.Data = append(foundRoles, resultGet.Data.([]*model.Role)...)
		}
		result.Data = resultGet.Data
		result.Err = resultGet.Err
	})
}

func (s RedisCacheRoleStore) Delete(roleId string) store.StoreChannel {
	return store.Do(func(result *store.StoreResult) {
		resultDelete := <-s.RoleStore.Delete(roleId)

		if resultDelete.Err == nil {
			defer func() {
				role := resultDelete.Data.(*model.Role)
				key := buildRedisKeyForRoleName(role.Name)

				if err := s.client.Del(key).Err(); err != nil {
					mlog.Error("Redis failed to remove key " + key + " Error: " + err.Error())
				}
			}()
		}
		result.Data = resultDelete.Data
		result.Err = resultDelete.Err
	})
}

func (s RedisCacheRoleStore) PermanentDeleteAll() store.StoreChannel {
	defer func() {
		if keys, err := s.client.Keys("roles:*").Result(); err != nil {
			mlog.Error("Redis encountered an error on read: " + err.Error())
		} else {
			if err := s.client.Del(keys...).Err(); err != nil {
				mlog.Error("Redis encountered an error on delete: " + err.Error())
			}
		}
	}()

	return s.RoleStore.PermanentDeleteAll()
}

func buildRedisKeyForRoleName(roleName string) string {
	return fmt.Sprintf("roles:%s", roleName)
}
