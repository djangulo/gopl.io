/*
Exercise 5.4:	Extend the visit function so that it extracts other kinds of
				links from the document, such as images, scripts, and style
				sheets.
*/
package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/pkg/errors"
	"golang.org/x/net/html"
)

const (
	anchors     = "anchors"
	images      = "images"
	stylesheets = "stylesheets"
	scripts     = "scripts"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks5.01: %v\n", err)
		os.Exit(1)
	}

	links := map[string][]string{
		anchors:     make([]string, 0),
		images:      make([]string, 0),
		stylesheets: make([]string, 0),
		scripts:     make([]string, 0),
	}

	visit(links, doc)

	for _, t := range []string{anchors, images, stylesheets, scripts} {
		fmt.Println(t)
		for _, l := range links[t] {
			fmt.Printf("\t%s\n", l)
		}

	}

}

// visit appends to links each link found in n and returns the result.
func visit(links map[string][]string, n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links[anchors] = append(links[anchors], a.Val)
			}
		}
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		for _, a := range n.Attr {
			if a.Key == "src" {
				links[images] = append(links[images], a.Val)
			}
		}
	}

	if n.Type == html.ElementNode && n.Data == "link" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links[stylesheets] = append(links[stylesheets], a.Val)
			}
		}
	}

	if n.Type == html.ElementNode && n.Data == "script" {
		for _, a := range n.Attr {
			if a.Key == "src" {
				links[scripts] = append(links[scripts], a.Val)
			}
		}
	}

	if n.FirstChild != nil {
		visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		visit(links, n.NextSibling)
	}
}
