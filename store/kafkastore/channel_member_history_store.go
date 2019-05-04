package kafkastore

import "github.com/mattermost/mattermost-server/store"

type KafkaChannelMemberHistoryStore struct {
	baseStore store.ChannelMemberHistoryStore
	root      *KafkaStore
}

func (s KafkaChannelMemberHistoryStore) LogJoinEvent(userId string, channelId string, joinTime int64) store.StoreChannel {
	return s.baseStore.LogJoinEvent(userId, channelId, joinTime)
}

func (s KafkaChannelMemberHistoryStore) LogLeaveEvent(userId string, channelId string, leaveTime int64) store.StoreChannel {
	return s.baseStore.LogLeaveEvent(userId, channelId, leaveTime)
}

func (s KafkaChannelMemberHistoryStore) GetUsersInChannelDuring(startTime int64, endTime int64, channelId string) store.StoreChannel {
	return s.baseStore.GetUsersInChannelDuring(startTime, endTime, channelId)
}

func (s KafkaChannelMemberHistoryStore) PermanentDeleteBatch(endTime int64, limit int64) store.StoreChannel {
	return s.baseStore.PermanentDeleteBatch(endTime, limit)
}
