package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func server() {

	listener, err := net.Listen("tcp", "localhost:8081")
	handleError(err)
	fmt.Println("started server")
	for {
		conn, err := listener.Accept()
		handleError(err)
		fmt.Println("received connection")

		go func(conn net.Conn) {
			handleConnection(conn)
		}(conn)
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func handleConnection(conn net.Conn) {

	for {
		reader := bufio.NewReader(conn)
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
			conn.Close()
		}

		fmt.Println("Message Received:", string(message))
		resp := strings.ToUpper(message)
		conn.Write([]byte(resp + "\n"))
	}

}

func main() {
	server()
}
