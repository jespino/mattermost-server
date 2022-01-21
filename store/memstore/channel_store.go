// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package memstore

import (
	"context"
	"strings"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/store"
)

type allChannelMember struct {
	ChannelId                     string
	Roles                         string
	SchemeGuest                   bool
	SchemeUser                    bool
	SchemeAdmin                   bool
	TeamSchemeDefaultGuestRole    string
	TeamSchemeDefaultUserRole     string
	TeamSchemeDefaultAdminRole    string
	ChannelSchemeDefaultGuestRole string
	ChannelSchemeDefaultUserRole  string
	ChannelSchemeDefaultAdminRole string
}

type allChannelMembers []allChannelMember

func (db allChannelMember) Process() (string, string) {
	roles := strings.Fields(db.Roles)

	// Add any scheme derived roles that are not in the Roles field due to being Implicit from the Scheme, and add
	// them to the Roles field for backwards compatibility reasons.
	var schemeImpliedRoles []string
	if db.SchemeGuest {
		if db.ChannelSchemeDefaultGuestRole != "" {
			schemeImpliedRoles = append(schemeImpliedRoles, db.ChannelSchemeDefaultGuestRole)
		} else if db.TeamSchemeDefaultGuestRole != "" {
			schemeImpliedRoles = append(schemeImpliedRoles, db.TeamSchemeDefaultGuestRole)
		} else {
			schemeImpliedRoles = append(schemeImpliedRoles, model.ChannelGuestRoleId)
		}
	}
	if db.SchemeUser {
		if db.ChannelSchemeDefaultUserRole != "" {
			schemeImpliedRoles = append(schemeImpliedRoles, db.ChannelSchemeDefaultUserRole)
		} else if db.TeamSchemeDefaultUserRole != "" {
			schemeImpliedRoles = append(schemeImpliedRoles, db.TeamSchemeDefaultUserRole)
		} else {
			schemeImpliedRoles = append(schemeImpliedRoles, model.ChannelUserRoleId)
		}
	}
	if db.SchemeAdmin {
		if db.ChannelSchemeDefaultAdminRole != "" {
			schemeImpliedRoles = append(schemeImpliedRoles, db.ChannelSchemeDefaultAdminRole)
		} else if db.TeamSchemeDefaultAdminRole != "" {
			schemeImpliedRoles = append(schemeImpliedRoles, db.TeamSchemeDefaultAdminRole)
		} else {
			schemeImpliedRoles = append(schemeImpliedRoles, model.ChannelAdminRoleId)
		}
	}
	for _, impliedRole := range schemeImpliedRoles {
		alreadyThere := false
		for _, role := range roles {
			if role == impliedRole {
				alreadyThere = true
			}
		}
		if !alreadyThere {
			roles = append(roles, impliedRole)
		}
	}

	return db.ChannelId, strings.Join(roles, " ")
}

func (db allChannelMembers) ToMapStringString() map[string]string {
	result := make(map[string]string)

	for _, item := range db {
		key, value := item.Process()
		result[key] = value
	}

	return result
}

type MemChannelStore struct {
	MemStore *MemStore
	channels []*model.Channel
	members  []*model.ChannelMember
}

func (s *MemChannelStore) ClearCaches() {}

func newMemChannelStore(memStore *MemStore) store.ChannelStore {
	return &MemChannelStore{MemStore: memStore}
}
func (s *MemChannelStore) ClearSidebarOnTeamLeave(userId, teamId string) error {
	// TODO: Implement this
	return nil
}

func (s *MemChannelStore) CreateInitialSidebarCategories(userId, teamId string) (*model.OrderedSidebarCategories, error) {
	// TODO: Implement this
	return &model.OrderedSidebarCategories{}, nil
}

func (s *MemChannelStore) MigrateFavoritesToSidebarChannels(lastUserId string, runningOrder int64) (map[string]interface{}, error) {
	// TODO: Implement this
	return map[string]interface{}{}, nil
}

