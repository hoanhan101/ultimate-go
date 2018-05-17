// --------------------------
// Goroutines and parallelism
// --------------------------

// This programs show how Goroutines run in parallel.
// We are going to have 2 P with 2 m, and 2 Gorouines running in parallel on each m.
// This is still the same program that we are starting with. The only difference is that we are
// getting rid of the lowercase and uppercase function and putting their code directly inside Go's
// anonymous functions.

// Looking at the output, we can see a mix of uppercase of lowercase characters. These Goroutines
// are running in parallel now.

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	// Allocate 2 logical processors for the scheduler to use.
	runtime.GOMAXPROCS(2)
}

func main() {
	// wg is used to wait for the program to finish.
	// Add a count of two, one for each goroutine.
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// Declare an anonymous function and create a goroutine.
	go func() {
		// Display the alphabet three times.
		for count := 0; count < 3; count++ {
			for r := 'a'; r <= 'z'; r++ {
				fmt.Printf("%c ", r)
			}
		}

		// Tell main we are done.
		wg.Done()
	}()

	// Declare an anonymous function and create a goroutine.
	go func() {
		// Display the alphabet three times.
		for count := 0; count < 3; count++ {
			for r := 'A'; r <= 'Z'; r++ {
				fmt.Printf("%c ", r)
			}
		}

		// Tell main we are done.
		wg.Done()
	}()

	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
