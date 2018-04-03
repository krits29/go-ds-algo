package main

import "container/list"
import "fmt"

//FIFO
func queueSample() {
	queue := list.New()
	queue.PushBack(1) // put at the end of the queue
	queue.PushBack(2)
	fmt.Println(queue.Front().Value) //  query the from of the queue
	queue.Remove(queue.Front())
	fmt.Println(queue.Front().Value) //  query the from of the queue
	queue.Remove(queue.Front())

	for queue.Len() != 0 {
		queue.Remove(queue.Front())
	}
}

//FILO
func stackSample() {
	stack := list.New()
	stack.PushBack(20)
	stack.PushBack(30)                      // put at the end of the stack
	fmt.Println("Back", stack.Back().Value) //  query the from of the stack
	fmt.Println("Front", stack.Front().Value)
	stack.Remove(stack.Back())
	fmt.Println(stack.Back().Value)
	stack.Remove(stack.Back())

	for stack.Len() != 0 {
		stack.Remove(stack.Back())
	}

}

func main() {
	fmt.Println("Hello")
	queueSample()
	stackSample()
}