func (s *MemChannelStore) CreateSidebarCategory(userId, teamId string, newCategory *model.SidebarCategoryWithChannels) (*model.SidebarCategoryWithChannels, error) {
	// TODO: Implement this
	return &model.SidebarCategoryWithChannels{}, nil
}

func (s *MemChannelStore) GetSidebarCategory(categoryId string) (*model.SidebarCategoryWithChannels, error) {
	// TODO: Implement this
	return &model.SidebarCategoryWithChannels{}, nil
}

func (s *MemChannelStore) GetSidebarCategories(userId, teamId string) (*model.OrderedSidebarCategories, error) {
	// TODO: Implement this
	return &model.OrderedSidebarCategories{}, nil
}

func (s *MemChannelStore) GetSidebarCategoryOrder(userId, teamId string) ([]string, error) {
	// TODO: Implement this
	return []string{}, nil
}

func (s *MemChannelStore) UpdateSidebarCategoryOrder(userId, teamId string, categoryOrder []string) error {
	// TODO: Implement this
	return nil
}

func (s *MemChannelStore) UpdateSidebarCategories(userId, teamId string, categories []*model.SidebarCategoryWithChannels) ([]*model.SidebarCategoryWithChannels, []*model.SidebarCategoryWithChannels, error) {
	// TODO: Implement this
	return []*model.SidebarCategoryWithChannels{}, []*model.SidebarCategoryWithChannels{}, nil
}

func (s *MemChannelStore) UpdateSidebarChannelsByPreferences(preferences model.Preferences) error {
	// TODO: Implement this
	return nil
}

func (s *MemChannelStore) DeleteSidebarChannelsByPreferences(preferences model.Preferences) error {
	// TODO: Implement this
	return nil
}

func (s *MemChannelStore) UpdateSidebarChannelCategoryOnMove(channel *model.Channel, newTeamId string) error {
	// TODO: Implement this
	return nil
}

func (s *MemChannelStore) DeleteSidebarCategory(categoryId string) error {
	// TODO: Implement this
	return nil
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

	s.channels = append(s.channels, channel)

	return channel, nil
}

func (s *MemChannelStore) CreateDirectChannel(user *model.User, otherUser *model.User, channelOptions ...model.ChannelOption) (*model.Channel, error) {
	channel := new(model.Channel)

	for _, option := range channelOptions {
		option(channel)
	}

	channel.DisplayName = ""
	channel.Name = model.GetDMNameFromIds(otherUser.Id, user.Id)

	channel.Header = ""
	channel.Type = model.ChannelTypeDirect
	channel.Shared = model.NewBool(user.IsRemote() || otherUser.IsRemote())
	channel.CreatorId = user.Id

	cm1 := &model.ChannelMember{
		UserId:      user.Id,
		NotifyProps: model.GetDefaultChannelNotifyProps(),
		SchemeGuest: user.IsGuest(),
		SchemeUser:  !user.IsGuest(),
	}
	cm2 := &model.ChannelMember{
		UserId:      otherUser.Id,
		NotifyProps: model.GetDefaultChannelNotifyProps(),
		SchemeGuest: otherUser.IsGuest(),
		SchemeUser:  !otherUser.IsGuest(),
	}

	return s.SaveDirectChannel(channel, cm1, cm2)
}

func (s *MemChannelStore) SaveDirectChannel(directChannel *model.Channel, member1 *model.ChannelMember, member2 *model.ChannelMember) (*model.Channel, error) {
	if directChannel.DeleteAt != 0 {
		return nil, store.NewErrInvalidInput("Channel", "DeleteAt", directChannel.DeleteAt)
	}

	if directChannel.Type != model.ChannelTypeDirect {
		return nil, store.NewErrInvalidInput("Channel", "Type", directChannel.Type)
	}

	directChannel.TeamId = ""
	newChannel, err := s.Save(directChannel, 0)
	if err != nil {
		return newChannel, err
	}

	// Members need new channel ID
	member1.ChannelId = newChannel.Id
	member2.ChannelId = newChannel.Id

	if member1.UserId != member2.UserId {
		_, err = s.SaveMultipleMembers([]*model.ChannelMember{member1, member2})
	} else {
		_, err = s.SaveMember(member2)
	}
	if err != nil {
		return nil, err
	}
	return newChannel, nil
}

