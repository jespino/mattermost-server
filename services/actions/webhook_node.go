package actions

import (
	"fmt"

	"github.com/mattermost/mattermost-server/v6/model"
)

const NodeTypeWebhook = "webhook"

func (e *WebhookNode) Run(g *Graph, data map[string]string) error {
	fmt.Println("EVENT NODE FOUND, RUNNING INSIDE", g.getEdgesFrom(e.ID()), g.edges)
	g.RunEdgesForNode(e, data)
	return nil
}

type WebhookNode struct {
	BaseNode
	secret string
}

func NewWebhookNode(secret string) *WebhookNode {
	return &WebhookNode{
		BaseNode: BaseNode{id: model.NewId()},
		secret:   secret,
	}
}

func (e *WebhookNode) Type() string {
	return NodeTypeWebhook
}

func (e *WebhookNode) Inputs() []string {
	return []string{}
}

func (e *WebhookNode) Outputs() []string {
	return []string{"out"}
}
