package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaGroupStore struct {
	baseStore store.GroupStore
	root      *KafkaStore
}

func (s KafkaGroupStore) Create(group *model.Group) store.StoreChannel {
	return s.baseStore.Create(group)
}

func (s KafkaGroupStore) Get(groupID string) store.StoreChannel {
	return s.baseStore.Get(groupID)
}

func (s KafkaGroupStore) GetByRemoteID(remoteID string, groupSource model.GroupSource) store.StoreChannel {
	return s.baseStore.GetByRemoteID(remoteID, groupSource)
}

func (s KafkaGroupStore) GetAllBySource(groupSource model.GroupSource) store.StoreChannel {
	return s.baseStore.GetAllBySource(groupSource)
}

func (s KafkaGroupStore) Update(group *model.Group) store.StoreChannel {
	return s.baseStore.Update(group)
}

func (s KafkaGroupStore) Delete(groupID string) store.StoreChannel {
	return s.baseStore.Delete(groupID)
}

func (s KafkaGroupStore) GetMemberUsers(groupID string) store.StoreChannel {
	return s.baseStore.GetMemberUsers(groupID)
}

func (s KafkaGroupStore) GetMemberUsersPage(groupID string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetMemberUsersPage(groupID, offset, limit)
}

func (s KafkaGroupStore) GetMemberCount(groupID string) store.StoreChannel {
	return s.baseStore.GetMemberCount(groupID)
}

func (s KafkaGroupStore) CreateOrRestoreMember(groupID string, userID string) store.StoreChannel {
	return s.baseStore.CreateOrRestoreMember(groupID, userID)
}

func (s KafkaGroupStore) DeleteMember(groupID string, userID string) store.StoreChannel {
	return s.baseStore.DeleteMember(groupID, userID)
}

func (s KafkaGroupStore) CreateGroupSyncable(groupSyncable *model.GroupSyncable) store.StoreChannel {
	return s.baseStore.CreateGroupSyncable(groupSyncable)
}

func (s KafkaGroupStore) GetGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) store.StoreChannel {
	return s.baseStore.GetGroupSyncable(groupID, syncableID, syncableType)
}

func (s KafkaGroupStore) GetAllGroupSyncablesByGroupId(groupID string, syncableType model.GroupSyncableType) store.StoreChannel {
	return s.baseStore.GetAllGroupSyncablesByGroupId(groupID, syncableType)
}

func (s KafkaGroupStore) UpdateGroupSyncable(groupSyncable *model.GroupSyncable) store.StoreChannel {
	return s.baseStore.UpdateGroupSyncable(groupSyncable)
}

func (s KafkaGroupStore) DeleteGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) store.StoreChannel {
	return s.baseStore.DeleteGroupSyncable(groupID, syncableID, syncableType)
}

func (s KafkaGroupStore) TeamMembersToAdd(since int64) store.StoreChannel {
	return s.baseStore.TeamMembersToAdd(since)
}

func (s KafkaGroupStore) ChannelMembersToAdd(since int64) store.StoreChannel {
	return s.baseStore.ChannelMembersToAdd(since)
}

func (s KafkaGroupStore) TeamMembersToRemove() store.StoreChannel {
	return s.baseStore.TeamMembersToRemove()
}

func (s KafkaGroupStore) ChannelMembersToRemove() store.StoreChannel {
	return s.baseStore.ChannelMembersToRemove()
}

func (s KafkaGroupStore) GetGroupsByChannel(channelId string, page int, perPage int) store.StoreChannel {
	return s.baseStore.GetGroupsByChannel(channelId, page, perPage)
}

func (s KafkaGroupStore) GetGroupsByTeam(teamId string, page int, perPage int) store.StoreChannel {
	return s.baseStore.GetGroupsByTeam(teamId, page, perPage)
}
