package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentGroupStore struct {
	baseStore store.GroupStore
}

func (s TransparentGroupStore) Create(group *model.Group) store.StoreChannel {
	return s.baseStore.Create(group)
}

func (s TransparentGroupStore) Get(groupID string) store.StoreChannel {
	return s.baseStore.Get(groupID)
}

func (s TransparentGroupStore) GetByRemoteID(remoteID string, groupSource model.GroupSource) store.StoreChannel {
	return s.baseStore.GetByRemoteID(remoteID, groupSource)
}

func (s TransparentGroupStore) GetAllBySource(groupSource model.GroupSource) store.StoreChannel {
	return s.baseStore.GetAllBySource(groupSource)
}

func (s TransparentGroupStore) Update(group *model.Group) store.StoreChannel {
	return s.baseStore.Update(group)
}

func (s TransparentGroupStore) Delete(groupID string) store.StoreChannel {
	return s.baseStore.Delete(groupID)
}

func (s TransparentGroupStore) GetMemberUsers(groupID string) store.StoreChannel {
	return s.baseStore.GetMemberUsers(groupID)
}

func (s TransparentGroupStore) GetMemberUsersPage(groupID string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetMemberUsersPage(groupID, offset, limit)
}

func (s TransparentGroupStore) GetMemberCount(groupID string) store.StoreChannel {
	return s.baseStore.GetMemberCount(groupID)
}

func (s TransparentGroupStore) CreateOrRestoreMember(groupID string, userID string) store.StoreChannel {
	return s.baseStore.CreateOrRestoreMember(groupID, userID)
}

func (s TransparentGroupStore) DeleteMember(groupID string, userID string) store.StoreChannel {
	return s.baseStore.DeleteMember(groupID, userID)
}

func (s TransparentGroupStore) CreateGroupSyncable(groupSyncable *model.GroupSyncable) store.StoreChannel {
	return s.baseStore.CreateGroupSyncable(groupSyncable)
}

func (s TransparentGroupStore) GetGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) store.StoreChannel {
	return s.baseStore.GetGroupSyncable(groupID, syncableID, syncableType)
}

func (s TransparentGroupStore) GetAllGroupSyncablesByGroupId(groupID string, syncableType model.GroupSyncableType) store.StoreChannel {
	return s.baseStore.GetAllGroupSyncablesByGroupId(groupID, syncableType)
}

func (s TransparentGroupStore) UpdateGroupSyncable(groupSyncable *model.GroupSyncable) store.StoreChannel {
	return s.baseStore.UpdateGroupSyncable(groupSyncable)
}

func (s TransparentGroupStore) DeleteGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) store.StoreChannel {
	return s.baseStore.DeleteGroupSyncable(groupID, syncableID, syncableType)
}

func (s TransparentGroupStore) TeamMembersToAdd(since int64) store.StoreChannel {
	return s.baseStore.TeamMembersToAdd(since)
}

func (s TransparentGroupStore) ChannelMembersToAdd(since int64) store.StoreChannel {
	return s.baseStore.ChannelMembersToAdd(since)
}

func (s TransparentGroupStore) TeamMembersToRemove() store.StoreChannel {
	return s.baseStore.TeamMembersToRemove()
}

func (s TransparentGroupStore) ChannelMembersToRemove() store.StoreChannel {
	return s.baseStore.ChannelMembersToRemove()
}

func (s TransparentGroupStore) GetGroupsByChannel(channelId string, page int, perPage int) store.StoreChannel {
	return s.baseStore.GetGroupsByChannel(channelId, page, perPage)
}

func (s TransparentGroupStore) GetGroupsByTeam(teamId string, page int, perPage int) store.StoreChannel {
	return s.baseStore.GetGroupsByTeam(teamId, page, perPage)
}
