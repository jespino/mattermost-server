// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package transparentstore_test

import (
	"github.com/mattermost/mattermost-server/store/transparentstore"
	"testing"

	"github.com/mattermost/mattermost-server/testlib"
)

var mainHelper *testlib.MainHelper

func TestMain(m *testing.M) {
	mainHelper = testlib.NewMainHelperWithOptions(nil)
	defer mainHelper.Close()

	transparentstore.InitTest()

	mainHelper.Main(m)
	transparentstore.TearDownTest()
}
