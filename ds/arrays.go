package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func isArraySymretric(arr []int) bool {
	l := len(arr)
	if l%2 != 0 {
		return false
	}
	for i := 0; i < l/2; i++ {
		if arr[i] != arr[l-1-i] {
			return false
		}
	}
	return true
}

// rotate left
func rotateLeft(nums []int, k int) {
	l := len(nums)
	k %= l
	k = l - k
	reverse(nums, 0, l-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, l-1)
}

// rotate right
func rotateRight(nums []int, k int) {
	l := len(nums)
	k %= l
	reverse(nums, 0, l-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, l-1)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// find two number sum to target in the given array
func twoSum(nums []int, target int) []int {

	list := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		comp := target - nums[i]
		// if the other half is present in the list - return
		if v, ok := list[comp]; ok {
			return []int{v, i}
		}
		//if not add the current number to the map
		list[nums[i]] = i
	}
	return nil
}

// find count of pairs with given sum
func getPairCounts(nums []int, target int) int {

	list := make(map[int]int)

	for i := 0; i < len(nums); i++ {

		if _, ok := list[nums[i]]; !ok {
			list[nums[i]] = 0
		}
		list[nums[i]] = list[nums[i]] + 1
	}
	var count int
	for i := 0; i < len(nums); i++ {
		other := target - nums[i]
		if r, ok := list[other]; ok { // other half present
			count += r
		}
		if other == nums[i] { // decrement if the same
			count--
		}
	}

	return count / 2
}

func threeSum(nums []int) [][]int {

	res := [][]int{}
	// if array size is less than 3 return empty list
	if len(nums) < 3 {
		return res
	}
	// lets sort the array first
	sort.Ints(nums)
	fmt.Println(nums)

	// loop through the entire array, stop if only 2 are left
	for i := 0; i < len(nums)-2; i++ {

		// take the next and the last
		j := i + 1
		k := len(nums) - 1
		// loop until the two pointer crosses each other
		for j < k {
			// check the sum of all three
			sum := nums[i] + nums[j] + nums[k]
			//fmt.Println("SUM", sum)
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				//res = append(res, make([]int, {nums[i] + nums[j] + nums[k]}))
				j++
				k--
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}

	return res
}

func findDuplicate(nums []int) int {

	if len(nums) <= 1 {
		return -1
	}

	for i, n := range nums {
		for j, d := range nums {
			if i != j && n == d {
				return n
			}
		}
	}
	return -1
}

func TestfindDuplicate(t *testing.T) {
	dup := findDuplicate([]int{1, 2, 3, 4, 56, 3, 5, 3, 7})
	if dup == -1 {
		t.Error("Expected Duplicates")
	}
}

// Remove Duplicates from Sorted Array
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// To store index of next unique element
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[j] != nums[i] { // increment j when you find a new number and replace
			j++
			if i != j { // replace jth position by the new number
				nums[j] = nums[i]
			}
		}
	}
	nums = nums[:j+1]
	fmt.Println(nums)
	return j + 1
}

// Quick - Select method
func findKthLargest(list []int, k int) int {

	lo := 0
	hi := len(list) - 1
	if hi < lo || k > len(list) {
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

//Given an array nums and a target value k,
//find the maximum length of a subarray that sums to k. If there isn't one, return 0 instead.
//Given nums = [1, -1, 5, -2, 3], k = 3,
//return 4. (because the subarray [1, -1, 5, -2] sums to 3 and is the longest)
func maxSubArrayLen(nums []int, k int) int {

	mp := map[int]int{0: -1} // initialize map with 0:-1

	count, sum := 0, 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		// look for the difference in sum so far, if its not e
		if index, ok := mp[sum-k]; ok {
			if i-index > count {
				count = i - index
			}
		}
		// add the sum to the map if it doesn't exists already
		if _, ok := mp[sum]; !ok {
			mp[sum] = i
		}
	}
	return count
}

// A Dequeue (Double ended queue) based method for printing maixmum element of
// all subarrays of size k
func getMaxSlidingWindow(arr []int, k int) int {

	if len(arr) <= 0 {
		return -1
	}

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
	//arr := []int{0, 1, 2, 3, 4, 5, 6}
	arr := []int{1, 2, 2, 3, 4, 4, 4, 5, 5}
	fmt.Println("final count", removeDuplicates(arr))
	rotateLeft(arr, 4)
	rotateRight(arr, 4)
	fmt.Println(arr)

	fmt.Println(getPairCounts(arr, 6))
}
