package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

var (
	omdbKey      string
	omdbKeyUsage = `key to query the OMBD API with, will attempt to read
from the environment variable OMDBAPIKEY if unset`
	gopath  = os.Getenv("GOPATH")
	saveDir = filepath.Join(
		gopath,
		"src",
		"github.com",
		"djangulo",
		"gopl.io",
		"exercises",
		"ch4",
		"4.13",
		"omdb-data",
	)
)

func init() {
	flag.StringVar(&omdbKey, "k", os.Getenv("OMDBAPIKEY"), omdbKeyUsage)
	flag.StringVar(&omdbKey, "omdb-apikey", os.Getenv("OMDBAPIKEY"), omdbKeyUsage)
}

type movie struct {
	Title  string
	Year   string
	Poster string
}

func main() {
	flag.Parse()
	fmt.Println("OMDB api at work!")
	mov, err := searchMovie(flag.Args())
	if err != nil {
		panic(err)
	}

	err = downloadPoster(mov)
	if err != nil {
		panic(err)
	}
	fmt.Printf("poster downloaded at %s\n", saveDir)
}

const baseURL = "https://www.omdbapi.com"

func searchMovie(terms []string) (*movie, error) {
	q := url.QueryEscape(strings.Join(terms, " "))

	url := fmt.Sprintf("%s/?t=%s&apikey=%s", baseURL, q, omdbKey)

	fmt.Println(url)

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("cannot GET movie: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong result: %v", res.Status)
	}

	var mov movie
	if err := json.NewDecoder(res.Body).Decode(&mov); err != nil {
		return nil, err
	}

	return &mov, nil
}

func downloadPoster(mov *movie) error {
	if err := os.Mkdir(saveDir, os.ModeDir); os.IsExist(err) {
		fmt.Println("download directory exists, skipping creation")
	}

	ext := filepath.Ext((*mov).Poster)
	title := strings.ReplaceAll(strings.ToLower((*mov).Title), " ", "-")

	path := filepath.Join(saveDir, title+ext)

	file, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "cannot create file")
	}
	defer file.Close()

	res, err := http.Get((*mov).Poster)
	if err != nil {
		return errors.Wrap(err, "cannot GET poster")
	}
	defer res.Body.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return errors.Wrap(err, "cannot save poster")
	}

	return nil

}
