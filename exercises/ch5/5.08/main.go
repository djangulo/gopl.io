// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

/* Modify forEachNode so that the pre and post functions return a boolean resulrt
indicating whether to continue the traversal. Use it to write a function ElementByID
with the following signature that finds the first HTML element with the pecified id
attribute. The function should stop the traversal as soon as a match is found.

	func ElementByID(doc *html.Node, id string) *html.Node
*/

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var (
	id string
)

func main() {
	id = os.Args[2]
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	forEachNode(doc, thePre, nil)
	//!-call

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	var keepGoing = true
	if pre != nil {
		keepGoing = pre(n)
	}
	if !keepGoing {
		printTheThing(os.Stdout, n)
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		keepGoing = post(n)
	}
	if !keepGoing {
		printTheThing(os.Stdout, n)
		return
	}
}

func printTheThing(out io.Writer, n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Fprintf(out, "<%s", n.Data)
	}
	for _, a := range n.Attr {
		fmt.Fprintf(out, " %s=\"%s\"", a.Key, a.Val)

	}
	fmt.Fprint(out, ">")
	if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
		fmt.Fprintf(out, "\n\t%s\n", n.FirstChild.Data)
	}
	if n.Type == html.ElementNode {
		fmt.Fprintf(out, "</%s>\n", n.Data)
	}
}

func ElementByID(doc *html.Node, id string) *html.Node {
	for _, a := range doc.Attr {
		if a.Key == "id" && a.Val == id {
			return doc
		}
	}
	return nil
}

func thePre(n *html.Node) bool {
	test := ElementByID(n, id)
	return test == nil
}

//!-forEachNode

//!+startend
var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

//!-startend
