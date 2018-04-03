package main

import (
	"fmt"
	"math"
)

type Tupil struct {
	X int
	Y int
	D float64
}

func printKClosest(list [][]int, k int) {

	if k > len(list) {
		return
	}
	tupils := getTupilList(list)
	createHeap(tupils[:k])

	for i := k; i < len(tupils); i++ {
		fmt.Println("max heap", tupils[0].D)
		if tupils[i].D < tupils[0].D {
			fmt.Println("swap",i)
			swap(tupils, i, 0)
			heapify(tupils, 0)
			printTupils(tupils)
		}
	}
	//printTupils(tupils[:k])
}

func printTupils(list []*Tupil) {

	for _, v := range list {
		fmt.Println(v.D)
	}
}

func createHeap(list []*Tupil) {
	if len(list) <= 0 {
		return
	}

	for i := len(list)/2; i >=0 ; i-- {
		heapify(list, i)
	}
	//printTupils(list)
}

func heapify(list []*Tupil, i int) {
	// base condition
	if i < 0 {
		return
	}
	l := 2*i + 1 // first child node
	r := 2*i + 2 // second child node
	max := i

	if l < len(list) && list[l].D > list[max].D {
		max = l
	}
	if r < len(list) && list[r].D > list[max].D {
		max = r
	}
	if max != i {
		swap(list, i, max)
		heapify(list, max)
	}

}

func swap(list []*Tupil, i, j int) {
	list[i], list[j] = list[j], list[i]
}

func getTupilList(list [][]int) []*Tupil {
	if len(list) <= 0 {
		return nil
	}

	res := []*Tupil{}
	var t *Tupil
	for _, r := range list {
		t = &Tupil{X: r[0], Y: r[1]}
		t.D = getDistance(t)
		res = append(res, t)
	}
	printTupils(res)
	return res
}

func getDistance(t *Tupil) float64 {
	return math.Sqrt(float64(t.X*t.X) + float64(t.Y*t.Y))
}

func main() {
	fmt.Println("Hello, playground")
	list := [][]int{{12, 2}, {25, 3}, {4, 5}, {6, 7}, {8, 9}}
	fmt.Println(list)
	printKClosest(list, 3)

}
