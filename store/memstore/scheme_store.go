// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package memstore

import (
	"fmt"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/store"
	"github.com/pkg/errors"
)

type MemSchemeStore struct {
	MemStore *MemStore
	schemes  []*model.Scheme
}

func newMemSchemeStore(memStore *MemStore) store.SchemeStore {
	return &MemSchemeStore{MemStore: memStore}
}

func (s *MemSchemeStore) Save(scheme *model.Scheme) (*model.Scheme, error) {
	if scheme.Id == "" {
		for _, r := range s.schemes {
			if r.Name == scheme.Name {
				return nil, errors.New("duplicated name")
			}
		}

		scheme.Id = model.NewId()
		scheme.CreateAt = model.GetMillis()
		scheme.UpdateAt = scheme.CreateAt
		s.schemes = append(s.schemes, scheme)
		return scheme, nil
	}

	if !scheme.IsValid() {
		return nil, store.NewErrInvalidInput("Role", "<any>", fmt.Sprintf("%v", scheme))
	}

	sch, err := s.Get(scheme.Id)
	if err != nil {
		return nil, err
	}

	scheme.UpdateAt = model.GetMillis()
	*sch = *scheme

	return sch, nil
}

func (s *MemSchemeStore) Get(schemeId string) (*model.Scheme, error) {
	for _, r := range s.schemes {
		if r.Id == schemeId {
			return r, nil
		}
	}
	return nil, store.NewErrNotFound("Scheme", schemeId)
}

func (s *MemSchemeStore) GetByName(schemeName string) (*model.Scheme, error) {
	for _, r := range s.schemes {
		if r.Name == schemeName {
			return r, nil
		}
	}
	return nil, store.NewErrNotFound("Scheme", schemeName)
}

func (s *MemSchemeStore) Delete(schemeId string) (*model.Scheme, error) {
	sch, _ := s.Get(schemeId)
	if sch != nil && sch.DeleteAt == 0 {
		now := model.GetMillis()
		sch.DeleteAt = now
		sch.UpdateAt = now
	}
	return nil, store.NewErrNotFound("Scheme", schemeId)
}

func (s *MemSchemeStore) GetAllPage(scope string, offset int, limit int) ([]*model.Scheme, error) {
	schemes := []*model.Scheme{}
	counter := 0
	for _, sch := range s.schemes {
		if sch.DeleteAt == 0 && (scope == "" || sch.Scope == scope) {
			counter++
			if counter > offset {
				schemes = append(schemes, sch)
			}

			if counter >= offset+limit {
				return schemes, nil
			}
		}
	}
	return schemes, nil
}

func (s *MemSchemeStore) PermanentDeleteAll() error {
	s.schemes = []*model.Scheme{}
	return nil
}

func (s *MemSchemeStore) CountByScope(scope string) (int64, error) {
	var counter int64 = 0
	for _, sch := range s.schemes {
		if sch.DeleteAt == 0 && sch.Scope == scope {
			counter++
		}
	}
	return counter, nil
}

func (s *MemSchemeStore) CountWithoutPermission(schemeScope, permissionID string, roleScope model.RoleScope, roleType model.RoleType) (int64, error) {
	panic("not implemented")
}

