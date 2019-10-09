// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	reader := bytes.NewReader(p)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		*w++
	}

	return int(*w), nil
}

type LineCounter int

func (w *LineCounter) Write(p []byte) (int, error) {
	reader := bytes.NewReader(p)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		*w++
	}

	return int(*w), nil
}

func main() {
	//!+main
	var w WordCounter
	w.Write([]byte("hello darkness my old friend"))
	fmt.Println(w) // "5"

	w = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&w, "hello, %s", name)
	fmt.Println(w) // "2", = len("hello, Dolly")
	//!-main

	var l LineCounter
	l.Write([]byte(`one
four
three
four, i like counting up to four`))
	fmt.Println("lines should be 4: ", l)

}
