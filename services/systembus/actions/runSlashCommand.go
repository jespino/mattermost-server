package actions

import (
	"github.com/mattermost/mattermost-server/v6/app/request"
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/services/systembus"
	"github.com/mattermost/mattermost-server/v6/shared/i18n"
)

const RunSlashCommandID = "create-channel"

type CommandExecutor interface {
	GetChannel(channelId string) (*model.Channel, *model.AppError)
	ExecuteCommand(c *request.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError)
}

func NewRunSlashCommand(commandExecutor CommandExecutor, ctx *request.Context) *systembus.ActionDefinition {
	runSlashCommandActionHandler := func(event *systembus.Event, config map[string]string) (*systembus.Event, error) {
		command, err := applyTemplate(config["command"], event.Data)
		if err != nil {
			return nil, err
		}

		channelID, err := applyTemplate(config["channelId"], event.Data)
		if err != nil {
			return nil, err
		}

		userID, err := applyTemplate(config["userId"], event.Data)
		if err != nil {
			return nil, err
		}

		commandArgs := model.CommandArgs{
			Command:         command,
			UserId:          userID,
			ChannelId:       channelID,
			T:               i18n.GetUserTranslations(""),
			UserMentions:    model.UserMentionMap{},
			ChannelMentions: model.ChannelMentionMap{},
			Session:         model.Session{},
		}
		channel, appErr := commandExecutor.GetChannel(channelID)
		if appErr != nil {
			return nil, appErr
		}
		if channel.Type != model.ChannelTypeDirect && channel.Type != model.ChannelTypeGroup {
			commandArgs.TeamId = channel.TeamId
		}

		// TODO
		// commandArgs.SiteURL = ctx.GetSiteURLHeader()

		_, appErr = commandExecutor.ExecuteCommand(ctx, &commandArgs)
		if appErr != nil {
			return nil, appErr
		}

		return nil, nil
	}

	runSlashCommandAction := systembus.ActionDefinition{
		ID:               RunSlashCommandID,
		Name:             "Run Slash Command",
		Description:      "Run a slash command",
		ConfigDefinition: map[string]string{"command": "string", "userId": "string", "channelId": "string"},
		Handler:          runSlashCommandActionHandler,
	}
	return &runSlashCommandAction
}