func (s *MemSchemeStore) createScheme(scheme *model.Scheme) (*model.Scheme, error) {
	// Fetch the default system scheme roles to populate default permissions.
	defaultRoleNames := []string{
		model.TeamAdminRoleId,
		model.TeamUserRoleId,
		model.TeamGuestRoleId,
		model.ChannelAdminRoleId,
		model.ChannelUserRoleId,
		model.ChannelGuestRoleId,
		model.PlaybookAdminRoleId,
		model.PlaybookMemberRoleId,
		model.RunAdminRoleId,
		model.RunMemberRoleId,
	}
	defaultRoles := make(map[string]*model.Role)
	roles, err := s.MemStore.Role().GetByNames(defaultRoleNames)
	if err != nil {
		return nil, err
	}

	for _, role := range roles {
		defaultRoles[role.Name] = role
	}

	if len(defaultRoles) != len(defaultRoleNames) {
		return nil, errors.New("createScheme: unable to retrieve default scheme roles")
	}

	// Create the appropriate default roles for the scheme.
	if scheme.Scope == model.SchemeScopeTeam {
		// Team Admin Role
		teamAdminRole := &model.Role{
			Name:          model.NewId(),
			DisplayName:   fmt.Sprintf("Team Admin Role for Scheme %s", scheme.Name),
			Permissions:   defaultRoles[model.TeamAdminRoleId].Permissions,
			SchemeManaged: true,
		}

		savedRole, err := s.MemStore.Role().(*MemRoleStore).createRole(teamAdminRole)
		if err != nil {
			return nil, err
		}
		scheme.DefaultTeamAdminRole = savedRole.Name

		// Team User Role
		teamUserRole := &model.Role{
			Name:          model.NewId(),
			DisplayName:   fmt.Sprintf("Team User Role for Scheme %s", scheme.Name),
			Permissions:   defaultRoles[model.TeamUserRoleId].Permissions,
			SchemeManaged: true,
		}

		savedRole, err = s.MemStore.Role().(*MemRoleStore).createRole(teamUserRole)
		if err != nil {
			return nil, err
		}
		scheme.DefaultTeamUserRole = savedRole.Name

		// Team Guest Role
		teamGuestRole := &model.Role{
			Name:          model.NewId(),
			DisplayName:   fmt.Sprintf("Team Guest Role for Scheme %s", scheme.Name),
			Permissions:   defaultRoles[model.TeamGuestRoleId].Permissions,
			SchemeManaged: true,
		}

		savedRole, err = s.MemStore.Role().(*MemRoleStore).createRole(teamGuestRole)
		if err != nil {
			return nil, err
		}
		scheme.DefaultTeamGuestRole = savedRole.Name

		// playbook admin role
		playbookAdminRole := &model.Role{
			Name:          model.NewId(),
			DisplayName:   fmt.Sprintf("Playbook Admin Role for Scheme %s", scheme.Name),
			Permissions:   defaultRoles[model.PlaybookAdminRoleId].Permissions,
			SchemeManaged: true,
		}
		savedRole, err = s.MemStore.Role().(*MemRoleStore).createRole(playbookAdminRole)
		if err != nil {
			return nil, err
		}
		scheme.DefaultPlaybookAdminRole = savedRole.Name

		// playbook member role
		playbookMemberRole := &model.Role{
			Name:          model.NewId(),
			DisplayName:   fmt.Sprintf("Playbook Member Role for Scheme %s", scheme.Name),
			Permissions:   defaultRoles[model.PlaybookMemberRoleId].Permissions,
			SchemeManaged: true,
		}
		savedRole, err = s.MemStore.Role().(*MemRoleStore).createRole(playbookMemberRole)
		if err != nil {
			return nil, err
		}
		scheme.DefaultPlaybookMemberRole = savedRole.Name

		// run admin role
		runAdminRole := &model.Role{
			Name:          model.NewId(),
			DisplayName:   fmt.Sprintf("Run Admin Role for Scheme %s", scheme.Name),
			Permissions:   defaultRoles[model.RunAdminRoleId].Permissions,
			SchemeManaged: true,
		}
		savedRole, err = s.MemStore.Role().(*MemRoleStore).createRole(runAdminRole)
		if err != nil {
			return nil, err
		}
		scheme.DefaultRunAdminRole = savedRole.Name

		// run member role
		runMemberRole := &model.Role{
			Name:          model.NewId(),
			DisplayName:   fmt.Sprintf("Run Member Role for Scheme %s", scheme.Name),
			Permissions:   defaultRoles[model.RunMemberRoleId].Permissions,
			SchemeManaged: true,
		}
		savedRole, err = s.MemStore.Role().(*MemRoleStore).createRole(runMemberRole)
		if err != nil {
			return nil, err
		}
		scheme.DefaultRunMemberRole = savedRole.Name
	}

	if scheme.Scope == model.SchemeScopeTeam || scheme.Scope == model.SchemeScopeChannel {
		// Channel Admin Role
		channelAdminRole := &model.Role{
			Name:          model.NewId(),
			DisplayName:   fmt.Sprintf("Channel Admin Role for Scheme %s", scheme.Name),
			Permissions:   defaultRoles[model.ChannelAdminRoleId].Permissions,
			SchemeManaged: true,
		}

		if scheme.Scope == model.SchemeScopeChannel {
			channelAdminRole.Permissions = []string{}
		}

		savedRole, err := s.MemStore.Role().(*MemRoleStore).createRole(channelAdminRole)
		if err != nil {
			return nil, err
		}
		scheme.DefaultChannelAdminRole = savedRole.Name

		// Channel User Role
		channelUserRole := &model.Role{
			Name:          model.NewId(),
			DisplayName:   fmt.Sprintf("Channel User Role for Scheme %s", scheme.Name),
			Permissions:   defaultRoles[model.ChannelUserRoleId].Permissions,
			SchemeManaged: true,
		}

		if scheme.Scope == model.SchemeScopeChannel {
			channelUserRole.Permissions = filterModerated(channelUserRole.Permissions)
		}

		savedRole, err = s.MemStore.Role().(*MemRoleStore).createRole(channelUserRole)
		if err != nil {
			return nil, err
		}
		scheme.DefaultChannelUserRole = savedRole.Name

		// Channel Guest Role
		channelGuestRole := &model.Role{
			Name:          model.NewId(),
			DisplayName:   fmt.Sprintf("Channel Guest Role for Scheme %s", scheme.Name),
			Permissions:   defaultRoles[model.ChannelGuestRoleId].Permissions,
			SchemeManaged: true,
		}

		if scheme.Scope == model.SchemeScopeChannel {
			channelGuestRole.Permissions = filterModerated(channelGuestRole.Permissions)
		}

		savedRole, err = s.MemStore.Role().(*MemRoleStore).createRole(channelGuestRole)
		if err != nil {
			return nil, err
		}
		scheme.DefaultChannelGuestRole = savedRole.Name
	}

	scheme.Id = model.NewId()
	if scheme.Name == "" {
		scheme.Name = model.NewId()
	}
	scheme.CreateAt = model.GetMillis()
	scheme.UpdateAt = scheme.CreateAt

	// Validate the scheme
	if !scheme.IsValidForCreate() {
		return nil, store.NewErrInvalidInput("Scheme", "<any>", fmt.Sprintf("%v", scheme))
	}

	s.schemes = append(s.schemes, scheme)

	return scheme, nil
}

func filterModerated(permissions []string) []string {
	filteredPermissions := []string{}
	for _, perm := range permissions {
		if _, ok := model.ChannelModeratedPermissionsMap[perm]; ok {
			filteredPermissions = append(filteredPermissions, perm)
		}
	}
	return filteredPermissions
}
