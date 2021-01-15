// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package file_content_extract

import (
	"github.com/mattermost/mattermost-server/v5/app"
	tjobs "github.com/mattermost/mattermost-server/v5/jobs/interfaces"
)

type FileContentExtractJobInterfaceImpl struct {
	App *app.App
}

func init() {
	app.RegisterFileContentExtractJobInterface(func(a *app.App) tjobs.FileContentExtractJobInterface {
		return &FileContentExtractJobInterfaceImpl{a}
	})
}
