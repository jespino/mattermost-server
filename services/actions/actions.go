package actions

import (
	"errors"
	"fmt"
	"sync"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/services/systembus"
	"github.com/mattermost/mattermost-server/v6/shared/mlog"
)

type ActionDefinition struct {
	ID               string                                                  `json:"id"`
	Name             string                                                  `json:"name"`
	Description      string                                                  `json:"description"`
	ConfigDefinition map[string]string                                       `json:"config_definition"`
	Handler          func(data map[string]string) (map[string]string, error) `json:"-"`
}

type Actions struct {
	systemBus         systembus.SystemBus
	registerCommand   func(*model.Command, func(*model.CommandArgs, string) *model.CommandResponse)
	unregisterCommand func(string)
	actions           map[string]*ActionDefinition
	graphs            map[string]*Graph
	graphsByEvent     map[string]map[string]*Graph
	graphsByHook      map[string]*Graph
	mutex             sync.RWMutex
	logger            *mlog.Logger
}

func EventToGraphSubscription(a *Actions) func(event *systembus.Event) (*systembus.Event, error) {
	return func(event *systembus.Event) (*systembus.Event, error) {
		a.mutex.RLock()
		defer a.mutex.RUnlock()

		fmt.Println("RUNNING EVENT", a.graphs, a.graphsByEvent, event.ID)
		graphs, ok := a.graphsByEvent[event.ID]
		if !ok {
			fmt.Println("NO GRAPH FOUND")
			return event, nil
		}

		fmt.Println("GRAPHS FOUND")
		for _, graph := range graphs {
			fmt.Println("RUNNING GRAPH")
			go graph.RunEvent(event)
		}
		return nil, nil
	}
}

func New(logger *mlog.Logger, systemBus systembus.SystemBus, registerCommand func(*model.Command, func(*model.CommandArgs, string) *model.CommandResponse), unregisterCommand func(string)) *Actions {
	actions := &Actions{
		systemBus:         systemBus,
		registerCommand:   registerCommand,
		unregisterCommand: unregisterCommand,
		actions:           map[string]*ActionDefinition{},
		graphs:            map[string]*Graph{},
		graphsByEvent:     map[string]map[string]*Graph{},
		graphsByHook:      map[string]*Graph{},
		logger:            logger,
	}
	systemBus.Subscribe(EventToGraphSubscription(actions))
	return actions
}

func (a *Actions) RegisterAction(action *ActionDefinition) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.actions[action.ID] = action
	return nil
}

func (a *Actions) GetAction(actionID string) *ActionDefinition {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	return a.actions[actionID]
}

func (a *Actions) ListActions() ([]*ActionDefinition, error) {
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	actions := []*ActionDefinition{}
	for _, action := range a.actions {
		actions = append(actions, action)
	}
	return actions, nil
}

func (a *Actions) ListGraphs() ([]*Graph, error) {
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	graphs := []*Graph{}
	for _, graph := range a.graphs {
		graphs = append(graphs, graph)
	}
	return graphs, nil
}

func (a *Actions) Run(actionID string, data map[string]string) (map[string]string, error) {
	action, ok := a.actions[actionID]
	if !ok {
		return nil, fmt.Errorf("Action %s not found", actionID)
	}
	return action.Handler(data)
}

func (a *Actions) RunHook(hookID string, data map[string]string) error {
	graph, ok := a.graphsByHook[hookID]
	if !ok {
		fmt.Println("NO GRAPH FOUND")
		return errors.New("hook not found")
	}

	fmt.Println("RUNNING GRAPH")
	for _, node := range graph.nodes {
		fmt.Println("LOOKING FOR EVENT NODES")
		if node.Type() == NodeTypeWebhook && node.ID() == hookID {
			fmt.Println("EVENT NODE FOUND, RUNNING")
			if node.(*WebhookNode).secret != "" && data["secret"] != node.(*WebhookNode).secret {
				return errors.New("invalid secret")
			}

			err := node.(*WebhookNode).Run(graph, data)
			if err != nil {
				fmt.Println("*************** ERROR ****************")
				fmt.Println(err)
				fmt.Println("**************************************")
				return err
			}
		}
	}
	return nil
}

func (a *Actions) AddGraphData(g *GraphData) {
	nodes := []Node{}
	nodesById := map[string]Node{}
	for _, nodeData := range g.Nodes {
		var node Node
		switch nodeData.Type {
		case NodeTypeAction:
			node = NewActionNode(a.GetAction(nodeData.ActionName))
			node.(*ActionNode).id = nodeData.ID
		case NodeTypeEvent:
			node = NewEventNode(nodeData.EventName)
			node.(*EventNode).id = nodeData.ID
		case NodeTypeWebhook:
			node = NewWebhookNode(nodeData.Secret)
			node.(*WebhookNode).id = nodeData.ID
		case NodeTypeSlashCommand:
			node = nodeData.Command
			node.(*SlashCommandNode).id = nodeData.ID
		}
		nodes = append(nodes, node)
		nodesById[node.ID()] = node
	}

	edges := []*Edge{}
	for _, edgeData := range g.Edges {
		edge := NewEdge(nodesById[edgeData.From], nodesById[edgeData.To], edgeData.Config)
		edge.id = edgeData.ID
		edge.fromOutput = edgeData.FromOutput
		edges = append(edges, edge)
	}

	a.AddGraph(&Graph{
		id:    g.ID,
		name:  g.Name,
		nodes: nodes,
		edges: edges,
	})
}

func (a *Actions) AddGraph(graph *Graph) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	a.graphs[graph.id] = graph
	for _, node := range graph.nodes {
		if node.Type() == NodeTypeEvent {
			if _, ok := a.graphsByEvent[node.(*EventNode).eventName]; ok {
				a.graphsByEvent[node.(*EventNode).eventName][graph.id] = graph
			} else {
				a.graphsByEvent[node.(*EventNode).eventName] = map[string]*Graph{graph.id: graph}
			}
		} else if node.Type() == NodeTypeWebhook {
			a.graphsByHook[node.ID()] = graph
		} else if node.Type() == NodeTypeSlashCommand {
			realDoCommand := node.(*SlashCommandNode).DoCommand
			doCommand := func(args *model.CommandArgs, message string) *model.CommandResponse {
				return realDoCommand(graph, args, message)
			}
			a.registerCommand(node.(*SlashCommandNode).GetCommand(), doCommand)
		}
	}
}

func (a *Actions) DeleteGraph(graphID string) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	graph := a.graphs[graphID]
	for _, node := range graph.nodes {
		if node.Type() == "event" {
			delete(a.graphsByEvent[node.(*EventNode).eventName], graph.id)
		} else if node.Type() == "slash-command" {
			a.unregisterCommand(node.(*SlashCommandNode).Command)
		}
	}
	delete(a.graphs, graph.id)
	return nil
}

func (a *Actions) Start() error {
	return nil
}

func (a *Actions) Shutdown() error {
	return nil
}
