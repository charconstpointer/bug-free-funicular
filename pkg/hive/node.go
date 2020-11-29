package hive

import "log"

//Node represents single node in a cluster
type Node struct {
	addr string
}

//NewNode creates new node, communicating with other nodes on given port
func NewNode(addr string) *Node {
	return &Node{addr}
}

//From cretes nodes from given pool of addresses
func NodesFrom(nodes []string) []*Node {
	n := make([]*Node, 0)
	for _, addr := range nodes {
		log.Println(addr)
		n = append(n, NewNode(addr))
	}
	return n
}
