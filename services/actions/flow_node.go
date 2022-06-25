package actions

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/mattermost/mattermost-server/v6/model"
)

const NodeTypeFlowTypeIf = "if"
const NodeTypeFlowTypeSwitch = "switch"
const NodeTypeFlowTypeRandom = "random"
const NodeTypeFlow = "flow"

type FlowNode struct {
	BaseNode
	controlType   string
	ifValue       string
	ifComparison  string
	caseValues    []string
	randomOptions int
}

func NewFlowIfNode(controlType string, ifValue string, ifComparison string) *FlowNode {
	return &FlowNode{
		BaseNode:     BaseNode{id: model.NewId()},
		controlType:  controlType,
		ifValue:      ifValue,
		ifComparison: ifComparison,
	}
}

func NewFlowSwitchNode(controlType string, caseValues []string) *FlowNode {
	return &FlowNode{
		BaseNode:    BaseNode{id: model.NewId()},
		controlType: controlType,
		caseValues:  caseValues,
	}
}

func NewFlowRandomNode(controlType string, options int) *FlowNode {
	return &FlowNode{
		BaseNode:      BaseNode{id: model.NewId()},
		controlType:   controlType,
		randomOptions: options,
	}
}

func (e *FlowNode) Type() string {
	return NodeTypeFlow
}

func (e *FlowNode) Inputs() []string {
	return []string{"in"}
}

func (e *FlowNode) Outputs() []string {
	switch e.controlType {
	case NodeTypeFlowTypeIf:
		return []string{"then", "else"}
	case NodeTypeFlowTypeSwitch:
		return e.caseValues
	case NodeTypeFlowTypeRandom:
		result := []string{}
		for x := 0; x < e.randomOptions; x++ {
			result = append(result, fmt.Sprintf("out-%d", x))
		}
		return result
	}
	return []string{}
}

func (e *FlowNode) Run(g *Graph, data map[string]string) error {
	fmt.Println("FLOW NODE FOUND, RUNNING INSIDE", g.getEdgesFrom(e.ID()), g.edges)

	switch e.controlType {
	case NodeTypeFlowTypeIf:
		fmt.Println("RUNNING HERE -3")
		switch e.ifComparison {
		case "":
			if data["value"] == e.ifValue {
				data["Output"] = "then"
			} else {
				data["Output"] = "else"
			}
		case "eq":
			if data["value"] == e.ifValue {
				data["Output"] = "then"
			} else {
				data["Output"] = "else"
			}
		case "gt":
			if data["value"] > e.ifValue {
				data["Output"] = "then"
			} else {
				data["Output"] = "else"
			}
		case "lt":
			if data["value"] < e.ifValue {
				data["Output"] = "then"
			} else {
				data["Output"] = "else"
			}
		case "gte":
			if data["value"] >= e.ifValue {
				data["Output"] = "then"
			} else {
				data["Output"] = "else"
			}
		case "lte":
			if data["value"] <= e.ifValue {
				data["Output"] = "then"
			} else {
				data["Output"] = "else"
			}
		case "contains":
			if strings.Contains(data["value"], e.ifValue) {
				data["Output"] = "then"
			} else {
				data["Output"] = "else"
			}
		case "prefix":
			if strings.HasPrefix(data["value"], e.ifValue) {
				data["Output"] = "then"
			} else {
				data["Output"] = "else"
			}
		case "suffix":
			if strings.HasSuffix(data["value"], e.ifValue) {
				data["Output"] = "then"
			} else {
				data["Output"] = "else"
			}
		}
	case NodeTypeFlowTypeSwitch:
		fmt.Println("RUNNING HERE -2")
		for _, caseValue := range e.caseValues {
			if data["value"] == caseValue {
				data["Output"] = caseValue
				break
			}
		}
	case NodeTypeFlowTypeRandom:
		fmt.Println("RUNNING HERE -1")
		i := rand.Intn(e.randomOptions)
		data["Output"] = fmt.Sprintf("out-%d", i)
	}
	fmt.Println("RUNNING HERE 1")
	g.RunEdgesForNode(e, data)
	fmt.Println("RUNNING HERE 2")
	return nil
}
