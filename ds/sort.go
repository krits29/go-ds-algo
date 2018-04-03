package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func QuickSort(s []int, low int, high int) {

	if high <= low {
		return
	}
	//fmt.Println(s)
	pivot := partition(s, low, high)

	QuickSort(s, low, pivot-1)
	QuickSort(s, pivot+1, high)

}

func partition(s []int, low int, high int) int {
	//take the first element
	i := low + 1
	j := high
	for { // break when the pointer crosses each other
		for s[i] < s[low] {
			i++
			if i == high {
				break
			}
		}

		for s[j] > s[low] {
			j--
			if j == low {
				break
			}
		}
		if i >= j {
			break
		}
		//swap the two
		s[i], s[j] = s[j], s[i]
	}
	// swap with the jth position
	s[low], s[j] = s[j], s[low]
	return j
}

func Merge(l, r []int) []int {
	ret := make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(ret, r...)
		}
		if len(r) == 0 {
			return append(ret, l...)
		}
		if l[0] <= r[0] {
			ret = append(ret, l[0])
			l = l[1:]
		} else {
			ret = append(ret, r[0])
			r = r[1:]
		}
	}
	return ret
}

//Plain merge sort performs log(n) = k iterations.
//On iteration i the algorithm merges 2^(k-i) blocks, each of size 2^i.

//Thus iteration i of merge sort performs:
//Calls to Less  O(2^(k-i) * 2^i) = O(2^k) = O(2^log(n)) = O(n)
//Calls to Swap  O(2^(k-i) * 2^i * log(2^i)) = O(2^k * i) = O(n*i)

//In total k = log(n) iterations are performed; so in total:
//Calls to Less O(log(n) * n)
//Calls to Swap O(n + 2*n + 3*n + ... + (k-1)*n + k*n)
//   = O((k/2) * k * n) = O(n * k^2) = O(n * log^2(n))
func MergeSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	n := len(s) / 2
	l := MergeSort(s[:n])
	r := MergeSort(s[n:])
	return Merge(l, r)
}

// sort a list containing duplicates 3-way quick sort
func SortDuplicates(list []int, lo int, hi int) {

	if hi <= lo {
		return
	}
	v := list[lo] // partitioning item
	lt := lo // keys on the left are less than the partition
	gt := hi // keys on the right are more than the partition and in between equal to the partition key
	i := lo
	for i <= gt {
		if list[i] < v { // Less than partitioning item
			list[i], list[lt] = list[lt], list[i] // Swap
			lt++
			i++
		} else if list[i] > v {
			list[i], list[gt] = list[gt], list[i]
			gt--
		} else {
			i++
		}
	}
	SortDuplicates(list, lo, lt-1)
	SortDuplicates(list, gt+1, hi)

}

// find kth element in an array
func QuickSelect(list []int, k int) int {

	lo := 0
	hi := len(list) - 1
	if hi <= lo || k >= len(list) {
		return -1
	}

	Shuffle(list)

	for hi > lo {
		pivot := partition(list, lo, hi)

		if pivot < k { // kth is on the right side
			lo = pivot + 1
		} else if pivot > k {
			hi = pivot - 1
		} else {
			return list[k]
		}
	}
	return list[k]
}
func Shuffle(a []int) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func TestSum(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	QuickSort(s, 0, len(s)-1)
	for i := range s {
		if s[i] != res[i] {
			t.Errorf("not equal")
		}

	}
}

func main() {
	s := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}
	fmt.Printf("%v\n%v\n", s, MergeSort(s))
	Shuffle(s)
	fmt.Println(s)
	QuickSort(s, 0, len(s)-1)
	fmt.Println(s)
	fmt.Println(QuickSelect(s, 8))
	d := []int{1, 3, 4, 7, 8, 1, 3, 4, 0, 7, 6, 4, 5, 8, 4, 1, 6}
	fmt.Println(d)
	SortDuplicates(d, 0, len(d)-1)
	fmt.Println(d)
}
