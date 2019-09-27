// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	neturl "net/url"
	"os"
	"path/filepath"
	"strings"

	"gopl.io/ch5/links"
)

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//!-breadthFirst

var (
	gopath  = os.Getenv("GOPATH")
	dataDir = filepath.Join(
		gopath,
		"src",
		"github.com",
		"djangulo",
		"gopl.io",
		"exercises",
		"ch5",
		"5.13",
		"data",
	)
)

var (
	originalURLs []string
	checker      = map[string]bool{}
)

//!+crawl
func crawl(rawurl string) []string {

	for _, o := range originalURLs {
		if strings.HasPrefix(rawurl, o) {
			url, err := neturl.Parse(rawurl)
			if err != nil {
				panic(err)
			}
			dir, file := filepath.Split(url.Path)
			fullpath := filepath.Join(dataDir, url.Host, dir)
			os.MkdirAll(fullpath, os.ModeDir)

			var path string
			if file == "" {
				path = filepath.Join(fullpath, "index.html")
			} else {
				path = filepath.Join(fullpath, file)
			}

			fmt.Printf("saving %s => %s\n", rawurl, path)

			f, err := os.Create(path)
			if err != nil {
				panic(err)
			}
			defer f.Close()

			res, err := http.Get(rawurl)
			if err != nil {
				panic(err)
			}
			if res.StatusCode != http.StatusOK {
				fmt.Printf("received %d from %s", res.StatusCode, rawurl)
			}
			defer res.Body.Close()

			_, err = io.Copy(f, res.Body)
			if err != nil {
				panic(err)
			}
		}
	}

	list, err := links.Extract(rawurl)
	if err != nil {
		log.Print(err)
	}
	return list
}

func isOriginal(url string) bool {
	if checker[url] {
		return true
	}
	return false
}

//!+main
func main() {
	originalURLs = os.Args[1:]
	for _, url := range originalURLs {
		checker[url] = true
	}
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, originalURLs)
}

//!-main
