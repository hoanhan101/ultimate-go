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
		// This shows us that this checking is at runtime, not compile time.
		if v, ok := mvs[rn].(cloud); ok {
			fmt.Println("Got Lucky:", v)
			continue
		}

		// We have to guarantee that variable in question (x in `x.(T)`) can always be asserted correctly as T type

		// Or else, We wouldn't want to use that ok variable because we want it to panic if there is an integrity
		// issue. We must shut it down immediately if that happens if we cannot recover from a
		// panic and guarantee that we are back at 100% integrity, the software has to be restarted. 
		// Shutting down means you have to call log.Fatal, os.exit, or panic for stack trace.
		// When we use type assertion, we need to understand when it is okay that whatever
		// we are asking for is not there.

		// Important note:
		// ---------------
		// If the type assertion is causing us to call the concrete value out, that should raise a big
		// flag. We are using interface to maintain a level of decoupling and now we are using type
		// assertion to go back to the concrete.
		// When we are in the concrete, we are putting our codes in the situation where cascading
		// changes can cause widespread refactoring. What we want with interface is the opposite,
		// internal changes minimize cascading changes.
		fmt.Println("Got Unlucky")
	}
}
