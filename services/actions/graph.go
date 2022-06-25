package actions

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/services/systembus"
)

const NodeTypeAction = "action"
const NodeTypeEvent = "event"
const NodeTypeSlashCommand = "slash-command"
const NodeTypeWebhook = "webhook"

type NodeData struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	X          int               `json:"x"`
	Y          int               `json:"y"`
	Inputs     []string          `json:"inputs"`
	Outputs    []string          `json:"outputs"`
	EventName  string            `json:"eventName"`
	ActionName string            `json:"actionName"`
	Command    *SlashCommandNode `json:"command"`
	Secret     string            `json:"secret"`
}

type EdgeData struct {
	ID         string            `json:"id"`
	From       string            `json:"from"`
	FromOutput string            `json:"fromOutput"`
	To         string            `json:"to"`
	Config     map[string]string `json:"config"`
}

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
			nodeData.ActionName = node.(*ActionNode).action.ID
		case NodeTypeEvent:
			nodeData.EventName = node.(*EventNode).eventName
		case NodeTypeSlashCommand:
			nodeData.Command = node.(*SlashCommandNode)
		case NodeTypeWebhook:
			nodeData.Secret = node.(*WebhookNode).secret
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

type Edge struct {
	id         string
	from       Node
	fromOutput string
	to         Node
	config     map[string]string
}

func NewEdge(from Node, to Node, config map[string]string) *Edge {
	return &Edge{
		id:         model.NewId(),
		from:       from,
		fromOutput: "",
		to:         to,
		config:     config,
	}
}

func (e *Edge) SetFromOutput(fromOutput string) {
	e.fromOutput = fromOutput
}

type Node interface {
	ID() string
	Type() string
	X() int
	Y() int
	SetPos(x int, y int)
	Inputs() []string
	Outputs() []string
	Run(g *Graph, data map[string]string) error
}

