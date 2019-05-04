// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package kafkastore

import (
	"testing"

	"github.com/mattermost/mattermost-server/store/storetest"
)

func TestUserStore(t *testing.T) {
	StoreTest(t, storetest.TestUserStore)
}
