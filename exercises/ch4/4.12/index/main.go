//Package index creates an xkcd index
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	// "time"
)

const (
	comicCount = 2205 // max 2205
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

const (
	bufferSize = 30
)

func main() {
	if err := os.Mkdir(dataDir, os.ModeDir); os.IsExist(err) {
		log.Fatalf(
			"dir %s already exists, refusing to download everything again",
			dataDir,
		)
	}
	pipe := make(chan io.ReadCloser, bufferSize)
	var wg sync.WaitGroup
	for i := 1; i <= comicCount; i++ {
		wg.Add(1)
		go func(i int) {
			err := fetch(i, pipe)
			if err == nil {
				save(i, pipe)
			}
			wg.Done()
		}(i)
		// time.Sleep(10 * time.Millisecond)
	}
	wg.Wait()
}

func fetch(number int, pipe chan<- io.ReadCloser) error {
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", number)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	// defer res.Body.Close() fails if this happens here
	log.Printf("Downloading %s\n", url)

	// return error, send nothing down pipe
	if res.StatusCode != 200 {
		return fmt.Errorf("response error %d", res.StatusCode)
	}
	pipe <- res.Body
	return nil

}

func save(number int, data <-chan io.ReadCloser) {
	path := filepath.Join(dataDir, fmt.Sprintf("%4.4d.json", number))

	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	d := <-data
	defer d.Close()
	scanner := bufio.NewScanner(d)

	log.Printf("saving %s\n", path)
	for scanner.Scan() {
		_, err := w.WriteString(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
	}
}
