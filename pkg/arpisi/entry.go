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
func Run(a *Arpi, port int) {
	a.listenRPC(port)
	// a.pingNodes(a.nodes)
}
func (a *Arpi) pingNodes(nodes []*Node) {
	for {
		log.Println("pingnodes", len(nodes))
		for _, node := range nodes {
			log.Println(node)
			client, err := rpc.DialHTTP("tcp", node.Addr)
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
