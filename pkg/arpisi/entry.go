package arpisi

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

//Run is
func NewArpi(nodes []string) *Arpi {
	ns := make([]*Node, 0)
	log.Println("initializing known nodes")
	for _, node := range nodes {
		log.Printf("adding node %s", node)
		ns = append(ns, &Node{addr: node, state: Unhealthy})
	}
	arpi := Arpi{nodes: ns}
	return &arpi
}

//Run is
func Run(a *Arpi, port int) {
	a.listenRPC(port)
	// a.pingNodes(a.nodes)
}
func (a *Arpi) pingNodes(nodes []*Node) {
	for {
		log.Println("pingnodes", len(nodes))
		for _, node := range nodes {
			log.Println(node)
			client, err := rpc.DialHTTP("tcp", node.addr)
			if err != nil {
				log.Println("Connection error: ", err)
				continue
			}
			var reply Reply
			client.Call("Arpi.Commit", "hello", &reply)
			log.Printf("got response %v", reply)
		}
		time.Sleep(5 * time.Second)
	}
}
func (a *Arpi) listenRPC(port int) {
	rpc.Register(a)
	rpc.HandleHTTP()

	listener, e := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if e != nil {
		log.Fatal("Listen error: ", e)
	}
	log.Printf("Serving RPC server on port %d", port)
	err := http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving: ", err)
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
