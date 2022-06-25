package actions

import (
	"errors"
	"fmt"

	"github.com/mattermost/mattermost-server/v6/model"
)

const NodeTypeAction = "action"

type ActionNode struct {
	BaseNode
	action *ActionDefinition
}

func NewActionNode(action *ActionDefinition) *ActionNode {
	return &ActionNode{
		BaseNode: BaseNode{id: model.NewId()},
		action:   action,
	}
}

func (e *ActionNode) Type() string {
	return NodeTypeAction
}

func (e *ActionNode) Inputs() []string {
	return []string{"in"}
}

func (e *ActionNode) Outputs() []string {
	return []string{"out"}
}

func (a *ActionNode) Run(g *Graph, data map[string]string) error {
	if a.action == nil {
		return errors.New("Action handler doesn't exists")
	}
	fmt.Println("RUNNING ACTION INSIDE", a.action.ID, a.action.Name, data)
	result, err := a.action.Handler(data)
	if err != nil {
		return err
	}
	if result == nil {
		return nil
	}
	fmt.Println("ACTION EXECUTED, PASSING RESULT TO EDGES")
	g.RunEdgesForNode(a, data)
	return nil
}
