/*
Exercise 5.2: 	Write a function to populate a mapping from element names -
				p, div, span and so on - to the number of elements with
				that name in an HTML document tree.
*/
package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/pkg/errors"
	"golang.org/x/net/html"
)

var (
	errNilNode     = errors.New("node is nil")
	errNoMoreNodes = errors.New("no more nodes")
	mu             sync.Mutex
	wg             sync.WaitGroup
)

type counter struct {
	elems map[string]int
	mu    sync.Mutex
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks5.02: %v\n", err)
		os.Exit(1)
	}

	var ctr counter
	ctr.elems = map[string]int{}
	mapElems(&ctr, doc)

	// linkChan := make(chan *string)
	// go visitConcur(linkChan, doc)
	for k, v := range ctr.elems {
		fmt.Printf("%10s %-d\n", k, v)
	}

	// for {
	// 	select {
	// 	case x := <-linkChan:
	// 		fmt.Println(*x)
	// 	}
	// }
	// wg.Wait()

}

// elems appends to links each link found in n and returns the result.
func mapElems(ctr *counter, n *html.Node) {
	if n.Type == html.ElementNode {
		ctr.mu.Lock()
		ctr.elems[n.Data]++
		ctr.mu.Unlock()
	}

	if n.FirstChild != nil {
		mapElems(ctr, n.FirstChild)
	}
	if n.NextSibling != nil {
		mapElems(ctr, n.NextSibling)
	}
}
