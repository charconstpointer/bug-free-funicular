package hive

type RPC struct {
	rpcCh chan Command
}

func NewRPC(rpcCh chan Command) *RPC {
	return &RPC{
		rpcCh: rpcCh,
	}
}

type Reply struct {
	Value string
}
type Command struct {
	Value interface{}
}

func (r *RPC) Commit(command Command, reply *Reply) error {
	r.rpcCh <- command

	*reply = Reply{Value: "reply"}
	return nil
}
