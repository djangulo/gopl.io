package main

import (
	"bufio"
	"bytes"
	"golang.org/x/net/html"
	"testing"
)

func TestPretty(t *testing.T) {
	t.Run("golang.org", func(t *testing.T) {
		out := new(bytes.Buffer)

		w := bufio.NewWriter(out)
		r := bufio.NewReader(out)
		rw := bufio.NewReadWriter(r, w)
		Outline(rw, "https://golang.org")

		_, err := html.Parse(rw)
		if err != nil {
			t.Errorf("didn't want an error but got one: %v", err)
		}
	})
}
