package testuru

import "log"

type Command int
type Reply int
type Test struct {
}

func (t *Test) Hello(cmd Command, reply *Reply) error {
	*reply = -1
	return nil
}

func (t *Test) Dont() {
	log.Println("please ignore me")
}
