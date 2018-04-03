package main

import "fmt"

// explicit character and only 3 nodes
type Node struct {
	char          byte
	Left          *Node
	Right         *Node
	eq            *Node
}

type Ternary struct {
	root *Node
}

func newNode() *Node {
	return &Node{}
}

func (t *Ternary) insert(word string) {
	t.root = t.insertUtil(t.root, word, 0)
}

func (t *Ternary) insertUtil(node *Node, word string, index int) *Node {

	r := word[index]
	if node == nil {
		node = newNode()
		node.char = r
	}

	if r < node.char { // go left
		node.Left = t.insertUtil(node.Left, word, index)
	} else if r > node.char { // go right
		node.Right = t.insertUtil(node.Right, word, index)
	} else if index < len(word)-1 { // equal
		node.eq = t.insertUtil(node.eq, word, index+1)
	}
	return node
}

func (t *Ternary) search(word string) bool {

	return t.searchUtil(t.root, word, 0)
}

func (t *Ternary) searchUtil(node *Node, word string, index int) bool {

	if node == nil {
		return false
	}

	r := word[index]
	switch {
	case r < node.char: // go left
		return t.searchUtil(node.Left, word, index)
	case r > node.char: // go right
		return t.searchUtil(node.Right, word, index)
	case index < len(word)-1: // go eq
		return t.searchUtil(node.eq, word, index+1)
	default:
		return true
	}

}

func main() {

	fmt.Println("hello")
	list := []string{"the", "a", "there", "answer", "any",
		"by", "bye", "their"}
	t := &Ternary{}
	for _, str := range list {
		t.insert(str)
	}
	if t.search("answer") {
		fmt.Println("Present")
	}
}
