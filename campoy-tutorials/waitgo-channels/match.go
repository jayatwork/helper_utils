package main

import (
	"fmt"
	"net"
)

var partners = make(chan net.Conn)

func match(conn net.Conn) {
	fmt.Println(conn, "Waiting for a partner....")
	select {
	case partners <- conn:
		//goodbye
	case p := <-partners:
		chat(conn, p)
	}
}

func chat(a, b net.Conn)
