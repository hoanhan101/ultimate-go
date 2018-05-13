// ------------
// Find the bug
// ------------

package main

import "log"

// customError is just an empty struct.
type customError struct{}

// Error implements the error interface.
func (c *customError) Error() string {
	return "Find the bug."
}

// fail returns nil values for both return types.
func fail() ([]byte, *customError) {
	return nil, nil
}

func main() {
	// This set the err to its zero value.
	//  -----
	// | nil |
	//  -----
	// | nil |
	//  -----
	var err error

	// When we call fail, it returns the value of nil. However, we have the nil value of type
	// *customError. We always want to use the error intarface as the return value. The customError
	// type is just an artifact, a value that we store inside. We cannot use the custom type
	// directly. We must use the error interface, like so func fail() ([]byte, error)
	if _, err = fail(); err != nil {
		log.Fatal("Why did this fail?")
	}

	log.Println("No Error")
}
