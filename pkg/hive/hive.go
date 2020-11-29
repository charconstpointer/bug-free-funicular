package hive

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type Hive struct {
	nodes         []*Node
	log           []string
	commitCh      chan string
	rpcCh         chan string
	routinesGroup sync.WaitGroup
	port          int
}

func NewHive(nodes []*Node, port int) *Hive {
	hive := &Hive{
		nodes:    nodes,
		log:      make([]string, 0),
		commitCh: make(chan string),
		rpcCh:    make(chan string),
		port:     port,
	}
	hive.goFunc(hive.run)
	hive.goFunc(hive.runServer)
	return hive
}

type Message struct {
	Value string `json:"value"`
}

func (h *Hive) do(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var msg Message
		b, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(b, &msg)
		h.rpcCh <- msg.Value
	}
	fmt.Fprintf(w, "%s", "ok")
}

func (h *Hive) runServer() {
	http.HandleFunc("/", h.do)
	err := http.ListenAndServe(fmt.Sprintf(":%d", h.port), nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (h *Hive) run() {
	for {
		select {
		case c := <-h.commitCh:
			msg := Message{Value: c}
			b, err := json.Marshal(msg)
			if err != nil {
				log.Println(err.Error())
			}
			for _, node := range h.nodes {
				log.Println("what", node.addr)
				_, err := http.Post(node.addr, "application/json", bytes.NewReader(b))
				if err != nil {
					log.Println(err.Error())
				}
			}
		case r := <-h.rpcCh:
			h.log = append(h.log, r)
			log.Println("current log", h.log)
		}
	}
}
func (h *Hive) goFunc(f func()) {
	h.routinesGroup.Add(1)
	go func() {
		defer h.routinesGroup.Done()
		f()
	}()
}
func (h *Hive) Commit(value string) error {
	h.log = append(h.log, value)
	h.commitCh <- value
	return nil
}

type Node struct {
	addr string
}

func NewNode(addr string) *Node {
	return &Node{addr}
}
