package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	Data    interface{}
	Next    *Node
	visited bool
}

type LinkedList struct {
	count int
	head  *Node
}

// Graph is represented as an array of LinkedList containing all the vertices
// represents a directed graph using adjacency list representation
type Graph struct {
	list []*LinkedList // list of all the vertices
	V    int           // number of Vertices
}

func (list *LinkedList) add(data interface{}) *Node {

	node := &Node{data, nil, false}
	if list.head == nil {
		list.head = node
	} else {
		node.Next = list.head
		list.head = node
	}
	list.count++
	return node
}

func (list *LinkedList) print() {

	current := list.head
	for current != nil {
		fmt.Print(current.Data, " ")
		current = current.Next
	}
	fmt.Println(" ")
}

func (g *Graph) init(v int) {
	g.list = make([]*LinkedList, v)
	g.V = v

	for i := range g.list {
		g.list[i] = &LinkedList{}
	}

}

// returns list( of vertices) adjacent to the given vertex
func (g *Graph) adj(v int) *LinkedList {
	if v < 0 || v > len(g.list)-1 {
		return nil
	}
	return g.list[v]
}

func (g *Graph) addEdge(src int, dest int) {
	g.list[src].add(dest)
	g.list[dest].add(src) // undirected Graph
}

func (g *Graph) print() {
	for i := range g.list {
		fmt.Print(i, " ")
		g.list[i].print()
	}
}

// print degree - depth of a vertex
func (g *Graph) degree(v int) int {
	list := g.adj(v)
	degree := 0
	if list != nil {
		current := list.head
		for current != nil {
			degree++
			current = current.Next
		}
	}
	return degree
}

// max degree of the graph
func (g *Graph) maxDegree() int {

	maxDegree := 0

	for _, list := range g.list {
		degree := 0
		if list != nil {
			current := list.head
			for current != nil {
				degree++
				current = current.Next
			}
		}
		maxDegree = max(maxDegree, degree)
	}
	return maxDegree
}

// Breath First Traversal from source s
func (paths *Paths) BFS(s int) {

	queue := list.New()
	paths.marked = make([]bool, len(paths.g.list))
	paths.edgeTo = make([]int, len(paths.g.list))

	queue.PushBack(s)
	paths.marked[s] = true
	//fmt.Println("BFS visiting", s)

	for queue.Len() != 0 {
		//v := queue.Front().Value.(int)
		v := queue.Remove(queue.Front()).(int)

		current := paths.g.adj(v).head
		for current != nil {
			n := current.Data.(int)
			if !paths.marked[n] {
				queue.PushBack(n)
				paths.marked[n] = true
				//fmt.Println("BFS visiting", n)
				paths.edgeTo[n] = v
			}
			current = current.Next
		}
	}

}

func (paths *Paths) DFS(s int) {

	paths.marked = make([]bool, len(paths.g.list))
	paths.edgeTo = make([]int, len(paths.g.list))

	paths.DFSUtil(s)
}

func (paths *Paths) DFSUtil(s int) {

	//fmt.Println("DFS visiting", s)
	paths.marked[s] = true // mark this node visited

	current := paths.g.adj(s).head
	for current != nil {
		if !paths.marked[current.Data.(int)] {
			paths.DFSUtil(current.Data.(int))
			paths.edgeTo[current.Data.(int)] = s
		}
		current = current.Next
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// find paths in G from source s
type Paths struct {
	g      *Graph
	s      int
	marked []bool
	edgeTo []int // edgeTo[v] = previous vertex on path from s to v
}

// print all vertices connected to s
func (paths *Paths) printPaths(g *Graph, s int) {

	for i := range g.list {
		if paths.hasPathTo(i) {
			fmt.Println(i)
		}
	}
}

// is there a path from s to v?
func (paths *Paths) hasPathTo(v int) bool {

	return paths.marked[v]
}

// path from s to v; null if no such path
func (paths *Paths) pathTo(v int) *list.List {
	if !paths.hasPathTo(v) {
		return nil
	}
	stack := list.New()
	for i := v; i != paths.s; i = paths.edgeTo[i] {
		stack.PushFront(i)
	}
	stack.PushFront(paths.s)
	return stack
}

func printPath(list *list.List) {
	fmt.Print("path ")
	for list.Len() != 0 {
		fmt.Print(list.Front().Value, " ")
		list.Remove(list.Front())
	}
	fmt.Println()
}

// Topological sorting - dependencies/sequencing tasks etc..
type DepthFirstOrder struct {
	marked []bool
	stack  *list.List // reverse DFS postorder
	g      *Graph
}

func (d *DepthFirstOrder) init() {
	if d.g == nil {
		return
	}
	d.marked = make([]bool, len(d.g.list))
	d.stack = list.New()
}

func (d *DepthFirstOrder) DFS(s int) {
	fmt.Println("DepthFirstOrder DFS visiting", s)
	d.marked[s] = true // mark this node visited

	current := d.g.adj(s).head
	for current != nil {
		if !d.marked[current.Data.(int)] {
			d.DFS(current.Data.(int))
		}
		current = current.Next
	}
	d.stack.PushFront(s)

}

// connected components using DFS
type CC struct {
	g      *Graph
	marked []bool
	cc     []int
}

func (c *CC) createCC() {
	c.marked = make([]bool, len(c.g.list))
	c.cc = make([]int, len(c.g.list))

	ccID := 0
	for i := range c.g.list {
		if !c.marked[i] {
			c.connectedDFS(i, ccID)
			ccID++
		}

	}
}

func (c *CC) connectedDFS(s int, ccID int) {

	c.marked[s] = true // mark this node visited
	c.cc[s] = ccID     // set the connected component ID
	//fmt.Println("CC", s, c.cc[s])
	current := c.g.adj(s).head
	for current != nil {
		if !c.marked[current.Data.(int)] {
			c.connectedDFS(current.Data.(int), ccID)
		}
		current = current.Next
	}
}

func (c *CC) connected(v int, w int) bool {
	return c.cc[v] == c.cc[w]
}

func main() {
	graph := &Graph{}
	graph.init(10)
	graph.addEdge(0, 1)
	graph.addEdge(0, 4)
	graph.addEdge(1, 2)
	graph.addEdge(1, 3)
	graph.addEdge(1, 4)
	graph.addEdge(2, 3)
	graph.addEdge(3, 4)
	graph.addEdge(5, 6)
	graph.addEdge(6, 7)
	graph.addEdge(8, 9)

	graph.print()
	// paths from source "3"
	paths := &Paths{g: graph, s: 3}
	paths.DFS(3)
	printPath(paths.pathTo(1))

	paths.BFS(3)

	topo := &DepthFirstOrder{g: graph}
	topo.init()
	topo.DFS(3)
	printPath(topo.stack)

	// create connected components
	cc := &CC{g: graph}
	cc.createCC()
	fmt.Println("Connected", cc.connected(1, 4))
	fmt.Println("degree", graph.degree(3))
	fmt.Println("maxDegree", graph.maxDegree())
}
