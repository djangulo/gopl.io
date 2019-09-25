/*
Exercise 5.1: Change the findlinks program to traverse the n.FirstChild linked
list using recursive calls to visit instead of a loop
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

type link string

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks5.01: %v\n", err)
		os.Exit(1)
	}

	links := make([]*string, 0)
	visit(&links, doc)

	// linkChan := make(chan *string)
	// go visitConcur(linkChan, doc)
	for _, l := range links {
		fmt.Println(*l)
	}

	// for {
	// 	select {
	// 	case x := <-linkChan:
	// 		fmt.Println(*x)
	// 	}
	// }
	// wg.Wait()

}

// visit appends to links each link found in n and returns the result.
func visit(links *[]*string, n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				l := a.Val
				*links = append(*links, &l)
			}
		}
	}

	// var childLinks, siblingLinks []*string
	if n.FirstChild != nil {
		visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		visit(links, n.NextSibling)
	}
}

// visitConcur appends to links each link found in n and returns the result.
func visitConcur(linkChan chan<- *string, n *html.Node) {
	wg.Add(1)
	defer wg.Done()
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				l := a.Val
				linkChan <- &l
			}
		}
	}

	// var childLinks, siblingLinks []*string
	if n.FirstChild != nil {
		go visitConcur(linkChan, n.FirstChild)
	}
	if n.NextSibling != nil {
		go visitConcur(linkChan, n.NextSibling)
	}
}