func (s *MemChannelStore) Update(channel *model.Channel) (*model.Channel, error) {
	channel.PreUpdate()

	if channel.DeleteAt != 0 {
		return nil, store.NewErrInvalidInput("Channel", "DeleteAt", channel.DeleteAt)
	}

	if err := channel.IsValid(); err != nil {
		return nil, err
	}

	existing, _ := s.GetByName(channel.TeamId, channel.Name, false)
	if existing != nil {
		return nil, store.NewErrInvalidInput("Channel", "Id", channel.Id)
	}
	if existing.Id != channel.Id {
		return nil, store.NewErrInvalidInput("Channel", "Id", channel.Id)
	}
	channel.UpdateAt = model.GetMillis()
	*existing = *channel

	return channel, nil
}

func (s *MemChannelStore) GetChannelUnread(channelId, userId string) (*model.ChannelUnread, error) {
	// TODO: Implement this
	return &model.ChannelUnread{}, nil
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
	// TODO: Implement this
	return &model.PostList{}, nil
}

func (s *MemChannelStore) GetFromMaster(id string) (*model.Channel, error) {
	return s.Get(id, false)
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
	// TODO: Implement this
	return nil
}

func (s *MemChannelStore) SetDeleteAt(channelId string, deleteAt, updateAt int64) error {
	// TODO: Implement this
	return nil
}

func (s *MemChannelStore) PermanentDeleteByTeam(teamId string) error {
	result := []*model.Channel{}
	for _, c := range s.channels {
		if c.TeamId != teamId {
			result = append(result, c)
		}
	}
	s.channels = result
	return nil
}

func (s *MemChannelStore) PermanentDelete(channelId string) error {
	result := []*model.Channel{}
	for _, c := range s.channels {
		if c.Id != channelId {
			result = append(result, c)
		}
	}
	s.channels = result
	return nil
}

func (s *MemChannelStore) PermanentDeleteMembersByChannel(channelId string) error {
	result := []*model.ChannelMember{}
	for _, m := range s.members {
		if m.ChannelId != channelId {
			result = append(result, m)
		}
	}
	s.members = result
	return nil
}

func (s *MemChannelStore) GetChannels(teamId string, userId string, includeDeleted bool, lastDeleteAt int) (model.ChannelList, error) {
	result := model.ChannelList{}
	for _, m := range s.members {
		if m.UserId != userId {
			continue
		}
		c, err := s.Get(m.ChannelId, false)
		if err != nil {
			return nil, err
		}
		if c.TeamId != "" && c.TeamId != teamId {
			continue
		}
		if !includeDeleted && c.DeleteAt != 0 {
			continue
		}
		if includeDeleted && lastDeleteAt != 0 && c.DeleteAt < int64(lastDeleteAt) {
			continue
		}
		result = append(result, c)
	}
	return result, nil
}

func (s *MemChannelStore) GetChannelsByUser(userId string, includeDeleted bool, lastDeleteAt, pageSize int, fromChannelID string) (model.ChannelList, error) {
	// TODO: Implement this
	return model.ChannelList{}, nil
}

func (s *MemChannelStore) GetAllChannelMembersById(channelID string) ([]string, error) {
	// TODO: Implement this
	return []string{}, nil
}

func (s *MemChannelStore) GetAllChannels(offset, limit int, opts store.ChannelSearchOpts) (model.ChannelListWithTeamData, error) {
	// TODO: Implement this
	return model.ChannelListWithTeamData{}, nil
}

func (s *MemChannelStore) GetAllChannelsCount(opts store.ChannelSearchOpts) (int64, error) {
	// TODO: Implement this
	return 0, nil
}

