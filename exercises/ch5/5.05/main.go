/*
Exercise 5.4:	Extend the visit function so that it extracts other kinds of
				links from the document, such as images, scripts, and style
				sheets.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks5.01: %v\n", err)
		os.Exit(1)
	}

	words := map[string]int{}
	images := countWordsAndImages(words, doc)

	for _, image := range images {
		fmt.Println(image)
	}

	for word, count := range words {
		fmt.Printf("%20s: %10d\n", word, count)
	}

}

var re = regexp.MustCompile(`\r?\n`)

func countWordsAndImages(words map[string]int, n *html.Node) (images []string) {

	if n.Type == html.ElementNode && n.Data == "script" || n.Data == "style" || n.Data == "link" {
		return
	}
	if n.Type == html.TextNode {
		d := re.ReplaceAllString(n.Data, "")
		r := strings.NewReader(d)
		scanner := bufio.NewScanner(r)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			text := strings.ToLower(scanner.Text())
			words[text]++
		}

	}

	if n.Type == html.ElementNode && n.Data == "img" {
		for _, a := range n.Attr {
			if a.Key == "src" {
				images = append(images, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		img := countWordsAndImages(words, c)
		images = append(images, img...)
	}
	return images

}
