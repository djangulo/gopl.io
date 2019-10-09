package intset

import (
	"fmt"
	"reflect"
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

func TestUnionWith(t *testing.T) {
	var x, y, phi IntSet
	x.AddAll(1, 2, 3)
	y.AddAll(3, 4, 5)

	t.Run("non empty set", func(t *testing.T) {
		z := x.Copy()
		z.UnionWith(&y)

		got, want := z.Len(), 5
		if got != want {
			t.Errorf("got len %d want len %d for union of %s and %s", got, want, x.String(), y.String())
		}
	})

	t.Run("empty set", func(t *testing.T) {
		z := x.Copy()
		z.UnionWith(&phi)

		got, want := z.Len(), x.Len()
		if got != want {
			t.Errorf("got len %d want len %d for union of %s and %s", got, want, x.String(), phi.String())
		}
	})
}

func TestIntersectWith(t *testing.T) {
	var x, y, phi IntSet
	x.AddAll(1, 2, 3)
	y.AddAll(3, 4, 5)

	t.Run("non empty set", func(t *testing.T) {
		z := x.Copy()
		z.IntersectWith(&y)

		got, want := z.Len(), 1
		if got != want {
			t.Errorf("got len %d want len %d for intersection of %s and %s", got, want, x.String(), y.String())
		}
	})

	t.Run("empty set", func(t *testing.T) {
		z := x.Copy()
		z.IntersectWith(&phi)

		got, want := z.Len(), 0
		if got != want {
			t.Errorf("got len %d want len %d for intersection of %s and %s", got, want, x.String(), phi.String())
		}
	})

}

func TestDifferenceWith(t *testing.T) {
	var x, y, phi IntSet
	x.AddAll(1, 2, 3)
	y.AddAll(3, 4, 5)

	t.Run("non empty set", func(t *testing.T) {
		z := x.Copy()
		z.DifferenceWith(&y)

		got, want := z.Len(), 2
		if got != want {
			t.Errorf("got len %d want len %d for %s.DifferenceWith(%s) = %s", got, want, x.String(), y.String(), z.String())
		}
	})

	t.Run("empty set", func(t *testing.T) {
		z := x.Copy()
		z.DifferenceWith(&phi)

		got, want := z.Len(), x.Len()
		if got != want {
			t.Errorf("got len %d want len %d for %s.DifferenceWith(%s) = %s", got, want, x.String(), phi.String(), z.String())
		}
	})

}

func TestSymmetricDifference(t *testing.T) {
	var x, y, phi IntSet
	x.AddAll(1, 2, 3)
	y.AddAll(3, 4, 5)

	t.Run("non empty set", func(t *testing.T) {
		z := x.Copy()
		z.SymmetricDifference(&y)

		got, want := z.Len(), 4
		if got != want {
			t.Errorf("got len %d want len %d for %s.SymmetricDifference(%s) = %s", got, want, x.String(), y.String(), z.String())
		}
	})

	t.Run("empty set", func(t *testing.T) {
		z := x.Copy()
		z.SymmetricDifference(&phi)

		got, want := z.Len(), x.Len()
		if got != want {
			t.Errorf("got len %d want len %d for %s.SymmetricDifference(%s) = %s", got, want, x.String(), phi.String(), z.String())
		}
	})

}

func TestElems(t *testing.T) {
	var x IntSet
	x.AddAll(7, 42, 108)

	want := []int{7, 42, 108}

	got := x.Elems()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v for %s.Elems()", got, want, x.String())
	}

}
