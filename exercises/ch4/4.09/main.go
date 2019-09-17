// wordfreq counts the distinc words in a file, filenames passed in as stdin args
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	counts := newWordCount(os.Args[1:])
	for k := range counts {
		f, err := os.Open(k)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot read file %s, %v\n", k, err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			text := scanner.Text()
			clean := ""
			for _, r := range text {
				if unicode.IsLetter(rune(r)) {
					clean += string(r)
				}
			}
			clean = strings.ToLower(clean)
			counts[k][clean]++
		}
	}
	fmt.Println(counts)

}

type wordCount map[string]map[string]int

func newWordCount(names []string) wordCount {
	counts := wordCount{}

	for _, n := range names {
		counts[n] = map[string]int{}
	}
	return counts
}

func (w wordCount) String() string {
	var buf bytes.Buffer
	for filename := range w {
		fmt.Fprintf(&buf, "%s\n", filename)
		fmt.Fprintf(&buf, "\t%-18s\tcount\n", "word")
		fmt.Fprintf(&buf, "\t%s\n", strings.Repeat("-", 24))
		for w, n := range w[filename] {
			fmt.Fprintf(&buf, "\t%-18s\t%4d\n", w, n)
		}
	}
	return buf.String()
}
