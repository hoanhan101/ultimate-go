// --------------
// Race Detection
// --------------

// As soon as we add another Goroutine to our program, we add a huge amount of complexity. We can't
// always let the Goroutine run stateless. There has to be coordination. There are, in fact, 2
// things that we can do with multithread software.
// (1) We either have to synchronize access to share state like that WaitGroup is done with Add, Done
// and Wait.
// (2) Or we have to coordinate these Goroutines to behave in a predictable or responsible manner.
// Up until the use of channel, we have to use atomic function, mutex, to do both. The channel
// gives us a simple way to do orchestration. However, in many cases, using atomic function, mutex,
// and synchronizing access to shared state is the best way to go.

// Atomic instructions are the fastest way to go because deep down in memory, Go is synchronizaing
// 4-8 bytes at a time.
// Mutexes are the next fastest. Channels are very slow because not only they are mutexes, there
// are all data structure and logic that go with them.

// Data races is when we have multiple Goroutines trying to access the same memory location.
// For example, in the simplest case, we have a integer that is a counter. We have 2 Goroutines
// that want to read and write to that variable at the same time. If they are actually doing it at
// the same time, they are going to trash each other read and write. Therefore, this type of
// synchronizing access to the shared state has to be coordinated.

// The problem with data races is that they always appear random.

// Sample program to show how to create race conditions in our programs. We don't want to do this.
// To identify race condition : go run -race <file_name>

package main

import (
	"fmt"
	"runtime"
	"sync"
)

// counter is a variable incremented by all Goroutines.
var counter int

func main() {
	// Number of Goroutines to use.
	const grs = 2

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(grs)

	// Create two Goroutines.
	// They loop twice: perform a read to a local counter, increase by 1, write it back to the shared state
	// Every time we run the program, the output should be 4.
	// The data races that we have here is that: at any given time, both Gooutines could be reading
	// and writing at the same time. However, we are very lucky in this case. What we are seeing it
	// that, each Goroutine is executing the 3 statements atomically completely by accident every
	// time this code run.
	// If we put the line runtime.Gosched(), it will tell the scheduler to be part of the
	// cooperation here and yield my time on that m. This will force the data race to happen. Once
	// we read the value out of that shared state, we are gonna force the context switch. Then we
	// come back, we are not getting 4 as frequent.
	for i := 0; i < grs; i++ {
		go func() {
			for count := 0; count < 2; count++ {
				// Capture the value of Counter.
				value := counter

				// Yield the thread and be placed back in queue.
				// FOR TESTING ONLY! DO NOT USE IN PRODUCTION CODE!
				runtime.Gosched()

				// Increment our local value of Counter.
				value++

				// Store the value back into Counter.
				counter = value
			}
			wg.Done()
		}()
	}

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
