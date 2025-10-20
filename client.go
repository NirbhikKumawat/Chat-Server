package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

var (
	cyan  = "\033[36m"
	green = "\033[32m"
	reset = "\033[0m"
)

func main() {
	conn, err := net.Dial("tcp4", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println(cyan + "Connected to the chat server!" + reset)
	fmt.Print("Enter your name: ")

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)
	fmt.Fprintf(conn, "%s joined the chat\n", name)

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		timestamp := time.Now().Format("15:04:05")
		msg := fmt.Sprintf("%s[%s][%s]: %s%s\n", green, timestamp, name, text, reset)
		fmt.Fprintf(conn, msg)
	}
}
