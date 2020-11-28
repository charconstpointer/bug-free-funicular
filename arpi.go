package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

var (
	nodes = flag.String("nodes", "", "adresses of other nodes in the network")
	port  = flag.Int("port", 7777, "port to listen on")
)

type Node struct {
	conn net.Conn
}

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal("cannot listen on port ", *port)
	}
	go func() {
		log.Println("waiting for connections")
		for {

			conn, err := listener.Accept()
			if err != nil {
				log.Fatalf(err.Error())
			}
			log.Println("accepted conn from ", conn.RemoteAddr())
		}
	}()

	ns := make([]*Node, 0)
	flag.Parse()
	tokens := strings.Split(*nodes, ",")
	for _, n := range tokens {
		conn, err := net.Dial("tcp", n)
		if err != nil {
			continue
		}
		ns = append(ns, &Node{conn})
	}

	time.Sleep(time.Second * 99)

}
