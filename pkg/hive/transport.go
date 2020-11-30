package hive

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

//Transport is a way to propagate your action through to the other nodes
type Transport interface {
	Call(node *Node, method string, cmd Command) Reply
}

//RPCTransport uses net/rpc package to communicate with other nodes
type RPCTransport struct {
	r    *RPC
	port int
}

//Call lets you invoke any valid net/rpc method
func (t *RPCTransport) Call(node *Node, method string, cmd Command) Reply {
	log.Println("RPC.Commit", cmd)

	client, err := rpc.DialHTTP("tcp", node.addr)
	if err != nil {
		log.Println("Connection error: ", err)
	}
	var reply Reply
	err = client.Call(method, cmd, &reply)
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("got response %v", reply)

	return Reply{}
}

//NewRPCTransport is
func NewRPCTransport(port int, rpcCh chan Command) *RPCTransport {
	r := NewRPC(rpcCh)
	rpc.Register(r)
	rpc.HandleHTTP()

	go func() {
		listener, e := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if e != nil {
			log.Fatal("Listen error: ", e)
		}
		log.Printf("Serving RPC server on port %d", port)
		err := http.Serve(listener, nil)
		if err != nil {
			log.Fatal("Error serving: ", err)
		}
	}()
	return &RPCTransport{
		r:    r,
		port: port,
	}
}
