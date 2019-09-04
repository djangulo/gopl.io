package main

import (
	"fmt"
	"github.com/djangulo/gopl.io/ch3/exercises/surface2"
	// "image"
	"image/color"
	// "image/gif"
	// "io"
	// "math"
	// "math/rand"
	"log"
	"net/http"
	"os"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		var width, height, cells int
		q := r.URL.Query()
		qWidth, qWidthOK := q["width"]
		qHeight, qHeightOK := q["height"]
		qCells, qCellsOK := q["cells"]
		// qPalette, qPaletteOK := q["palette"]

		var line string

		if qWidthOK {
			inty, err := strconv.Atoi(qWidth[0]) // only parsing the first one
			if err != nil {
				fmt.Fprintf(os.Stdout, "error parsing url query: %v", err)
			}
			width = inty
			line += fmt.Sprintf("width: %d", width)
		} else {
			width = 600
			line += fmt.Sprintf("width: %d (default)", width)
		}
		if qHeightOK {
			inty, err := strconv.Atoi(qHeight[0]) // only parsing the first one
			if err != nil {
				fmt.Fprintf(os.Stdout, "error parsing url query: %v", err)
			}
			height = inty
			line += fmt.Sprintf(", height: %d", height)
		} else {
			height = 320
			line += fmt.Sprintf(", height: %d (default)", height)
		}
		if qCellsOK {
			inty, err := strconv.Atoi(qCells[0]) // only parsing the first one
			if err != nil {
				fmt.Fprintf(os.Stdout, "error parsing url query: %v", err)
			}
			cells = inty
			line += fmt.Sprintf(", cells: %d", cells)
		} else {
			cells = 100
			line += fmt.Sprintf(", cells: %d (default)", cells)
		}
		log.Printf("received %s\n", line)
		surface2.SurfacePlot(w, width, height, cells)
	})
	log.Println("Listening on port :8000")
	http.ListenAndServe(":8000", nil)
}
