// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package memstore

import (
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/store"
)

type MemProductNoticesStore struct {
	productNoticies []*model.ProductNotice
}

func newMemProductNoticesStore() store.ProductNoticesStore {
	return &MemProductNoticesStore{}
}

func (s *MemProductNoticesStore) Clear(notices []string) error {
	panic("not implemented")
}

func (s *MemProductNoticesStore) ClearOldNotices(currentNotices model.ProductNotices) error {
	panic("not implemented")
}

func (s *MemProductNoticesStore) View(userId string, notices []string) error {
	// TODO: Implement this
	return nil
}

func (s *MemProductNoticesStore) GetViews(userId string) ([]model.ProductNoticeViewState, error) {
	panic("not implemented")
}
