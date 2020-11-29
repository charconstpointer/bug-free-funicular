package hive

type RPC struct {
	cmdCh chan Command
}

func NewRPC(cmdCh chan Command) *RPC {
	return &RPC{
		cmdCh: cmdCh,
	}
}

type Reply struct {
	Value string
}
type Command struct {
	Value interface{}
}

func (r *RPC) Commit(command Command, reply *Reply) error {
	r.cmdCh <- command
	*reply = Reply{Value: "reply"}
	return nil
}
