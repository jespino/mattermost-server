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
	// TODO: Implement this
	return &model.Group{}, nil
}

func (s *MemGroupStore) Get(groupId string) (*model.Group, error) {
	// TODO: Implement this
	return &model.Group{}, nil
}

func (s *MemGroupStore) GetByName(name string, opts model.GroupSearchOpts) (*model.Group, error) {
	// TODO: Implement this
	return &model.Group{}, nil
}

func (s *MemGroupStore) GetByIDs(groupIDs []string) ([]*model.Group, error) {
	// TODO: Implement this
	return []*model.Group{}, nil
}

func (s *MemGroupStore) GetByRemoteID(remoteID string, groupSource model.GroupSource) (*model.Group, error) {
	// TODO: Implement this
	return &model.Group{}, nil
}

func (s *MemGroupStore) GetAllBySource(groupSource model.GroupSource) ([]*model.Group, error) {
	// TODO: Implement this
	return []*model.Group{}, nil
}

func (s *MemGroupStore) GetByUser(userId string) ([]*model.Group, error) {
	// TODO: Implement this
	return []*model.Group{}, nil
}

func (s *MemGroupStore) Update(group *model.Group) (*model.Group, error) {
	// TODO: Implement this
	return &model.Group{}, nil
}

func (s *MemGroupStore) Delete(groupID string) (*model.Group, error) {
	// TODO: Implement this
	return &model.Group{}, nil
}

func (s *MemGroupStore) GetMemberUsers(groupID string) ([]*model.User, error) {
	// TODO: Implement this
	return []*model.User{}, nil
}

func (s *MemGroupStore) GetMemberUsersPage(groupID string, page int, perPage int) ([]*model.User, error) {
	// TODO: Implement this
	return []*model.User{}, nil
}

func (s *MemGroupStore) GetMemberCount(groupID string) (int64, error) {
	// TODO: Implement this
	return 0, nil
}

func (s *MemGroupStore) GetMemberUsersInTeam(groupID string, teamID string) ([]*model.User, error) {
	// TODO: Implement this
	return []*model.User{}, nil
}

func (s *MemGroupStore) GetMemberUsersNotInChannel(groupID string, channelID string) ([]*model.User, error) {
	// TODO: Implement this
	return []*model.User{}, nil
}

func (s *MemGroupStore) UpsertMember(groupID string, userID string) (*model.GroupMember, error) {
	// TODO: Implement this
	return &model.GroupMember{}, nil
}

func (s *MemGroupStore) DeleteMember(groupID string, userID string) (*model.GroupMember, error) {
	// TODO: Implement this
	return &model.GroupMember{}, nil
}

func (s *MemGroupStore) PermanentDeleteMembersByUser(userId string) error {
	// TODO: Implement this
	return nil
}

func (s *MemGroupStore) CreateGroupSyncable(groupSyncable *model.GroupSyncable) (*model.GroupSyncable, error) {
	// TODO: Implement this
	return &model.GroupSyncable{}, nil
}

func (s *MemGroupStore) GetGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) (*model.GroupSyncable, error) {
	// TODO: Implement this
	return &model.GroupSyncable{}, nil
}

func (s *MemGroupStore) getGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) (*model.GroupSyncable, error) {
	// TODO: Implement this
	return &model.GroupSyncable{}, nil
}

func (s *MemGroupStore) GetAllGroupSyncablesByGroupId(groupID string, syncableType model.GroupSyncableType) ([]*model.GroupSyncable, error) {
	// TODO: Implement this
	return []*model.GroupSyncable{}, nil
}

func (s *MemGroupStore) UpdateGroupSyncable(groupSyncable *model.GroupSyncable) (*model.GroupSyncable, error) {
	// TODO: Implement this
	return &model.GroupSyncable{}, nil
}

func (s *MemGroupStore) DeleteGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) (*model.GroupSyncable, error) {
	// TODO: Implement this
	return &model.GroupSyncable{}, nil
}

func (s *MemGroupStore) TeamMembersToAdd(since int64, teamID *string, includeRemovedMembers bool) ([]*model.UserTeamIDPair, error) {
	// TODO: Implement this
	return []*model.UserTeamIDPair{}, nil
}

