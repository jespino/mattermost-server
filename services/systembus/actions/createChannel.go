package actions

import (
	"github.com/mattermost/mattermost-server/v6/app/request"
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/services/systembus"
)

const CreateChannelID = "create-channel"

type ChannelCreator interface {
	CreateChannelWithUser(c *request.Context, channel *model.Channel, userID string) (*model.Channel, *model.AppError)
}

func NewCreateChannel(channelCreator ChannelCreator, ctx *request.Context) *systembus.ActionDefinition {
	createChannelActionHandler := func(event *systembus.Event, config map[string]string) (*systembus.Event, error) {
		channelName, err := applyTemplate(config["name"], event.Data)
		if err != nil {
			return nil, err
		}

		channelDisplayName, err := applyTemplate(config["display-name"], event.Data)
		if err != nil {
			return nil, err
		}

		teamID, err := applyTemplate(config["team-id"], event.Data)
		if err != nil {
			return nil, err
		}

		creatorID, err := applyTemplate(config["creator-id"], event.Data)
		if err != nil {
			return nil, err
		}

		channelType, err := applyTemplate(config["type"], event.Data)
		if err != nil {
			return nil, err
		}

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

	createChannelAction := systembus.ActionDefinition{
		ID:               CreateChannelID,
		Name:             "Create channel",
		Description:      "Create a new channel in a team",
		ConfigDefinition: map[string]string{"name": "string", "display-name": "string", "team-id": "string", "creator-id": "string", "type": "string"},
		Handler:          createChannelActionHandler,
	}
	return &createChannelAction
}
