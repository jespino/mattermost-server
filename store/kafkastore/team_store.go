package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaTeamStore struct {
	baseStore store.TeamStore
	root      *KafkaStore
}

func (s KafkaTeamStore) Save(team *model.Team) (*model.Team, *model.AppError) {
	return s.baseStore.Save(team)
}

func (s KafkaTeamStore) Update(team *model.Team) (*model.Team, *model.AppError) {
	return s.baseStore.Update(team)
}

func (s KafkaTeamStore) UpdateDisplayName(name string, teamId string) store.StoreChannel {
	return s.baseStore.UpdateDisplayName(name, teamId)
}

func (s KafkaTeamStore) Get(id string) (*model.Team, *model.AppError) {
	return s.baseStore.Get(id)
}

func (s KafkaTeamStore) GetByName(name string) store.StoreChannel {
	return s.baseStore.GetByName(name)
}

func (s KafkaTeamStore) SearchByName(name string) store.StoreChannel {
	return s.baseStore.SearchByName(name)
}

func (s KafkaTeamStore) SearchAll(term string) store.StoreChannel {
	return s.baseStore.SearchAll(term)
}

func (s KafkaTeamStore) SearchOpen(term string) store.StoreChannel {
	return s.baseStore.SearchOpen(term)
}

func (s KafkaTeamStore) SearchPrivate(term string) store.StoreChannel {
	return s.baseStore.SearchPrivate(term)
}

func (s KafkaTeamStore) GetAll() store.StoreChannel {
	return s.baseStore.GetAll()
}

func (s KafkaTeamStore) GetAllPage(offset int, limit int) store.StoreChannel {
	return s.baseStore.GetAllPage(offset, limit)
}

func (s KafkaTeamStore) GetAllPrivateTeamListing() store.StoreChannel {
	return s.baseStore.GetAllPrivateTeamListing()
}

func (s KafkaTeamStore) GetAllPrivateTeamPageListing(offset int, limit int) store.StoreChannel {
	return s.baseStore.GetAllPrivateTeamPageListing(offset, limit)
}

func (s KafkaTeamStore) GetAllTeamListing() store.StoreChannel {
	return s.baseStore.GetAllTeamListing()
}

func (s KafkaTeamStore) GetAllTeamPageListing(offset int, limit int) store.StoreChannel {
	return s.baseStore.GetAllTeamPageListing(offset, limit)
}

func (s KafkaTeamStore) GetTeamsByUserId(userId string) store.StoreChannel {
	return s.baseStore.GetTeamsByUserId(userId)
}

func (s KafkaTeamStore) GetByInviteId(inviteId string) store.StoreChannel {
	return s.baseStore.GetByInviteId(inviteId)
}

func (s KafkaTeamStore) PermanentDelete(teamId string) store.StoreChannel {
	return s.baseStore.PermanentDelete(teamId)
}

func (s KafkaTeamStore) AnalyticsTeamCount() store.StoreChannel {
	return s.baseStore.AnalyticsTeamCount()
}

func (s KafkaTeamStore) SaveMember(member *model.TeamMember, maxUsersPerTeam int) store.StoreChannel {
	return s.baseStore.SaveMember(member, maxUsersPerTeam)
}

func (s KafkaTeamStore) UpdateMember(member *model.TeamMember) store.StoreChannel {
	return s.baseStore.UpdateMember(member)
}

func (s KafkaTeamStore) GetMember(teamId string, userId string) store.StoreChannel {
	return s.baseStore.GetMember(teamId, userId)
}

func (s KafkaTeamStore) GetMembers(teamId string, offset int, limit int, restrictions *model.ViewUsersRestrictions) store.StoreChannel {
	return s.baseStore.GetMembers(teamId, offset, limit, restrictions)
}

func (s KafkaTeamStore) GetMembersByIds(teamId string, userIds []string, restrictions *model.ViewUsersRestrictions) store.StoreChannel {
	return s.baseStore.GetMembersByIds(teamId, userIds, restrictions)
}