func (s *MemGroupStore) ChannelMembersToAdd(since int64, channelID *string, includeRemovedMembers bool) ([]*model.UserChannelIDPair, error) {
	// TODO: Implement this
	return []*model.UserChannelIDPair{}, nil
}

func (s *MemGroupStore) TeamMembersToRemove(teamID *string) ([]*model.TeamMember, error) {
	// TODO: Implement this
	return []*model.TeamMember{}, nil
}

func (s *MemGroupStore) CountGroupsByChannel(channelId string, opts model.GroupSearchOpts) (int64, error) {
	// TODO: Implement this
	return 0, nil
}

func (s *MemGroupStore) GetGroupsByChannel(channelId string, opts model.GroupSearchOpts) ([]*model.GroupWithSchemeAdmin, error) {
	// TODO: Implement this
	return []*model.GroupWithSchemeAdmin{}, nil
}

func (s *MemGroupStore) ChannelMembersToRemove(channelID *string) ([]*model.ChannelMember, error) {
	// TODO: Implement this
	return []*model.ChannelMember{}, nil
}

func (s *MemGroupStore) CountGroupsByTeam(teamId string, opts model.GroupSearchOpts) (int64, error) {
	// TODO: Implement this
	return 0, nil
}

func (s *MemGroupStore) GetGroupsByTeam(teamId string, opts model.GroupSearchOpts) ([]*model.GroupWithSchemeAdmin, error) {
	// TODO: Implement this
	return []*model.GroupWithSchemeAdmin{}, nil
}

func (s *MemGroupStore) GetGroupsAssociatedToChannelsByTeam(teamId string, opts model.GroupSearchOpts) (map[string][]*model.GroupWithSchemeAdmin, error) {
	// TODO: Implement this
	return map[string][]*model.GroupWithSchemeAdmin{}, nil
}

func (s *MemGroupStore) GetGroups(page, perPage int, opts model.GroupSearchOpts) ([]*model.Group, error) {
	// TODO: Implement this
	return []*model.Group{}, nil
}

func (s *MemGroupStore) TeamMembersMinusGroupMembers(teamID string, groupIDs []string, page, perPage int) ([]*model.UserWithGroups, error) {
	// TODO: Implement this
	return []*model.UserWithGroups{}, nil
}

func (s *MemGroupStore) CountTeamMembersMinusGroupMembers(teamID string, groupIDs []string) (int64, error) {
	// TODO: Implement this
	return 0, nil
}

func (s *MemGroupStore) ChannelMembersMinusGroupMembers(channelID string, groupIDs []string, page, perPage int) ([]*model.UserWithGroups, error) {
	// TODO: Implement this
	return []*model.UserWithGroups{}, nil
}

func (s *MemGroupStore) CountChannelMembersMinusGroupMembers(channelID string, groupIDs []string) (int64, error) {
	// TODO: Implement this
	return 0, nil
}

func (s *MemGroupStore) AdminRoleGroupsForSyncableMember(userID, syncableID string, syncableType model.GroupSyncableType) ([]string, error) {
	// TODO: Implement this
	return []string{}, nil
}

func (s *MemGroupStore) PermittedSyncableAdmins(syncableID string, syncableType model.GroupSyncableType) ([]string, error) {
	// TODO: Implement this
	return []string{}, nil
}

func (s *MemGroupStore) GroupCount() (int64, error) {
	// TODO: Implement this
	return 0, nil
}

func (s *MemGroupStore) GroupTeamCount() (int64, error) {
	// TODO: Implement this
	return 0, nil
}

func (s *MemGroupStore) GroupChannelCount() (int64, error) {
	// TODO: Implement this
	return 0, nil
}

func (s *MemGroupStore) GroupMemberCount() (int64, error) {
	// TODO: Implement this
	return 0, nil
}

func (s *MemGroupStore) DistinctGroupMemberCount() (int64, error) {
	// TODO: Implement this
	return 0, nil
}

func (s *MemGroupStore) GroupCountWithAllowReference() (int64, error) {
	// TODO: Implement this
	return 0, nil
}
