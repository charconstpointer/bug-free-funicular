package arpisi

//State is
type State uint32

//Healthy is
const (
	Healthy State = iota
	Unhealthy
)

func (s State) String() string {
	switch s {
	case Healthy:
		return "Healthy"
	case Unhealthy:
		return "Unhealthy"
	default:
		return "Unknown"
	}
}

//Reply is
type Reply struct {
	Value string
}

//Arpi is
type Arpi struct {
	nodes []*Node
	log   []string
}

//Commit is
func (a *Arpi) Commit(message string, reply *Reply) error {
	*reply = Reply{Value: "hello back"}
	return nil
}

//Node is
type Node struct {
	addr  string
	state State
}
