package main

import (
	"fmt"
	"io"
	"os"
)

type countWriter struct {
	w io.Writer
	c int64
}

func (c *countWriter) Write(p []byte) (int, error) {
	n, err := (*c).w.Write(p)
	c.c += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := countWriter{w, 0}
	return &c, &c.c
}

func main() {
	w, c := CountingWriter(os.Stdout)

	fmt.Fprint(w, "hey there answer\n")
	fmt.Println(*c) //  17

	fmt.Fprint(w, "one two\n")
	fmt.Println(*c) // 17 + 8 = 25
}
