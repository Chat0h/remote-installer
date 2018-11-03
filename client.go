package main

import (
	"log"
	"net"
)

type Client struct {
	Connections map[net.Addr]*net.Conn
}

func NewClient() *Client {
	c := new(Client)
	c.Connections = make(map[net.Addr]*net.Conn)
	return c
}

func (c *Client) NewConnection(network string, addr string) {
	conn, err := net.Dial(network, addr)
	defer conn.Close()
	if err != nil {
		log.Printf("%v", err)
	}
	c.Connections[conn.RemoteAddr()] = &conn
	log.Printf("Connected to: %s\n", conn.RemoteAddr())
}

func main() {
	c := NewClient()
	c.NewConnection("tcp", ":7000")
}
