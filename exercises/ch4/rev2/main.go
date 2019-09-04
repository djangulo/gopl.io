// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev2 reverses a slice using an array pointer instead of a slice (exercise 4.3).
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		i   int = 1
		ii  int = 2
		iii int = 3
		iv  int = 4
		v   int = 5
	)
	//!+array
	a := [...]*int{&i, &ii, &iii, &iv, &v}
	reverse(a[:])
	fmt.Println(prettyPrint(a[:])) // "[5 4 3 2 1 0]"
	//!-array

	//!+slice
	s := []*int{&i, &ii, &iii, &iv, &v}
	// Rotate s left by two positions.
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(prettyPrint(s[:])) // "[2 3 4 5 0 1]"
	//!-slice

	// Interactive test of reverse.
	input := bufio.NewScanner(os.Stdin)
outer:
	for input.Scan() {
		var ints []*int
		for _, s := range strings.Fields(input.Text()) {
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue outer
			}
			xint := int(x)
			ints = append(ints, &xint)
		}
		reverse(ints)
		fmt.Printf("%v\n", prettyPrint(ints))
	}
	// NOTE: ignoring potential errors from input.Err()
}

func prettyPrint(in []*int) (str string) {
	str += "["
	for i := range in {
		switch i {
		case 0:
			str += fmt.Sprintf("%d", *in[i])
		case len(in) - 1:
			str += fmt.Sprintf(" %d", *in[i])
		default:
			str += fmt.Sprintf(" %d", *in[i])
		}
	}
	str += "]"
	return
}

//!+rev
// reverse reverses a slice of ints in place.
func reverse(s []*int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//!-rev
