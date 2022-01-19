package memstore

import (
	"testing"

	"github.com/mattermost/mattermost-server/v6/store/storetest"
)

func TestOAuthStore(t *testing.T) {
	storetest.TestOAuthStore(t, New())
}
