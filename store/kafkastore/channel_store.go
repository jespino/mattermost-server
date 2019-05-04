package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaChannelStore struct {
	baseStore store.ChannelStore
	root      *KafkaStore
}

func (s KafkaChannelStore) Save(channel *model.Channel, maxChannelsPerTeam int64) store.StoreChannel {
	return s.baseStore.Save(channel, maxChannelsPerTeam)
}

func (s KafkaChannelStore) CreateDirectChannel(userId string, otherUserId string) store.StoreChannel {
	return s.baseStore.CreateDirectChannel(userId, otherUserId)
}

func (s KafkaChannelStore) SaveDirectChannel(channel *model.Channel, member1 *model.ChannelMember, member2 *model.ChannelMember) store.StoreChannel {
	return s.baseStore.SaveDirectChannel(channel, member1, member2)
}

func (s KafkaChannelStore) Update(channel *model.Channel) store.StoreChannel {
	return s.baseStore.Update(channel)
}

func (s KafkaChannelStore) Get(id string, allowFromCache bool) (*model.Channel, *model.AppError) {
	return s.baseStore.Get(id, allowFromCache)
}

func (s KafkaChannelStore) InvalidateChannel(id string) {
	s.baseStore.InvalidateChannel(id)
}

func (s KafkaChannelStore) InvalidateChannelByName(teamId string, name string) {
	s.baseStore.InvalidateChannelByName(teamId, name)
}

func (s KafkaChannelStore) GetFromMaster(id string) (*model.Channel, *model.AppError) {
	return s.baseStore.GetFromMaster(id)
}

func (s KafkaChannelStore) Delete(channelId string, time int64) store.StoreChannel {
	return s.baseStore.Delete(channelId, time)
}

func (s KafkaChannelStore) Restore(channelId string, time int64) store.StoreChannel {
	return s.baseStore.Restore(channelId, time)
}

func (s KafkaChannelStore) SetDeleteAt(channelId string, deleteAt int64, updateAt int64) store.StoreChannel {
	return s.baseStore.SetDeleteAt(channelId, deleteAt, updateAt)
}

func (s KafkaChannelStore) PermanentDeleteByTeam(teamId string) store.StoreChannel {
	return s.baseStore.PermanentDeleteByTeam(teamId)
}

func (s KafkaChannelStore) PermanentDelete(channelId string) store.StoreChannel {
	return s.baseStore.PermanentDelete(channelId)
}

func (s KafkaChannelStore) GetByName(team_id string, name string, allowFromCache bool) store.StoreChannel {
	return s.baseStore.GetByName(team_id, name, allowFromCache)
}

func (s KafkaChannelStore) GetByNames(team_id string, names []string, allowFromCache bool) store.StoreChannel {
	return s.baseStore.GetByNames(team_id, names, allowFromCache)
}

func (s KafkaChannelStore) GetByNameIncludeDeleted(team_id string, name string, allowFromCache bool) store.StoreChannel {
	return s.baseStore.GetByNameIncludeDeleted(team_id, name, allowFromCache)
}

func (s KafkaChannelStore) GetDeletedByName(team_id string, name string) store.StoreChannel {
	return s.baseStore.GetDeletedByName(team_id, name)
}

func (s KafkaChannelStore) GetDeleted(team_id string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetDeleted(team_id, offset, limit)
}

func (s KafkaChannelStore) GetChannels(teamId string, userId string, includeDeleted bool) store.StoreChannel {
	return s.baseStore.GetChannels(teamId, userId, includeDeleted)
}

func (s KafkaChannelStore) GetAllChannels(page int, perPage int, includeDeleted bool) store.StoreChannel {
	return s.baseStore.GetAllChannels(page, perPage, includeDeleted)
}

func (s KafkaChannelStore) GetMoreChannels(teamId string, userId string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetMoreChannels(teamId, userId, offset, limit)
}

func (s KafkaChannelStore) GetPublicChannelsForTeam(teamId string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetPublicChannelsForTeam(teamId, offset, limit)
}

func (s KafkaChannelStore) GetPublicChannelsByIdsForTeam(teamId string, channelIds []string) store.StoreChannel {
	return s.baseStore.GetPublicChannelsByIdsForTeam(teamId, channelIds)
}

