// -------------
// Sub benchmark
// -------------

// Like sub test, we can also do sub benchmark.

// Sample available commands:
// go test -run none -bench . -benchtime 3s -benchmem
// go test -run none -bench BenchmarkSprintSub/none -benchtime 3s -benchmem
// go test -run none -bench BenchmarkSprintSub/format -benchtime 3s -benchmem

package main

import (
	"fmt"
	"testing"
)

// BenchmarkSprint tests all the Sprint related benchmarks as sub benchmarks.
func BenchmarkSprintSub(b *testing.B) {
	b.Run("none", benchSprint)
	b.Run("format", benchSprintf)
}

// benchSprint tests the performance of using Sprint.
func benchSprint(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = fmt.Sprint("hello")
	}

	gs = s
}

// benchSprintf tests the performance of using Sprintf.
func benchSprintf(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("hello")
	}

	gs = s
}
