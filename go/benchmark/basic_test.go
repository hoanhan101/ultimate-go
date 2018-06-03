// ---------------
// Basic benchmark
// ---------------

// Benchmark file's have to have <file_name>_test.go and use the Benchmark functions like below.
// The goal is to know what perform better and what allocate more or less between Sprint and Sprintf.

// Our guess is that Sprint is gonna be better because it doesn't have any overhead doing the
// formatting. However, this is not true. Remember we have to optimize for correctness so we don't
// want to guess.

// Run benchmark:
// go test -run none -bench . -benchmem -benchtime 3s

// Sample output:
// BenchmarkSprintBasic-8       50000000                78.7 ns/op             5 B/op          1 allocs/op
// BenchmarkSprintfBasic-8      100000000               60.5 ns/op             5 B/op          1 allocs/op

package main

import (
	"fmt"
	"testing"
)

var gs string

// BenchmarkSprint tests the performance of using Sprint.
// All the code we want to benchmark need to be inside the b.N for loop.
// The first time the tool call it, b.N is equal to 1. It will keep increasing the value of N and
// run long enough based on our bench time.
// fmt.Sprint returns a value and we want to capture this value so it doesn't look like dead code.
// We assign it to the global variable gs.
func BenchmarkSprintBasic(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = fmt.Sprint("hello")
	}

	gs = s
}

// BenchmarkSprint tests the performance of using Sprintf.
func BenchmarkSprintfBasic(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("hello")
	}

	gs = s
}
