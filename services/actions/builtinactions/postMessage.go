package builtinactions

import (
	"github.com/mattermost/mattermost-server/v6/app/request"
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/services/actions"
)

const PostMessageID = "post-message"

type MessagePoster interface {
	GetChannel(string) (*model.Channel, *model.AppError)
	CreatePost(*request.Context, *model.Post, *model.Channel, bool, bool) (*model.Post, *model.AppError)
}

func NewPostMessage(messagePoster MessagePoster, ctx *request.Context) *actions.ActionDefinition {
	postActionHandler := func(config map[string]string, data map[string]string) (map[string]string, error) {
		message := config["template"]
		channelID := config["channel-id"]
		userID := config["user-id"]

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

	postMessageAction := actions.ActionDefinition{
		ID:               PostMessageID,
		Name:             "Post Message",
		Description:      "Post a message into a channel",
		ConfigDefinition: map[string]string{"template": "string", "channel-id": "string", "user-id": "string"},
		Handler:          postActionHandler,
	}
	return &postMessageAction
}
