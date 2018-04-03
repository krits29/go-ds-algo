package main

import (
	"fmt"
)

type Heap struct {
	list  []int
	count int
}
// Max heap
func (h *Heap) buildHeap() {

	if h.list == nil || len(h.list) == 0 {
		return
	}
	// build from bottom half
	for i := h.count / 2; i >= 0; i-- {
		h.heapify(i)
	}

}

// both the child nodes are smaller than the parent node
func (h *Heap) heapify(i int) {

	l := 2*i + 1
	r := 2*i + 2
	max := i

	if l < h.count && h.list[l] > h.list[max] {
		max = l
	}
	if r < h.count && h.list[r] > h.list[max] {
		max = r
	}

	if max != i {
		h.list[max], h.list[i] = h.list[i], h.list[max]
		h.heapify(max)
	}
}

func (h *Heap) heapSort() {

	for i := len(h.list) - 1; i >= 1; i-- {
		// swap the first one with the last
		h.list[0], h.list[i] = h.list[i], h.list[0]
		h.count--
		h.heapify(0)
	}

}

func main() {

	l := []int{2, 4, 6, 12, 43, 0, 9}
	heap := &Heap{
		list:  l,
		count: len(l),
	}
	fmt.Println(heap.list)
	heap.buildHeap()
	fmt.Println(heap.list)
	heap.heapSort()
	fmt.Println(heap.list)
}
