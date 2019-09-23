// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 110.
//!+

// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import (
	"flag"
	"time"
)

var (
	owner, repo string
	CommandLine = NewFlagSet(os.Args[0], ExitOnError)
	Usage       func()
)

func init() {
	const (
		ownerUsage = "owner of the repo"
		repoUsage  = "repo to create the issue"
	)
	flag.StringVar(&owner, "owner", "", ownerUsage)
	flag.StringVar(&owner, "o", "", ownerUsage+" (shorthand)")
	flag.StringVar(&owner, "repo", "", repoUsage)
	flag.StringVar(&owner, "r", "", repoUsage+" (shorthand)")

	usageStr = fmt.Sprintf(`Usage of %s
	issues ACTION [options]

	Where ACTION is one of
		N, new		create new issues
		R, read		read issues
		U, update	update existing issues
		C, close	close an issue

	N, new usage:
	issues new <owner> <repo> <title> -b "issue body"
		-o, -owner	Owner of the repo. Required.
		-r, -repo	Repository to create the issue for. Required.
		-t, -title	Issue title. Required.
		-b, -body	The contents of the issue.

	Options for read:

	`, os.Args[0])
}

func main() {
	flag.Parse()
	action := os.Args[1]

}

func validateAction(action string) error {
	action = strings.Lower(action)
	valid := map[string]struct{}{
		"new":    struct{}{},
		"N":      struct{}{},
		"read":   struct{}{},
		"R":      struct{}{},
		"update": struct{}{},
		"U":      struct{}{},
		"close":  struct{}{},
		"C":      struct{}{},
	}
	if _, ok valid[action]; !ok {
		return fmt.Errorf("'%s' is not a valid action", action)
	}
	return nil
}

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	// resp, err := http.Get(IssuesURL + "?q=" + q)
	// if err != nil {
	// 	return nil, err
	// }
	// For long-term stability, instead of http.Get, use the
	// variant below which adds an HTTP request header indicating
	// that only version 3 of the GitHub API is acceptable.

	req, err := http.NewRequest("GET", IssuesURL+"?q="+q, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(
		"Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

// CreateIssue queries the GitHub issue tracker.
func CreateIssue(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	// resp, err := http.Get(IssuesURL + "?q=" + q)
	// if err != nil {
	// 	return nil, err
	// }
	// For long-term stability, instead of http.Get, use the
	// variant below which adds an HTTP request header indicating
	// that only version 3 of the GitHub API is acceptable.

	req, err := http.NewRequest("GET", IssuesURL+"?q="+q, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(
		"Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

//!-
