package builtinactions

import (
	"fmt"

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
	postActionHandler := func(data map[string]string) (map[string]string, error) {
		message := data["template"]
		channelID := data["channel-id"]
		userID := data["user-id"]

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
		p, appErr := messagePoster.CreatePost(ctx, &post, channel, false, false)
		if appErr != nil {
			return nil, appErr
		}

		result := map[string]string{}
		for key, value := range data {
			result[key] = value
		}
		result["id"] = p.Id
		result["create-at"] = fmt.Sprintf("%d", p.CreateAt)
		result["update-at"] = fmt.Sprintf("%d", p.UpdateAt)

		return result, nil
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
