package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func createServer() {
	listerner, err := net.Listen("tcp", "localhost:8080")
	log.Println("Created new server")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listerner.Accept()
		if err != nil {
			log.Fatal(err)
			continue // continue to accept other connections
		}
		go handleConnection(conn)
	}

}

func broadcaster() {
	// list of all connected clients
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages: // accept a message from the connection and broadcase it to all connections
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true // add the new connection
		case cli := <-leaving:
			delete(clients, cli)
			close(cli) // close the leaving connection
		}
	}
}

func handleConnection(conn net.Conn) {
	ch := make(chan string, 100) // channel for outgiong client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "Welcome! You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		msg := input.Text()
		fmt.Println(msg)
		messages <- msg
	}
	leaving <- ch
	messages <- who + " is leaving"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	createServer()
}
