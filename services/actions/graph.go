package actions

import (
	"fmt"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/services/systembus"
)

type GraphData struct {
	ID    string     `json:"id"`
	Name  string     `json:"name"`
	Nodes []NodeData `json:"nodes"`
	Edges []EdgeData `json:"edges"`
}

type Graph struct {
	id    string
	name  string
	nodes []Node
	edges []*Edge
}

func (g *Graph) ToGraphData() *GraphData {
	nodes := []NodeData{}
	for _, node := range g.nodes {
		nodeData := NodeData{
			ID:      node.ID(),
			Type:    node.Type(),
			X:       node.X(),
			Y:       node.Y(),
			Inputs:  node.Inputs(),
			Outputs: node.Outputs(),
		}

		switch node.Type() {
		case NodeTypeAction:
			if node.(*ActionNode).action != nil {
				nodeData.ActionName = node.(*ActionNode).action.ID
			}
		case NodeTypeEvent:
			nodeData.EventName = node.(*EventNode).eventName
		case NodeTypeSlashCommand:
			nodeData.Command = node.(*SlashCommandNode)
		case NodeTypeWebhook:
			nodeData.Secret = node.(*WebhookNode).secret
		case NodeTypeFlow:
			nodeData.ControlType = node.(*FlowNode).controlType
			nodeData.IfValue = node.(*FlowNode).ifValue
			nodeData.IfComparison = node.(*FlowNode).ifComparison
			nodeData.CaseValues = node.(*FlowNode).caseValues
			nodeData.RandomOptions = node.(*FlowNode).randomOptions
		}
		nodes = append(nodes, nodeData)
	}

	edges := []EdgeData{}
	for _, edge := range g.edges {
		edges = append(edges, EdgeData{
			ID:         edge.id,
			From:       edge.from.ID(),
			FromOutput: edge.fromOutput,
			To:         edge.to.ID(),
			Config:     edge.config,
		})
	}

	return &GraphData{
		ID:    g.id,
		Name:  g.name,
		Nodes: nodes,
		Edges: edges,
	}
}

func NewGraph(name string) *Graph {
	return &Graph{
		id:   model.NewId(),
		name: name,
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

func (g *Graph) RunEdgesForNode(n Node, data map[string]string) error {
	for _, edge := range g.getEdgesFrom(n.ID()) {
		if edge.fromOutput != "" && edge.fromOutput != data["Output"] {
			fmt.Println("SKIPPING EDGE", edge, data)
			continue
		}

		newData, err := edge.ApplyConfig(data)
		if err != nil {
			return err
		}
		err = edge.to.Run(g, newData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Graph) RunEvent(event *systembus.Event) {
	fmt.Println("RUNNING GRAPH INSIDE")
	for _, node := range g.nodes {
		fmt.Println("LOOKING FOR EVENT NODES")
		if node.Type() == NodeTypeEvent && node.(*EventNode).eventName == event.ID {
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
