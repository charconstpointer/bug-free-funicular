// package main

// import (
// 	"log"
// 	"net"
// 	"strings"
// 	"time"
// )

// type Role uint32
// type State uint32

// const (
// 	Master Role = iota
// 	Slave
// )
// const (
// 	Healthy State = iota
// 	Unhealthy
// )

// func (s State) String() string {
// 	switch s {
// 	case Healthy:
// 		return "Healthy"

// 	case Unhealthy:
// 		return "Unhealthy"
// 	default:
// 		return "Unknown"
// 	}
// }

// type Client struct {
// 	conn  net.Conn
// 	state State
// 	role  Role
// }

// func (c *Client) BecomeMaster() {
// 	c.conn.Write([]byte("master"))
// }

// type Hub struct {
// 	clients []*Client
// }

// func main() {
// 	l, err := net.Listen("tcp", ":7777")
// 	hub := Hub{clients: make([]*Client, 0)}
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	for {
// 		log.Printf("current clients %d", len(hub.clients))
// 		conn, err := l.Accept()
// 		if err != nil {
// 			log.Fatalf(err.Error())
// 		}
// 		c := &Client{conn, Healthy, Slave}
// 		go handleClient(c)
// 		hub.clients = append(hub.clients, c)

// 	}
// }

// func handleClient(c *Client) {
// 	log.Printf("handle client for %s", c.conn.RemoteAddr().String())
// 	go checkLiveness(c)
// 	buffer := make([]byte, 1024)
// 	for {
// 		switch c.state {
// 		case Healthy:
// 			_, err := c.conn.Read(buffer)
// 			if err != nil {
// 				c.state = Unhealthy
// 			}
// 			if strings.TrimSpace(string(buffer)) != "" {
// 				log.Println(string(buffer))
// 			}
// 			break
// 		default:
// 			log.Println("waiting for client to become healthy again")
// 			time.Sleep(time.Second * 3)
// 		}

// 	}
// }

// func checkLiveness(c *Client) {
// 	for {
// 		log.Printf(">%s", c.state.String())
// 		c.conn.SetWriteDeadline(time.Now().Add(time.Second * 1))
// 		_, err := c.conn.Write([]byte("ping"))
// 		if err != nil {
// 			c.state = Unhealthy
// 		}
// 		time.Sleep(time.Second * 1)
// 	}
// }
