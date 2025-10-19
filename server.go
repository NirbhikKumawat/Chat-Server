package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("client connected:", conn.RemoteAddr())
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("client disconnected:", conn.RemoteAddr())
			return
		}
		fmt.Printf("[%s] %s", conn.RemoteAddr(), msg)
	}
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
