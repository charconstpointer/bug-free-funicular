package main

import (
	"flag"
	"log"
	"strings"
	"time"

	"github.com/charconstpointer/arpisi/pkg/hive"
)

var (
	nodes = flag.String("nodes", "", "other nodes")
	port  = flag.Int("port", 7999, "port to listen on")
)

func main() {
	flag.Parse()
	tokens := strings.Split(*nodes, ",")
	nodes := make([]*hive.Node, 0)
	for _, addr := range tokens {
		log.Println(addr)
		nodes = append(nodes, hive.NewNode(addr))
	}
	hive := hive.NewHive(nodes, *port)

	err := hive.Commit("hello")
	if err != nil {
		log.Println(err.Error())
	}

	time.Sleep(time.Second * 100)
}
