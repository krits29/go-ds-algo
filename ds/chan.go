package main

import "fmt"

func torun(c chan string) {
	fmt.Println(<-c)

}

func main() {
	c := make(chan string)

	num_cores := 1
	for i := 0; i < num_cores; i++ {
		go torun(c)
	}

	message := "x"

	for i := 0; i < 10000000; i++ {
		c <- message
	}
}
