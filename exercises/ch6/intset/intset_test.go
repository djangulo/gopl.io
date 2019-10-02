package intset

import (
	"fmt"
	"testing"
)

func TestLen(t *testing.T) {
	for _, test := range []struct {
		name string
		in   []int
		want int
	}{
		{"empty set", []int{}, 0},
		{"length 1", []int{10}, 1},
		{"lenght 3", []int{10, 9, 32}, 3},
	} {
		t.Run(test.name, func(t *testing.T) {
			var x IntSet
			for _, v := range test.in {
				x.Add(v)
			}
			got := x.Len()
			if got != test.want {
				t.Errorf("got length %d want length %d on set %s", got, test.want, x.String())
			}
		})
	}
}

func TestRemove(t *testing.T) {
	std := []int{10, 9, 32, 23, 44, 45}
	for _, test := range []struct {
		drop int
		want bool
	}{
		{9, false},
		{33, false},
		{32, false},
		{45, false},
		{0, false},
	} {
		t.Run(fmt.Sprintf("drop %d", test.drop), func(t *testing.T) {
			var x IntSet
			for _, v := range std {
				x.Add(v)
			}
			x.Remove(test.drop)
			got := x.Has(test.drop)
			if got != test.want {
				t.Errorf("got \"%t\" want \"%t\" after IntSet.Remove(%d) on set %s", got, test.want, test.drop, x.String())
			}
		})
	}
}

func TestClear(t *testing.T) {
	std := []int{10, 9, 32, 23, 44, 45}
	var x IntSet
	for _, v := range std {
		x.Add(v)
	}
	x.Clear()
	got := x.Len()
	if got != 0 {
		t.Errorf("received len %d should be 0 after clear", got)
	}
}

func TestCopy(t *testing.T) {
	std := []int{10, 9, 32, 23, 44, 45}
	var x IntSet
	for _, v := range std {
		x.Add(v)
	}
	y := x.Copy()
	if y.String() != x.String() {
		t.Errorf("copy failed:\noriginal: %s\ncopy    : %s", x.String(), y.String())
	}
}

func TestAddAll(t *testing.T) {
	var x IntSet
	x.AddAll(10, 9, 32, 23, 44, 45)
	got := x.Len()
	if got != 6 {
		t.Errorf("received length %d want length 6", got)
	}
}
