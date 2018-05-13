// -----------------------
// Runtime Type Assertions
// -----------------------

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// car represents something you drive.
type car struct{}

// String implements the fmt.Stringer interface.
func (car) String() string {
	return "Vroom!"
}

// cloud represents somewhere you store information.
type cloud struct{}

// String implements the fmt.Stringer interface.
func (cloud) String() string {
	return "Big Data!"
}

func main() {
	// Seed the number random generator.
	rand.Seed(time.Now().UnixNano())

	// Create a slice of the Stringer interface values.

	//  ---------------------
	// |   car   |   cloud   |
	//  ---------------------
	// |    *    |     *     |
	//  ---------------------
	//      A          A
	//      |          |
	//     car       cloud
	//    -----      -----
	//   |     |    |     |
	//    -----      -----

	mvs := []fmt.Stringer{
		car{},
		cloud{},
	}

	// Let's run this experiment ten times.
	for i := 0; i < 10; i++ {
		// Choose a random number from 0 to 1.
		rn := rand.Intn(2)

		// Perform a type assertion that we have a concrete type of cloud in the interface
		// value we randomly chose.
		// This shows us that this checking is at runtime, not compiled time.
		if v, ok := mvs[rn].(cloud); ok {
			fmt.Println("Got Lucky:", v)
			continue
		}

		// This shows us that this checking is at runtime, not compiled time. We have to decide
		// if there should always be a value of some type that should never change.
		// We wouldn't want to use that ok because we want it to panic if there is a integrity
		// issue. We must shut it down immediately if that happen. If we cannot recover from a
		// panic and give the guarantee that you are back at 100% integrity, the software has to be
		// restarted. Shutting downs means you have to either call log.Fatal or os.exit, or panic
		// for stack trace.
		// When we use type assertion, we need to understand where or not it is okay that whatever
		// we are asking for is not there.

		// Important note:
		// ---------------
		// When we are using type assertion, it raises a big red flag.
		// If the type assertion is causing us to call the concrete value out, that should raise a
		// flag. We are using interface to maintain a level of decoupling and now we are using type
		// asserting to go back to the concrete.
		// When we are in the concrete, we are putting our codes in the situation where cascading
		// changes can cause widespread refactoring. What we want with interface is the opposite,
		// internal changes minimize cascading changes.
		fmt.Println("Got Unlucky")
	}
}
