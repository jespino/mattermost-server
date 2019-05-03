package transparentstore

import "github.com/mattermost/mattermost-server/store"

type TransparentChannelMemberHistoryStore struct {
	baseStore store.ChannelMemberHistoryStore
}

func (s TransparentChannelMemberHistoryStore) LogJoinEvent(userId string, channelId string, joinTime int64) store.StoreChannel {
	return s.baseStore.LogJoinEvent(userId, channelId, joinTime)
}

func (s TransparentChannelMemberHistoryStore) LogLeaveEvent(userId string, channelId string, leaveTime int64) store.StoreChannel {
	return s.baseStore.LogLeaveEvent(userId, channelId, leaveTime)
}

func (s TransparentChannelMemberHistoryStore) GetUsersInChannelDuring(startTime int64, endTime int64, channelId string) store.StoreChannel {
	return s.baseStore.GetUsersInChannelDuring(startTime, endTime, channelId)
}

func (s TransparentChannelMemberHistoryStore) PermanentDeleteBatch(endTime int64, limit int64) store.StoreChannel {
	return s.baseStore.PermanentDeleteBatch(endTime, limit)
}
