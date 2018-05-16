// ------------------
// Language Mechanics
// ------------------

// One of the most important thing that we must do from day one is to write software that can
// startup and shutdown cleanly. This is very very important.

package main

import (
	"fmt"
	"runtime"
	"sync"
)

// init calls a function from the runtime package called GOMAXPROCS. This is also an environment
// variable, which is why is all capitalized.
// Prior to 1.5, when our Go program came up for the first time, it came up with just a single P,
// regardless of how many cores. The improvement that we made to the garbage collector and
// scheduler changed all that.
func init() {
	// Allocate one logical processor for the scheduler to use.
	runtime.GOMAXPROCS(1)
}

func main() {
	// wg is used to manage concurrency.
	// wg is set to its zero value. This is one of the very special types in Go that are usable in
	// its zero value state.
	// It is also called Asynchronous Counting Semaphore. It has three methods: Add, Done and Wait.
	// n number of Goroutines can call this method at the same time and it's all get serialized.
	// - Add keeps a count of how many Goroutines out there.
	// - Done decrements that count because some Goroutines are about to terminated.
	// - Wait holds the program until that count goes back down to zero.
	var wg sync.WaitGroup

	// We are creating 2 Gorouines.
	// We rather call Add(1) and call it over and over again to increment by 1. If we don't how
	// many Goroutines that we are going to create, that is a smell.
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// Create a Goroutine from the uppercase function using anonymous function.
	// We have a function decoration here with no name and being called by the () in the end. We
	// are declaring and calling this function right here, inside of main. The big thing here is
	// the keyword go in front of func().
	// We don't execute this function right now in series here. Go schedules that function to be a
	// G, say G1, and load in some LRQ for our P. This is our first G.
	// Remember, we want to think that every G that is in runnable state is running at the same time.
	// Even though we have a single P, even though we have a single thread, we don't care.
	// We are having 2 Goroutines running at the same time: main and G1.
	go func() {
		lowercase()
		wg.Done()
	}()

	// Create a Goroutine from the lowercase function.
	// We are doing it again. We are now having 3 Goroutines running at the same time.
	go func() {
		uppercase()
		wg.Done()
	}()

	// Wait for the Goroutines to finish.
	// This is holding main from terminating because when the main terminates, our program
	// terminates, regardless of what any other Goroutine is doing.
	// There is a golden rule here: We are not allowed to create a Goroutine unless we can tell
	// when and how it terminates.
	// Wait allows us to hold the program until the two other Goroutines report that they are done.
	// It is gonna wait, count from 2 to 0. When it reaches 0, the scheduler will wake up the main
	// Goroutine again and allow it to be terminated.
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}

// lowercase displays the set of lowercase letters three times.
func lowercase() {
	// Display the alphabet three times
	for count := 0; count < 3; count++ {
		for r := 'a'; r <= 'z'; r++ {
			fmt.Printf("%c ", r)
		}
	}
}

// uppercase displays the set of uppercase letters three times.
func uppercase() {
	// Display the alphabet three times
	for count := 0; count < 3; count++ {
		for r := 'A'; r <= 'Z'; r++ {
			fmt.Printf("%c ", r)
		}
	}
}

// Sequence
// --------
// We call the uppercase after lowercase but Go's scheduler chooses to call the lowercase first.
// Remember we are running on a single thread so there is only one Goroutine is executed at a given
// time here. We can't see that we are running concurrently that the uppercase runs before the
// lowercase. Everything starts and completes cleanly.

// What if we forget to hold Wait?
// -------------------------------
// We would see no output of uppercase and lowercase. This is pretty much a data race. It's a race
// to see the program terminates before the scheduler stops it and schedules another Goroutine to
// run. By not waiting, these Goroutine never get a chance to execute at all.

// What if we forget to call Done?
// -------------------------------
// Deadlock!
// This is a very special thing in Go. When the runtime determines that all the Goroutines are
// there can no longer move forward, it's gonna panic.
