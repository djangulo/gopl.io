// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 101.

// Package treesort provides insertion sort using an unbalanced binary tree.
// package treesort
package main

import (
	"bytes"
	"fmt"
)

func main() {
	t := new(tree)
	add(t, 5)
	add(t, 6)
	add(t, 1)
	add(t, 10)
	add(t, 11)
	add(t, 12)
	add(t, 122)
	add(t, -1)
	fmt.Println(t)
}

//!+
type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	buf := new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("%d\n", t.value))

	if t.left != nil && t.left.value != 0 {
		buf.WriteString(fmt.Sprintf("%*sL: %s\n", 2, "", t.left.String()))
	}

	if t.right != nil && t.right.value != 0 {
		buf.WriteString(fmt.Sprintf("%*sR: %s\n", 2, "", t.right.String()))
	}

	return buf.String()
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

//!-
