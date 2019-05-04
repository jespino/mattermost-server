package localcachestore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type LocalCacheGroupStore struct {
	baseStore store.GroupStore
	rootStore *LocalCacheStore
}

func (s LocalCacheGroupStore) Create(group *model.Group) store.StoreChannel {
	return s.baseStore.Create(group)
}

func (s LocalCacheGroupStore) Get(groupID string) store.StoreChannel {
	return s.baseStore.Get(groupID)
}

func (s LocalCacheGroupStore) GetByRemoteID(remoteID string, groupSource model.GroupSource) store.StoreChannel {
	return s.baseStore.GetByRemoteID(remoteID, groupSource)
}

func (s LocalCacheGroupStore) GetAllBySource(groupSource model.GroupSource) store.StoreChannel {
	return s.baseStore.GetAllBySource(groupSource)
}

func (s LocalCacheGroupStore) Update(group *model.Group) store.StoreChannel {
	return s.baseStore.Update(group)
}

func (s LocalCacheGroupStore) Delete(groupID string) store.StoreChannel {
	return s.baseStore.Delete(groupID)
}

func (s LocalCacheGroupStore) GetMemberUsers(groupID string) store.StoreChannel {
	return s.baseStore.GetMemberUsers(groupID)
}

func (s LocalCacheGroupStore) GetMemberUsersPage(groupID string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetMemberUsersPage(groupID, offset, limit)
}

func (s LocalCacheGroupStore) GetMemberCount(groupID string) store.StoreChannel {
	return s.baseStore.GetMemberCount(groupID)
}

func (s LocalCacheGroupStore) CreateOrRestoreMember(groupID string, userID string) store.StoreChannel {
	return s.baseStore.CreateOrRestoreMember(groupID, userID)
}

func (s LocalCacheGroupStore) DeleteMember(groupID string, userID string) store.StoreChannel {
	return s.baseStore.DeleteMember(groupID, userID)
}

func (s LocalCacheGroupStore) CreateGroupSyncable(groupSyncable *model.GroupSyncable) store.StoreChannel {
	return s.baseStore.CreateGroupSyncable(groupSyncable)
}

func (s LocalCacheGroupStore) GetGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) store.StoreChannel {
	return s.baseStore.GetGroupSyncable(groupID, syncableID, syncableType)
}

func (s LocalCacheGroupStore) GetAllGroupSyncablesByGroupId(groupID string, syncableType model.GroupSyncableType) store.StoreChannel {
	return s.baseStore.GetAllGroupSyncablesByGroupId(groupID, syncableType)
}

func (s LocalCacheGroupStore) UpdateGroupSyncable(groupSyncable *model.GroupSyncable) store.StoreChannel {
	return s.baseStore.UpdateGroupSyncable(groupSyncable)
}

func (s LocalCacheGroupStore) DeleteGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) store.StoreChannel {
	return s.baseStore.DeleteGroupSyncable(groupID, syncableID, syncableType)
}

func (s LocalCacheGroupStore) TeamMembersToAdd(since int64) store.StoreChannel {
	return s.baseStore.TeamMembersToAdd(since)
}

func (s LocalCacheGroupStore) ChannelMembersToAdd(since int64) store.StoreChannel {
	return s.baseStore.ChannelMembersToAdd(since)
}

func (s LocalCacheGroupStore) TeamMembersToRemove() store.StoreChannel {
	return s.baseStore.TeamMembersToRemove()
}

func (s LocalCacheGroupStore) ChannelMembersToRemove() store.StoreChannel {
	return s.baseStore.ChannelMembersToRemove()
}

func (s LocalCacheGroupStore) GetGroupsByChannel(channelId string, page int, perPage int) store.StoreChannel {
	return s.baseStore.GetGroupsByChannel(channelId, page, perPage)
}

func (s LocalCacheGroupStore) GetGroupsByTeam(teamId string, page int, perPage int) store.StoreChannel {
	return s.baseStore.GetGroupsByTeam(teamId, page, perPage)
}
