package main

import (
	"strings"
	"unsafe"
)

// JoinVariadicUnsafe concatenates the elements of a to create a single string. The separator string
// sep is placed between elements in the resulting string.
func JoinVariadicUnsafe(sep string, a ...string) string {
	lena := len(a)
	switch lena {
	case 0:
		return ""
	case 1:
		return a[0]
	}

	n := len(sep) * (lena - 1)
	for i := 0; i < lena; i++ {
		n += len(a[i])
	}

	b := make([]byte, n)
	loc := 0
	for i, s := range a {
		if i > 0 {
			for _, r := range sep {
				b[loc] = byte(r)
				loc++
			}
		}
		for _, r := range s {
			b[loc] = byte(r)
			loc++
		}
	}
	return *(*string)(unsafe.Pointer(&b))
}

// JoinVariadicBuilder concatenates the elements of a to create a single string. The separator string
// sep is placed between elements in the resulting string.
func JoinVariadicBuilder(sep string, a ...string) string {
	lena := len(a)
	switch lena {
	case 0:
		return ""
	case 1:
		return a[0]
	}

	n := len(sep) * (lena - 1)
	for i := 0; i < lena; i++ {
		n += len(a[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(a[0])
	for _, s := range a[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}
	return b.String()
}

func main() {

}

/*
Exeriment to check allocations and performance for standard `strings` implementation
vs an unsafe-ly built one mimicking the strings.Builder implementation. Negligible difference.

goos: windows
goarch: amd64
pkg: github.com/djangulo/gopl.io/exercises/ch5/5.16
BenchmarkVariadicUnsafe-2    	 8570938	       127 ns/op	      32 B/op	       1 allocs/op
BenchmarkVariadicBuilder-2   	 9676872	       112 ns/op	      32 B/op	       1 allocs/op
PASS
ok  	github.com/djangulo/gopl.io/exercises/ch5/5.16	2.602s
*/
