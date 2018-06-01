// ------------------
// Review Stack Trace
// ------------------

// How to read stack traces?

package main

func main() {
	// We are making a slice of length 2, capacity 4 and then passing that slice value into a
	// function call example.
	// example takes a slice, a string, and an integer.
	example(make([]string, 2, 4), "hello", 10)
}

// examples call the built-in function panic to demonstrate the stack traces.
func example(slice []string, str string, i int) {
	panic("Want stack trace")
}

// This is the output that we get:
/*
   panic: Want stack trace

   goroutine 1 [running]:
   main.example(0xc420053f38, 0x2, 0x4, 0x1066c02, 0x5, 0xa)
           /Users/hoanhan/go/src/github.com/hoanhan101/ultimate-go/go/profiling/stack_trace.go:18 +0x39
   main.main()
           /Users/hoanhan/go/src/github.com/hoanhan101/ultimate-go/go/profiling/stack_trace.go:13 +0x72
   exit status 2
*/

// Analysis:
// ---------
// We already know that the compiler tells us the lines of problems. That's good.
// What is even better is that we know exactly what values are passed to the function at the time
// stack traces occur. Stack traces show words of data at a time.

// We know that a slice is a 3-word data structure. In our case, 1st word is a pointer, 2nd is 2
// (length) and 3rd is 4 (capacity). String is a 2-word data structure: a pointer and length of 5
// because there are 5 bytes in string "hello". Then we have a 1 word integer of value 10.

// In the stack traces, main.example(0xc420053f38, 0x2, 0x4, 0x1066c02, 0x5, 0xa),
// the corresponding values in the function are address, 2, 4, address, 5, a (which is 10 in base 2).

// If we ask for the data we need, this is a benefit that we can get just by looking at the stack
// traces and see the values that are going in. If we work with the error package from Dave, wrap
// it and add more context, and log package, we have more than enough information to debug a
// problem.
