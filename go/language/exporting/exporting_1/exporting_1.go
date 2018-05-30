// -------------------
// Exported identifier
// -------------------

package main

import (
	"fmt"

	// This is a relative patth to a physical location on our disk - relative to GOPATH
	"github.com/hoanhan101/ultimate-go/go/language/exporting/exporting_1/counters"
)

func main() {
	// Create a variable of the exported type and initialize the value to 10.
	counter := counters.AlertCounter(10)

	// However, when we create a variable of the unexported type and initialize the value to 10:
	// counter := counters.alertCounter(10)
	// The compiler will say:
	// - cannot refer to unexported name counters.alertCounter
	// - undefined: counters.alertCounter

	fmt.Printf("Counter: %d\n", counter)
}
