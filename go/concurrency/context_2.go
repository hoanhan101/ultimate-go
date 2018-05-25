// ----------
// WithCancel
// ----------

// Different ways we can do cancellation, timeout in Go.

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a context that is cancellable only manually.
	// The cancel function must be called regardless of the outcome.
	// WithCancel allows us to create a context and provides us a cancel function that can be
	// called in order to report a signal, a signal without data, that we want whatever that
	// Goroutine is doing to stop right away. Again, we are using Background as our parents context.
	ctx, cancel := context.WithCancel(context.Background())

	// The cancel function must be called regardless of the outcome.
	// The Goroutine that creates the context must always call cancel. These are things that have
	// to be cleaned up. It's the responsibility that the Goroutine creates the context the first
	// time to make sure to call cancel after everything is done.
	// The use of the defer keyword is perfect here for this use case.
	defer cancel()

	// We launch a Goroutine to do some work for us.
	// It is gonna sleep for 50 milliseconds and then call cancel. It is reporting that it want to
	// signal a cancel without data.
	go func() {
		// Simulate work.
		// If we run the program using 50 ms, we expect the work to be complete. But if it is 150
		// ms, then we move on.
		time.Sleep(50 * time.Millisecond)

		// Report the work is done.
		cancel()
	}()

	// The original Goroutine that creates that channel is in its select case. It is gonna receive
	// after time.After. We are gonna wait 100 milliseconds for something to happen. We are also
	// waiting on context.Done. We are just gonna sit here, and if we are told to Done, we know
	// that work up there is complete.
	select {
	case <-time.After(100 * time.Millisecond):
		fmt.Println("moving on")
	case <-ctx.Done():
		fmt.Println("work complete")
	}
}
