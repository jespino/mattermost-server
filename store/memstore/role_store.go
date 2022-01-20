// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package memstore

import (
	"context"
	"fmt"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/store"
	"github.com/pkg/errors"
)

type MemRoleStore struct {
	roles []*model.Role
}

func newMemRoleStore() store.RoleStore {
	return &MemRoleStore{}
}

func (s *MemRoleStore) Save(role *model.Role) (*model.Role, error) {
	if !role.IsValidWithoutId() {
		return nil, store.NewErrInvalidInput("Role", "<any>", fmt.Sprintf("%v", role))
	}

	if role.Id == "" {
		return s.createRole(role)
	}

	r, err := s.Get(role.Id)
	if err != nil {
		return nil, err
	}

	role.UpdateAt = model.GetMillis()
	*r = *role

	return r, nil
}

func (s *MemRoleStore) Get(roleId string) (*model.Role, error) {
	for _, r := range s.roles {
		if r.Id == roleId {
			return r, nil
		}
	}
	return nil, store.NewErrNotFound("Role", roleId)
}

func (s *MemRoleStore) GetAll() ([]*model.Role, error) {
	result := []*model.Role{}
	for _, r := range s.roles {
		if r.DeleteAt == 0 {
			result = append(result, r)
		}
	}
	return result, nil
}

func (s *MemRoleStore) GetByName(ctx context.Context, name string) (*model.Role, error) {
	for _, r := range s.roles {
		if r.Name == name {
			return r, nil
		}
	}
	return nil, store.NewErrNotFound("Role", name)
}

func (s *MemRoleStore) GetByNames(names []string) ([]*model.Role, error) {
	result := []*model.Role{}
	for _, r := range s.roles {
		for _, name := range names {
			if r.DeleteAt == 0 && r.Name == name {
				result = append(result, r)
			}
		}
	}
	return result, nil
}

func (s *MemRoleStore) Delete(roleId string) (*model.Role, error) {
	r, _ := s.Get(roleId)
	if r != nil && r.DeleteAt == 0 {
		now := model.GetMillis()
		r.DeleteAt = now
		r.UpdateAt = now
	}
	return nil, store.NewErrNotFound("Role", roleId)
}

func (s *MemRoleStore) PermanentDeleteAll() error {
	s.roles = []*model.Role{}
	return nil
}

func (s *MemRoleStore) ChannelHigherScopedPermissions(roleNames []string) (map[string]*model.RolePermissions, error) {
	// TODO: not implemented
	roleNameHigherScopedPermissions := map[string]*model.RolePermissions{}
	return roleNameHigherScopedPermissions, nil
}

func (s *MemRoleStore) AllChannelSchemeRoles() ([]*model.Role, error) {
	// TODO: Implement this
	return []*model.Role{}, nil
}

func (s *MemRoleStore) ChannelRolesUnderTeamRole(roleName string) ([]*model.Role, error) {
	// TODO: Implement this
	return []*model.Role{}, nil
}

func (s *MemRoleStore) createRole(role *model.Role) (*model.Role, error) {
	for _, r := range s.roles {
		if r.Name == role.Name {
			return nil, errors.New("duplicated name")
		}
	}

	role.Id = model.NewId()
	role.CreateAt = model.GetMillis()
	role.UpdateAt = role.CreateAt
	s.roles = append(s.roles, role)
	return role, nil
}
