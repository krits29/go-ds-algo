package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	left, right *Node
	key         int
	value       interface{}
	count       int
}

type Tree struct {
	root *Node
}

func size(node *Node) int {
	if node == nil {
		return 0
	}
	return node.count
}

func height(node *Node) int {

	//base condition
	if node == nil {
		return 0
	}

	leftHeight := height(node.left)
	rightHeight := height(node.right)

	return 1 + max(leftHeight, rightHeight)
}

//2-1-3
func levelOrder(node *Node) {

	if node == nil {
		return
	}
	height := height(node)

	for i := 1; i <= height; i++ {
		printLevel(node, i)
	}

}

func printLevel(node *Node, level int) {

	if node == nil {
		return
	}
	if level == 1 {
		fmt.Print(node.key, " ")
	}
	if level > 1 {
		printLevel(node.left, level-1)
		printLevel(node.right, level-1)
	}

}

// 1-2-3
func inOrder(node *Node) {

	//base condition
	if node == nil {
		return
	}

	inOrder(node.left)
	fmt.Println(node.key, node.value)
	inOrder(node.right)

}

// Morris Traversal without recurrsion
func inOrderMorrisTraversal(node *Node) {
	if node == nil {
		return
	}

	curr := node
	var left *Node
	for curr != nil {

		if curr.left == nil { // if left is nil
			fmt.Println(curr.key) // print the current for 1-2-3 in-order
			curr = curr.right     // go to the right
		} else {
			left = curr.left                              // save the left child
			for left.right != nil && left.right != curr { // right is not empty and is not parent
				left = left.right
			}
			if left.right == nil { // if parent does not have right, put the parent as the right node
				left.right = curr
				curr = curr.left
			} else {
				left.right = nil
				fmt.Println(curr.key) // this is the right node
				curr = curr.right
			}

		}
	}
}

//2-1-3
func preOrder(node *Node) {

	//base condition
	if node == nil {
		return
	}
	fmt.Println(node.key, node.value)
	preOrder(node.left)
	preOrder(node.right)

}

//2-3-1
func postOrder(node *Node) {

	//base condition
	if node == nil {
		return
	}
	postOrder(node.left)
	postOrder(node.right)
	fmt.Println(node.key, node.value)

}

func (t *Tree) insert(node *Node, key int, value interface{}) *Node {
	if node == nil {
		return &Node{key: key, value: value, count: 1}
	}
	switch {
	case key < node.key: // go left side
		node.left = t.insert(node.left, key, value)
	case key > node.key: // go right side
		node.right = t.insert(node.right, key, value)
	default: // update the root
		node.value = value
	}
	// increase the count
	node.count = 1 + size(node.left) + size(node.right)
	return node
}

// linear search
func (t *Tree) get(key int) interface{} {

	x := t.root

	for x != nil {
		switch {
		case key < x.key: // go left
			x = x.left
		case key > x.key: // go right
			x = x.right
		default:
			return x.value
		}
	}
	return nil
}

// find an index which has the same value as index itself
func findMagicIndex(list []int, lo int, hi int) int {

	//base condition
	if hi < lo || lo < 0 || hi > len(list) {
		return -1
	}

	midIndex := lo + (hi-lo)/2
	midValue := list[midIndex]
	if midIndex == midValue {
		return midIndex
	} else if midValue > midIndex {
		return findMagicIndex(list, lo, midIndex-1)
	} else {
		return findMagicIndex(list, midIndex+1, hi)
	}

}

// find an index which has the same value as index itself
func findMagicIndexDuplicate(list []int, lo int, hi int) int {

	//base condition
	if hi < lo || lo < 0 || hi > len(list) {
		return -1
	}

	midIndex := lo + (hi-lo)/2
	midValue := list[midIndex]
	if midIndex == midValue {
		return midIndex
	}

	// search left // find minimum of midIndex and midvalue
	left := findMagicIndexDuplicate(list, lo, min(midIndex-1, midValue))
	if left >= 0 {
		return left
	}

	// search right // max of midIndex+1 and midValue
	right := findMagicIndexDuplicate(list, max(midIndex+1, midValue), hi)

	return right
}

