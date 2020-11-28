package main

import (
	"net"
	"time"
)

func main() {
	_, err := net.Dial("tcp", ":7777")
	if err != nil {
		panic(err.Error())
	}
	time.Sleep(20 * time.Second)
}
