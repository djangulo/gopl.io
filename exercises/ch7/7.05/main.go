package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	in := "This is a sample string"
	r := strings.NewReader(in)

	lr := io.LimitReader(r, 11)

	buf := make([]byte, len(in))

	lr.Read(buf)

	fmt.Println(string(buf)) // This is a s
}
