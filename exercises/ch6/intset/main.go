// This file contains modifications that satisfy the requirements in exercises
// 6.1 through 6.5

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

const uintEffectiveSize = 32 << (^uint(0) >> 63)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/uintEffectiveSize, uint(x%uintEffectiveSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/uintEffectiveSize, uint(x%uintEffectiveSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	if t.Len() == 0 {
		return //unchanged
	}
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith sets s to the intersection of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	if t.Len() == 0 || s.Len() == 0 { // empty set intersection is empty set
		s.Clear()
		return
	}
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// DifferenceWith sets s to the intersection of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	if t.Len() == 0 { // empty set difference is unchanged
		return
	}
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// SymmetricDifference sets s to the intersection of s and t.
func (s *IntSet) SymmetricDifference(t *IntSet) {
	if t.Len() == 0 { // empty set difference is unchanged
		return
	}
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uintEffectiveSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", uintEffectiveSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Elems returns the number of elements in the intset
func (s *IntSet) Elems() (elems []int) {
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uintEffectiveSize; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, uintEffectiveSize*i+j)
			}
		}
	}
	return
}

// Len returns the number of elements in the intset
func (s *IntSet) Len() (sum int) {
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uintEffectiveSize; j++ {
			if word&(1<<uint(j)) != 0 {
				sum++
			}
		}
	}
	return
}

// Remove removes (inserts zero value) x from the set
func (s *IntSet) Remove(x int) {
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uintEffectiveSize; j++ {
			if word&(1<<uint(j)) != 0 {
				s.words[i] = 0
			}
		}
	}
	return
}

// Clear removes all words from the set
func (s *IntSet) Clear() {
	s.words = make([]uint, 0)
}

// Copy returns a copy of the original intset
func (s *IntSet) Copy() *IntSet {
	var copy IntSet
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uintEffectiveSize; j++ {
			if word&(1<<uint(j)) != 0 {
				copy.Add(uintEffectiveSize*i + j)
			}
		}
	}
	return &copy
}

// AddAll adds the non-negative values x to the set.
func (s *IntSet) AddAll(values ...int) {
	if len(values) == 0 {
		return
	}
	for _, x := range values {
		word, bit := x/uintEffectiveSize, uint(x%uintEffectiveSize)
		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] |= 1 << bit
	}
}
