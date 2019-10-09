// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 123.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"io"
	"os"
	"regexp"

	"golang.org/x/net/html"
)

//!+
func main() {

	in := `
<!DOCTYPE html>
<html>
 <head>
	<title>Sample html</title>
 </head>
<body>
	<div class=\"the-div\">
		<p>Some text</p>
	</div>
</body>
</html>
`

	r := newReader(in)

	outline(os.Stdout, r)
}

//!-

type reader struct {
	s        string
	i        int64
	prevRune int
}

func newReader(s string) *reader { return &reader{s, 0, -1} }

func (r *reader) Read(b []byte) (n int, err error) {

	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}

	r.prevRune = -1
	n = copy(b, r.s[r.i:])
	r.i += int64(n)

	return
}

func outline(w io.Writer, r io.Reader) error {
	doc, err := html.Parse(r)
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