func (s KafkaChannelStore) GetChannelCounts(teamId string, userId string) store.StoreChannel {
	return s.baseStore.GetChannelCounts(teamId, userId)
}

func (s KafkaChannelStore) GetTeamChannels(teamId string) store.StoreChannel {
	return s.baseStore.GetTeamChannels(teamId)
}

func (s KafkaChannelStore) GetAll(teamId string) store.StoreChannel {
	return s.baseStore.GetAll(teamId)
}

func (s KafkaChannelStore) GetChannelsByIds(channelIds []string) store.StoreChannel {
	return s.baseStore.GetChannelsByIds(channelIds)
}

func (s KafkaChannelStore) GetForPost(postId string) store.StoreChannel {
	return s.baseStore.GetForPost(postId)
}

func (s KafkaChannelStore) SaveMember(member *model.ChannelMember) store.StoreChannel {
	return s.baseStore.SaveMember(member)
}

func (s KafkaChannelStore) UpdateMember(member *model.ChannelMember) store.StoreChannel {
	return s.baseStore.UpdateMember(member)
}

func (s KafkaChannelStore) GetMembers(channelId string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetMembers(channelId, offset, limit)
}

func (s KafkaChannelStore) GetMember(channelId string, userId string) (*model.ChannelMember, *model.AppError) {
	return s.baseStore.GetMember(channelId, userId)
}

func (s KafkaChannelStore) GetChannelMembersTimezones(channelId string) store.StoreChannel {
	return s.baseStore.GetChannelMembersTimezones(channelId)
}

func (s KafkaChannelStore) GetAllChannelMembersForUser(userId string, allowFromCache bool, includeDeleted bool) store.StoreChannel {
	return s.baseStore.GetAllChannelMembersForUser(userId, allowFromCache, includeDeleted)
}

func (s KafkaChannelStore) InvalidateAllChannelMembersForUser(userId string) {
	s.baseStore.InvalidateAllChannelMembersForUser(userId)
}

func (s KafkaChannelStore) IsUserInChannelUseCache(userId string, channelId string) bool {
	return s.baseStore.IsUserInChannelUseCache(userId, channelId)
}

func (s KafkaChannelStore) GetAllChannelMembersNotifyPropsForChannel(channelId string, allowFromCache bool) store.StoreChannel {
	return s.baseStore.GetAllChannelMembersNotifyPropsForChannel(channelId, allowFromCache)
}

func (s KafkaChannelStore) InvalidateCacheForChannelMembersNotifyProps(channelId string) {
	s.baseStore.InvalidateCacheForChannelMembersNotifyProps(channelId)
}

func (s KafkaChannelStore) GetMemberForPost(postId string, userId string) store.StoreChannel {
	return s.baseStore.GetMemberForPost(postId, userId)
}

func (s KafkaChannelStore) InvalidateMemberCount(channelId string) {
	s.baseStore.InvalidateMemberCount(channelId)
}

func (s KafkaChannelStore) GetMemberCountFromCache(channelId string) int64 {
	return s.baseStore.GetMemberCountFromCache(channelId)
}

func (s KafkaChannelStore) GetMemberCount(channelId string, allowFromCache bool) store.StoreChannel {
	return s.baseStore.GetMemberCount(channelId, allowFromCache)
}

func (s KafkaChannelStore) GetPinnedPosts(channelId string) store.StoreChannel {
	return s.baseStore.GetPinnedPosts(channelId)
}

func (s KafkaChannelStore) RemoveMember(channelId string, userId string) store.StoreChannel {
	return s.baseStore.RemoveMember(channelId, userId)
}

func (s KafkaChannelStore) PermanentDeleteMembersByUser(userId string) store.StoreChannel {
	return s.baseStore.PermanentDeleteMembersByUser(userId)
}

func (s KafkaChannelStore) PermanentDeleteMembersByChannel(channelId string) store.StoreChannel {
	return s.baseStore.PermanentDeleteMembersByChannel(channelId)
}

func (s KafkaChannelStore) UpdateLastViewedAt(channelIds []string, userId string) store.StoreChannel {
	return s.baseStore.UpdateLastViewedAt(channelIds, userId)
}

func (s KafkaChannelStore) IncrementMentionCount(channelId string, userId string) store.StoreChannel {
	return s.baseStore.IncrementMentionCount(channelId, userId)
}

