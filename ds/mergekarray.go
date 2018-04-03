package main

import (
	"fmt"
)

// Constant definitions
const MaxUint = ^uint(0) // XOR
const MinUint = 0

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

type Heap struct {
	list  []*HeapNode
	count int
}

type HeapNode struct {
	element int
	index   int // index of the array
	next    int // index of the next element in the array
}

func mergeKArrays(list [][]int) []int {

	k := len(list)
	n := len(list[0]) // length of the array's
	res := make([]int, k*n)

	// create an array of HeapNode
	arr := make([]*HeapNode, k)

	// put first element from each array
	for i := 0; i < k; i++ {
		arr[i] = &HeapNode{list[i][0], i, 1}
	}

	// create a heap with the given array
	heap := &Heap{
		list:  arr,
		count: k,
	}

	// build the min heap
	heap.buildMinHeap()

	// Now one by one get the minimum element from min
	// heap and replace it with next element of its array

	for i := 0; i < n*k; i++ {
		root := heap.list[0]  // get the min element, put it in a temp root
		res[i] = root.element // put it in the output array

		if root.next < n {
			root.element = list[root.index][root.next] // put the next element from the array
			root.next += 1
		} else { // end of array
			root.element = MaxInt
		}
		heap.list[0] = root
		heap.minHeapify(0)

	}

	return res
}

func (h *Heap) buildMinHeap() {

	if h.list == nil || len(h.list) == 0 {
		return
	}
	// build from bottom half
	for i := h.count / 2; i >= 0; i-- {
		h.minHeapify(i)
	}

}

func (h *Heap) minHeapify(i int) {

	l := 2*i + 1
	r := 2*i + 2
	min := i

	if l < h.count && h.list[l].element < h.list[min].element {
		min = l
	}
	if r < h.count && h.list[r].element < h.list[min].element {
		min = r
	}

	if min != i {
		h.list[min], h.list[i] = h.list[i], h.list[min]
		h.minHeapify(min)
	}

}

func (h *Heap) maxHeapify(i int) {

	l := 2*i + 1
	r := 2*i + 2
	max := i

	if l < h.count && h.list[l].element > h.list[max].element {
		max = l
	}
	if r < h.count && h.list[r].element > h.list[max].element {
		max = r
	}

	if max != i {
		h.list[max], h.list[i] = h.list[i], h.list[max]
		h.maxHeapify(max)
	}

}

func main() {
	arr := [][]int{{2, 6, 12, 34, 45},
		{1, 9, 20, 1000, 1234},
		{23, 34, 90, 2000, 2344}}

	res := mergeKArrays(arr)
	fmt.Println(res)

}
