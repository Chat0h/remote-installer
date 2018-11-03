package main

import (
	"flag"
	"log"
	"net"
)

type Agent struct {
	Network string
	Addr    string
}

func (a *Agent) HandleConn(conn net.Conn) {
	log.Printf("New connection: %s", conn.RemoteAddr())
}

func (a *Agent) Listen() {
	srv, err := net.Listen(a.Network, a.Addr)
	log.Printf("Server agent started on %s", a.Addr)
	defer srv.Close()
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := srv.Accept()
		if err != nil {
			log.Printf("%v", err)
		}
		go a.HandleConn(conn)
	}
}

func main() {
	network := flag.String("network", "tcp", "Network")
	addr := flag.String("bind", "127.0.0.1:7000", "Bind address")
	flag.Parse()
	a := Agent{*network, *addr}
	a.Listen()
}
