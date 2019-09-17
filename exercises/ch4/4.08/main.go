// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

type countss map[string]map[rune]int

func (c Counts) String() string {
	var buf bytes.Buffer
	for cat := range c {
		fmt.Fprintf(&buf, "%s\n", cat)
		fmt.Fprintf(&buf, "\trune\tcount\n")
		for c, n := range c[cat] {
			fmt.Fprintf(&buf, "\t%q\t%d\n", c, n)
		}
	}
	return buf.String()
}

func main() {
	counts := countss{
		"simbols": map[rune]int{},
		"spaces":  map[rune]int{},
		"puncts":  map[rune]int{},
		"marks":   map[rune]int{},
		"letters": map[rune]int{},
		"numbers": map[rune]int{},
	} // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == '1' {
			break
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		switch {
		case unicode.IsSymbol(r):
			counts["simbols"][r]++
		case unicode.IsSpace(r):
			counts["spaces"][r]++
		case unicode.IsPunct(r):
			counts["puncts"][r]++
		case unicode.IsMark(r):
			counts["marks"][r]++
		case unicode.IsLetter(r):
			counts["letters"][r]++
		case unicode.IsNumber(r):
			counts["numbers"][r]++
		}
		utflen[n]++
	}
	fmt.Println(counts)
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

//!-
