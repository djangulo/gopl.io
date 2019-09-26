package main

import (
	"fmt"
	"github.com/djangulo/gopl.io/exercises/ch5/5.09/expand"
	"os"
	"strings"
)

func main() {
	input := strings.Join(os.Args[1:], " ")
	fmt.Println(expand.Expand(input, os.Getenv))

}
