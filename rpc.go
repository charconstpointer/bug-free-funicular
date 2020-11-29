// package arpisi

// import (
// 	"flag"
// 	"fmt"
// 	"log"
// 	"net"
// 	"net/http"
// 	"net/rpc"
// )

// type Server struct {
// }
// type Reply struct {
// 	Value int
// }

// func (s *Server) Ping(arg int, reply *Reply) error {
// 	log.Println(arg)
// 	*reply = Reply{Value: 13}
// 	return nil
// }

// var (
// 	clientType = flag.String("clientType", "server", "who are you")
// )

// func _main() {
// 	flag.Parse()
// 	if *clientType != "server" {
// 		client, err := rpc.DialHTTP("tcp", "localhost:1234")
// 		if err != nil {
// 			log.Fatal("Connection error: ", err)
// 		}
// 		var reply Reply
// 		client.Call("Server.Ping", 1, &reply)
// 		log.Printf("got response %v", reply)
// 	} else {
// 		server := new(Server)
// 		rpc.Register(server)
// 		rpc.HandleHTTP()

// 		listener, e := net.Listen("tcp", ":1234")
// 		if e != nil {
// 			log.Fatal("Listen error: ", e)
// 		}
// 		log.Printf("Serving RPC server on port %d", 1234)
// 		err := http.Serve(listener, nil)
// 		if err != nil {
// 			log.Fatal("Error serving: ", err)
// 		}
// 	}

// }

// func sayHi(user string) string {
// 	return fmt.Sprintf("Hi, '%s' :D", user)
// }
