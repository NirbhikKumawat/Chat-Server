package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

var (
	clients = make(map[net.Conn]bool)
	mutex   sync.Mutex
)

func broadcast(sender net.Conn, msg string) {
	mutex.Lock()
	defer mutex.Unlock()
	for client := range clients {
		if client != sender {
			client.Write([]byte(msg))
		}
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()
	fmt.Println("client connected:", conn.RemoteAddr())
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("client disconnected:", conn.RemoteAddr())
			mutex.Lock()
			delete(clients, conn)
			mutex.Unlock()
			return
		}
		fmt.Printf("[%s] %s", conn.RemoteAddr(), msg)
		fmt.Println()
		broadcast(conn, msg)
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
