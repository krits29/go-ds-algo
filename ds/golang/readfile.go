package main

import (
	"bufio"
	"fmt"
	"os"
)

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
	}

}

func readFile(path string) {
	inFile, err := os.Open(path)
	handleError(err)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func writeFile(filename string) {

	fl, err := os.Create(filename)
	handleError(err)
	defer fl.Close()

	for i := 0; i < 100; i++ {
		_, err = fl.WriteString(fmt.Sprintf("%d\n", i)) // writing...
		if err != nil {
			fmt.Printf("error writing string: %v", err)
		}
	}

}

func main() {
	writeFile("newfile.txt")
	readFile("newfile.txt")
}
