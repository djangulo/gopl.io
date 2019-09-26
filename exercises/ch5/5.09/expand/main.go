//Package expand has a single utility, Expand
package expand

import (
	// "fmt"
	// "os"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`\$[a-zA-Z0-9]*`)

// Expand replaces each substring "$foo" within s by calling f("foo")
func Expand(s string, f func(string) string) string {
	words := re.FindAllString(s, -1)
	for _, w := range words {
		w = w[1:]
		s = strings.Replace(s, "$"+w, f(w), 1)
	}

	return s
}

// func main() {
// 	input := strings.Join(os.Args[1:], " ")
// 	fmt.Println(Expand(input, os.Getenv))

// }
