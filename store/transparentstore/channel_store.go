package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentChannelStore struct {
	baseStore store.ChannelStore
}

func (s TransparentChannelStore) Save(channel *model.Channel, maxChannelsPerTeam int64) store.StoreChannel {
	return s.baseStore.Save(channel, maxChannelsPerTeam)
}

func (s TransparentChannelStore) CreateDirectChannel(userId string, otherUserId string) store.StoreChannel {
	return s.baseStore.CreateDirectChannel(userId, otherUserId)
}

func (s TransparentChannelStore) SaveDirectChannel(channel *model.Channel, member1 *model.ChannelMember, member2 *model.ChannelMember) store.StoreChannel {
	return s.baseStore.SaveDirectChannel(channel, member1, member2)
}

func (s TransparentChannelStore) Update(channel *model.Channel) store.StoreChannel {
	return s.baseStore.Update(channel)
}

func (s TransparentChannelStore) Get(id string, allowFromCache bool) (*model.Channel, *model.AppError) {
	return s.baseStore.Get(id, allowFromCache)
}

func (s TransparentChannelStore) InvalidateChannel(id string) {
	s.baseStore.InvalidateChannel(id)
}

func (s TransparentChannelStore) InvalidateChannelByName(teamId string, name string) {
	s.baseStore.InvalidateChannelByName(teamId, name)
}

func (s TransparentChannelStore) GetFromMaster(id string) (*model.Channel, *model.AppError) {
	return s.baseStore.GetFromMaster(id)
}

func (s TransparentChannelStore) Delete(channelId string, time int64) store.StoreChannel {
	return s.baseStore.Delete(channelId, time)
}

func (s TransparentChannelStore) Restore(channelId string, time int64) store.StoreChannel {
	return s.baseStore.Restore(channelId, time)
}

func (s TransparentChannelStore) SetDeleteAt(channelId string, deleteAt int64, updateAt int64) store.StoreChannel {
	return s.baseStore.SetDeleteAt(channelId, deleteAt, updateAt)
}

func (s TransparentChannelStore) PermanentDeleteByTeam(teamId string) store.StoreChannel {
	return s.baseStore.PermanentDeleteByTeam(teamId)
}

func (s TransparentChannelStore) PermanentDelete(channelId string) store.StoreChannel {
	return s.baseStore.PermanentDelete(channelId)
}

func (s TransparentChannelStore) GetByName(team_id string, name string, allowFromCache bool) store.StoreChannel {
	return s.baseStore.GetByName(team_id, name, allowFromCache)
}

func (s TransparentChannelStore) GetByNames(team_id string, names []string, allowFromCache bool) store.StoreChannel {
	return s.baseStore.GetByNames(team_id, names, allowFromCache)
}

func (s TransparentChannelStore) GetByNameIncludeDeleted(team_id string, name string, allowFromCache bool) store.StoreChannel {
	return s.baseStore.GetByNameIncludeDeleted(team_id, name, allowFromCache)
}

func (s TransparentChannelStore) GetDeletedByName(team_id string, name string) store.StoreChannel {
	return s.baseStore.GetDeletedByName(team_id, name)
}

func (s TransparentChannelStore) GetDeleted(team_id string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetDeleted(team_id, offset, limit)
}

func (s TransparentChannelStore) GetChannels(teamId string, userId string, includeDeleted bool) store.StoreChannel {
	return s.baseStore.GetChannels(teamId, userId, includeDeleted)
}

func (s TransparentChannelStore) GetAllChannels(page int, perPage int, includeDeleted bool) store.StoreChannel {
	return s.baseStore.GetAllChannels(page, perPage, includeDeleted)
}

func (s TransparentChannelStore) GetMoreChannels(teamId string, userId string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetMoreChannels(teamId, userId, offset, limit)
}

func (s TransparentChannelStore) GetPublicChannelsForTeam(teamId string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetPublicChannelsForTeam(teamId, offset, limit)
}

func (s TransparentChannelStore) GetPublicChannelsByIdsForTeam(teamId string, channelIds []string) store.StoreChannel {
	return s.baseStore.GetPublicChannelsByIdsForTeam(teamId, channelIds)
}

func (s TransparentChannelStore) GetChannelCounts(teamId string, userId string) store.StoreChannel {
	return s.baseStore.GetChannelCounts(teamId, userId)
}