func (s *MemChannelStore) GetMoreChannels(teamId string, userId string, offset int, limit int) (model.ChannelList, error) {
	// TODO: Implement this
	return model.ChannelList{}, nil
}

func (s *MemChannelStore) GetPrivateChannelsForTeam(teamId string, offset int, limit int) (model.ChannelList, error) {
	// TODO: Implement this
	return model.ChannelList{}, nil
}

func (s *MemChannelStore) GetPublicChannelsForTeam(teamId string, offset int, limit int) (model.ChannelList, error) {
	// TODO: Implement this
	return model.ChannelList{}, nil
}

func (s *MemChannelStore) GetPublicChannelsByIdsForTeam(teamId string, channelIds []string) (model.ChannelList, error) {
	// TODO: Implement this
	return model.ChannelList{}, nil
}

func (s *MemChannelStore) GetChannelCounts(teamId string, userId string) (*model.ChannelCounts, error) {
	// TODO: Implement this
	return &model.ChannelCounts{}, nil
}

func (s *MemChannelStore) GetTeamChannels(teamId string) (model.ChannelList, error) {
	result := model.ChannelList{}
	for _, c := range s.channels {
		if c.TeamId == teamId && c.Type != "D" {
			result = append(result, c)
		}
	}
	return result, nil
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
	// TODO: Implement this
	return model.ChannelList{}, nil
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
	// TODO: Implement this
	return []*model.ChannelMember{}, nil
}

func (s *MemChannelStore) UpdateMember(member *model.ChannelMember) (*model.ChannelMember, error) {
	// TODO: Implement this
	return &model.ChannelMember{}, nil
}

func (s *MemChannelStore) UpdateMemberNotifyProps(channelID, userID string, props map[string]string) (*model.ChannelMember, error) {
	// TODO: Implement this
	return &model.ChannelMember{}, nil
}

func (s *MemChannelStore) GetMembers(channelId string, offset, limit int) (model.ChannelMembers, error) {
	result := model.ChannelMembers{}
	for _, m := range s.members {
		counter := 0
		if m.ChannelId == channelId {
			if counter >= offset && counter < offset+limit {
				result = append(result, *m)
			}
			counter++
		}
	}
	return result, nil
}

func (s *MemChannelStore) GetChannelMembersTimezones(channelId string) ([]model.StringMap, error) {
	// TODO: Implement this
	return []model.StringMap{}, nil
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
	// TODO: Implement this
	return false
}

func (s *MemChannelStore) GetMemberForPost(postId string, userId string) (*model.ChannelMember, error) {
	// TODO: Implement this
	return &model.ChannelMember{}, nil
}

func (s *MemChannelStore) GetAllChannelMembersForUser(userId string, allowFromCache bool, includeDeleted bool) (map[string]string, error) {
	var data allChannelMembers
	for _, m := range s.members {
		if m.UserId != userId {
			channel, err := s.Get(m.ChannelId, false)
			if err != nil {
				return nil, err
			}
			channelMember := allChannelMember{
				ChannelId:   m.ChannelId,
				Roles:       m.Roles,
				SchemeGuest: m.SchemeGuest,
				SchemeUser:  m.SchemeUser,
				SchemeAdmin: m.SchemeAdmin,
			}
			if channel.SchemeId != nil {
				channelScheme, err := s.MemStore.Scheme().Get(*channel.SchemeId)
				if err != nil {
					return nil, err
				}
				channelMember.ChannelSchemeDefaultGuestRole = channelScheme.DefaultChannelGuestRole
				channelMember.ChannelSchemeDefaultUserRole = channelScheme.DefaultChannelUserRole
				channelMember.ChannelSchemeDefaultAdminRole = channelScheme.DefaultChannelAdminRole
			}
			if channel.TeamId != "" {
				team, err := s.Get(channel.TeamId, false)
				if err != nil {
					return nil, err
				}
				if team.SchemeId != nil {
					teamScheme, err := s.MemStore.Scheme().Get(*team.SchemeId)
					if err != nil {
						return nil, err
					}
					channelMember.TeamSchemeDefaultGuestRole = teamScheme.DefaultTeamGuestRole
					channelMember.TeamSchemeDefaultUserRole = teamScheme.DefaultTeamUserRole
					channelMember.TeamSchemeDefaultAdminRole = teamScheme.DefaultTeamAdminRole
				}
			}
			if includeDeleted || channel.DeleteAt == 0 {
				data = append(data, channelMember)
			}
		}
	}
	return data.ToMapStringString(), nil
}

