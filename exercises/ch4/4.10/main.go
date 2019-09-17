// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 4.10. Modify 'issues' to report the results in age categories, say
// less than a month old, less than a year old and more than a year old

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"gopl.io/ch4/github"
)

var (
	OneYearAgo  time.Time
	OneMonthAgo time.Time
)

func init() {
	now := time.Now()
	OneYearAgo = now.AddDate(-1, 0, 0)
	OneMonthAgo = now.AddDate(0, -1, 0)
}

//!+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	// sort by date
	sort.SliceStable(result.Items, func(i, j int) bool { return result.Items[i].CreatedAt.Before(result.Items[j].CreatedAt) })

	now := time.Now()
	prunt := 0 // past tense of print
	for _, item := range result.Items {
		if now.Sub(item.CreatedAt) > now.Sub(OneYearAgo) {
			if prunt < 1 {
				fmt.Printf("More than a year old\n")
				prunt++
			}
		} else if now.Sub(item.CreatedAt) > now.Sub(OneMonthAgo) {
			if prunt < 2 {
				fmt.Printf("More than a month old\n")
				prunt++
			}
		} else {
			if prunt < 3 {
				fmt.Printf("Less than a month old\n")
				prunt++
			}
		}
		fmt.Printf("\t#%-5d %9.9s  %16s %.55s\n",
			item.Number, item.User.Login, item.CreatedAt.Format(time.UnixDate), item.Title)
	}
}

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/
