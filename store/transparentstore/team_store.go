package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentTeamStore struct {
	baseStore store.TeamStore
}

func (s TransparentTeamStore) Save(team *model.Team) (*model.Team, *model.AppError) {
	return s.baseStore.Save(team)
}

func (s TransparentTeamStore) Update(team *model.Team) (*model.Team, *model.AppError) {
	return s.baseStore.Update(team)
}

func (s TransparentTeamStore) UpdateDisplayName(name string, teamId string) store.StoreChannel {
	return s.baseStore.UpdateDisplayName(name, teamId)
}

func (s TransparentTeamStore) Get(id string) (*model.Team, *model.AppError) {
	return s.baseStore.Get(id)
}

func (s TransparentTeamStore) GetByName(name string) store.StoreChannel {
	return s.baseStore.GetByName(name)
}

func (s TransparentTeamStore) SearchByName(name string) store.StoreChannel {
	return s.baseStore.SearchByName(name)
}

func (s TransparentTeamStore) SearchAll(term string) store.StoreChannel {
	return s.baseStore.SearchAll(term)
}

func (s TransparentTeamStore) SearchOpen(term string) store.StoreChannel {
	return s.baseStore.SearchOpen(term)
}

func (s TransparentTeamStore) SearchPrivate(term string) store.StoreChannel {
	return s.baseStore.SearchPrivate(term)
}

func (s TransparentTeamStore) GetAll() store.StoreChannel {
	return s.baseStore.GetAll()
}

func (s TransparentTeamStore) GetAllPage(offset int, limit int) store.StoreChannel {
	return s.baseStore.GetAllPage(offset, limit)
}

func (s TransparentTeamStore) GetAllPrivateTeamListing() store.StoreChannel {
	return s.baseStore.GetAllPrivateTeamListing()
}

func (s TransparentTeamStore) GetAllPrivateTeamPageListing(offset int, limit int) store.StoreChannel {
	return s.baseStore.GetAllPrivateTeamPageListing(offset, limit)
}

func (s TransparentTeamStore) GetAllTeamListing() store.StoreChannel {
	return s.baseStore.GetAllTeamListing()
}

func (s TransparentTeamStore) GetAllTeamPageListing(offset int, limit int) store.StoreChannel {
	return s.baseStore.GetAllTeamPageListing(offset, limit)
}

func (s TransparentTeamStore) GetTeamsByUserId(userId string) store.StoreChannel {
	return s.baseStore.GetTeamsByUserId(userId)
}

func (s TransparentTeamStore) GetByInviteId(inviteId string) store.StoreChannel {
	return s.baseStore.GetByInviteId(inviteId)
}

func (s TransparentTeamStore) PermanentDelete(teamId string) store.StoreChannel {
	return s.baseStore.PermanentDelete(teamId)
}

func (s TransparentTeamStore) AnalyticsTeamCount() store.StoreChannel {
	return s.baseStore.AnalyticsTeamCount()
}

func (s TransparentTeamStore) SaveMember(member *model.TeamMember, maxUsersPerTeam int) store.StoreChannel {
	return s.baseStore.SaveMember(member, maxUsersPerTeam)
}

func (s TransparentTeamStore) UpdateMember(member *model.TeamMember) store.StoreChannel {
	return s.baseStore.UpdateMember(member)
}

func (s TransparentTeamStore) GetMember(teamId string, userId string) store.StoreChannel {
	return s.baseStore.GetMember(teamId, userId)
}

func (s TransparentTeamStore) GetMembers(teamId string, offset int, limit int, restrictions *model.ViewUsersRestrictions) store.StoreChannel {
	return s.baseStore.GetMembers(teamId, offset, limit, restrictions)
}

func (s TransparentTeamStore) GetMembersByIds(teamId string, userIds []string, restrictions *model.ViewUsersRestrictions) store.StoreChannel {
	return s.baseStore.GetMembersByIds(teamId, userIds, restrictions)
}

func (s TransparentTeamStore) GetTotalMemberCount(teamId string) store.StoreChannel {
	return s.baseStore.GetTotalMemberCount(teamId)
}

func (s TransparentTeamStore) GetActiveMemberCount(teamId string) store.StoreChannel {
	return s.baseStore.GetActiveMemberCount(teamId)
}

func (s TransparentTeamStore) GetTeamsForUser(userId string) store.StoreChannel {
	return s.baseStore.GetTeamsForUser(userId)
}

func (s TransparentTeamStore) GetTeamsForUserWithPagination(userId string, page int, perPage int) store.StoreChannel {
	return s.baseStore.GetTeamsForUserWithPagination(userId, page, perPage)
}

func (s TransparentTeamStore) GetChannelUnreadsForAllTeams(excludeTeamId string, userId string) store.StoreChannel {
	return s.baseStore.GetChannelUnreadsForAllTeams(excludeTeamId, userId)
}

func (s TransparentTeamStore) GetChannelUnreadsForTeam(teamId string, userId string) store.StoreChannel {
	return s.baseStore.GetChannelUnreadsForTeam(teamId, userId)
}

func (s TransparentTeamStore) RemoveMember(teamId string, userId string) store.StoreChannel {
	return s.baseStore.RemoveMember(teamId, userId)
}

func (s TransparentTeamStore) RemoveAllMembersByTeam(teamId string) store.StoreChannel {
	return s.baseStore.RemoveAllMembersByTeam(teamId)
}

func (s TransparentTeamStore) RemoveAllMembersByUser(userId string) store.StoreChannel {
	return s.baseStore.RemoveAllMembersByUser(userId)
}

func (s TransparentTeamStore) UpdateLastTeamIconUpdate(teamId string, curTime int64) store.StoreChannel {
	return s.baseStore.UpdateLastTeamIconUpdate(teamId, curTime)
}

func (s TransparentTeamStore) GetTeamsByScheme(schemeId string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetTeamsByScheme(schemeId, offset, limit)
}

func (s TransparentTeamStore) MigrateTeamMembers(fromTeamId string, fromUserId string) store.StoreChannel {
	return s.baseStore.MigrateTeamMembers(fromTeamId, fromUserId)
}

func (s TransparentTeamStore) ResetAllTeamSchemes() store.StoreChannel {
	return s.baseStore.ResetAllTeamSchemes()
}

func (s TransparentTeamStore) ClearAllCustomRoleAssignments() store.StoreChannel {
	return s.baseStore.ClearAllCustomRoleAssignments()
}

func (s TransparentTeamStore) AnalyticsGetTeamCountForScheme(schemeId string) store.StoreChannel {
	return s.baseStore.AnalyticsGetTeamCountForScheme(schemeId)
}

func (s TransparentTeamStore) GetAllForExportAfter(limit int, afterId string) store.StoreChannel {
	return s.baseStore.GetAllForExportAfter(limit, afterId)
}

func (s TransparentTeamStore) GetTeamMembersForExport(userId string) store.StoreChannel {
	return s.baseStore.GetTeamMembersForExport(userId)
}

func (s TransparentTeamStore) UserBelongsToTeams(userId string, teamIds []string) store.StoreChannel {
	return s.baseStore.UserBelongsToTeams(userId, teamIds)
}

func (s TransparentTeamStore) GetUserTeamIds(userId string, allowFromCache bool) store.StoreChannel {
	return s.baseStore.GetUserTeamIds(userId, allowFromCache)
}

func (s TransparentTeamStore) InvalidateAllTeamIdsForUser(userId string) {
	s.baseStore.InvalidateAllTeamIdsForUser(userId)
}

func (s TransparentTeamStore) ClearCaches() {
	s.baseStore.ClearCaches()
}
