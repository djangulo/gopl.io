// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"fmt"
)

func main() {
	//!+array
	a := []byte("Im a pretty little pony")
	reverse(a)
	fmt.Println(string(a)) // "[5 4 3 2 1 0]"
	//!-array

	//!+slice
	s := []byte("I enjoy evening walks.")
	// Rotate s left by two positions.
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(string(s)) // "[2 3 4 5 0 1]"
	//!-slice

}

//!+rev
// reverse reverses a slice of ints in place.
func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//!-rev
