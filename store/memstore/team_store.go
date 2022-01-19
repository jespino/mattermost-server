// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package memstore

import (
	"context"
	"sync"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/store"
)

type MemTeamStore struct {
	teams   []*model.Team
	members []*model.TeamMember
	mutex   sync.RWMutex
}

func newMemTeamStore() store.TeamStore {
	return &MemTeamStore{}
}
func (s *MemTeamStore) Save(team *model.Team) (*model.Team, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if team.Id != "" {
		return nil, store.NewErrInvalidInput("Team", "id", team.Id)
	}

	team.PreSave()

	if err := team.IsValid(); err != nil {
		return nil, err
	}

	s.teams = append(s.teams, team)

	return team, nil
}

func (s *MemTeamStore) Update(team *model.Team) (*model.Team, error) {
	panic("not implemented")
}

func (s *MemTeamStore) Get(id string) (*model.Team, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for _, t := range s.teams {
		if t.Id == id {
			return t, nil
		}
	}
	return nil, store.NewErrNotFound("Team", id)
}

func (s *MemTeamStore) GetByInviteId(inviteId string) (*model.Team, error) {
	panic("not implemented")
}

func (s *MemTeamStore) GetByName(name string) (*model.Team, error) {
	panic("not implemented")
}

func (s *MemTeamStore) GetByNames(names []string) ([]*model.Team, error) {
	panic("not implemented")
}

func (s *MemTeamStore) SearchAll(opts *model.TeamSearch) ([]*model.Team, error) {
	panic("not implemented")
}

func (s *MemTeamStore) SearchAllPaged(opts *model.TeamSearch) ([]*model.Team, int64, error) {
	panic("not implemented")
}

func (s *MemTeamStore) SearchOpen(opts *model.TeamSearch) ([]*model.Team, error) {
	panic("not implemented")
}

func (s *MemTeamStore) SearchPrivate(opts *model.TeamSearch) ([]*model.Team, error) {
	panic("not implemented")
}

func (s *MemTeamStore) GetAll() ([]*model.Team, error) {
	panic("not implemented")
}

func (s *MemTeamStore) GetAllPage(offset int, limit int, opts *model.TeamSearch) ([]*model.Team, error) {
	panic("not implemented")
}

func (s *MemTeamStore) GetTeamsByUserId(userId string) ([]*model.Team, error) {
	panic("not implemented")
}

func (s *MemTeamStore) GetAllPrivateTeamListing() ([]*model.Team, error) {
	panic("not implemented")
}

func (s *MemTeamStore) GetAllTeamListing() ([]*model.Team, error) {
	panic("not implemented")
}

func (s *MemTeamStore) PermanentDelete(teamId string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	result := []*model.Team{}
	for _, t := range s.teams {
		if t.Id != teamId {
			result = append(result, t)
		}
	}
	s.teams = result

	return nil
}

func (s *MemTeamStore) AnalyticsTeamCount(opts *model.TeamSearch) (int64, error) {
	panic("not implemented")
}

func (s *MemTeamStore) SaveMultipleMembers(members []*model.TeamMember, maxUsersPerTeam int) ([]*model.TeamMember, error) {
	newTeamMembers := map[string]int{}
	users := map[string]bool{}
	for _, member := range members {
		newTeamMembers[member.TeamId] = 0
	}

	for _, member := range members {
		newTeamMembers[member.TeamId]++
		users[member.UserId] = true

		if err := member.IsValid(); err != nil {
			return nil, err
		}
	}

	s.members = append(s.members, members...)

	return members, nil
}

func (s *MemTeamStore) SaveMember(member *model.TeamMember, maxUsersPerTeam int) (*model.TeamMember, error) {
	members, err := s.SaveMultipleMembers([]*model.TeamMember{member}, maxUsersPerTeam)
	if err != nil {
		return nil, err
	}
	return members[0], nil
}

func (s *MemTeamStore) UpdateMultipleMembers(members []*model.TeamMember) ([]*model.TeamMember, error) {
	panic("not implemented")
}

func (s *MemTeamStore) UpdateMember(member *model.TeamMember) (*model.TeamMember, error) {
	panic("not implemented")
}

