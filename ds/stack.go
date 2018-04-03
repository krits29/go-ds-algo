package main

import (
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
}

// LIFO
type Stack struct {
	head *Node
}

func (s *Stack) push(d interface{}) *Node {
	//create a new Node
	node := &Node{data: d}

	// if stack is empty
	if s.head == nil {
		s.head = node
		return node
	}
	node.next = s.head
	s.head = node
	return node
}

func (s *Stack) pop() *Node {

	if s.head == nil {
		return nil
	}

	//reset the head
	temp := s.head
	s.head = temp.next
	if temp != nil {
		temp.next = nil

	}
	return temp
}

func main() {
	s := &Stack{}
	s.push(4)
	s.push(5)
	s.push(6)
	s.push(7)
	s.push(8)
	fmt.Println(s.pop()) // should return 4
	fmt.Println(s.pop())
	fmt.Println(s.pop())

	s.push(9)
	fmt.Println(s.pop())
}
