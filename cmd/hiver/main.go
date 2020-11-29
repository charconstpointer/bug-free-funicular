package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charconstpointer/arpisi/pkg/hive"
)

var (
	nodes = flag.String("nodes", "", "other nodes")
	port  = flag.Int("port", 7999, "port to listen on")
)

func main() {
	flag.Parse()
	sc := bufio.NewScanner(os.Stdin)

	tokens := strings.Split(*nodes, ",")
	nodes := make([]*hive.Node, 0)
	for _, addr := range tokens {
		log.Println(addr)
		nodes = append(nodes, hive.NewNode(addr))
	}
	hive := hive.NewHive(nodes, *port)

	for {
		fmt.Printf("%s", ">")
		sc.Scan()
		command := sc.Text()
		cmdTokens := strings.Split(command, " ")
		if len(cmdTokens) != 2 {
			log.Println("wrong cmd")
			continue
		}
		switch cmdTokens[0] {
		case "cmd":
			err := hive.Commit(cmdTokens[1])
			if err != nil {
				log.Println(err.Error())
			}
		}
	}
}
