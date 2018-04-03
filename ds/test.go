package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sumArray(n int, arr []int) int {

	sum := 0
	for _, r := range arr {
		sum += r
	}

	return sum
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var input string
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')      // read until newline
	input = strings.TrimSuffix(input, "\n") // trim trailing newline
	n, err := strconv.Atoi(input)           // read the length
	if err != nil {
		// handle error
		fmt.Println("enter number", err)
		os.Exit(3)
	}
	//reader = bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	commands := strings.Split(input, " ") // get all values
	arr := make([]int, n)                 // create the int array

	for _, r := range commands {
		num, err := strconv.Atoi(r)
		if err != nil {
			// handle error
			fmt.Println("enter number")
			os.Exit(3)
		}
		arr = append(arr, num)
	}

	fmt.Println(sumArray(n, arr))
}
