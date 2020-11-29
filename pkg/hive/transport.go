package hive

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Transport interface {
	Commit(command Command) Reply
}

type RPCTransport struct {
	r     *RPC
	port  int
	cmdCh chan Command
}

func (t *RPCTransport) Commit(command Command) Reply {
	log.Println("RPC.Commit", command)
	return Reply{}
}
func NewRPCTransport(port int) *RPCTransport {
	cmdCh := make(chan Command)
	r := NewRPC(cmdCh)
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
	go func() {
		for {
			select {
			case c := <-cmdCh:
				log.Println("c", c)
			}
		}
	}()
	return &RPCTransport{
		r:     r,
		port:  port,
		cmdCh: cmdCh,
	}
}
