package main

import "fmt"

type Node struct {
	value      interface{}
	next, prev *Node
}

type DoublyLinkedList struct {
	head, tail *Node
	count      int
}

func (list *DoublyLinkedList) insert(val interface{}) {

	node := &Node{value: val}
	if list.head == nil {
		list.head = node
		list.tail = node
		return
	}
	node.next = list.head
	list.head.prev = node
	list.head = node
}

func (list *DoublyLinkedList) delete(n *Node) {

}

func (list *DoublyLinkedList) print() {

	cur := list.head
	for cur != nil {
		fmt.Println(cur.value)
		cur = cur.next
	}
}

func (list *DoublyLinkedList) printReverse() {

	cur := list.tail
	for cur != nil {
		fmt.Println(cur.value)
		cur = cur.prev
	}
}
func (list *DoublyLinkedList) reverse() {

	cur := list.head
	list.tail = cur
	var prev *Node
	for cur != nil {
		cur.next, cur.prev = cur.prev, cur.next // swap the two pointers
		prev = cur                              // set the prev to the current
		cur = cur.prev                          // move to the next node
	}

	list.head = prev

}

func main() {
	list := &DoublyLinkedList{}
	list.insert(5)
	list.insert(6)
	list.insert(7)
	list.insert(8)
	list.insert(9)
	list.insert(10)
	list.print()
	//list.printReverse()
	list.reverse()
	list.print()
	//list.printReverse()
}
