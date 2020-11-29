package main

import (
	"flag"
	"strings"

	"github.com/charconstpointer/arpisi/pkg/arpisi"
)

var (
	nodes = flag.String("nodes", "", "adresses of other nodes in the network")
	port  = flag.Int("port", 7777, "port to listen on")
)

func main() {
	flag.Parse()
	tokens := strings.Split(*nodes, ",")
	arpisi.Run(tokens, *port)
}
