package actions

import (
	"github.com/mattermost/mattermost-server/v6/model"
)

const NodeTypeSchedTypeCron = "cron"
const NodeTypeSchedTypeInterval = "interval"
const NodeTypeSched = "sched"

type SchedNode struct {
	BaseNode
	controlType string
	seconds     int
	cron        string
}

func NewSchedCronNode(cron string) *SchedNode {
	return &SchedNode{
		BaseNode:    BaseNode{id: model.NewId()},
		controlType: NodeTypeSchedTypeCron,
		cron:        cron,
	}
}

func NewSchedIntervalNode(seconds int) *SchedNode {
	return &SchedNode{
		BaseNode:    BaseNode{id: model.NewId()},
		controlType: NodeTypeSchedTypeInterval,
		seconds:     seconds,
	}
}

func (e *SchedNode) Type() string {
	return NodeTypeSched
}

func (e *SchedNode) Inputs() []string {
	return []string{""}
}

func (e *SchedNode) Outputs() []string {
	return []string{"out"}
}

func (e *SchedNode) Run(g *Graph, data map[string]string) error {
	return g.RunEdgesForNode(e, data)
}
