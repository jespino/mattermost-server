package actions

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
