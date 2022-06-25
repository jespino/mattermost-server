package builtinactions

import (
	"fmt"

	"github.com/mattermost/mattermost-server/v6/app/request"
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/services/actions"
)

const CreateChannelID = "create-channel"

type ChannelCreator interface {
	GetUserByUsername(string) (*model.User, *model.AppError)
	GetTeamByName(string) (*model.Team, *model.AppError)
	CreateChannelWithUser(c *request.Context, channel *model.Channel, userID string) (*model.Channel, *model.AppError)
}

func NewCreateChannel(channelCreator ChannelCreator, ctx *request.Context) *actions.ActionDefinition {
	createChannelActionHandler := func(data map[string]string) (map[string]string, error) {
		now := model.GetMillis()
		creatorID := data["creator-id"]
		teamID := data["team-id"]
		creatorUsername := data["creator-username"]
		teamName := data["team-name"]

		if creatorID == "" {
			creator, appErr := channelCreator.GetUserByUsername(creatorUsername)
			if appErr != nil {
				return nil, appErr
			}
			creatorID = creator.Id
		}

		if teamID == "" {
			team, appErr := channelCreator.GetTeamByName(teamName)
			if appErr != nil {
				return nil, appErr
			}
			teamID = team.Id
		}

		channel := model.Channel{
			Name:        data["name"],
			DisplayName: data["display-name"],
			TeamId:      teamID,
			CreatorId:   creatorID,
			CreateAt:    now,
			UpdateAt:    now,
			Type:        model.ChannelType(data["type"]),
		}
		c, appErr := channelCreator.CreateChannelWithUser(ctx, &channel, data["creator-id"])
		if appErr != nil {
			return nil, appErr
		}

		newData := map[string]string{}
		for key, value := range data {
			newData[key] = value
		}

		newData["id"] = c.Id
		newData["create-at"] = fmt.Sprintf("%d", c.CreateAt)
		newData["update-at"] = fmt.Sprintf("%d", c.UpdateAt)

		return newData, nil
	}

	createChannelAction := actions.ActionDefinition{
		ID:               CreateChannelID,
		Name:             "Create channel",
		Description:      "Create a new channel in a team",
		ConfigDefinition: map[string]string{"name": "string", "display-name": "string", "team-name": "string", "team-id": "string", "creator-id": "string", "creator-username": "string", "type": "string"},
		Handler:          createChannelActionHandler,
	}
	return &createChannelAction
}
