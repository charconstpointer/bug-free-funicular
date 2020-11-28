package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	nodes = flag.String("nodes", "", "adresses of other nodes in the network")
	port  = flag.Int("port", 7777, "port to listen on")
)

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

type Node struct {
	addr  string
	state State
}

func main() {
	flag.Parse()

	ns := make([]*Node, 0)
	tokens := strings.Split(*nodes, ",")
	log.Println("initializing known nodes")
	for _, node := range tokens {
		log.Printf("adding node %s", node)
		ns = append(ns, &Node{addr: node, state: Unhealthy})
	}

	go serve(*port)
	go watchNodes(ns)

	for {
		for _, node := range ns {
			log.Printf("node %s is now %s", node.addr, node.state.String())
		}
		time.Sleep(time.Second * 2)
	}
}
func serve(port int) {
	log.Println("waiting for connections")
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, port)
	})
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("cannot listen on port ", port)
	}
}

func watchNodes(nodes []*Node) {
	for {
		for _, node := range nodes {
			_, err := http.Get(node.addr)
			if err != nil {
				node.state = Unhealthy
				continue
			}
			node.state = Healthy

		}
		time.Sleep(time.Second * 2)
	}
}
