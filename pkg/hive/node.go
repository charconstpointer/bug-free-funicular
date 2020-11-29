package hive

//Node represents single node in a cluster
type Node struct {
	addr string
}

//NewNode creates new node, communicating with other nodes on given port
func NewNode(addr string) *Node {
	return &Node{addr}
}
