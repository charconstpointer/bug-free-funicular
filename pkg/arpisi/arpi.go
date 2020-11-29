package arpisi

import (
	"log"
	"strings"
)

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
	Log   []string
}

//Commit is
func (a *Arpi) Commit(message string, reply *Reply) error {
	log.Println("received new commit", message)
	a.Log = append(a.Log, message)
	*reply = Reply{Value: strings.ToUpper(message)}
	return nil
}

//Node is
type Node struct {
	addr  string
	state State
}
