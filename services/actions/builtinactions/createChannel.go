package builtinactions

import (
	"github.com/mattermost/mattermost-server/v6/app/request"
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/services/actions"
)

const CreateChannelID = "create-channel"

type ChannelCreator interface {
	CreateChannelWithUser(c *request.Context, channel *model.Channel, userID string) (*model.Channel, *model.AppError)
}

func NewCreateChannel(channelCreator ChannelCreator, ctx *request.Context) *actions.ActionDefinition {
	createChannelActionHandler := func(config map[string]string, data map[string]string) (map[string]string, error) {
		channelName := config["name"]
		channelDisplayName := config["display-name"]
		teamID := config["team-id"]
		creatorID := config["creator-id"]
		channelType := config["type"]

		now := model.GetMillis()
		channel := model.Channel{
			Name:        channelName,
			DisplayName: channelDisplayName,
			TeamId:      teamID,
			CreatorId:   creatorID,
			CreateAt:    now,
			UpdateAt:    now,
			Type:        model.ChannelType(channelType),
		}
		_, appErr := channelCreator.CreateChannelWithUser(ctx, &channel, creatorID)
		if appErr != nil {
			return nil, appErr
		}

		return nil, nil
	}

	createChannelAction := actions.ActionDefinition{
		ID:               CreateChannelID,
		Name:             "Create channel",
		Description:      "Create a new channel in a team",
		ConfigDefinition: map[string]string{"name": "string", "display-name": "string", "team-id": "string", "creator-id": "string", "type": "string"},
		Handler:          createChannelActionHandler,
	}
	return &createChannelAction
}
