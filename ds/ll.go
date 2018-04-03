package main

import (
	"fmt"
)

type Node struct {
	Data interface{}
	Next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) add(data interface{}) *Node {

	node := &Node{data, nil}
	if list.head == nil {
		list.head = node
	} else {
		node.Next = list.head
		list.head = node
	}
	return node
}

func (list *LinkedList) addNode(node *Node) *Node {

	if list.head == nil {
		list.head = node
	} else {
		node.Next = list.head
		list.head = node
	}
	return node
}

func (list *LinkedList) print() {

	current := list.head
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}

func printNodeList(node *Node) {

	current := node
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}

func (list *LinkedList) reverse() {

	current := list.head
	var prev, next *Node

	for current != nil {
		next = current.Next // store nextt pointer
		current.Next = prev // set next to prev
		prev = current      // store current in prev
		current = next      // make next current
	}
	list.head = prev // set the head to the last prev
}

func (list *LinkedList) getKthFromLast(k int) interface{} {

	if list.head == nil || k <= 0 {
		return nil
	}
	main_pointer := list.head
	ref_pointer := list.head
	var count int

	//move the reference pointer k times
	for count < k {
		ref_pointer = ref_pointer.Next
		count++
	}

	if count != k {
		// not enough elements return nil
		return nil
	}

	for ref_pointer != nil {
		ref_pointer = ref_pointer.Next
		main_pointer = main_pointer.Next

	}

	return main_pointer.Data

}

func (list *LinkedList) detectLoopUsingSet() bool {
	set := make(map[*Node]bool)

	current := list.head

	for current != nil {

		if set[current] {
			// found, there is a loop
			return true
		}
		set[current] = true
		current = current.Next
	}
	return false
}

// take two pointers and move 1-step, 2-sep resp, and see if they cross
func (list *LinkedList) detectLoop() bool {
	//set := make(map[*Node]bool)

	if list.head == nil {
		return false
	}
	slow_pointer := list.head
	fast_pointer := list.head

	for slow_pointer != nil && fast_pointer != nil && fast_pointer.Next != nil {
		slow_pointer = slow_pointer.Next
		fast_pointer = fast_pointer.Next.Next
		if slow_pointer == fast_pointer {
			return true
		}
	}
	return false
}

func mergeSortedList(a *Node, b *Node) *Node {

	var newRoot *Node

	current := &Node{} // dummy node
	for {

		if a == nil {
			current.Next = b
			break
		} else if b == nil {
			current.Next = a
			break
		}
		if a.Data.(int) < b.Data.(int) {
			current.Next = a
			a = a.Next
		} else {
			current.Next = b
			b = b.Next
		}
		if newRoot == nil {
			newRoot = current.Next
		}
		current = current.Next
	}

	return newRoot
}

// see if two linked list intersect
func getIntersectionNode(headA *Node, headB *Node) int {

	if headA == nil || headB == nil {
		return -1
	}

	curA := headA
	curB := headB
	for curA != curB {

		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}
		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
		if curA == headA && curB == headB {
			return -1

		}
	}
	return curA.Data.(int)
}

func main() {
	fmt.Println("hello")
	ll := &LinkedList{}
	ll.add(47)
	ll.add(26)
	ll.add(14)
	ll.add(5)
	//ll.addNode(n)
	ll.print()
	//ll.reverse()
	//ll.print()
	ll2 := &LinkedList{}
	ll2.add(33)
	ll2.add(24)
	ll2.add(22)
	ll2.add(12)
	ll2.add(2)
	ll2.print()
	//fmt.Println("Print kthe element", ll.getKthFromLast(0))
	//fmt.Println("Is Loop", ll.detectLoopUsingSet())
	printNodeList(mergeSortedList(ll.head, ll2.head))
}