type EventNode struct {
	id        string
	x         int
	y         int
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

func (e *EventNode) X() int {
	return e.x
}

func (e *EventNode) Y() int {
	return e.y
}

func (e *EventNode) SetPos(x, y int) {
	e.x = x
	e.y = y
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

type WebhookNode struct {
	id     string
	x      int
	y      int
	secret string
}

func NewWebhookNode(secret string) *WebhookNode {
	return &WebhookNode{
		id:     model.NewId(),
		secret: secret,
	}
}

func (e *WebhookNode) ID() string {
	return e.id
}

func (e *WebhookNode) X() int {
	return e.x
}

func (e *WebhookNode) Y() int {
	return e.y
}

func (e *WebhookNode) SetPos(x, y int) {
	e.x = x
	e.y = y
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

type ActionNode struct {
	id     string
	x      int
	y      int
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

func (e *ActionNode) X() int {
	return e.x
}

func (e *ActionNode) Y() int {
	return e.y
}

func (e *ActionNode) SetPos(x, y int) {
	e.x = x
	e.y = y
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

type SubCommand struct {
	SubCommand  string            `json:"subcommand"`
	Description string            `json:"description"`
	Hint        string            `json:"hint"`
	Name        string            `json:"name"`
	Flags       map[string]string `json:"flags"`
}

type SlashCommandNode struct {
	id          string `json:"id"`
	x           int
	y           int
	Command     string            `json:"command"`
	Description string            `json:"description"`
	Hint        string            `json:"hint"`
	Name        string            `json:"name"`
	SubCommands []SubCommand      `json:"subCommands"`
	Flags       map[string]string `json:"flags"`
}

func NewSlashCommandNode(command string, description string, hint string, name string) *SlashCommandNode {
	return &SlashCommandNode{
		id:          model.NewId(),
		Command:     command,
		Flags:       map[string]string{},
		SubCommands: []SubCommand{},
		Description: description,
		Hint:        hint,
		Name:        name,
	}
}

func (s *SlashCommandNode) AddFlag(flagName string, flagType string) {
	s.Flags[flagName] = flagType
}

func (s *SlashCommandNode) AddSubCommand(subcommand SubCommand) {
	s.SubCommands = append(s.SubCommands, subcommand)
}

func (s *SlashCommandNode) GetCommand() *model.Command {
	autocomplete := model.NewAutocompleteData(s.Command, s.Hint, s.Description)

	for _, subcommand := range s.SubCommands {
		subAutocompleteData := model.NewAutocompleteData(subcommand.SubCommand, "", "")
		for flagName, flagType := range subcommand.Flags {
			if flagType == "boolean" {
				subAutocompleteData.AddNamedTextArgument(flagName, "", "", "Y|N|y|n", false)
			}
			if flagType == "text" {
				subAutocompleteData.AddNamedTextArgument(flagName, "", "", "", false)
			}
		}
		autocomplete.AddCommand(subAutocompleteData)
	}

	for flagName, flagType := range s.Flags {
		if flagType == "boolean" {
			autocomplete.AddNamedTextArgument(flagName, "", "", "Y|N|y|n", false)
		}
		if flagType == "text" {
			autocomplete.AddNamedTextArgument(flagName, "", "", "", false)
		}
	}
	return &model.Command{
		Id:               s.id,
		DeleteAt:         0,
		Trigger:          s.Command,
		AutoComplete:     true,
		AutoCompleteDesc: s.Description,
		AutoCompleteHint: s.Hint,
		DisplayName:      s.Name,
		AutocompleteData: autocomplete,
	}
}

func (s *SlashCommandNode) DoCommand(g *Graph, args *model.CommandArgs, message string) *model.CommandResponse {
	trimSpaceAndQuotes := func(s string) string {
		trimmed := strings.TrimSpace(s)
		trimmed = strings.TrimPrefix(trimmed, "\"")
		trimmed = strings.TrimPrefix(trimmed, "'")
		trimmed = strings.TrimSuffix(trimmed, "\"")
		trimmed = strings.TrimSuffix(trimmed, "'")
		return trimmed
	}

	parseNamedArgs := func(cmd string) map[string]string {
		m := make(map[string]string)

		split := strings.Fields(cmd)

		// check for optional action
		if len(split) >= 2 && !strings.HasPrefix(split[1], "--") {
			m["-action"] = split[1] // prefix with hyphen to avoid collision with arg named "action"
		}

		for i := 0; i < len(split); i++ {
			if !strings.HasPrefix(split[i], "--") {
				continue
			}
			var val string
			arg := trimSpaceAndQuotes(strings.Trim(split[i], "-"))
			if i < len(split)-1 && !strings.HasPrefix(split[i+1], "--") {
				val = trimSpaceAndQuotes(split[i+1])
			}
			if arg != "" {
				m[arg] = val
			}
		}
		return m
	}

	margs := parseNamedArgs(args.Command)
	action, ok := margs["-action"]

	data := margs
	data["UserId"] = args.UserId
	data["ChannelId"] = args.ChannelId
	data["TeamId"] = args.TeamId
	data["RootId"] = args.RootId
	data["ParentId"] = args.ParentId
	data["TriggerId"] = args.TriggerId

	if !ok {
		s.Run(g, data)
		data["Output"] = "0"
		return &model.CommandResponse{}
	}

	for idx, subCommand := range s.SubCommands {
		if action == subCommand.SubCommand {
			data["Output"] = fmt.Sprintf("%d", idx+1)
			s.Run(g, data)
			return &model.CommandResponse{}
		}
	}

	return &model.CommandResponse{
		ResponseType: model.CommandResponseTypeEphemeral,
		Text:         fmt.Sprintf("Command action %s not found.", action),
		Type:         model.PostTypeDefault,
	}
}

func (s *SlashCommandNode) ID() string {
	return s.id
}

func (s *SlashCommandNode) X() int {
	return s.x
}

func (s *SlashCommandNode) Y() int {
	return s.y
}

func (s *SlashCommandNode) SetPos(x, y int) {
	s.x = x
	s.y = y
}

func (s *SlashCommandNode) Type() string {
	return NodeTypeSlashCommand
}

func (s *SlashCommandNode) Inputs() []string {
	return []string{}
}

func (s *SlashCommandNode) Outputs() []string {
	result := []string{"main"}
	for _, command := range s.SubCommands {
		result = append(result, fmt.Sprintf("subcommand:%s", command.SubCommand))
	}

	return result
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

func (e *EventNode) Run(g *Graph, data map[string]string) error {
	fmt.Println("EVENT NODE FOUND, RUNNING INSIDE", g.getEdgesFrom(e.ID()), g.edges)
	g.RunEdgesForNode(e, data)
	return nil
}

func (e *WebhookNode) Run(g *Graph, data map[string]string) error {
	fmt.Println("EVENT NODE FOUND, RUNNING INSIDE", g.getEdgesFrom(e.ID()), g.edges)
	g.RunEdgesForNode(e, data)
	return nil
}

func (s *SlashCommandNode) Run(g *Graph, data map[string]string) error {
	fmt.Println("SLASH COMMAND NODE FOUND, RUNNING INSIDE", g.getEdgesFrom(s.ID()), g.edges)
	g.RunEdgesForNode(s, data)
	return nil
}

func (a *ActionNode) Run(g *Graph, data map[string]string) error {
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
