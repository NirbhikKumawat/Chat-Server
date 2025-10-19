package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	fmt.Println("client connected:", conn.RemoteAddr())
	conn.Close()
}

func main() {
	ln, err := net.Listen("tcp4", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	fmt.Println("Listening on :8080")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("error accepting:", err)
			continue
		}
		go handleConnection(conn)
	}
}