func (s KafkaTeamStore) GetTotalMemberCount(teamId string) store.StoreChannel {
	return s.baseStore.GetTotalMemberCount(teamId)
}

func (s KafkaTeamStore) GetActiveMemberCount(teamId string) store.StoreChannel {
	return s.baseStore.GetActiveMemberCount(teamId)
}

func (s KafkaTeamStore) GetTeamsForUser(userId string) store.StoreChannel {
	return s.baseStore.GetTeamsForUser(userId)
}

func (s KafkaTeamStore) GetTeamsForUserWithPagination(userId string, page int, perPage int) store.StoreChannel {
	return s.baseStore.GetTeamsForUserWithPagination(userId, page, perPage)
}

func (s KafkaTeamStore) GetChannelUnreadsForAllTeams(excludeTeamId string, userId string) store.StoreChannel {
	return s.baseStore.GetChannelUnreadsForAllTeams(excludeTeamId, userId)
}

func (s KafkaTeamStore) GetChannelUnreadsForTeam(teamId string, userId string) store.StoreChannel {
	return s.baseStore.GetChannelUnreadsForTeam(teamId, userId)
}

func (s KafkaTeamStore) RemoveMember(teamId string, userId string) store.StoreChannel {
	return s.baseStore.RemoveMember(teamId, userId)
}

func (s KafkaTeamStore) RemoveAllMembersByTeam(teamId string) store.StoreChannel {
	return s.baseStore.RemoveAllMembersByTeam(teamId)
}

func (s KafkaTeamStore) RemoveAllMembersByUser(userId string) store.StoreChannel {
	return s.baseStore.RemoveAllMembersByUser(userId)
}

func (s KafkaTeamStore) UpdateLastTeamIconUpdate(teamId string, curTime int64) store.StoreChannel {
	return s.baseStore.UpdateLastTeamIconUpdate(teamId, curTime)
}

func (s KafkaTeamStore) GetTeamsByScheme(schemeId string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetTeamsByScheme(schemeId, offset, limit)
}

func (s KafkaTeamStore) MigrateTeamMembers(fromTeamId string, fromUserId string) store.StoreChannel {
	return s.baseStore.MigrateTeamMembers(fromTeamId, fromUserId)
}

func (s KafkaTeamStore) ResetAllTeamSchemes() store.StoreChannel {
	return s.baseStore.ResetAllTeamSchemes()
}

func (s KafkaTeamStore) ClearAllCustomRoleAssignments() store.StoreChannel {
	return s.baseStore.ClearAllCustomRoleAssignments()
}

func (s KafkaTeamStore) AnalyticsGetTeamCountForScheme(schemeId string) store.StoreChannel {
	return s.baseStore.AnalyticsGetTeamCountForScheme(schemeId)
}

func (s KafkaTeamStore) GetAllForExportAfter(limit int, afterId string) store.StoreChannel {
	return s.baseStore.GetAllForExportAfter(limit, afterId)
}

func (s KafkaTeamStore) GetTeamMembersForExport(userId string) store.StoreChannel {
	return s.baseStore.GetTeamMembersForExport(userId)
}

func (s KafkaTeamStore) UserBelongsToTeams(userId string, teamIds []string) store.StoreChannel {
	return s.baseStore.UserBelongsToTeams(userId, teamIds)
}

func (s KafkaTeamStore) GetUserTeamIds(userId string, allowFromCache bool) store.StoreChannel {
	return s.baseStore.GetUserTeamIds(userId, allowFromCache)
}

func (s KafkaTeamStore) InvalidateAllTeamIdsForUser(userId string) {
	s.baseStore.InvalidateAllTeamIdsForUser(userId)
}

func (s KafkaTeamStore) ClearCaches() {
	s.baseStore.ClearCaches()
}
