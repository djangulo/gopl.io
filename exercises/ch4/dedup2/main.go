package main

import (
	"fmt"
)

func main() {
	in := []string{"a", "b", "b", "c", "d", "b"}
	fmt.Println(dedup2(in))

}

// dedup2 drops adjacent duplicates (i.e. only works if slice is sorted)
func dedup2(in []string) []string {
	j := 0
	for i := 1; i < len(in); i++ {
		if in[j] == in[i] {
			continue
		}
		j++
		in[j] = in[i]
	}
	return in[:j+1]
}