func (s *MemChannelStore) InvalidateCacheForChannelMembersNotifyProps(channelId string) {}

func (s *MemChannelStore) GetAllChannelMembersNotifyPropsForChannel(channelId string, allowFromCache bool) (map[string]model.StringMap, error) {
	// TODO: Implement this
	return map[string]model.StringMap{}, nil
}

func (s *MemChannelStore) InvalidateMemberCount(channelId string) {}

func (s *MemChannelStore) GetMemberCountFromCache(channelId string) int64 {
	count, _ := s.GetMemberCount(channelId, true)
	return count
}

func (s *MemChannelStore) GetMemberCount(channelId string, allowFromCache bool) (int64, error) {
	var count int64 = 0
	for _, m := range s.members {
		if m.ChannelId == channelId {
			user, err := s.MemStore.User().Get(context.Background(), m.UserId)
			if err != nil {
				return 0, err
			}
			if user.DeleteAt == 0 {
				count++
			}
		}
	}
	return count, nil
}

func (s *MemChannelStore) GetMemberCountsByGroup(ctx context.Context, channelID string, includeTimezones bool) ([]*model.ChannelMemberCountByGroup, error) {
	// TODO: Implement this
	return []*model.ChannelMemberCountByGroup{}, nil
}

func (s *MemChannelStore) InvalidatePinnedPostCount(channelId string) {}

func (s *MemChannelStore) GetPinnedPostCount(channelId string, allowFromCache bool) (int64, error) {
	// TODO: Implement this
	return 0, nil
}

func (s *MemChannelStore) InvalidateGuestCount(channelId string) {}

func (s *MemChannelStore) GetGuestCount(channelId string, allowFromCache bool) (int64, error) {
	// TODO: Implement this
	return 0, nil
}

func (s *MemChannelStore) RemoveMembers(channelId string, userIds []string) error {
	result := []*model.ChannelMember{}
	for _, m := range s.members {
		skip := false
		if m.ChannelId == channelId {
			for _, userId := range userIds {
				if m.UserId == userId {
					skip = true
					break
				}
			}
		}
		if !skip {
			result = append(result, m)
		}
	}
	s.members = result
	return nil
}

func (s *MemChannelStore) RemoveMember(channelId string, userId string) error {
	result := []*model.ChannelMember{}
	for _, m := range s.members {
		if m.ChannelId != channelId || m.UserId == userId {
			result = append(result, m)
		}
	}
	s.members = result
	return nil
}

func (s *MemChannelStore) RemoveAllDeactivatedMembers(channelId string) error {
	result := []*model.ChannelMember{}
	for _, m := range s.members {
		if m.ChannelId != channelId {
			result = append(result, m)
			continue
		}

		user, err := s.MemStore.User().Get(context.Background(), m.UserId)
		if err != nil {
			return err
		}
		if user.DeleteAt == 0 {
			result = append(result, m)
		}
	}
	s.members = result
	return nil
}

func (s *MemChannelStore) PermanentDeleteMembersByUser(userId string) error {
	// TODO: Implement this
	return nil
}

func (s *MemChannelStore) UpdateLastViewedAt(channelIds []string, userId string, updateThreads bool) (map[string]int64, error) {
	// TODO: Implement this
	return map[string]int64{}, nil
}

