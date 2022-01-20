// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package memstore

import (
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/store"
)

type MemChannelMemberHistoryStore struct {
	membersHistory []*model.ChannelMemberHistory
}

func newMemChannelMemberHistoryStore() store.ChannelMemberHistoryStore {
	return &MemChannelMemberHistoryStore{}
}

func (s *MemChannelMemberHistoryStore) LogJoinEvent(userId string, channelId string, joinTime int64) error {
	s.membersHistory = append(s.membersHistory, &model.ChannelMemberHistory{UserId: userId, ChannelId: channelId, JoinTime: joinTime})
	return nil
}

func (s *MemChannelMemberHistoryStore) LogLeaveEvent(userId string, channelId string, leaveTime int64) error {
	for _, h := range s.membersHistory {
		if h.UserId == userId && h.ChannelId == channelId {
			h.LeaveTime = &leaveTime
		}
	}

	return nil
}

func (s *MemChannelMemberHistoryStore) GetUsersInChannelDuring(startTime int64, endTime int64, channelId string) ([]*model.ChannelMemberHistoryResult, error) {
	panic("not implemented")
}

func (s *MemChannelMemberHistoryStore) PermanentDeleteBatchForRetentionPolicies(now, globalPolicyEndTime, limit int64, cursor model.RetentionPolicyCursor) (int64, model.RetentionPolicyCursor, error) {
	panic("not implemented")
}

func (s *MemChannelMemberHistoryStore) DeleteOrphanedRows(limit int) (deleted int64, err error) {
	panic("not implemented")
}

func (s *MemChannelMemberHistoryStore) PermanentDeleteBatch(endTime int64, limit int64) (int64, error) {
	panic("not implemented")
}
