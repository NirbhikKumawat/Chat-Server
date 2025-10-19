package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp4", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	fmt.Println("Listening on :8080")
	conn, err := ln.Accept()
	if err != nil {
		panic(err)
	}
	fmt.Println("Accepted connection from", conn.RemoteAddr())
	conn.Close()
}