func (s *MemChannelStore) CountPostsAfter(channelId string, timestamp int64, userId string) (int, int, error) {
	// TODO: Implement this
	return 0, 0, nil
}

func (s *MemChannelStore) UpdateLastViewedAtPost(unreadPost *model.Post, userID string, mentionCount, mentionCountRoot int, updateThreads bool, setUnreadCountRoot bool) (*model.ChannelUnreadAt, error) {
	// TODO: Implement this
	return &model.ChannelUnreadAt{}, nil
}

func (s *MemChannelStore) IncrementMentionCount(channelId string, userId string, updateThreads, isRoot bool) error {
	// TODO: Implement this
	return nil
}

func (s *MemChannelStore) GetAll(teamId string) ([]*model.Channel, error) {
	// TODO: Implement this
	return []*model.Channel{}, nil
}

func (s *MemChannelStore) GetChannelsByIds(channelIds []string, includeDeleted bool) ([]*model.Channel, error) {
	// TODO: Implement this
	return []*model.Channel{}, nil
}

func (s *MemChannelStore) GetChannelsWithTeamDataByIds(channelIDs []string, includeDeleted bool) ([]*model.ChannelWithTeamData, error) {
	// TODO: Implement this
	return []*model.ChannelWithTeamData{}, nil
}

func (s *MemChannelStore) GetForPost(postId string) (*model.Channel, error) {
	// TODO: Implement this
	return &model.Channel{}, nil
}

func (s *MemChannelStore) AnalyticsTypeCount(teamId string, channelType model.ChannelType) (int64, error) {
	// TODO: Implement this
	return 0, nil
}

func (s *MemChannelStore) AnalyticsDeletedTypeCount(teamId string, channelType string) (int64, error) {
	// TODO: Implement this
	return 0, nil
}

func (s *MemChannelStore) GetMembersForUser(teamId string, userId string) (model.ChannelMembers, error) {
	result := model.ChannelMembers{}
	for _, m := range s.members {
		if m.UserId == userId {
			channel, err := s.Get(m.ChannelId, false)
			if err != nil {
				return nil, err
			}
			if channel.TeamId == teamId {
				result = append(result, *m)
			}
		}
	}
	return result, nil
}

func (s *MemChannelStore) GetMembersForUserWithPagination(userId string, page, perPage int) (model.ChannelMembersWithTeamData, error) {
	result := model.ChannelMembersWithTeamData{}
	counter := 0
	offset := page * perPage
	for _, m := range s.members {
		if m.UserId == userId {
			if counter >= offset && counter < offset-perPage {
				channel, err := s.Get(m.ChannelId, false)
				if err != nil {
					return nil, err
				}

				team, err := s.MemStore.Team().Get(channel.TeamId)
				if err != nil {
					return nil, err
				}
				result = append(result, model.ChannelMemberWithTeamData{
					ChannelMember:   *m,
					TeamDisplayName: team.DisplayName,
					TeamName:        team.Name,
					TeamUpdateAt:    team.UpdateAt,
				})
			}
			counter++
		}
	}
	return result, nil
}

func (s *MemChannelStore) GetTeamMembersForChannel(channelID string) ([]string, error) {
	// TODO: Implement this
	return []string{}, nil
}

func (s *MemChannelStore) Autocomplete(userID, term string, includeDeleted bool) (model.ChannelListWithTeamData, error) {
	// TODO: Implement this
	return model.ChannelListWithTeamData{}, nil
}

func (s *MemChannelStore) AutocompleteInTeam(teamID, userID, term string, includeDeleted bool) (model.ChannelList, error) {
	// TODO: Implement this
	return model.ChannelList{}, nil
}

func (s *MemChannelStore) AutocompleteInTeamForSearch(teamId string, userId string, term string, includeDeleted bool) (model.ChannelList, error) {
	// TODO: Implement this
	return model.ChannelList{}, nil
}

