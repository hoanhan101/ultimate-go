// -------
// Mutexes
// -------

// We don't always have the luxury of using of 4-8 bytes of memory as a shared data. This is where
// the muxtex comes in. Mutex allows us to have the API like the WaitGroup (Add, Done and Wait)
// where any Goroutine can execute one at a time.

package main

import (
	"fmt"
	"sync"
)

var (
	// counter is a variable incremented by all Goroutines.
	counter int

	// mutex is used to define a critical section of code.
	// Picture mutex as a room where all Goroutines have to go through. However, only one Goroutine
	// can go at a time. The scheduler will decide who can get in and which one is next. We cannot
	// determine what the scheduler is gonna do. It's gonna be hopefully be fair. Just because one
	// Goroutine got to the door before another doesn't that Goroutine got to ended first. Nothing
	// here is predictable.

	// The key here is, once a Goroutine is allowed in, it must report that it's out.
	// All the Goroutines come in will ask for a lock and unlock when it leave for other one to get
	// in.

	// Two different functions can use the same mutex which means only one Goroutine can execute
	// any of given functions at a time.
	mutex sync.Mutex
)

func main() {
	// Number of Goroutines to use.
	const grs = 2

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(grs)

	// Create two Goroutines.
	for i := 0; i < grs; i++ {
		go func() {
			for count := 0; count < 2; count++ {
				// Only allow one Goroutine through this critical section at a time.
				// Creating these artificial curly brackets gives readability. We don't have to do
				// this but it is highly recommended.
				// The Lock and Unlock function must always be together in line of sight.
				mutex.Lock()
				{
					// Capture the value of counter.
					value := counter

					// Increment our local value of counter.
					value++

					// Store the value back into counter.
					counter = value
				}
				mutex.Unlock()
				// Release the lock and allow any waiting Goroutine through.
			}

			wg.Done()
		}()
	}

	// Wait for the Goroutines to finish.
	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}
