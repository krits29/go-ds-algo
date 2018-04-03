package main

// working with streams instead of arrays
// define two interface for both input and output stream
// Note: **************** untest code **************
type InputStream interface {
	poll() interface{}
}

type OutputStream interface {
	push(interface{})
}

type StreamNode struct {
	Val    interface{}
	stream InputStream
}

type StreamHeap struct {
	list []*StreamNode
	size int
}

func (s *StreamNode) isNotNull() bool {
	return s.Val != nil
}

func (h *StreamHeap) Size() int {
	return h.size

}

func (h *StreamHeap) poll() *StreamNode {
	tmp := h.list[0]
	//make sure to remove it from the heap and keep the size value updated
	return tmp

}
func (h *StreamHeap) buildHeap() {

	if h.list == nil || len(h.list) == 0 {
		return
	}
	// build from bottom half
	for i := len(h.list) / 2; i >= 0; i-- {
		h.minHeapify(i)
	}

}

func (h *StreamHeap) minHeapify(i int) {

	l := 2*i + 1
	r := 2*i + 2
	min := i

	if l < h.size && h.list[l].Val.(int) < h.list[min].Val.(int) {
		min = l
	}
	if r < h.size && h.list[r].Val.(int) < h.list[min].Val.(int) {
		min = r
	}

	if min != i {
		h.list[min], h.list[i] = h.list[i], h.list[min]
		h.minHeapify(min)
	}

}

func mergeStreams(list []InputStream, out OutputStream) {

	k := len(list)

	// create an array of StreamNode's
	arr := make([]*StreamNode, k)

	// put first element from each array
	for i := 0; i < k; i++ {
		arr[i] = &StreamNode{list[i].poll(), list[i]}
	}

	// create a heap with the given array
	heap := &StreamHeap{
		list: arr,
		size: k,
	}

	// build the min heap
	heap.buildHeap()

	// Now one by one get the minimum element from min
	// heap and replace it with next element of its array

	for heap.Size() > 0 {
		root := heap.poll() // get the min StreamNode, put it in a temp root
		out.push(root.Val)  // put it in the output Stream

		// make sure to delete the root

		new := &StreamNode{root.stream.poll(), root.stream}
		if new.isNotNull() {

			heap.list[0] = new
			heap.minHeapify(0)
		}

	}
}

func main() {

}
