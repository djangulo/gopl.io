//Package read reads github issues
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/djangulo/gopl.io/exercises/ch4/4.11/common"
)

const issuesURL = "https://api.github.com/search/issues"

var (
	sort, order      string
	flagSet          = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	validSortOptions = map[string]struct{}{
		"comments":                struct{}{},
		"reactions":               struct{}{},
		"reactions-+1":            struct{}{},
		"reactions--1":            struct{}{},
		"reactions-smile":         struct{}{},
		"reactions-thinking_face": struct{}{},
		"reactions-heart":         struct{}{},
		"reactions-tada":          struct{}{},
		"interactions":            struct{}{},
		"created":                 struct{}{},
		"updated":                 struct{}{},
	}
	validOrderOptions = map[string]struct{}{
		"asc":  struct{}{},
		"desc": struct{}{},
	}
	Usage func()
)

func init() {
	const (
		sortUsage = `Sorts the results of your query by the number of 'comments', 'reactions',
	'reactions-+1', 'reactions--1', 'reactions-smile', 'reactions-thinking_face',
	'reactions-heart', 'reactions-tada', or 'interactions'. You can also sort
	results	by how recently the items were 'created' or 'updated'.
	Default: 'created'.`
		orderUsage = `Determines whether the first search result returned is the highest
	number of matches ('desc') or lowest number of matches ('asc').
	Default: 'desc'.`
	)
	flagSet.StringVar(&sort, "sort", "created", sortUsage)
	flagSet.StringVar(&sort, "s", "created", sortUsage+" (shorthand)")
	flagSet.StringVar(&order, "order", "desc", orderUsage)
	flagSet.StringVar(&order, "o", "desc", orderUsage+" (shorthand)")

	usageStr := `Usage of %s:

read [OPTIONS] QUERIES[..]

  e.g.
	read repo:golang/go is:open json decoder

  QUERIES 	Required. Space separated list of query parameters.

  -S, -sort string
	%s
  -O, -order string
	%s

`

	Usage = func() {
		fmt.Fprintf(
			flagSet.Output(),
			usageStr,
			os.Args[0],
			sortUsage,
			orderUsage,
		)
	}
	flagSet.Usage = Usage
}

func main() {
	flagSet.Parse(os.Args[1:])
	if flagSet.NArg() < 1 {
		flagSet.Usage()
		os.Exit(1)
	}
	err := validateSort(sort)
	if err != nil {
		log.Fatal(err)
	}
	err = validateOrder(order)
	if err != nil {
		log.Fatal(err)
	}

	result, err := SearchIssues(flagSet.Args())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

}

func validateSort(sort string) error {
	if sort != "" {
		if _, ok := validSortOptions[sort]; !ok {
			return fmt.Errorf("invalid sort option '%s'", sort)
		}
	}
	return nil
}

func validateOrder(order string) error {
	if order != "" {
		if _, ok := validOrderOptions[order]; !ok {
			return fmt.Errorf("invalid order option '%s'", order)
		}
	}
	return nil
}

// IssuesSearchResult result type for issue search
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*common.Issue
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	// For long-term stability, instead of http.Get, use the
	// variant below which adds an HTTP request header indicating
	// that only version 3 of the GitHub API is acceptable.
	req, err := http.NewRequest(
		"GET",
		issuesURL+"?q="+q+"&sort="+sort+"&order="+order,
		nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(
		"Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("failed to get response: %v", err)
	}
	defer resp.Body.Close()

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	// bod, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(bod))

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
