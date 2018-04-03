package main

import (
	"fmt"
	"testing"
)

func feb(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}
	return feb(n-1) + feb(n-2)
}

func Testfeb(t *testing.T) bool {

	ret := feb(10)
	if ret != 55 {
		return t.Failed()
	}
	return true
}

func main() {
	fmt.Println(feb(5))
}
