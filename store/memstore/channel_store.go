// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package memstore

import (
	"context"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/store"
)

type MemChannelStore struct {
	channels []*model.Channel
	members  []*model.ChannelMember
}

func (s *MemChannelStore) ClearCaches() {}

func newMemChannelStore() store.ChannelStore {
	return &MemChannelStore{}
}
func (s *MemChannelStore) ClearSidebarOnTeamLeave(userId, teamId string) error {
	panic("not implemented")
}

func (s *MemChannelStore) CreateInitialSidebarCategories(userId, teamId string) (*model.OrderedSidebarCategories, error) {
	// TODO: Implement this
	return &model.OrderedSidebarCategories{}, nil
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
	if channel.DeleteAt != 0 {
		return nil, store.NewErrInvalidInput("Channel", "DeleteAt", channel.DeleteAt)
	}

	if channel.Type == model.ChannelTypeDirect {
		return nil, store.NewErrInvalidInput("Channel", "Type", channel.Type)
	}

	if channel.Id != "" && !channel.IsShared() {
		return nil, store.NewErrInvalidInput("Channel", "Id", channel.Id)
	}

	channel.PreSave()
	if err := channel.IsValid(); err != nil { // TODO: this needs to return plain error in v6.
		return nil, err // we just pass through the error as-is for now.
	}

	if channel.Type != model.ChannelTypeDirect && channel.Type != model.ChannelTypeGroup && maxChannelsPerTeam >= 0 {
		var count int64 = 0
		for _, c := range s.channels {
			if c.DeleteAt == 0 {
				count++
			}
		}
		if count >= maxChannelsPerTeam {
			return nil, store.NewErrLimitExceeded("channels_per_team", int(count), "teamId="+channel.TeamId)
		}
	}

	return channel, nil
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

func (s *MemChannelStore) InvalidateChannel(id string) {}

func (s *MemChannelStore) InvalidateChannelByName(teamId, name string) {}

func (s *MemChannelStore) Get(id string, allowFromCache bool) (*model.Channel, error) {
	for _, c := range s.channels {
		if c.Id == id {
			return c, nil
		}
	}
	return nil, store.NewErrNotFound("Channel", id)
}

func (s *MemChannelStore) GetPinnedPosts(channelId string) (*model.PostList, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetFromMaster(id string) (*model.Channel, error) {
	panic("not implemented")
}

func (s *MemChannelStore) Delete(channelId string, time int64) error {
	c, _ := s.Get(channelId, false)
	if c != nil && c.DeleteAt == 0 {
		c.DeleteAt = time
		c.UpdateAt = time
	}
	return store.NewErrNotFound("Channel", channelId)
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
	for _, c := range s.channels {
		if c.DeleteAt == 0 && c.TeamId == teamId && c.Name == name {
			return c, nil
		}
	}
	return nil, store.NewErrNotFound("Channel", teamId+"-"+name)
}

func (s *MemChannelStore) GetByNames(teamId string, names []string, allowFromCache bool) ([]*model.Channel, error) {
	result := []*model.Channel{}
	for _, c := range s.channels {
		if c.TeamId == teamId {
			for _, n := range names {
				if c.Name == n {
					result = append(result, c)
				}
			}
		}
	}
	return result, nil
}

func (s *MemChannelStore) GetByNameIncludeDeleted(teamId string, name string, allowFromCache bool) (*model.Channel, error) {
	for _, c := range s.channels {
		if c.TeamId == teamId && c.Name == name {
			return c, nil
		}
	}
	return nil, store.NewErrNotFound("Channel", teamId+"-"+name)
}

func (s *MemChannelStore) GetDeletedByName(teamId string, name string) (*model.Channel, error) {
	for _, c := range s.channels {
		if c.DeleteAt != 0 && c.TeamId == teamId && c.Name == name {
			return c, nil
		}
	}
	return nil, store.NewErrNotFound("Channel", teamId+"-"+name)
}

func (s *MemChannelStore) GetDeleted(teamId string, offset int, limit int, userId string) (model.ChannelList, error) {
	panic("not implemented")
}

func (s *MemChannelStore) SaveMultipleMembers(members []*model.ChannelMember) ([]*model.ChannelMember, error) {
	for _, member := range members {
		defer s.InvalidateAllChannelMembersForUser(member.UserId)
	}

	newChannelMembers := map[string]int{}
	users := map[string]bool{}
	for _, member := range members {
		newChannelMembers[member.ChannelId] = 0
	}

	for _, member := range members {
		newChannelMembers[member.ChannelId]++
		users[member.UserId] = true

		if err := member.IsValid(); err != nil {
			return nil, err
		}
	}

	s.members = append(s.members, members...)

	return members, nil
}

func (s *MemChannelStore) SaveMember(member *model.ChannelMember) (*model.ChannelMember, error) {
	newMembers, err := s.SaveMultipleMembers([]*model.ChannelMember{member})
	if err != nil {
		return nil, err
	}
	return newMembers[0], nil
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
	for _, m := range s.members {
		if m.ChannelId == channelId && m.UserId == userId {
			return m, nil
		}
	}
	return nil, store.NewErrNotFound("ChannelMember", channelId+"-"+userId)
}

func (s *MemChannelStore) InvalidateAllChannelMembersForUser(userId string) {}

func (s *MemChannelStore) IsUserInChannelUseCache(userId string, channelId string) bool {
	panic("not implemented")
}

func (s *MemChannelStore) GetMemberForPost(postId string, userId string) (*model.ChannelMember, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetAllChannelMembersForUser(userId string, allowFromCache bool, includeDeleted bool) (map[string]string, error) {
	panic("not implemented")
}

func (s *MemChannelStore) InvalidateCacheForChannelMembersNotifyProps(channelId string) {}

func (s *MemChannelStore) GetAllChannelMembersNotifyPropsForChannel(channelId string, allowFromCache bool) (map[string]model.StringMap, error) {
	// TODO: Implement this
	return map[string]model.StringMap{}, nil
}

func (s *MemChannelStore) InvalidateMemberCount(channelId string) {}

func (s *MemChannelStore) GetMemberCountFromCache(channelId string) int64 {
	panic("not implemented")
}

func (s *MemChannelStore) GetMemberCount(channelId string, allowFromCache bool) (int64, error) {
	panic("not implemented")
}

func (s *MemChannelStore) GetMemberCountsByGroup(ctx context.Context, channelID string, includeTimezones bool) ([]*model.ChannelMemberCountByGroup, error) {
	panic("not implemented")
}

func (s *MemChannelStore) InvalidatePinnedPostCount(channelId string) {}

func (s *MemChannelStore) GetPinnedPostCount(channelId string, allowFromCache bool) (int64, error) {
	panic("not implemented")
}

func (s *MemChannelStore) InvalidateGuestCount(channelId string) {}

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
