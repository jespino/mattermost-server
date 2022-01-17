// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package memstore

import (
	"context"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/store"
)

type MemChannelStore struct{}

func (s *MemChannelStore) ClearCaches() {}

func newMemChannelStore() store.ChannelStore {
	return &MemChannelStore{}
}
func (s *MemChannelStore) ClearSidebarOnTeamLeave(userId, teamId string) error {
	panic("not implemented")
}

func (s *MemChannelStore) CreateInitialSidebarCategories(userId, teamId string) (*model.OrderedSidebarCategories, error) {
	panic("not implemented")
}

func (s *MemChannelStore) MigrateFavoritesToSidebarChannels(lastUserId string, runningOrder int64) (map[string]interface{}, error) {
	panic("not implemented")
}

func (s *MemChannelStore) CreateSidebarCategory(userId, teamId string, newCategory *model.SidebarCategoryWithChannels) (*model.SidebarCategoryWithChannels, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetSidebarCategory(categoryId string) (*model.SidebarCategoryWithChannels, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetSidebarCategories(userId, teamId string) (*model.OrderedSidebarCategories, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetSidebarCategoryOrder(userId, teamId string) ([]string, error) {
	panic("not implemented")
}

func (s *MemChannelStore) UpdateSidebarCategoryOrder(userId, teamId string, categoryOrder []string) error {
	panic("not implemented")
}

func (s *MemChannelStore) UpdateSidebarCategories(userId, teamId string, categories []*model.SidebarCategoryWithChannels) ([]*model.SidebarCategoryWithChannels, []*model.SidebarCategoryWithChannels, error) {
	panic("not implemented")
}

func (s *MemChannelStore) UpdateSidebarChannelsByPreferences(preferences model.Preferences) error {
	panic("not implemented")
}

func (s *MemChannelStore) DeleteSidebarChannelsByPreferences(preferences model.Preferences) error {
	panic("not implemented")
}

func (s *MemChannelStore) UpdateSidebarChannelCategoryOnMove(channel *model.Channel, newTeamId string) error {
	panic("not implemented")
}

func (s *MemChannelStore) DeleteSidebarCategory(categoryId string) error {
	panic("not implemented")
}

func (s *MemChannelStore) Save(channel *model.Channel, maxChannelsPerTeam int64) (*model.Channel, error) {
	panic("not implemented")
}

func (s *MemChannelStore) CreateDirectChannel(user *model.User, otherUser *model.User, channelOptions ...model.ChannelOption) (*model.Channel, error) {
	panic("not implemented")
}

func (s *MemChannelStore) SaveDirectChannel(directChannel *model.Channel, member1 *model.ChannelMember, member2 *model.ChannelMember) (*model.Channel, error) {
	panic("not implemented")
}

func (s *MemChannelStore) Update(channel *model.Channel) (*model.Channel, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetChannelUnread(channelId, userId string) (*model.ChannelUnread, error) {
	panic("not implemented")
}

func (s *MemChannelStore) InvalidateChannel(id string) {
	panic("not implemented")
}

func (s *MemChannelStore) InvalidateChannelByName(teamId, name string) {
	panic("not implemented")
}

func (s *MemChannelStore) Get(id string, allowFromCache bool) (*model.Channel, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetPinnedPosts(channelId string) (*model.PostList, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetFromMaster(id string) (*model.Channel, error) {
	panic("not implemented")
}

func (s *MemChannelStore) Delete(channelId string, time int64) error {
	panic("not implemented")
}

func (s *MemChannelStore) Restore(channelId string, time int64) error {
	panic("not implemented")
}

func (s *MemChannelStore) SetDeleteAt(channelId string, deleteAt, updateAt int64) error {
	panic("not implemented")
}

func (s *MemChannelStore) PermanentDeleteByTeam(teamId string) error {
	panic("not implemented")
}

func (s *MemChannelStore) PermanentDelete(channelId string) error {
	panic("not implemented")
}

func (s *MemChannelStore) PermanentDeleteMembersByChannel(channelId string) error {
	panic("not implemented")
}

func (s *MemChannelStore) GetChannels(teamId string, userId string, includeDeleted bool, lastDeleteAt int) (model.ChannelList, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetChannelsByUser(userId string, includeDeleted bool, lastDeleteAt, pageSize int, fromChannelID string) (model.ChannelList, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetAllChannelMembersById(channelID string) ([]string, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetAllChannels(offset, limit int, opts store.ChannelSearchOpts) (model.ChannelListWithTeamData, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetAllChannelsCount(opts store.ChannelSearchOpts) (int64, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetMoreChannels(teamId string, userId string, offset int, limit int) (model.ChannelList, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetPrivateChannelsForTeam(teamId string, offset int, limit int) (model.ChannelList, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetPublicChannelsForTeam(teamId string, offset int, limit int) (model.ChannelList, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetPublicChannelsByIdsForTeam(teamId string, channelIds []string) (model.ChannelList, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetChannelCounts(teamId string, userId string) (*model.ChannelCounts, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetTeamChannels(teamId string) (model.ChannelList, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetByName(teamId string, name string, allowFromCache bool) (*model.Channel, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetByNames(teamId string, names []string, allowFromCache bool) ([]*model.Channel, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetByNameIncludeDeleted(teamId string, name string, allowFromCache bool) (*model.Channel, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetDeletedByName(teamId string, name string) (*model.Channel, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetDeleted(teamId string, offset int, limit int, userId string) (model.ChannelList, error) {
	panic("not implemented")
}

func (s *MemChannelStore) SaveMultipleMembers(members []*model.ChannelMember) ([]*model.ChannelMember, error) {
	panic("not implemented")
}

func (s *MemChannelStore) SaveMember(member *model.ChannelMember) (*model.ChannelMember, error) {
	panic("not implemented")
}

func (s *MemChannelStore) UpdateMultipleMembers(members []*model.ChannelMember) ([]*model.ChannelMember, error) {
	panic("not implemented")
}

func (s *MemChannelStore) UpdateMember(member *model.ChannelMember) (*model.ChannelMember, error) {
	panic("not implemented")
}

func (s *MemChannelStore) UpdateMemberNotifyProps(channelID, userID string, props map[string]string) (*model.ChannelMember, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetMembers(channelId string, offset, limit int) (model.ChannelMembers, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetChannelMembersTimezones(channelId string) ([]model.StringMap, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetMember(ctx context.Context, channelId string, userId string) (*model.ChannelMember, error) {
	panic("not implemented")
}

func (s *MemChannelStore) InvalidateAllChannelMembersForUser(userId string) {
	panic("not implemented")
}

func (s *MemChannelStore) IsUserInChannelUseCache(userId string, channelId string) bool {
	panic("not implemented")
}

func (s *MemChannelStore) GetMemberForPost(postId string, userId string) (*model.ChannelMember, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetAllChannelMembersForUser(userId string, allowFromCache bool, includeDeleted bool) (map[string]string, error) {
	panic("not implemented")
}

func (s *MemChannelStore) InvalidateCacheForChannelMembersNotifyProps(channelId string) {
	panic("not implemented")
}

func (s *MemChannelStore) GetAllChannelMembersNotifyPropsForChannel(channelId string, allowFromCache bool) (map[string]model.StringMap, error) {
	panic("not implemented")
}

func (s *MemChannelStore) InvalidateMemberCount(channelId string) {
	panic("not implemented")
}

func (s *MemChannelStore) GetMemberCountFromCache(channelId string) int64 {
	panic("not implemented")
}

func (s *MemChannelStore) GetMemberCount(channelId string, allowFromCache bool) (int64, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetMemberCountsByGroup(ctx context.Context, channelID string, includeTimezones bool) ([]*model.ChannelMemberCountByGroup, error) {
	panic("not implemented")
}

func (s *MemChannelStore) InvalidatePinnedPostCount(channelId string) {
	panic("not implemented")
}

func (s *MemChannelStore) GetPinnedPostCount(channelId string, allowFromCache bool) (int64, error) {
	panic("not implemented")
}

func (s *MemChannelStore) InvalidateGuestCount(channelId string) {
	panic("not implemented")
}

func (s *MemChannelStore) GetGuestCount(channelId string, allowFromCache bool) (int64, error) {
	panic("not implemented")
}

func (s *MemChannelStore) RemoveMembers(channelId string, userIds []string) error {
	panic("not implemented")
}

func (s *MemChannelStore) RemoveMember(channelId string, userId string) error {
	panic("not implemented")
}

func (s *MemChannelStore) RemoveAllDeactivatedMembers(channelId string) error {
	panic("not implemented")
}

func (s *MemChannelStore) PermanentDeleteMembersByUser(userId string) error {
	panic("not implemented")
}

func (s *MemChannelStore) UpdateLastViewedAt(channelIds []string, userId string, updateThreads bool) (map[string]int64, error) {
	panic("not implemented")
}

func (s *MemChannelStore) CountPostsAfter(channelId string, timestamp int64, userId string) (int, int, error) {
	panic("not implemented")
}

func (s *MemChannelStore) UpdateLastViewedAtPost(unreadPost *model.Post, userID string, mentionCount, mentionCountRoot int, updateThreads bool, setUnreadCountRoot bool) (*model.ChannelUnreadAt, error) {
	panic("not implemented")
}

func (s *MemChannelStore) IncrementMentionCount(channelId string, userId string, updateThreads, isRoot bool) error {
	panic("not implemented")
}

func (s *MemChannelStore) GetAll(teamId string) ([]*model.Channel, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetChannelsByIds(channelIds []string, includeDeleted bool) ([]*model.Channel, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetChannelsWithTeamDataByIds(channelIDs []string, includeDeleted bool) ([]*model.ChannelWithTeamData, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetForPost(postId string) (*model.Channel, error) {
	panic("not implemented")
}

func (s *MemChannelStore) AnalyticsTypeCount(teamId string, channelType model.ChannelType) (int64, error) {
	panic("not implemented")
}

func (s *MemChannelStore) AnalyticsDeletedTypeCount(teamId string, channelType string) (int64, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetMembersForUser(teamId string, userId string) (model.ChannelMembers, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetMembersForUserWithPagination(userId string, page, perPage int) (model.ChannelMembersWithTeamData, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetTeamMembersForChannel(channelID string) ([]string, error) {
	panic("not implemented")
}

func (s *MemChannelStore) Autocomplete(userID, term string, includeDeleted bool) (model.ChannelListWithTeamData, error) {
	panic("not implemented")
}

func (s *MemChannelStore) AutocompleteInTeam(teamID, userID, term string, includeDeleted bool) (model.ChannelList, error) {
	panic("not implemented")
}

func (s *MemChannelStore) AutocompleteInTeamForSearch(teamId string, userId string, term string, includeDeleted bool) (model.ChannelList, error) {
	panic("not implemented")
}

func (s *MemChannelStore) SearchInTeam(teamId string, term string, includeDeleted bool) (model.ChannelList, error) {
	panic("not implemented")
}

func (s *MemChannelStore) SearchArchivedInTeam(teamId string, term string, userId string) (model.ChannelList, error) {
	panic("not implemented")
}

func (s *MemChannelStore) SearchForUserInTeam(userId string, teamId string, term string, includeDeleted bool) (model.ChannelList, error) {
	panic("not implemented")
}

func (s *MemChannelStore) SearchAllChannels(term string, opts store.ChannelSearchOpts) (model.ChannelListWithTeamData, int64, error) {
	panic("not implemented")
}

func (s *MemChannelStore) SearchMore(userId string, teamId string, term string) (model.ChannelList, error) {
	panic("not implemented")
}

func (s *MemChannelStore) SearchGroupChannels(userId, term string) (model.ChannelList, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetMembersByIds(channelId string, userIds []string) (model.ChannelMembers, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetMembersByChannelIds(channelIds []string, userId string) (model.ChannelMembers, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetChannelsByScheme(schemeId string, offset int, limit int) (model.ChannelList, error) {
	panic("not implemented")
}

func (s *MemChannelStore) MigrateChannelMembers(fromChannelId string, fromUserId string) (map[string]string, error) {
	panic("not implemented")
}

func (s *MemChannelStore) ResetAllChannelSchemes() error {
	panic("not implemented")
}

func (s *MemChannelStore) ClearAllCustomRoleAssignments() error {
	panic("not implemented")
}

func (s *MemChannelStore) GetAllChannelsForExportAfter(limit int, afterId string) ([]*model.ChannelForExport, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetChannelMembersForExport(userId string, teamId string) ([]*model.ChannelMemberForExport, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetAllDirectChannelsForExportAfter(limit int, afterId string) ([]*model.DirectChannelForExport, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetChannelsBatchForIndexing(startTime, endTime int64, limit int) ([]*model.Channel, error) {
	panic("not implemented")
}

func (s *MemChannelStore) UserBelongsToChannels(userId string, channelIds []string) (bool, error) {
	panic("not implemented")
}

func (s *MemChannelStore) UpdateMembersRole(channelID string, userIDs []string) error {
	panic("not implemented")
}

func (s *MemChannelStore) GroupSyncedChannelCount() (int64, error) {
	panic("not implemented")
}

func (s *MemChannelStore) SetShared(channelId string, shared bool) error {
	panic("not implemented")
}

func (s *MemChannelStore) GetTeamForChannel(channelID string) (*model.Team, error) {
	panic("not implemented")
}
