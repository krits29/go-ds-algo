package main

import (
	"bytes"
	"container/list"
	"fmt"
	"io"
)

func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

func StreamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
}

// A Dequeue (Double ended queue) based method for printing maximum element of
// all subarrays of size k
func getMaxSlidingWindow(arr []int, k int) int {

	if len(arr) <= 0 {
		return -1
	}
	// elements are stored in a decreasing order
	deque := list.New() // add index's

	/* Process first k (or first window) elements of array */
	for i := 0; i < k; i++ {
		// For very element, the previous smaller elements are useless so remove them from deque
		for deque.Len() != 0 && arr[i] > arr[deque.Back().Value.(int)] {
			deque.Remove(deque.Back())
		}
		// Add new element at rear of queue
		deque.PushBack(i)
	}
	// Process rest of the elements, i.e., from arr[k] to arr[n-1]

	for i := k; i < len(arr); i++ {

		// The element at the front of the queue is the largest element of
		// previous window, so print it
		fmt.Println(arr[deque.Front().Value.(int)])

		// Remove the elements which are out of this window
		for deque.Len() != 0 && deque.Front().Value.(int) <= i-k {
			deque.Remove(deque.Front())

		}
		// Remove all elements smaller than the currently
		// being added element (remove useless elements)
		for deque.Len() != 0 && arr[i] >= arr[deque.Back().Value.(int)] {
			deque.Remove(deque.Back())

		}

		deque.PushBack(i)
	}
	fmt.Println(arr[deque.Front().Value.(int)])
	if deque.Len() != 0 {
		return deque.Front().Value.(int)
	}
	return -1
}

func main() {
	fmt.Println("hello sliding window")
	list := []int{12, 1, 78, 90, 57, 89, 56}
	getMaxSlidingWindow(list, 3)

}
