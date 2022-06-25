package builtinactions

import (
	"fmt"

	"github.com/mattermost/mattermost-server/v6/app/request"
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/services/actions"
)

const SendDMID = "send-direct-message"

type DMPoster interface {
	GetUserByUsername(string) (*model.User, *model.AppError)
	GetOrCreateDirectChannel(c *request.Context, userID, otherUserID string, channelOptions ...model.ChannelOption) (*model.Channel, *model.AppError)
	CreatePost(*request.Context, *model.Post, *model.Channel, bool, bool) (*model.Post, *model.AppError)
}

func NewSendDM(dmPoster DMPoster, ctx *request.Context) *actions.ActionDefinition {
	sendDMActionHandler := func(data map[string]string) (map[string]string, error) {
		message := data["template"]
		receiverID := data["receiver-id"]
		senderID := data["sender-id"]
		receiverUsername := data["receiver-username"]
		senderUsername := data["sender-username"]

		if receiverID == "" {
			receiver, appErr := dmPoster.GetUserByUsername(receiverUsername)
			if appErr != nil {
				return nil, appErr
			}
			receiverID = receiver.Id
		}

		if senderID == "" {
			sender, appErr := dmPoster.GetUserByUsername(senderUsername)
			if appErr != nil {
				return nil, appErr
			}
			senderID = sender.Id
		}

		channel, appErr := dmPoster.GetOrCreateDirectChannel(ctx, receiverID, senderID)
		if appErr != nil {
			return nil, appErr
		}
		now := model.GetMillis()
		post := model.Post{
			Message:   message,
			UserId:    senderID,
			ChannelId: channel.Id,
			CreateAt:  now,
			UpdateAt:  now,
			Type:      model.PostTypeDefault,
		}
		p, appErr := dmPoster.CreatePost(ctx, &post, channel, false, false)
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

	sendDMAction := actions.ActionDefinition{
		ID:               SendDMID,
		Name:             "Send Direct Message",
		Description:      "Send a directe message to a user",
		ConfigDefinition: map[string]string{"template": "string", "receiver-id": "string", "receiver-username": "string", "sender-id": "string", "sender-username": "string"},
		Handler:          sendDMActionHandler,
	}
	return &sendDMAction
}
