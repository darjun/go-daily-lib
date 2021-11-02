package fib

import "testing"

func BenchmarkFib1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib1(20)
	}
}

func BenchmarkFib2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib2(20)
	}
}

func BenchmarkFib3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib3(20)
	}
}
