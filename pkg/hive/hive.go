package hive

import (
	"log"
	"net/rpc"
	"sync"
)

type Hive struct {
	nodes         []*Node
	log           []string
	commitCh      chan string
	rpcCh         chan string
	routinesGroup sync.WaitGroup
	port          int
	transport     Transport
}

func NewHive(nodes []*Node, port int) *Hive {
	hive := &Hive{
		nodes:     nodes,
		log:       make([]string, 0),
		commitCh:  make(chan string),
		rpcCh:     make(chan string),
		port:      port,
		transport: NewRPCTransport(port),
	}
	hive.goFunc(hive.run)
	return hive
}

type Node struct {
	addr string
}

func NewNode(addr string) *Node {
	return &Node{addr}
}

func (h *Hive) run() {
	for {
		select {
		case c := <-h.commitCh:
			for _, node := range h.nodes {
				client, err := rpc.DialHTTP("tcp", node.addr)
				if err != nil {
					log.Println("Connection error: ", err)
					continue
				}
				var reply Reply
				err = client.Call("RPC.Commit", Command{Value: c}, &reply)
				if err != nil {
					log.Println(err.Error())
				}
				log.Printf("got response %v", reply)
			}
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
func (h *Hive) Commit(value string) error {
	h.log = append(h.log, value)
	h.commitCh <- value
	return nil
}
