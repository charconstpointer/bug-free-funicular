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

	tokens := strings.Split(*nodes, ",")
	nodes := hive.NodesFrom(tokens)
	h := hive.NewHive(nodes, *port)
	listenInput(h)
}

func listenInput(h *hive.Hive) {
	sc := bufio.NewScanner(os.Stdin)
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

			cmd := hive.Command{Value: cmdTokens[1]}
			err := h.Commit(cmd)
			if err != nil {
				log.Println(err.Error())
			}
		}
	}
}
