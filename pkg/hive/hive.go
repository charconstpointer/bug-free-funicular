package hive

import (
	"log"
	"sync"
)

//Hive is
type Hive struct {
	nodes []*Node
	log   []Command

	rpcCh         chan Command
	routinesGroup sync.WaitGroup
	port          int
	transport     Transport
}

//NewHive creates new hive duh
func NewHive(nodes []*Node, port int) *Hive {
	rpcCh := make(chan Command)
	hive := &Hive{
		nodes: nodes,
		log:   make([]Command, 0),

		rpcCh:     rpcCh,
		port:      port,
		transport: NewRPCTransport(port, rpcCh),
	}
	hive.goFunc(hive.run)
	return hive
}

//Commit save value to the log and propagates to other nodes in a cluster
func (h *Hive) Commit(command Command) error {
	h.log = append(h.log, command)
	for _, node := range h.nodes {
		h.transport.Call(node, "RPC.Commit", command)
	}
	// h.transport.Commit(command)
	return nil
}

func (h *Hive) Nodes() []*Node {
	return h.nodes
}

func (h *Hive) run() {
	for {
		select {
		case r := <-h.rpcCh:
			h.log = append(h.log, r)
			log.Println("current log", h.log)
		}
	}
}

func (h *Hive) goFunc(f func()) {
	h.routinesGroup.Add(1)
	go func() {
		defer h.routinesGroup.Done()
		f()
	}()
}