// Simple Binary Search
func binarySearch(list []int, lo int, hi int, k int) int {

	//base condition
	if hi < lo || lo < 0 || hi >= len(list) {
		return -1
	}

	midIndex := lo + (hi-lo)/2
	midValue := list[midIndex]

	switch {
	case midValue == k:
		return midIndex
	case midValue > k:
		return binarySearch(list, lo, midIndex-1, k)
	default:
		return binarySearch(list, midIndex+1, hi, k)
	}
	/*if midValue == k {
		return midIndex
	} else if midValue > k {
		return binarySearch(list, 0, midIndex-1, k)
	} else {
		return binarySearch(list, midIndex+1, hi, k)
	}*/

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// do inorder traversal and check the increasing trend
func isBST(node *Node, prev *Node) bool {

	if node == nil {
		return true
	}
	// if left side is not in order return false
	if !isBST(node.left, prev) {
		return false
	}

	// if prev is small continue otherwise return false
	if prev != nil && node.key <= prev.key {
		return false
	}
	prev = node

	return isBST(node.right, prev)
}

// do inorder traversal and decrement by 1 everytime
func kthSmallestUsingMorris(node *Node, k int) int {

	if node == nil {
		return -1
	}
	// Count to iterate over elements till we
	// get the kth smallest number
	count := 0
	res := -1
	curr := node
	var pre *Node
	for curr != nil {

		if curr.left == nil {
			count++ // count self node
			if count == k {
				res = curr.key
			}
			curr = curr.right

		} else {
			pre = curr.left // save the left child
			for pre.right != nil && pre.right != curr {
				pre = pre.right
			}
			if pre.right == nil {
				pre.right = curr
				curr = curr.left
			} else { // right child is the parent node
				pre.right = nil
				count++
				if count == k {
					res = curr.key
				}
				curr = curr.right
			}

		}
	}
	return res
}

// Breath First Traversal
func levelOrderQueue(root *Node) [][]int {

	arr := [][]int{}
	if root == nil {
		return arr
	}
	queue := list.New()  // create a new queue
	queue.PushBack(root) // push the root
	count := 0
	var node *Node
	for queue.Len() != 0 {
		count = queue.Len()
		subarr := []int{}
		// go through every element in the queue
		for i := 0; i < count; i++ {
			node = queue.Front().Value.(*Node) // peek into the front element
			// add the children to the queue if present
			if node.left != nil {
				queue.PushBack(node.left)
			}
			if node.right != nil {
				queue.PushBack(node.right)
			}

			subarr = append(subarr, node.key)
			queue.Remove(queue.Front()) // remove the element from the queue
		}
		arr = append(arr, subarr)
	}
	return arr
}

func isSymmetricR(root *Node) bool {

	return isMirror(root, root)
}

func isMirror(t1 *Node, t2 *Node) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return t1.key == t2.key && isMirror(t1.right, t2.left) && isMirror(t1.left, t2.right)
}

// non-recursive
func isSymmetric(root *Node) bool {

	queue := list.New()  // create a new queue
	queue.PushBack(root) // push the root
	queue.PushBack(root) // push the root again to make it symmetric
	for queue.Len() != 0 {
		t1 := queue.Remove(queue.Front()).(*Node)
		t2 := queue.Remove(queue.Front()).(*Node)
		if t1 == nil && t2 == nil {
			continue
		}
		if t1 == nil || t2 == nil {
			return false
		}
		if t1.key != t2.key {
			return false
		}
		queue.PushBack(t1.left)
		queue.PushBack(t1.right)
		queue.PushBack(t2.left)
		queue.PushBack(t2.right)

	}
	return true
}

func main() {
	t := &Tree{}
	t.root = t.insert(t.root, 4, "hello4")
	t.root = t.insert(t.root, 5, "hello5")
	t.root = t.insert(t.root, 6, "hello6")
	t.root = t.insert(t.root, 1, "hello1")
	t.root = t.insert(t.root, 2, "hello2")
	t.root = t.insert(t.root, 4, "hhh4")
	inOrderMorrisTraversal(t.root)
	levelOrder(t.root)
	fmt.Println(t.get(4), size(t.root.right))
	fmt.Println("isBST", isBST(t.root, nil))
	fmt.Println("kthSmallestUsingMorris", kthSmallestUsingMorris(t.root, 5))
	l := []int{-2, -1, 0, 4, 5, 6, 7, 8, 9, 10, 12, 14, 15, 19, 21, 24, 27}
	fmt.Println(findMagicIndex(l, 0, len(l)-1))
	fmt.Println(binarySearch(l, 0, len(l)-1, 24))
	//fmt.Println(height(t.root))
}
