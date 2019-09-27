package main

import (
	"fmt"
	"testing"
)

func TestJoinVariadicUnsafe(t *testing.T) {
	for _, test := range []struct {
		a    []string
		sep  string
		want string
	}{
		{[]string{"this", "is", "a", "sentence"}, " ", "this is a sentence"},
		{[]string{"this", "is", "a", "conjunction"}, ", ", "this, is, a, conjunction"},
	} {
		name := fmt.Sprintf("JoinVariadicUnsafe(\"%s\", \"%v\"...)", test.sep, test.a)
		t.Run(name, func(t *testing.T) {
			got := JoinVariadicUnsafe(test.sep, test.a...)
			if got != test.want {
				t.Errorf("JoinVariadicUnsafe(\"%s\", \"%v\"...) == \"%s\", want \"%s\"", test.sep, test.a, got, test.want)
			}
		})
	}
}

func TestJoinVariadicBuilder(t *testing.T) {
	for _, test := range []struct {
		a    []string
		sep  string
		want string
	}{
		{[]string{"this", "is", "a", "sentence"}, " ", "this is a sentence"},
		{[]string{"this", "is", "a", "conjunction"}, ", ", "this, is, a, conjunction"},
	} {
		name := fmt.Sprintf("JoinVariadicUnsafe(\"%s\", \"%v\"...)", test.sep, test.a)
		t.Run(name, func(t *testing.T) {
			got := JoinVariadicBuilder(test.sep, test.a...)
			if got != test.want {
				t.Errorf("JoinVariadicUnsafe(\"%s\", \"%v\"...) == \"%s\", want \"%s\"", test.sep, test.a, got, test.want)
			}
		})
	}
}

func BenchmarkVariadicUnsafe(b *testing.B) {
	in, sep := []string{"this", "is", "a", "sentence"}, ", "
	for i := 0; i < b.N; i++ {
		JoinVariadicUnsafe(sep, in...)
	}
}

func BenchmarkVariadicBuilder(b *testing.B) {
	in, sep := []string{"this", "is", "a", "sentence"}, ", "
	for i := 0; i < b.N; i++ {
		JoinVariadicBuilder(sep, in...)
	}
}
