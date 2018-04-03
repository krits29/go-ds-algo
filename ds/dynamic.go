package main

import (
	"fmt"
)

func feb(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}
	return feb(n-1) + feb(n-2)
}

type Map map[int]uint64
// memoized - dynamic programming
func febD(m Map, n int) uint64 {
	if n <= 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	} else if v, ok := m[n]; ok {
		return v
	}
	m[n] = febD(m, n-1) + febD(m, n-2)
	return m[n]
}

func main() {
	
	go func(){
		m := make(Map)
		fmt.Println("memoized done", febD(m, 1000000))
	}()
	fmt.Println(feb(50))
}
