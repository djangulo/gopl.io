// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		Outline(os.Stdout, url)
	}
}

func Outline(w io.Writer, url string) error {
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
	forEachNode(w, doc, startElement, endElement)
	//!-call

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(
	w io.Writer,
	n *html.Node,
	pre, post func(w io.Writer, n *html.Node),
) {
	if pre != nil {
		pre(w, n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(w, c, pre, post)
	}

	if post != nil {
		post(w, n)
	}
}

//!-forEachNode

//!+startend
var (
	depth int
	re    = regexp.MustCompile(`[\t\r]*`)
)

func startElement(w io.Writer, n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Fprintf(w, "%*s<%s", depth*2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Fprintf(w, " %s=\"%s\"", a.Key, a.Val)
		}
		if n.FirstChild == nil && n.Data != "script" {
			fmt.Fprint(w, " />\n")
			return
		}
		fmt.Print(">")
		depth++
	}
	if n.Type == html.TextNode {
		data := re.ReplaceAllString(n.Data, "")
		fmt.Fprintf(w, "%*s%s", depth*2, "", data)
	}
}

func endElement(w io.Writer, n *html.Node) {
	if n.Type == html.ElementNode {
		if n.FirstChild == nil && n.Data != "script" {
			return
		}
		depth--
		fmt.Fprintf(w, "%*s</%s>\n", depth*2, "", n.Data)
	}
}

//!-startend
