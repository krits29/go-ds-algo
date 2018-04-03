package main

import "fmt"

// search miss could be sublinear
// search hit is L characters - log(n)
// waste space for nil nodes
// used for SPELL CHECKING // 26 array trie ( all lowercase char)
// out of memory for UNICODE
type Node struct {
	isEndOfWord bool
	children    []*Node
}

const (
	R = 26
)

func newNode() *Node {
	return &Node{isEndOfWord: false, children: make([]*Node, R)} // 256 for extended ASCII
}

type Trie struct {
	root *Node
}

func (t *Trie) insert(s string) {
	t.root = t.insertUtil(t.root, s, 0) // index of the char in the string
}

func (t *Trie) insertUtil(node *Node, s string, index int) *Node {

	if node == nil {
		node = newNode() // if node is null, create a new node
	}
	if index == len(s) { // if this is the last character in the string
		node.isEndOfWord = true
		return node
	}

	pos := s[index] - 'a'
	node.children[pos] = t.insertUtil(node.children[pos], s, index+1) // recursive follow that node
	return node
}

func (t *Trie) search(s string) bool {
	if t.root == nil {
		return false
	}
	return t.searchUtil(t.root, s, 0)
}

func (t *Trie) searchUtil(node *Node, s string, index int) bool {
	if node == nil {
		return false
	}

	if index == len(s) {
		return node.isEndOfWord
	}
	pos := s[index] - 'a'
	return t.searchUtil(node.children[pos], s, index+1)
}

func (t *Trie) searchNonR(s string) bool {
	var index rune
	curr := t.root
	for _, r := range s {
		index = r - 'a'
		if curr.children[index] == nil {
			return false
		}
		curr = curr.children[index]
	}
	return curr != nil && curr.isEndOfWord
}

// english spell checker
// word completion
func main() {
	fmt.Println("Hello Trie")
	list := []string{"the", "a", "there", "answer", "any",
		"by", "bye", "their"}
	t := &Trie{}
	//t.root = newNode() // create a root
	//if t.root == nil {
	//	fmt.Println("oopsadfsdfdsfdsf")
	//}
	for _, str := range list {
		t.insert(str)
	}
	if t.search("answer") {
		fmt.Println("Present")
	}
}
