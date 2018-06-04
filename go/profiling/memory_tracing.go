// --------------
// Memory Tracing
// --------------

// Memory Tracing gives us a general idea if our software is healthy as related to the GC and
// memory in the heap that we are working with.

// We are using a special environmental variable called GODEBUG. It gives us the ability to do
// a memory trace and a scheduler trace. Below is a sample program that causes memory leak that we
// can use GODEBUG to understand what's going on.

// Here are the steps to build and run:
// Build the program by: go build memory_tracing.go
// Run the binary: GODEBUG=gctrace=1 ./memory_tracing

// Setting the GODEBUG=gctrace=1 causes the garbage collector to emit a single line to standard
// error at each collection, summarizing the amount of memory collected and the length of the pause.

// What we are gonna see are bad traces followed by this pattern:
/*
   gc {0} @{1}s {2}%: {3}+...+{4} ms clock, {5}+...+{6} ms cpu, {7}->{8}->{9} MB, {10} MB goal, {11} P

   where:
       {0} : The number of times gc run
       {1} : The amount of time the program has been running.
       {2} : The percentage of CPU the gc is taking away from us.
       {3} : Stop of wall clock time - a measure of the real time including time that passes due to programmed
            delays or waiting for resources to become available.
       {4} : Stop of wall clock. This is normally a more important number to look at.
       {5} : CPU clock
       {6} : CPU clock
       {7} : The size of the heap prior to the gc starting.
       {8} : The size of the heap after the gc run.
       {9} : The size of the live heap.
       {10}: The goal of the gc, pacing algorithm.
       {11}: The number of processes.
*/

// For example:
/*
   gc 1 @0.007s 0%: 0.010+0.13+0.030 ms clock, 0.080+0/0.058/0.15+0.24 ms cpu, 5->5->3 MB, 6 MB goal, 8 P
   gc 2 @0.013s 0%: 0.003+0.21+0.034 ms clock, 0.031+0/0.030/0.22+0.27 ms cpu, 9->9->7 MB, 10 MB goal, 8 P
   gc 3 @0.029s 0%: 0.003+0.23+0.030 ms clock, 0.029+0.050/0.016/0.25+0.24 ms cpu, 18->18->15 MB, 19 MB goal, 8 P
   gc 4 @0.062s 0%: 0.003+0.40+0.040 ms clock, 0.030+0/0.28/0.11+0.32 ms cpu, 36->36->30 MB, 37 MB goal, 8 P
   gc 5 @0.135s 0%: 0.003+0.63+0.045 ms clock, 0.027+0/0.026/0.64+0.36 ms cpu, 72->72->60 MB, 73 MB goal, 8 P
   gc 6 @0.302s 0%: 0.003+0.98+0.043 ms clock, 0.031+0.078/0.016/0.88+0.34 ms cpu, 65->66->42 MB, 120 MB goal, 8 P
   gc 7 @0.317s 0%: 0.003+1.2+0.080 ms clock, 0.026+0/1.1/0.13+0.64 ms cpu, 120->121->120 MB, 121 MB goal, 8 P
   gc 8 @0.685s 0%: 0.004+1.6+0.041 ms clock, 0.032+0/1.5/0.72+0.33 ms cpu, 288->288->241 MB, 289 MB goal, 8 P
   gc 9 @1.424s 0%: 0.004+4.0+0.081 ms clock, 0.033+0.027/3.8/0.53+0.65 ms cpu, 577->577->482 MB, 578 MB goal, 8 P
   gc 10 @2.592s 0%: 0.003+11+0.045 ms clock, 0.031+0/5.9/5.2+0.36 ms cpu, 499->499->317 MB, 964 MB goal, 8 P
*/

// It go really fast in the beginning and start to slow down. This is bad.
// The size of the heap is increasing every time the gc run. It shows that there is a memory leak.

package main

import (
	"os"
	"os/signal"
)

func main() {
	// Create a Goroutine that leaks memory. Dumping key-value pairs to put tons of allocation.
	go func() {
		m := make(map[int]int)

		for i := 0; ; i++ {
			m[i] = i
		}
	}()

	// Shutdown the program with Ctrl-C
	sig := make(chan os.Signal, 1)
	signal.Notify(sig)
	<-sig
}
