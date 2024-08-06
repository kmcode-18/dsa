package main

import "testing"

func BenchmarkFactorial(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Factorial(200)
	}
}

func BenchmarkFactorialTail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FactorialTail(200)
	}
}
