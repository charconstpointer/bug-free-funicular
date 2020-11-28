package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":7777")
	if err != nil {
		panic(err.Error())
	}
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Println(n)
	}
}
