package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp4", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Connected to the chat server!")
	fmt.Print("Enter your name: ")

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Fprintf(conn, "%s joined the chat\n", name)

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	for {
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, "%s: %s", name[:len(name)-1], text)
	}
}
