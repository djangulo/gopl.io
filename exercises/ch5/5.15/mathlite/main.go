/*Exercise 5.15: Write variadic functions max and min, analogous to sum.
What should these functions do when called with no arguments? Write varians
that require at least one argument.*/

//Package mathlite has functions for sumreduce, max and min
// arrays of ints
package mathlite

// original sum function from gopl.io/ch5/sum
// func sum(vals ...int) int {
// 	total := 0
// 	for _, val := range vals {
// 		total += val
// 	}
// 	return total
// }

func isNonEmpty(vals ...int) bool {
	return len(vals) > 0
}

func MaxVariadic(vals ...int) (max int) {
	if !isNonEmpty(vals...) {
		return
	}
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return
}

func MinVariadic(vals ...int) (min int) {
	if !isNonEmpty(vals...) {
		return
	}
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return
}

func Max(v int, vals ...int) (max int) {
	if !isNonEmpty(vals...) {
		return v
	}
	vals = append(vals, v)
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return
}

func Min(v int, vals ...int) (min int) {
	if !isNonEmpty(vals...) {
		return v
	}
	vals = append(vals, v)
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return
}