func (s *MemChannelStore) SearchInTeam(teamId string, term string, includeDeleted bool) (model.ChannelList, error) {
	// TODO: Implement this
	return model.ChannelList{}, nil
}

func (s *MemChannelStore) SearchArchivedInTeam(teamId string, term string, userId string) (model.ChannelList, error) {
	// TODO: Implement this
	return model.ChannelList{}, nil
}

func (s *MemChannelStore) SearchForUserInTeam(userId string, teamId string, term string, includeDeleted bool) (model.ChannelList, error) {
	// TODO: Implement this
	return model.ChannelList{}, nil
}

func (s *MemChannelStore) SearchAllChannels(term string, opts store.ChannelSearchOpts) (model.ChannelListWithTeamData, int64, error) {
	// TODO: Implement this
	return model.ChannelListWithTeamData{}, 0, nil
}

func (s *MemChannelStore) SearchMore(userId string, teamId string, term string) (model.ChannelList, error) {
	// TODO: Implement this
	return model.ChannelList{}, nil
}

func (s *MemChannelStore) SearchGroupChannels(userId, term string) (model.ChannelList, error) {
	// TODO: Implement this
	return model.ChannelList{}, nil
}

func (s *MemChannelStore) GetMembersByIds(channelId string, userIds []string) (model.ChannelMembers, error) {
	// TODO: Implement this
	return model.ChannelMembers{}, nil
}

func (s *MemChannelStore) GetMembersByChannelIds(channelIds []string, userId string) (model.ChannelMembers, error) {
	result := model.ChannelMembers{}
	for _, m := range s.members {
		if m.UserId == userId {
			for _, c := range channelIds {
				if m.ChannelId == c {
					result = append(result, *m)
					break
				}
			}
		}
	}
	return result, nil
}

func (s *MemChannelStore) GetChannelsByScheme(schemeId string, offset int, limit int) (model.ChannelList, error) {
	// TODO: Implement this
	return model.ChannelList{}, nil
}

func (s *MemChannelStore) MigrateChannelMembers(fromChannelId string, fromUserId string) (map[string]string, error) {
	// TODO: Implement this
	return map[string]string{}, nil
}

func (s *MemChannelStore) ResetAllChannelSchemes() error {
	// TODO: Implement this
	return nil
}

func (s *MemChannelStore) ClearAllCustomRoleAssignments() error {
	// TODO: Implement this
	return nil
}

func (s *MemChannelStore) GetAllChannelsForExportAfter(limit int, afterId string) ([]*model.ChannelForExport, error) {
	// TODO: Implement this
	return []*model.ChannelForExport{}, nil
}

func (s *MemChannelStore) GetChannelMembersForExport(userId string, teamId string) ([]*model.ChannelMemberForExport, error) {
	// TODO: Implement this
	return []*model.ChannelMemberForExport{}, nil
}

func (s *MemChannelStore) GetAllDirectChannelsForExportAfter(limit int, afterId string) ([]*model.DirectChannelForExport, error) {
	// TODO: Implement this
	return []*model.DirectChannelForExport{}, nil
}

func (s *MemChannelStore) GetChannelsBatchForIndexing(startTime, endTime int64, limit int) ([]*model.Channel, error) {
	// TODO: Implement this
	return []*model.Channel{}, nil
}

func (s *MemChannelStore) UserBelongsToChannels(userId string, channelIds []string) (bool, error) {
	// TODO: Implement this
	return true, nil
}

func (s *MemChannelStore) UpdateMembersRole(channelID string, userIDs []string) error {
	// TODO: Implement this
	return nil
}

func (s *MemChannelStore) GroupSyncedChannelCount() (int64, error) {
	// TODO: Implement this
	return 0, nil
}

func (s *MemChannelStore) SetShared(channelId string, shared bool) error {
	// TODO: Implement this
	return nil
}

func (s *MemChannelStore) GetTeamForChannel(channelID string) (*model.Team, error) {
	// TODO: Implement this
	return &model.Team{}, nil
}
