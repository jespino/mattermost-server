package actions

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/services/systembus"
)

type Graph struct {
	id    string
	nodes []Node
	edges []*Edge
}

func NewGraph() *Graph {
	return &Graph{
		id: model.NewId(),
	}
}

type Edge struct {
	id     string
	from   Node
	to     Node
	config map[string]string
}

func NewEdge(from Node, to Node, config map[string]string) *Edge {
	return &Edge{
		id:     model.NewId(),
		from:   from,
		to:     to,
		config: config,
	}
}

type Node interface {
	ID() string
	Type() string
	Inputs() int
	Outputs() int
	Run(g *Graph, data map[string]string) error
}

type EventNode struct {
	id        string
	eventName string
}

func NewEventNode(eventName string) *EventNode {
	return &EventNode{
		id:        model.NewId(),
		eventName: eventName,
	}
}

func (e *EventNode) ID() string {
	return e.id
}

func (e *EventNode) Type() string {
	return "event"
}

func (e *EventNode) Inputs() int {
	return 0
}

func (e *EventNode) Outputs() int {
	return 1
}

type ActionNode struct {
	id     string
	action *ActionDefinition
}

func NewActionNode(action *ActionDefinition) *ActionNode {
	return &ActionNode{
		id:     model.NewId(),
		action: action,
	}
}

func (e *ActionNode) ID() string {
	return e.id
}

func (e *ActionNode) Type() string {
	return "action"
}

func (e *ActionNode) Inputs() int {
	return 1
}

func (e *ActionNode) Outputs() int {
	return 1
}

type SlashCommandNode struct {
	id string
}

func (e *SlashCommandNode) ID() string {
	return e.id
}

func (e *SlashCommandNode) Type() string {
	return "slash-command"
}

func (e *SlashCommandNode) Inputs() int {
	return 0
}

func (e *SlashCommandNode) Outputs() int {
	return 1
}

func (g *Graph) RunEvent(event *systembus.Event) {
	fmt.Println("RUNNING GRAPH INSIDE")
	for _, node := range g.nodes {
		fmt.Println("LOOKING FOR EVENT NODES")
		if node.Type() == "event" && node.(*EventNode).eventName == event.ID {
			fmt.Println("EVENT NODE FOUND, RUNNING")
			err := node.(*EventNode).Run(g, event.Data)
			if err != nil {
				fmt.Println("*************** ERROR ****************")
				fmt.Println(err)
				fmt.Println("**************************************")
			}
		}
	}
}

func (g *Graph) getEdgesFrom(id string) []*Edge {
	edges := []*Edge{}
	for _, edge := range g.edges {
		if edge.from.ID() == id {
			edges = append(edges, edge)
		}
	}
	return edges
}

func (g *Graph) AddNode(node Node) {
	g.nodes = append(g.nodes, node)
}

func (g *Graph) AddEdge(edge *Edge) {
	g.edges = append(g.edges, edge)
}

func (n EventNode) Run(g *Graph, data map[string]string) error {
	fmt.Println("EVENT NODE FOUND, RUNNING INSIDE", g.getEdgesFrom(n.ID()), g.edges)
	for _, edge := range g.getEdgesFrom(n.ID()) {
		fmt.Println("APPLYING CONFIG")
		newData, err := edge.ApplyConfig(data)
		if err != nil {
			return err
		}
		fmt.Println("RUNNING ACTION", edge.to, newData)
		err = edge.to.Run(g, newData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (n *ActionNode) Run(g *Graph, data map[string]string) error {
	fmt.Println("RUNNING ACTION INSIDE", n.action.ID, n.action.Name, data)
	result, err := n.action.Handler(data)
	if err != nil {
		return err
	}
	if result == nil {
		return nil
	}
	fmt.Println("ACTION EXECUTED, PASSING RESULT TO EDGES")
	for _, edge := range g.getEdgesFrom(n.ID()) {
		fmt.Println("APPLYING EDGE CONFIG")
		newData, err := edge.ApplyConfig(result)
		if err != nil {
			return err
		}
		fmt.Println("RUNNING ACTION", edge.to, newData)
		err = edge.to.Run(g, newData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Edge) ApplyConfig(data map[string]string) (map[string]string, error) {
	result := map[string]string{}
	for key, value := range data {
		result[key] = value
	}
	for key, value := range e.config {
		tmpl, err := template.New("template").Parse(value)
		if err != nil {
			return nil, err
		}
		var buf bytes.Buffer
		tmpl.Execute(&buf, data)
		result[key] = buf.String()
	}
	return result, nil
}
