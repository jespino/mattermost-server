package memstore

import (
	"testing"

	"github.com/mattermost/mattermost-server/v6/store/storetest"
)

func TestSessionStore(t *testing.T) {
	storetest.TestSessionStore(t, New())
}
