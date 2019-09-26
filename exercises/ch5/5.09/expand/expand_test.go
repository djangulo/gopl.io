package expand_test

import (
	"fmt"
	"github.com/djangulo/gopl.io/exercises/ch5/5.09/expand"
	"os"
	"strings"
	"testing"
)

func TestExpand(t *testing.T) {
	for _, test := range []struct {
		name    string
		str     string
		funcStr string
		want    string
		f       func(string) string
	}{
		{"os.Getenv", "this is the $GOPATH", "os.Getenv", "this is the " + os.Getenv("GOPATH"), os.Getenv},
		{"reverse", "this is $reversed", "reverse", "this is desrever", reverse},
		{"multiple", "this is $reversed and $upsideDown", "reverse", "this is desrever and nwoDedispu", reverse},
	} {
		t.Run(test.name, func(t *testing.T) {
			got := expand.Expand(test.str, test.f)
			if got != test.want {
				t.Errorf("expand(\"%s\", %s) == \"%s\", want \"%s\"", test.str, test.funcStr, got, test.want)
			}
		})
	}
}

func reverse(x string) string {
	s := []byte(x)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func ExampleExpand() {

	reverse := func(x string) string {
		s := []byte(x)
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
		return string(s)
	}

	uppercase := func(s string) string {
		return strings.ToUpper(s)
	}

	fmt.Println(expand.Expand("reversed: $reversed", reverse))
	fmt.Println(expand.Expand("i am unchanged and $i $am $uppercased", uppercase))
	// Output:
	// reversed: desrever
	// i am unchanged and I AM UPPERCASED
}
