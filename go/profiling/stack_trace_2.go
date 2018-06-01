// -------
// Packing
// -------

// Sample stack races that pack values.

package main

func main() {
	// Passing values that are 1-byte values.
	example(true, false, true, 25)
}

func example(b1, b2, b3 bool, i uint8) {
	panic("Want stack trace")
}

// This is the output that we get:
/*
   panic: Want stack trace

   goroutine 1 [running]:
   main.example(0xc419010001)
           /Users/hoanhan/go/src/github.com/hoanhan101/ultimate-go/go/profiling/stack_trace_2.go:12 +0x39
   main.main()
           /Users/hoanhan/go/src/github.com/hoanhan101/ultimate-go/go/profiling/stack_trace_2.go:8 +0x29
   exit status 2
*/

// Analysis:
// ---------
// Since stack traces show 1 word at a time, all of these 4 bytes fit in a half-word on a 32-bit
// platform and a full word on 64-bit. Also, the system we are looking at is using little endian so
// we need to read from right to left. In our case, the word value 0xc419010001 can be represented as:
/*
   Bits    Binary      Hex   Value
   00-07   0000 0001   01    true
   08-15   0000 0000   00    false
   16-23   0000 0001   01    true
   24-31   0001 1001   19    25
*/
