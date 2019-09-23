//Package update edits a github issue
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/djangulo/gopl.io/exercises/ch4/4.11/common"
	"github.com/pkg/errors"
)

const (
	baseURL = "https://api.github.com/repos"
	minArgs = 3 // minumum number of positional arguments
)

var (
	owner, repo, title, body, state string
	milestone, number               int
	labels, assignees               common.StringSliceFlag
	flagSet                         = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	// Usage prints out usage
	Usage func()
)

func init() {
	const (
		assigneesUsage = `Logins for Users to assign to this issue.
	NOTE: Only users with push access can set assignees for new issues.
	Assignees are silently dropped otherwise.`
		labelsUsage = `Labels to associate with this issue.
	NOTE: Only users with push access can set labels for new issues.
	Labels are silently dropped otherwise.`
		milestoneUsage = `The number of the milestone to associate this issue with.
	NOTE: Only users with push access can set the milestone for new
	issues. The milestone is silently dropped otherwise.`
		orderUsage = `Determines whether the first search result returned is the highest
	number of matches ('desc') or lowest number of matches ('asc').
	Default: 'desc'.`
		bodyUsage  = "The contents of the issue."
		titleUsage = "The title of the issue."
		stateUsage = "State of the issue. Either 'open' or 'closed'."
	)
	flagSet.StringVar(&title, "title", "", titleUsage)
	flagSet.StringVar(&title, "T", "", titleUsage+" (shorthand)")
	flagSet.StringVar(&body, "body", "", bodyUsage)
	flagSet.StringVar(&body, "B", "", bodyUsage+" (shorthand)")
	flagSet.StringVar(&state, "state", "", stateUsage)
	flagSet.StringVar(&state, "S", "", stateUsage+" (shorthand)")
	flagSet.IntVar(&milestone, "milestone", 0, milestoneUsage)
	flagSet.IntVar(&milestone, "M", 0, milestoneUsage+" (shorthand)")
	flagSet.Var(&labels, "labels", labelsUsage)
	flagSet.Var(&labels, "L", labelsUsage+" (shorthand)")
	flagSet.Var(&assignees, "assignees", assigneesUsage)
	flagSet.Var(&assignees, "A", assigneesUsage+" (shorthand)")
	usageStr := `Usage of %s:

update [OPTIONS] OWNER REPO ISSUE_NUMBER

  OWNER 	Required. Owner of the repository.
  REPO		Required. Repository to create the issue at.
  ISSUE_NUMBER	Required. The title of the issue.

  -T, -title string
	%s
  -B, -body string
	%s
  -S, -state string
	%s
  -M, -milestone int
	%s
  -L, -labels []string
	%s
  -A, -assignees []string
	%s

`
	Usage = func() {
		fmt.Fprintf(
			flagSet.Output(),
			usageStr,
			os.Args[0],
			titleUsage,
			bodyUsage,
			stateUsage,
			milestoneUsage,
			labelsUsage,
			assigneesUsage,
		)
	}
	flagSet.Usage = Usage
}

func main() {
	flagSet.Parse(os.Args[1:])
	if flagSet.NArg() < minArgs {
		flagSet.Usage()
		os.Exit(1)
	}
	owner = flagSet.Arg(0)
	repo = flagSet.Arg(1)
	number, err := strconv.Atoi(flagSet.Arg(2))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(owner, repo, number)

	var issue common.Issue
	if title != "" {
		issue.Title = title
	}
	if milestone != 0 {
		issue.Milestone = common.NullInt{Valid: true, Int: milestone}
	} else {
		issue.Milestone = common.NullInt{Valid: false}
	}
	if body != "" {
		issue.Body = body
	}
	if state != "" {
		issue.State = state
	}
	if len(labels) > 0 {
		var labs []common.Label
		for _, label := range labels {
			label := common.Label{Name: label}
			labs = append(labs, label)
		}
		issue.Labels = labs
	}
	if len(assignees) > 0 {
		issue.Assignees = assignees
	}
	fmt.Printf("%+v\n", assignees)

	err = UpdateIssue(&issue, owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(os.Stdout, "\nOK")
	bytes, err := json.MarshalIndent(issue, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(os.Stdout, string(bytes))

}

// UpdateIssue creates a new issue
func UpdateIssue(issue *common.Issue, owner, repo string, number int) error {
	// For long-term stability, instead of http.Get, use the
	// variant below which adds an HTTP request header indicating
	// that only version 3 of the GitHub API is acceptable.
	data, err := json.Marshal(issue)
	if err != nil {
		return errors.Wrap(err, "failed to marshal issue")
	}
	fmt.Println(string(data))
	req, err := http.NewRequest(
		"PATCH",
		fmt.Sprintf("%s/%s/%s/issues/%d", baseURL, owner, repo, number),
		bytes.NewBuffer(data),
	)
	if err != nil {
		return errors.Wrap(err, "failed to create POST request")
	}
	username, password := common.Credentials()
	req.SetBasicAuth(username, password)
	req.Header.Set(
		"Accept", "application/vnd.github.v3.text+json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to POST")
	}
	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("update failed: %s", resp.Status))
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return errors.Wrap(err, "failed to decode response")
	}
	return nil
}