func (s *MemTeamStore) GetMember(ctx context.Context, teamId string, userId string) (*model.TeamMember, error) {
	for _, m := range s.members {
		if m.TeamId == teamId && m.UserId == userId {
			return m, nil
		}
	}
	return nil, store.NewErrNotFound("Team", teamId+"-"+userId)
}

func (s *MemTeamStore) GetMembers(teamId string, offset int, limit int, teamMembersGetOptions *model.TeamMembersGetOptions) ([]*model.TeamMember, error) {
	panic("not implemented")
}

func (s *MemTeamStore) GetTotalMemberCount(teamId string, restrictions *model.ViewUsersRestrictions) (int64, error) {
	panic("not implemented")
}

func (s *MemTeamStore) GetActiveMemberCount(teamId string, restrictions *model.ViewUsersRestrictions) (int64, error) {
	panic("not implemented")
}

func (s *MemTeamStore) GetMembersByIds(teamId string, userIds []string, restrictions *model.ViewUsersRestrictions) ([]*model.TeamMember, error) {
	panic("not implemented")
}

func (s *MemTeamStore) GetTeamsForUser(ctx context.Context, userId string) ([]*model.TeamMember, error) {
	panic("not implemented")
}

func (s *MemTeamStore) GetTeamsForUserWithPagination(userId string, page, perPage int) ([]*model.TeamMember, error) {
	panic("not implemented")
}

func (s *MemTeamStore) GetChannelUnreadsForAllTeams(excludeTeamId, userId string) ([]*model.ChannelUnread, error) {
	panic("not implemented")
}

func (s *MemTeamStore) GetChannelUnreadsForTeam(teamId, userId string) ([]*model.ChannelUnread, error) {
	panic("not implemented")
}

func (s *MemTeamStore) RemoveMembers(teamId string, userIds []string) error {
	panic("not implemented")
}

func (s *MemTeamStore) RemoveMember(teamId string, userId string) error {
	panic("not implemented")
}

func (s *MemTeamStore) RemoveAllMembersByTeam(teamId string) error {
	panic("not implemented")
}

func (s *MemTeamStore) RemoveAllMembersByUser(userId string) error {
	panic("not implemented")
}

func (s *MemTeamStore) UpdateLastTeamIconUpdate(teamId string, curTime int64) error {
	panic("not implemented")
}

func (s *MemTeamStore) GetTeamsByScheme(schemeId string, offset int, limit int) ([]*model.Team, error) {
	panic("not implemented")
}

func (s *MemTeamStore) MigrateTeamMembers(fromTeamId string, fromUserId string) (map[string]string, error) {
	panic("not implemented")
}

func (s *MemTeamStore) ResetAllTeamSchemes() error {
	panic("not implemented")
}

func (s *MemTeamStore) ClearCaches() {}

func (s *MemTeamStore) InvalidateAllTeamIdsForUser(userId string) {}

func (s *MemTeamStore) ClearAllCustomRoleAssignments() error {
	panic("not implemented")
}

func (s *MemTeamStore) AnalyticsGetTeamCountForScheme(schemeId string) (int64, error) {
	panic("not implemented")
}

func (s *MemTeamStore) GetAllForExportAfter(limit int, afterId string) ([]*model.TeamForExport, error) {
	panic("not implemented")
}

func (s *MemTeamStore) GetUserTeamIds(userId string, allowFromCache bool) ([]string, error) {
	panic("not implemented")
}

func (s *MemTeamStore) GetCommonTeamIDsForTwoUsers(userID, otherUserID string) ([]string, error) {
	panic("not implemented")
}

func (s *MemTeamStore) GetTeamMembersForExport(userId string) ([]*model.TeamMemberForExport, error) {
	panic("not implemented")
}

func (s *MemTeamStore) UserBelongsToTeams(userId string, teamIds []string) (bool, error) {
	panic("not implemented")
}

func (s *MemTeamStore) UpdateMembersRole(teamID string, userIDs []string) error {
	panic("not implemented")
}

func (s *MemTeamStore) GroupSyncedTeamCount() (int64, error) {
	panic("not implemented")
}
