// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package memstore

import (
	"testing"

	"github.com/mattermost/mattermost-server/v6/store/storetest"
)

func TestUserStore(t *testing.T) {
	storetest.TestUserStore(t, New(), nil)
}
