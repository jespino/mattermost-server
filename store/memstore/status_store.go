// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package memstore

import (
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/store"
)

type MemStatusStore struct {
	status []*model.Status
}

func newMemStatusStore() store.StatusStore {
	return &MemStatusStore{}
}

func (s *MemStatusStore) SaveOrUpdate(st *model.Status) error {
	for _, item := range s.status {
		if item.UserId == st.UserId {
			*item = *st
			return nil
		}
	}
	s.status = append(s.status, st)
	return nil
}

func (s *MemStatusStore) Get(userId string) (*model.Status, error) {
	for _, item := range s.status {
		if item.UserId == userId {
			return item, nil
		}
	}
	return nil, store.NewErrNotFound("Status", userId)
}

func (s *MemStatusStore) GetByIds(userIds []string) ([]*model.Status, error) {
	results := []*model.Status{}
	for _, item := range s.status {
		for _, id := range userIds {
			if item.UserId == id {
				results = append(results, item)
			}
		}
	}
	return results, nil
}

func (s *MemStatusStore) UpdateExpiredDNDStatuses() ([]*model.Status, error) {
	panic("not implemented")
}

func (s *MemStatusStore) ResetAll() error {
	s.status = []*model.Status{}
	return nil
}

func (s *MemStatusStore) GetTotalActiveUsersCount() (int64, error) {
	panic("not implemented")
}

func (s *MemStatusStore) UpdateLastActivityAt(userId string, lastActivityAt int64) error {
	panic("not implemented")
}
