package actions

type NodeData struct {
	ID            string            `json:"id"`
	Type          string            `json:"type"`
	X             int               `json:"x"`
	Y             int               `json:"y"`
	Inputs        []string          `json:"inputs"`
	Outputs       []string          `json:"outputs"`
	EventName     string            `json:"eventName,omitempty"`
	ActionName    string            `json:"actionName,omitempty"`
	Command       *SlashCommandNode `json:"command,omitempty"`
	Secret        string            `json:"secret,omitempty"`
	ControlType   string            `json:"controlType,omitempty"`
	IfValue       string            `json:"ifValue,omitempty"`
	IfComparison  string            `json:"ifComparison,omitempty"`
	CaseValues    []string          `json:"caseValues,omitempty"`
	RandomOptions int               `json:"randomOptions,omitempty"`
	Cron          string            `json:"cron,omitempty"`
	Seconds       int               `json:"seconds,omitempty"`
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

type BaseNode struct {
	id string
	x  int
	y  int
}

func (n *BaseNode) ID() string {
	return n.id
}

func (n *BaseNode) X() int {
	return n.x
}

func (n *BaseNode) Y() int {
	return n.y
}

func (n *BaseNode) SetPos(x, y int) {
	n.x = x
	n.y = y
}
