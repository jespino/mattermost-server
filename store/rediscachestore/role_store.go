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

func (s RedisCacheRoleStore) Save(role *model.Role) (*model.Role, *model.AppError) {
	key := buildRedisKeyForRoleName(role.Name)

	defer func() {
		if err := s.client.Del(key).Err(); err != nil {
			mlog.Error("Redis failed to remove key " + key + " Error: " + err.Error())
		}
	}()

	return s.RoleStore.Save(role)
}

func (s RedisCacheRoleStore) GetByName(name string) (*model.Role, *model.AppError) {
	key := buildRedisKeyForRoleName(name)

	var role *model.Role
	found, err := s.rootStore.load(key, &role)
	if err != nil {
		mlog.Error("Redis encountered an error on read: " + err.Error())
	} else if found {
		return role, nil
	}

	role, appErr := s.RoleStore.GetByName(name)
	if appErr != nil {
		return nil, appErr
	}

	if err := s.rootStore.save(key, role, REDIS_EXPIRY_TIME); err != nil {
		mlog.Error("Redis encountered and error on write: " + err.Error())
	}
	return role, nil
}

func (s RedisCacheRoleStore) GetByNames(names []string) ([]*model.Role, *model.AppError) {
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

	rolesFound, err := s.RoleStore.GetByNames(rolesToQuery)
	if err != nil {
		return nil, err
	}

	for _, role := range rolesFound {
		if err := s.rootStore.save(buildRedisKeyForRoleName(role.Name), role, REDIS_EXPIRY_TIME); err != nil {
			mlog.Error("Redis encountered and error on write: " + err.Error())
		}
	}
	return append(foundRoles, rolesFound...), nil
}

func (s RedisCacheRoleStore) Delete(roleId string) (*model.Role, *model.AppError) {
	role, err := s.RoleStore.Delete(roleId)

	if err == nil {
		defer func() {
			key := buildRedisKeyForRoleName(role.Name)

			if err := s.client.Del(key).Err(); err != nil {
				mlog.Error("Redis failed to remove key " + key + " Error: " + err.Error())
			}
		}()
	}
	return role, err
}

func (s RedisCacheRoleStore) PermanentDeleteAll() *model.AppError {
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
