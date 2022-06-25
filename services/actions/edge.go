package actions

import (
	"bytes"
	"text/template"

	"github.com/mattermost/mattermost-server/v6/model"
)

type EdgeData struct {
	ID         string            `json:"id"`
	From       string            `json:"from"`
	FromOutput string            `json:"fromOutput"`
	To         string            `json:"to"`
	Config     map[string]string `json:"config"`
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
