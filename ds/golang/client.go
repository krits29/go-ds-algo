package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func client() {
	conn, _ := net.Dial("tcp", "localhost:8081")

	for {
		// read from command line
		reader := bufio.NewReader(os.Stdin)
		msg, _ := reader.ReadString('\n')

		//send to socket
		fmt.Fprintln(conn, msg)

		//listen for reply
		reply, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(reply)
	}
}

func main() {
	client()
}
