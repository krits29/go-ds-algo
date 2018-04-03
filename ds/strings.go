package main

import (
	"fmt"
	"strings"
	"unicode"
)

func firstUniqChar(s string) int {

	// create a map of byte-count(int)
	list := make(map[byte]int)

	// go over the given string, increase the count in the map
	for i, _ := range s {
		list[s[i]] = list[s[i]] + 1
	}

	// go over the map again to see which one has only one count
	for i, _ := range s {
		if list[s[i]] == 1 {
			return i
		}
	}
	return -1
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseWords(str []byte) []byte {
	fmt.Println(string(str))
	words := strings.Split(string(str), " ")
	len := len(words)
	for i := 0; i < len/2; i++ {
		words[len-i-1], words[i] = words[i], words[len-i-1]
	}
	return []byte(strings.Join(words, " "))

}

func longestPalindrome(str string) string {
	res := ""
	curr := 0
	for i := 0; i < len(str); i++ {
		if isPalindrome(str, i-curr-1, i) {
			res = str[i-curr-1 : i+1]
			curr = curr + 2
		} else if isPalindrome(str, i-curr, i) {
			res = str[i-curr : i+1]
			curr = curr + 1
		}
	}
	return res
}

func isPalindrome(str string, begin int, end int) bool {
	if begin < 0 {
		return false
	}
	for begin < end {
		if str[begin] != str[end] {
			return false
		}
		begin++
		end--
	}
	return true
}

func isPalindromeAll(s string) bool {

	l := len(s)
	if l <= 1 {
		return true
	}
	s = strings.ToLower(s)
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {

		for i < j && !unicode.IsLetter(rune(s[i])) && !unicode.IsDigit(rune(s[i])) {
			i++
		}
		if i == j {
			break
		}
		for i < j && !unicode.IsLetter(rune(s[j])) && !unicode.IsDigit(rune(s[j])) {
			j--
		}
		if i == j {
			break
		}
		if s[i] != s[j] {
			return false
		}

	}
	return true
}

//Given an array of n integers where n > 1, nums,
//return an array output such that output[i] is equal to
//the product of all the elements of nums except nums[i].
// multiple everything on the left and everything on the right
func productExceptSelf(nums []int) []int {

	output := make([]int, len(nums))
	tmp := 1
	for i := 0; i < len(nums); i++ {
		output[i] = tmp
		tmp *= nums[i] // store uptil the previous product
	}
	tmp = 1
	for i := len(nums) - 1; i >= 0; i-- {
		output[i] *= tmp // store uptil the previous product and multipy with the existing value
		tmp *= nums[i]
	}
	return output
}

func compareVersion(version1 string, version2 string) int {

	i, j := 0, 0
	// go until the longest string
	for i < len(version1) || j < len(version2) {
		cur1, cur2 := 0, 0

		// stop if there is a '.' - get the first major version and so on...
		for ; i < len(version1) && version1[i] != '.'; i++ {

			cur1 = cur1*10 + int(version1[i]) - '0'
		}
		// stop if there is a '.' - get the first major version and so on...
		for ; j < len(version2) && version2[j] != '.'; j++ {

			cur2 = cur2*10 + int(version2[j]) - '0'
		}
		// compare the versions now
		if cur1 > cur2 {
			return 1
		} else if cur1 < cur2 {
			return -1
		} else { // if same...continue to the lower versions
			i++
			j++
		}
	}
	return 0
}

func numberToWords(num int) string {
	to19 := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve",
		"Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	tens := []string{"Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}

	switch {
	case num == 0:
		return "Zero"
	case num < 20:
		return strings.Join(to19[num-1:num], " ")
	case num < 100:
		n := num / 10
		return strings.Join(tens[n-2:n-1], " ") + " " + numberToWords(num%10)
	case num < 1000:
		n := num / 100
		return strings.Join(to19[n-1:n], " ") + " hundred " + numberToWords(num%100)
	default:
		return ""
	}
	return ""
}

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}
	s = strings.ToLower(s)
	t = strings.ToLower(t)
	counts := make([]int, 26)

	for i := range s {
		counts[s[i]-'a']++
		counts[t[i]-'a']--
	}

	for _, c := range counts {
		if c != 0 {
			return false
		}
	}
	return true
}

func IsAlphaNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) || !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// []byte(str) string --> byte[]
// string([]byte) byte[] -->str

// runes := []rune(str)

//string [] -> string
// strings.Join(words, " ") // exlicit separator

//string -->string[]
// strings.Split(string(str), " ") // explicit separator
// strings.Fields("Foo bar baz") // default space

/* random string --> string[] based on a function
f:= func(c rune) bool {
	return !unicode.IsLetter(c) && !unicode.IsNumeric(c)
}
strings.FieldsFunc("foo1, foo2; foo3, foo4", f)
*/

/*concatinate strings byte.Buffer
var buffer bytes.Buffer
buffer.writeString("x")
buffer.String()
*/

/*concatinate strings using copy
b := make([]byte)
bool := 0 // count
bl = copy(b[bl:], "x")
return string(b)
*/

// sort strings 3 way string quicksort
// 3-way partitioning using dth character
func SortStrings(list []string, lo int, hi int, d int) {

	if hi <= lo {
		return
	}
	lt := lo
	gt := hi
	v := list[lo][d]
	i := lo + 1
	for i <= gt {
		if list[i][d] < v { // pick out the dth character in the partition key - list[lt]
			list[i], list[lt] = list[lt], list[i]
			lt++
			i++
		} else if list[i][d] > v {
			list[i], list[gt] = list[gt], list[i]
			gt--
		} else {
			i++
		}
	}
	SortStrings(list, lo, lt-1, d)
	if v >= 0 {
		SortStrings(list, lt, gt, d+1)
	}
	SortStrings(list, gt+1, hi, d)

}

// string search
// brute force = len(str) x len(pat) compares
// not good if the data is a stream
func strStr1(str string, pat string) int {
	// go over the entire string
	i, j := 0, 0
	for ; i < len(str) && j < len(pat); i++ {

		if str[i] == pat[j] { // if not found break and continue
			j++
		} else {
			i -= j // explicit backup
			j = 0  // reset j
		}

	}
	if j == len(pat) {
		return i - len(pat)
	}
	return -1 // string not found
}

// Boyer-Moore  about ~ N / M character compares
// Precompute index of rightmost occurrence of the characters in pattern - array called 'right'
// Worst-case. Can be as bad as ~ M N
// but simple one pass through the pattern to build the initial array
func strStr2(str string, pat string) int {

	N := len(str)
	M := len(pat)
	R := 256

	if N < M { // boundary check
		return -1
	}

	right := make([]int, R)
	//Precompute index of rightmost occurrence of character c in pattern
	// -1 if char not in pattern
	for i := 0; i < R; i++ {
		right[i] = -1
	}
	for j := 0; j < M; j++ {
		right[pat[j]] = j
	}

	var skip int

	for i := 0; i < N-M; i += skip { // skip computed dynamically
		//reset skip
		skip = 0
		for j := M - 1; j >= 0; j-- { // go from right to left
			if pat[j] != str[i+j] {
				// compute skip value
				skip = max(1, j-right[str[i+j]]) // if its -1, we don't backup but skip by 1
				break
			}
		}
		if skip == 0 { // match
			return i
		}
	}
	return -1
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}
func main() {
	s := []byte("The sun did not shine. It was too wet to play.")
	fmt.Println(firstUniqChar("loveleetcode"))
	fmt.Println(string(reverseWords(s)))
	fmt.Println(strings.Fields("Foo bar baz"))
	fmt.Println(strStr2("hello how are you", "how are y"))
	str := []string{"hello", "give", "hi", "how", "am"}
	SortStrings(str, 0, len(str)-1, 0)
	fmt.Println(str)
}
