package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func printHeader(header string) {
	hashes := strings.Repeat("#", 30)
	dashes := strings.Repeat("-", 30)
	fmt.Printf("\n%s\n%s\n%s\n", hashes, header, dashes)
}

func timeWrapper(name, header string, fn func()) {
	printHeader("Exercise 1.1 Modify Echo to also print os.Args[0]")
	start := time.Now()
	fn()
	fmt.Printf("elapsed time for %s: %v\n", name, time.Since(start))
}

func Echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func Echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func Echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

// Echo4 Exercise 1.1 Modify Echo to also print os.Args[0]
func Echo4() {
	fmt.Println(strings.Join(os.Args, " "))
}

// Echo5 Exercise 1.2 Print indexes and values, one per line
func Echo5() {
	for i, v := range os.Args {
		fmt.Printf("index: %d\tvalue: %s\n", i, v)
	}
}

func main() {
	timeWrapper("Echo1", "", Echo1)
	timeWrapper("Echo2", "", Echo2)
	timeWrapper("Echo3", "", Echo3)
	timeWrapper("Echo4", "Exercise 1.1 Modify Echo to also print os.Args[0]", Echo4)
	timeWrapper("Echo5", "Exercise 1.2 Print indexes and values, one per line", Echo5)
}
