package main

import "testing"

var globalResult *int

// BenchmarkSquare tests how fast square() is and how much memory it uses.
func BenchmarkSquare(b *testing.B) {
	// b.N is a number that is automatically adjusted by Go at runtime
	// to make the test run for a meaningful amount of time (usually 1 second)
	var r *int
	for i := 0; i < b.N; i++ {
		r = square(4)
	}
	globalResult = r
}
