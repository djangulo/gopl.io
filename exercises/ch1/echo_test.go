package main

import (
	"testing"
)

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo1()
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo2()
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo3()
	}
}

func BenchmarkEcho4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo4()
	}
}

func BenchmarkEcho5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo5()
	}
}