func (s KafkaChannelStore) AnalyticsTypeCount(teamId string, channelType string) store.StoreChannel {
	return s.baseStore.AnalyticsTypeCount(teamId, channelType)
}

func (s KafkaChannelStore) GetMembersForUser(teamId string, userId string) store.StoreChannel {
	return s.baseStore.GetMembersForUser(teamId, userId)
}

func (s KafkaChannelStore) GetMembersForUserWithPagination(teamId string, userId string, page int, perPage int) store.StoreChannel {
	return s.baseStore.GetMembersForUserWithPagination(teamId, userId, page, perPage)
}

func (s KafkaChannelStore) AutocompleteInTeam(teamId string, term string, includeDeleted bool) store.StoreChannel {
	return s.baseStore.AutocompleteInTeam(teamId, term, includeDeleted)
}

func (s KafkaChannelStore) AutocompleteInTeamForSearch(teamId string, userId string, term string, includeDeleted bool) store.StoreChannel {
	return s.baseStore.AutocompleteInTeamForSearch(teamId, userId, term, includeDeleted)
}

func (s KafkaChannelStore) SearchAllChannels(term string, includeDeleted bool) store.StoreChannel {
	return s.baseStore.SearchAllChannels(term, includeDeleted)
}

func (s KafkaChannelStore) SearchInTeam(teamId string, term string, includeDeleted bool) store.StoreChannel {
	return s.baseStore.SearchInTeam(teamId, term, includeDeleted)
}

func (s KafkaChannelStore) SearchMore(userId string, teamId string, term string) store.StoreChannel {
	return s.baseStore.SearchMore(userId, teamId, term)
}

func (s KafkaChannelStore) GetMembersByIds(channelId string, userIds []string) store.StoreChannel {
	return s.baseStore.GetMembersByIds(channelId, userIds)
}

func (s KafkaChannelStore) AnalyticsDeletedTypeCount(teamId string, channelType string) store.StoreChannel {
	return s.baseStore.AnalyticsDeletedTypeCount(teamId, channelType)
}

func (s KafkaChannelStore) GetChannelUnread(channelId string, userId string) store.StoreChannel {
	return s.baseStore.GetChannelUnread(channelId, userId)
}

func (s KafkaChannelStore) ClearCaches() {
	s.baseStore.ClearCaches()
}

func (s KafkaChannelStore) GetChannelsByScheme(schemeId string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetChannelsByScheme(schemeId, offset, limit)
}

func (s KafkaChannelStore) MigrateChannelMembers(fromChannelId string, fromUserId string) store.StoreChannel {
	return s.baseStore.MigrateChannelMembers(fromChannelId, fromUserId)
}

func (s KafkaChannelStore) ResetAllChannelSchemes() store.StoreChannel {
	return s.baseStore.ResetAllChannelSchemes()
}

func (s KafkaChannelStore) ClearAllCustomRoleAssignments() store.StoreChannel {
	return s.baseStore.ClearAllCustomRoleAssignments()
}

func (s KafkaChannelStore) MigratePublicChannels() error {
	return s.baseStore.MigratePublicChannels()
}

func (s KafkaChannelStore) GetAllChannelsForExportAfter(limit int, afterId string) store.StoreChannel {
	return s.baseStore.GetAllChannelsForExportAfter(limit, afterId)
}

func (s KafkaChannelStore) GetAllDirectChannelsForExportAfter(limit int, afterId string) store.StoreChannel {
	return s.baseStore.GetAllDirectChannelsForExportAfter(limit, afterId)
}

func (s KafkaChannelStore) GetChannelMembersForExport(userId string, teamId string) store.StoreChannel {
	return s.baseStore.GetChannelMembersForExport(userId, teamId)
}

func (s KafkaChannelStore) RemoveAllDeactivatedMembers(channelId string) store.StoreChannel {
	return s.baseStore.RemoveAllDeactivatedMembers(channelId)
}

func (s KafkaChannelStore) GetChannelsBatchForIndexing(startTime int64, endTime int64, limit int) store.StoreChannel {
	return s.baseStore.GetChannelsBatchForIndexing(startTime, endTime, limit)
}

func (s KafkaChannelStore) UserBelongsToChannels(userId string, channelIds []string) store.StoreChannel {
	return s.baseStore.UserBelongsToChannels(userId, channelIds)
}
