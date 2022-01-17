// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package memstore

import (
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/store"
)

type MemGroupStore struct{}

func newMemGroupStore() store.GroupStore {
	return &MemGroupStore{}
}

func (s *MemGroupStore) Create(group *model.Group) (*model.Group, error) {
	panic("not implemented")
}

func (s *MemGroupStore) Get(groupId string) (*model.Group, error) {
	panic("not implemented")
}

func (s *MemGroupStore) GetByName(name string, opts model.GroupSearchOpts) (*model.Group, error) {
	panic("not implemented")
}

func (s *MemGroupStore) GetByIDs(groupIDs []string) ([]*model.Group, error) {
	panic("not implemented")
}

func (s *MemGroupStore) GetByRemoteID(remoteID string, groupSource model.GroupSource) (*model.Group, error) {
	panic("not implemented")
}

func (s *MemGroupStore) GetAllBySource(groupSource model.GroupSource) ([]*model.Group, error) {
	panic("not implemented")
}

func (s *MemGroupStore) GetByUser(userId string) ([]*model.Group, error) {
	panic("not implemented")
}

func (s *MemGroupStore) Update(group *model.Group) (*model.Group, error) {
	panic("not implemented")
}

func (s *MemGroupStore) Delete(groupID string) (*model.Group, error) {
	panic("not implemented")
}

func (s *MemGroupStore) GetMemberUsers(groupID string) ([]*model.User, error) {
	panic("not implemented")
}

func (s *MemGroupStore) GetMemberUsersPage(groupID string, page int, perPage int) ([]*model.User, error) {
	panic("not implemented")
}

func (s *MemGroupStore) GetMemberCount(groupID string) (int64, error) {
	panic("not implemented")
}

func (s *MemGroupStore) GetMemberUsersInTeam(groupID string, teamID string) ([]*model.User, error) {
	panic("not implemented")
}

func (s *MemGroupStore) GetMemberUsersNotInChannel(groupID string, channelID string) ([]*model.User, error) {
	panic("not implemented")
}

func (s *MemGroupStore) UpsertMember(groupID string, userID string) (*model.GroupMember, error) {
	panic("not implemented")
}

func (s *MemGroupStore) DeleteMember(groupID string, userID string) (*model.GroupMember, error) {
	panic("not implemented")
}

func (s *MemGroupStore) PermanentDeleteMembersByUser(userId string) error {
	panic("not implemented")
}

func (s *MemGroupStore) CreateGroupSyncable(groupSyncable *model.GroupSyncable) (*model.GroupSyncable, error) {
	panic("not implemented")
}

func (s *MemGroupStore) GetGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) (*model.GroupSyncable, error) {
	panic("not implemented")
}

func (s *MemGroupStore) getGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) (*model.GroupSyncable, error) {
	panic("not implemented")
}

func (s *MemGroupStore) GetAllGroupSyncablesByGroupId(groupID string, syncableType model.GroupSyncableType) ([]*model.GroupSyncable, error) {
	panic("not implemented")
}

func (s *MemGroupStore) UpdateGroupSyncable(groupSyncable *model.GroupSyncable) (*model.GroupSyncable, error) {
	panic("not implemented")
}

func (s *MemGroupStore) DeleteGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) (*model.GroupSyncable, error) {
	panic("not implemented")
}

func (s *MemGroupStore) TeamMembersToAdd(since int64, teamID *string, includeRemovedMembers bool) ([]*model.UserTeamIDPair, error) {
	panic("not implemented")
}

func (s *MemGroupStore) ChannelMembersToAdd(since int64, channelID *string, includeRemovedMembers bool) ([]*model.UserChannelIDPair, error) {
	panic("not implemented")
}

func (s *MemGroupStore) TeamMembersToRemove(teamID *string) ([]*model.TeamMember, error) {
	panic("not implemented")
}

func (s *MemGroupStore) CountGroupsByChannel(channelId string, opts model.GroupSearchOpts) (int64, error) {
	panic("not implemented")
}

func (s *MemGroupStore) GetGroupsByChannel(channelId string, opts model.GroupSearchOpts) ([]*model.GroupWithSchemeAdmin, error) {
	panic("not implemented")
}

func (s *MemGroupStore) ChannelMembersToRemove(channelID *string) ([]*model.ChannelMember, error) {
	panic("not implemented")
}

func (s *MemGroupStore) CountGroupsByTeam(teamId string, opts model.GroupSearchOpts) (int64, error) {
	panic("not implemented")
}

func (s *MemGroupStore) GetGroupsByTeam(teamId string, opts model.GroupSearchOpts) ([]*model.GroupWithSchemeAdmin, error) {
	panic("not implemented")
}

func (s *MemGroupStore) GetGroupsAssociatedToChannelsByTeam(teamId string, opts model.GroupSearchOpts) (map[string][]*model.GroupWithSchemeAdmin, error) {
	panic("not implemented")
}

func (s *MemGroupStore) GetGroups(page, perPage int, opts model.GroupSearchOpts) ([]*model.Group, error) {
	panic("not implemented")
}

func (s *MemGroupStore) TeamMembersMinusGroupMembers(teamID string, groupIDs []string, page, perPage int) ([]*model.UserWithGroups, error) {
	panic("not implemented")
}

func (s *MemGroupStore) CountTeamMembersMinusGroupMembers(teamID string, groupIDs []string) (int64, error) {
	panic("not implemented")
}

func (s *MemGroupStore) ChannelMembersMinusGroupMembers(channelID string, groupIDs []string, page, perPage int) ([]*model.UserWithGroups, error) {
	panic("not implemented")
}

func (s *MemGroupStore) CountChannelMembersMinusGroupMembers(channelID string, groupIDs []string) (int64, error) {
	panic("not implemented")
}

func (s *MemGroupStore) AdminRoleGroupsForSyncableMember(userID, syncableID string, syncableType model.GroupSyncableType) ([]string, error) {
	panic("not implemented")
}

func (s *MemGroupStore) PermittedSyncableAdmins(syncableID string, syncableType model.GroupSyncableType) ([]string, error) {
	panic("not implemented")
}

func (s *MemGroupStore) GroupCount() (int64, error) {
	panic("not implemented")
}

func (s *MemGroupStore) GroupTeamCount() (int64, error) {
	panic("not implemented")
}

func (s *MemGroupStore) GroupChannelCount() (int64, error) {
	panic("not implemented")
}

func (s *MemGroupStore) GroupMemberCount() (int64, error) {
	panic("not implemented")
}

func (s *MemGroupStore) DistinctGroupMemberCount() (int64, error) {
	panic("not implemented")
}

func (s *MemGroupStore) GroupCountWithAllowReference() (int64, error) {
	panic("not implemented")
}
