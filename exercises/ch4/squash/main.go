package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("%s\n%s\n%s\n\n",
		"Exercise 4.6: Write an in-place function that squashes",
		" each run of adjacent unicode spaces (see unicode.IsSpace)",
		" in a UTF-8-encoded []byte slice into a single ASCII space.",
	)
	in := []byte("many    spaces    into one")
	on := squashSpaces(in)
	fmt.Println(string(on))
}

func squashSpaces(in []byte) []byte {
	j := 0
	for i := 1; i < len(in); i++ {
		if in[j] == in[i] && unicode.IsSpace(rune(in[j])) && unicode.IsSpace(rune(in[i])) {
			continue
		}
		j++
		in[j] = in[i]
	}
	return in[:j+1]
}