func (s TransparentChannelStore) GetTeamChannels(teamId string) store.StoreChannel {
	return s.baseStore.GetTeamChannels(teamId)
}

func (s TransparentChannelStore) GetAll(teamId string) store.StoreChannel {
	return s.baseStore.GetAll(teamId)
}

func (s TransparentChannelStore) GetChannelsByIds(channelIds []string) store.StoreChannel {
	return s.baseStore.GetChannelsByIds(channelIds)
}

func (s TransparentChannelStore) GetForPost(postId string) store.StoreChannel {
	return s.baseStore.GetForPost(postId)
}

func (s TransparentChannelStore) SaveMember(member *model.ChannelMember) store.StoreChannel {
	return s.baseStore.SaveMember(member)
}

func (s TransparentChannelStore) UpdateMember(member *model.ChannelMember) store.StoreChannel {
	return s.baseStore.UpdateMember(member)
}

func (s TransparentChannelStore) GetMembers(channelId string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetMembers(channelId, offset, limit)
}

func (s TransparentChannelStore) GetMember(channelId string, userId string) (*model.ChannelMember, *model.AppError) {
	return s.baseStore.GetMember(channelId, userId)
}

func (s TransparentChannelStore) GetChannelMembersTimezones(channelId string) store.StoreChannel {
	return s.baseStore.GetChannelMembersTimezones(channelId)
}

func (s TransparentChannelStore) GetAllChannelMembersForUser(userId string, allowFromCache bool, includeDeleted bool) store.StoreChannel {
	return s.baseStore.GetAllChannelMembersForUser(userId, allowFromCache, includeDeleted)
}

func (s TransparentChannelStore) InvalidateAllChannelMembersForUser(userId string) {
	s.baseStore.InvalidateAllChannelMembersForUser(userId)
}

func (s TransparentChannelStore) IsUserInChannelUseCache(userId string, channelId string) bool {
	return s.baseStore.IsUserInChannelUseCache(userId, channelId)
}

func (s TransparentChannelStore) GetAllChannelMembersNotifyPropsForChannel(channelId string, allowFromCache bool) store.StoreChannel {
	return s.baseStore.GetAllChannelMembersNotifyPropsForChannel(channelId, allowFromCache)
}

func (s TransparentChannelStore) InvalidateCacheForChannelMembersNotifyProps(channelId string) {
	s.baseStore.InvalidateCacheForChannelMembersNotifyProps(channelId)
}

func (s TransparentChannelStore) GetMemberForPost(postId string, userId string) store.StoreChannel {
	return s.baseStore.GetMemberForPost(postId, userId)
}

func (s TransparentChannelStore) InvalidateMemberCount(channelId string) {
	s.baseStore.InvalidateMemberCount(channelId)
}

func (s TransparentChannelStore) GetMemberCountFromCache(channelId string) int64 {
	return s.baseStore.GetMemberCountFromCache(channelId)
}

func (s TransparentChannelStore) GetMemberCount(channelId string, allowFromCache bool) store.StoreChannel {
	return s.baseStore.GetMemberCount(channelId, allowFromCache)
}

func (s TransparentChannelStore) GetPinnedPosts(channelId string) store.StoreChannel {
	return s.baseStore.GetPinnedPosts(channelId)
}

func (s TransparentChannelStore) RemoveMember(channelId string, userId string) store.StoreChannel {
	return s.baseStore.RemoveMember(channelId, userId)
}

func (s TransparentChannelStore) PermanentDeleteMembersByUser(userId string) store.StoreChannel {
	return s.baseStore.PermanentDeleteMembersByUser(userId)
}

func (s TransparentChannelStore) PermanentDeleteMembersByChannel(channelId string) store.StoreChannel {
	return s.baseStore.PermanentDeleteMembersByChannel(channelId)
}

func (s TransparentChannelStore) UpdateLastViewedAt(channelIds []string, userId string) store.StoreChannel {
	return s.baseStore.UpdateLastViewedAt(channelIds, userId)
}

func (s TransparentChannelStore) IncrementMentionCount(channelId string, userId string) store.StoreChannel {
	return s.baseStore.IncrementMentionCount(channelId, userId)
}

