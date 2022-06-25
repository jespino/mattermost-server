package actions

import (
	"fmt"

	"github.com/mattermost/mattermost-server/v6/model"
)

const NodeTypeEvent = "event"

type EventNode struct {
	BaseNode
	eventName string
}

func NewEventNode(eventName string) *EventNode {
	return &EventNode{
		BaseNode:  BaseNode{id: model.NewId()},
		eventName: eventName,
	}
}

func (e *EventNode) Type() string {
	return NodeTypeEvent
}

func (e *EventNode) Inputs() []string {
	return []string{}
}

func (e *EventNode) Outputs() []string {
	return []string{"out"}
}

func (e *EventNode) Run(g *Graph, data map[string]string) error {
	fmt.Println("EVENT NODE FOUND, RUNNING INSIDE", g.getEdgesFrom(e.ID()), g.edges)
	g.RunEdgesForNode(e, data)
	return nil
}
