// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package app

import (
	goi18n "github.com/mattermost/go-i18n/i18n"
	"github.com/mattermost/mattermost-server/v5/model"
)

type ShrugProvider struct {
}

const (
	CMD_SHRUG = "shrug"
)

func init() {
	registerCommandProvider(&ShrugProvider{})
}

func (me *ShrugProvider) getTrigger() string {
	return CMD_SHRUG
}

func (me *ShrugProvider) getCommand(a *App, T goi18n.TranslateFunc) *model.Command {
	return &model.Command{
		Trigger:          CMD_SHRUG,
		AutoComplete:     true,
		AutoCompleteDesc: T("api.command_shrug.desc"),
		AutoCompleteHint: T("api.command_shrug.hint"),
		DisplayName:      T("api.command_shrug.name"),
	}
}

func (me *ShrugProvider) doCommand(a *App, args *model.CommandArgs, message string) *model.CommandResponse {
	rmsg := `¯\\\_(ツ)\_/¯`
	if len(message) > 0 {
		rmsg = message + " " + rmsg
	}

	return &model.CommandResponse{ResponseType: model.COMMAND_RESPONSE_TYPE_IN_CHANNEL, Text: rmsg}
}