func (s TransparentChannelStore) AnalyticsTypeCount(teamId string, channelType string) store.StoreChannel {
	return s.baseStore.AnalyticsTypeCount(teamId, channelType)
}

func (s TransparentChannelStore) GetMembersForUser(teamId string, userId string) store.StoreChannel {
	return s.baseStore.GetMembersForUser(teamId, userId)
}

func (s TransparentChannelStore) GetMembersForUserWithPagination(teamId string, userId string, page int, perPage int) store.StoreChannel {
	return s.baseStore.GetMembersForUserWithPagination(teamId, userId, page, perPage)
}

func (s TransparentChannelStore) AutocompleteInTeam(teamId string, term string, includeDeleted bool) store.StoreChannel {
	return s.baseStore.AutocompleteInTeam(teamId, term, includeDeleted)
}

func (s TransparentChannelStore) AutocompleteInTeamForSearch(teamId string, userId string, term string, includeDeleted bool) store.StoreChannel {
	return s.baseStore.AutocompleteInTeamForSearch(teamId, userId, term, includeDeleted)
}

func (s TransparentChannelStore) SearchAllChannels(term string, includeDeleted bool) store.StoreChannel {
	return s.baseStore.SearchAllChannels(term, includeDeleted)
}

func (s TransparentChannelStore) SearchInTeam(teamId string, term string, includeDeleted bool) store.StoreChannel {
	return s.baseStore.SearchInTeam(teamId, term, includeDeleted)
}

func (s TransparentChannelStore) SearchMore(userId string, teamId string, term string) store.StoreChannel {
	return s.baseStore.SearchMore(userId, teamId, term)
}

func (s TransparentChannelStore) GetMembersByIds(channelId string, userIds []string) store.StoreChannel {
	return s.baseStore.GetMembersByIds(channelId, userIds)
}

func (s TransparentChannelStore) AnalyticsDeletedTypeCount(teamId string, channelType string) store.StoreChannel {
	return s.baseStore.AnalyticsDeletedTypeCount(teamId, channelType)
}

func (s TransparentChannelStore) GetChannelUnread(channelId string, userId string) store.StoreChannel {
	return s.baseStore.GetChannelUnread(channelId, userId)
}

func (s TransparentChannelStore) ClearCaches() {
	s.baseStore.ClearCaches()
}

func (s TransparentChannelStore) GetChannelsByScheme(schemeId string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetChannelsByScheme(schemeId, offset, limit)
}

func (s TransparentChannelStore) MigrateChannelMembers(fromChannelId string, fromUserId string) store.StoreChannel {
	return s.baseStore.MigrateChannelMembers(fromChannelId, fromUserId)
}

func (s TransparentChannelStore) ResetAllChannelSchemes() store.StoreChannel {
	return s.baseStore.ResetAllChannelSchemes()
}

func (s TransparentChannelStore) ClearAllCustomRoleAssignments() store.StoreChannel {
	return s.baseStore.ClearAllCustomRoleAssignments()
}

func (s TransparentChannelStore) MigratePublicChannels() error {
	return s.baseStore.MigratePublicChannels()
}

func (s TransparentChannelStore) GetAllChannelsForExportAfter(limit int, afterId string) store.StoreChannel {
	return s.baseStore.GetAllChannelsForExportAfter(limit, afterId)
}

func (s TransparentChannelStore) GetAllDirectChannelsForExportAfter(limit int, afterId string) store.StoreChannel {
	return s.baseStore.GetAllDirectChannelsForExportAfter(limit, afterId)
}

func (s TransparentChannelStore) GetChannelMembersForExport(userId string, teamId string) store.StoreChannel {
	return s.baseStore.GetChannelMembersForExport(userId, teamId)
}

func (s TransparentChannelStore) RemoveAllDeactivatedMembers(channelId string) store.StoreChannel {
	return s.baseStore.RemoveAllDeactivatedMembers(channelId)
}

func (s TransparentChannelStore) GetChannelsBatchForIndexing(startTime int64, endTime int64, limit int) store.StoreChannel {
	return s.baseStore.GetChannelsBatchForIndexing(startTime, endTime, limit)
}

func (s TransparentChannelStore) UserBelongsToChannels(userId string, channelIds []string) store.StoreChannel {
	return s.baseStore.UserBelongsToChannels(userId, channelIds)
}
