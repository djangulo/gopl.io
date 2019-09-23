package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

var (
	gopath  = os.Getenv("GOPATH")
	dataDir = filepath.Join(
		gopath,
		"src",
		"github.com",
		"djangulo",
		"gopl.io",
		"exercises",
		"ch4",
		"4.12",
		"xkcd-data",
	)
)

type comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

// URL returns the url for the given comic
func (c *comic) URL() string {
	return fmt.Sprintf("https://xkcd.com/%d/", (*c).Num)
}

// URL returns the url for the given comic
func (c *comic) Date() time.Time {
	year, err := strconv.Atoi((*c).Year)
	if err != nil {
		panic(err)
	}
	month, err := strconv.Atoi((*c).Month)
	if err != nil {
		panic(err)
	}
	day, err := strconv.Atoi((*c).Day)
	if err != nil {
		panic(err)
	}
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return t
}

// URL returns the url for the given comic
func (c *comic) StrDate() string {
	year, err := strconv.Atoi((*c).Year)
	if err != nil {
		panic(err)
	}
	month, err := strconv.Atoi((*c).Month)
	if err != nil {
		panic(err)
	}
	day, err := strconv.Atoi((*c).Day)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf(
		"%4.4d-%2.2d-%2.2d",
		year,
		month,
		day,
	)
}

func readIndex(comics *[]*comic) error {

	err := filepath.Walk(dataDir, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return errors.Wrap(err, "cannot open file")
		}

		data, err := ioutil.ReadAll(file)
		if err != nil {
			return errors.Wrap(err, "cannot read file")
		}

		var c comic
		err = json.Unmarshal(data, &c)
		if err != nil {
			fmt.Println("error at file ", info.Name())
			return errors.Wrap(err, "cannot unmarshal comic")
		}
		*comics = append(*comics, &c)
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "cannot walk data dir")
	}
	return nil
}

func search(comics *[]*comic, terms ...string) []*comic {
	results := make(map[int]*comic)
	for _, c := range *comics {
		for _, t := range terms {
			if strings.Contains(strings.ToLower((*c).Alt), strings.ToLower(t)) ||
				strings.Contains(strings.ToLower((*c).Title), strings.ToLower(t)) ||
				strings.Contains(strings.ToLower((*c).Transcript), strings.ToLower(t)) {
				if _, ok := results[(*c).Num]; !ok {
					results[(*c).Num] = c
				}
			}
		}
	}
	r := make([]*comic, 0)
	for _, v := range results {
		r = append(r, v)
	}
	return r
}

const (
	padding = 10
)

func main() {
	comics := make([]*comic, 0)

	err := readIndex(&comics)
	if err != nil {
		log.Fatal(err)
	}
	results := search(&comics, os.Args[1:]...)

	sort.SliceStable(results, func(i, j int) bool {
		return (*results[i]).Date().Before((*results[j]).Date())
	})

	w := tabwriter.NewWriter(os.Stdout, 12, 0, padding, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "Date\tTitle\tLink")
	fmt.Fprintln(w, "----------------------------------------------------")
	for _, r := range results {
		// fmt.Println(r.Title, r.URL())
		fmt.Fprintf(w, "%v\t%s\t%s\n", (*r).StrDate(), (*r).Title, (*r).URL())
	}
	w.Flush()

}
