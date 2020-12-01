package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"strings"

	"github.com/charconstpointer/arpisi/pkg/testuru"
)

var (
	nodes  = flag.String("nodes", "", "other nodes")
	port   = flag.Int("port", 7999, "port to listen on")
	cursor = flag.String("cursor", ">", "cursor")
)

func main() {
	flag.Parse()
	log.Println(*port)
	ns := strings.Split(*nodes, ",")

	go func() {
		r := new(testuru.Test)
		rpc.Register(r)
		rpc.HandleHTTP()
		listener, e := net.Listen("tcp", fmt.Sprintf(":%d", *port))
		if e != nil {
			log.Fatal("Listen error: ", e)
		}
		log.Printf("Serving RPC server on port %d", *port)
		err := http.Serve(listener, nil)
		if err != nil {
			log.Fatal("Error serving: ", err)
		}
	}()

	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s ", *cursor)
		sc.Scan()
		for _, node := range ns {
			client, err := rpc.DialHTTP("tcp", node)
			if err != nil {
				log.Println("Connection error: ", err)
			}
			var reply testuru.Reply
			err = client.Call("Test.Hello", 1, &reply)
			if err != nil {
				log.Println(err.Error())
			}
			log.Printf("got response %v", reply)
		}
	}
}
