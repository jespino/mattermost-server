package actions

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-server/v6/model"
)

const NodeTypeSlashCommand = "slash-command"

type SlashCommandNode struct {
	BaseNode
	Command     string            `json:"command"`
	Description string            `json:"description"`
	Hint        string            `json:"hint"`
	Name        string            `json:"name"`
	SubCommands []SubCommand      `json:"subCommands"`
	Flags       map[string]string `json:"flags"`
}

type SubCommand struct {
	SubCommand  string            `json:"subcommand"`
	Description string            `json:"description"`
	Hint        string            `json:"hint"`
	Name        string            `json:"name"`
	Flags       map[string]string `json:"flags"`
}

func NewSlashCommandNode(command string, description string, hint string, name string) *SlashCommandNode {
	return &SlashCommandNode{
		BaseNode:    BaseNode{id: model.NewId()},
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

func (s *SlashCommandNode) Run(g *Graph, data map[string]string) error {
	fmt.Println("SLASH COMMAND NODE FOUND, RUNNING INSIDE", g.getEdgesFrom(s.ID()), g.edges)
	g.RunEdgesForNode(s, data)
	return nil
}
