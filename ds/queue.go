package main

import (
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
	prev *Node
}

// FIFO
type Queue struct {
	head *Node
	tail *Node
}

func (q *Queue) enqueue(d interface{}) *Node {
	//create a new Node
	node := &Node{data: d}

	// if queue is empty
	if q.head == nil {
		q.head, q.tail = node, node
		return node
	}
	q.head.prev = node
	node.next = q.head
	q.head = node
	return node
}

func (q *Queue) dequeue() *Node {

	if q.tail == nil {
		return nil
	}
	// set the tail
	temp := q.tail
	q.tail = q.tail.prev
	if q.tail == nil {
		q.head = nil
	} else {
		q.tail.next = nil
	}

	//reset the node
	temp.next = nil
	temp.prev = nil
	return temp
}

func main() {
	q := &Queue{}
	q.enqueue(4)
	q.enqueue(5)
	q.enqueue(6)
	q.enqueue(7)
	q.enqueue(8)
	fmt.Println(q.dequeue()) // should return 4
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	q.enqueue(9)
	fmt.Println(q.dequeue())
}
