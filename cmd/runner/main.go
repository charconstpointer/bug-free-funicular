package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/charconstpointer/arpisi/pkg/arpisi"
)

var (
	nodes = flag.String("nodes", "", "adresses of other nodes in the network")
	port  = flag.Int("port", 7777, "port to listen on")
)
var arpi *arpisi.Arpi

type Message struct {
	Value string `json:"value"`
}

func do(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var msg Message
		b, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(b, &msg)
		var reply arpisi.Reply
		log.Println("commiting", msg.Value)
		arpi.Commit(msg.Value, &reply)
		log.Println("reply", reply.Value)
	}
	fmt.Fprintf(w, "%s", strings.Join(arpi.Log, ", "))
}

func main() {
	flag.Parse()
	tokens := strings.Split(*nodes, ",")
	arpi = arpisi.NewArpi(tokens)
	go arpisi.Run(arpi, *port)
	http.HandleFunc("/", do)
	http.ListenAndServe(":8080", nil)
	time.Sleep(time.Second * 100)
}
