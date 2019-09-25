/*
Exercise 5.3:	Write a function to print the contents of all text nodes
				in an HTML document tree. Do not descend into <script> or
				<style> elements, since tehir contents are not visible
				in a web browser
*/
package main

import (
	"fmt"
	"os"
	"regexp"
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

	str := make([]*string, 0)

	get(&str, doc)

	for _, x := range str {
		fmt.Printf("%s\n", *x)
	}

}

var re = regexp.MustCompile(`\r?\n|<br>`)

// visitConcur appends to links each link found in n and returns the result.
func get(str *[]*string, n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "script" || n.Data == "style" || n.Data == "link" {
		return
	}
	if n.Type == html.TextNode {
		d := re.ReplaceAllString(n.Data, "")
		*str = append(*str, &d)
	}

	// var childLinks, siblingLinks []*string
	if n.FirstChild != nil {
		get(str, n.FirstChild)
	}
	if n.NextSibling != nil {
		get(str, n.NextSibling)
	}
}
