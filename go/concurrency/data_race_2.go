// ----------------
// Atomic Functions
// ----------------

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// counter is a variable incremented by all Goroutines.
// Notice that it's not just an int but int64. We are being very specific about the precision
// because the atomic function requires us to do so.
var counter int64

func main() {
	// Number of Goroutines to use.
	const grs = 2

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(grs)

	// Create two goroutines.
	for i := 0; i < grs; i++ {
		go func() {
			for count := 0; count < 2; count++ {
				// Safely add one to counter.
				// Add the atomic functions that we have take an address as the first parameter and
				// that is being synchronized, no matter many Goroutines they are. If we call one
				// of these function on the same location, they will get serialized. This is the
				// fastest way to serialization.
				// We can run this program all day long and still get 4 every time.
				atomic.AddInt64(&counter, 1)

				// This call is now irrelevant because by the time AddInt64 function complete,
				// counter is incremented.
				runtime.Gosched()
			}

			wg.Done()
		}()
	}

	// Wait for the Goroutines to finish.
	wg.Wait()

	// Display the final value.
	fmt.Println("Final Counter:", counter)
}
