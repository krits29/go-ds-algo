package main

import "fmt"

type Node struct {
	Key   interface{}
	Value interface{}
	Next  *Node
}

type LinkedList struct {
	count int
	head  *Node
}

type Map struct {
	list []*LinkedList
	size int
}

func (m *Map) init(size int) {
	m.list = make([]*LinkedList, size)
	m.size = size

	for i := range m.list {
		m.list[i] = &LinkedList{}
	}

}

func (m *Map) add(key interface{}, value interface{}) {

	index := hash(key.(string)) % m.size

	m.list[index].add(key, value)
}

func (m *Map) get(key interface{}) interface{} {

	index := hash(key.(string)) % m.size

	list := m.list[index]
	if list == nil {
		return nil
	}
	return list.get(key)

}

func (list *LinkedList) add(key interface{}, value interface{}) *Node {

	node := &Node{key, value, nil}
	if list.head == nil {
		list.head = node
	} else {
		node.Next = list.head
		list.head = node
	}
	return node
}

func (list *LinkedList) get(key interface{}) interface{} {
	current := list.head
	for current != nil {
		if current.Key == key {
			return current.Value
		}
		current = current.Next
	}
	return nil
}

func (list *LinkedList) print() {

	current := list.head
	for current != nil {
		fmt.Print(current.Key, " ", current.Value)
		current = current.Next
	}
	fmt.Println(" ")
}

func (m *Map) print() {
	for i := range m.list {
		m.list[i].print()
	}
}

func hash(s string) int {
	h := 0
	for _, r := range s {
		h = 31*h + int(r)
	}
	return h
}

func main() {
	fmt.Println(hash("foobar"))
	m := &Map{}
	m.init(10)
	m.add("one", 1)
	fmt.Println(m.get("one"))
	fmt.Println(m.get("two"))
	m.add("two", 2)
	fmt.Println(m.get("two"))
	m.print()
}
