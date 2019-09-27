/*Exercise 5.15: Write variadic functions max and min, analogous to sum.
What should these functions do when called with no arguments? Write varians
that require at least one argument.*/

//Package mathlite has functions for sumreduce, max and min
// arrays of ints
package main

import (
	"fmt"

	"github.com/djangulo/gopl.io/exercises/ch5/5.15/mathlite"
)

func main() {
	//!+main
	for _, f := range []struct {
		name string
		f    func(...int) int
	}{
		{"MaxVariadic", mathlite.MaxVariadic},
		{"MinVariadic", mathlite.MinVariadic},
	} {
		fmt.Printf("%s: %d\n", f.name, f.f())           //  "0"
		fmt.Printf("%s: %d\n", f.name, f.f(3))          //  "3"
		fmt.Printf("%s: %d\n", f.name, f.f(1, 2, 3, 4)) //  "10"
		values := []int{1, 2, 3, 4}
		fmt.Printf("%s: %d\n", f.name, f.f(values...)) // "10"
	}

	for _, f := range []struct {
		name string
		f    func(int, ...int) int
	}{
		{"Max", mathlite.Max},
		{"Min", mathlite.Min},
	} {
		fmt.Printf("%s: %d\n", f.name, f.f(3))          //  "3"
		fmt.Printf("%s: %d\n", f.name, f.f(1, 2, 3, 4)) //  "10"
		values := []int{1, 2, 3, 4}
		fmt.Printf("%s: %d\n", f.name, f.f(values[0], values[1:]...)) // "10"
	}
	//!-main

	//!+slice
	//!-slice
}
