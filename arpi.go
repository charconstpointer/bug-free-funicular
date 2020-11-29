package arpisi

type State uint32

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

type Arpi struct {
	nodes []*Node
	log   []string
}

func (a *Arpi) Commit(message string) error {
	return nil
}

type Node struct {
	addr  string
	state State
}
