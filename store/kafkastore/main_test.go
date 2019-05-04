// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package kafkastore_test

import (
	"github.com/mattermost/mattermost-server/store/kafkastore"
	"testing"

	"github.com/mattermost/mattermost-server/testlib"
)

var mainHelper *testlib.MainHelper

func TestMain(m *testing.M) {
	mainHelper = testlib.NewMainHelperWithOptions(nil)
	defer mainHelper.Close()

	kafkastore.InitTest()

	mainHelper.Main(m)
	kafkastore.TearDownTest()
}
