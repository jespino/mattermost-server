package actions

import (
	"github.com/mattermost/mattermost-server/v6/app/request"
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/services/systembus"
)

const PostMessageID = "post-message"

type MessagePoster interface {
	GetChannel(string) (*model.Channel, *model.AppError)
	CreatePost(*request.Context, *model.Post, *model.Channel, bool, bool) (*model.Post, *model.AppError)
}

func NewPostMessage(messagePoster MessagePoster, ctx *request.Context) *systembus.ActionDefinition {
	postActionHandler := func(event *systembus.Event, config map[string]string) (*systembus.Event, error) {
		message, err := applyTemplate(config["template"], event.Data)
		if err != nil {
			return nil, err
		}

		channelID, err := applyTemplate(config["channel-id"], event.Data)
		if err != nil {
			return nil, err
		}

		userID, err := applyTemplate(config["user-id"], event.Data)
		if err != nil {
			return nil, err
		}

		channel, appErr := messagePoster.GetChannel(channelID)
		if appErr != nil {
			return nil, appErr
		}
		now := model.GetMillis()
		post := model.Post{
			Message:   message,
			UserId:    userID,
			ChannelId: channelID,
			CreateAt:  now,
			UpdateAt:  now,
			Type:      model.PostTypeDefault,
		}
		_, appErr = messagePoster.CreatePost(ctx, &post, channel, false, false)
		if appErr != nil {
			return nil, appErr
		}

		return nil, nil
	}

	postMessageAction := systembus.ActionDefinition{
		ID:               PostMessageID,
		Name:             "Post Message",
		Description:      "Post a message into a channel",
		ConfigDefinition: map[string]string{"template": "string", "channel-id": "string", "user-id": "string"},
		Handler:          postActionHandler,
	}
	return &postMessageAction
}
