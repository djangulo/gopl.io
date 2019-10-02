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

	resp, err := http.Get(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	nodes := ElementsByTagName(doc, os.Args[2:]...)

	for _, node := range nodes {
		printNode(os.Stdout, node)
	}

}

var (
	depth int
	re    = regexp.MustCompile(`[\t\r\n]*`)
	re2   = regexp.MustCompile(`\b{2,}`)
)

func printNode(w io.Writer, n *html.Node) {
	startElement(w, n)
	endElement(w, n)
}

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
	if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
		data := re.ReplaceAllString(n.FirstChild.Data, "")
		data = re2.ReplaceAllString(data, "")
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

func ElementsByTagName(n *html.Node, name ...string) []*html.Node {
	nodes := make([]*html.Node, 0)

	if len(name) == 0 {
		return nodes
	}

	set := make(map[string]struct{}, 0)
	for _, n := range name {
		set[n] = struct{}{}
	}

	if n.Type == html.ElementNode {
		if _, ok := set[n.Data]; ok {
			nodes = append(nodes, n)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		next := ElementsByTagName(c, name...)
		nodes = append(nodes, next...)
	}

	return nodes

}
