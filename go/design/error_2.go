// ---------------
// Error variables
// ---------------

// Sample program to show how to use error variables to help the caller determine
// the exact error being returned.

package main

import (
	"errors"
	"fmt"
)

// We want these to be on the top of the source code file.
// Naming convention: starting with Err
// They have to be exported because our user need to access to them.
// These are all error interfaces that we have discussed in the last file, with variables tied to
// them. The contexts for these errors are the variables themselves. This allows us to continue
// using the default error type, that unexported type with unexported field to maintain a level of
// decoupling through error handling.
var (
	// ErrBadRequest is returned when there are problems with the request.
	ErrBadRequest = errors.New("Bad Request")

	// ErrPageMoved is returned when a 301/302 is returned.
	ErrPageMoved = errors.New("Page Moved")
)

func main() {
	if err := webCall(true); err != nil {
		switch err {
		case ErrBadRequest:
			fmt.Println("Bad Request Occurred")
			return

		case ErrPageMoved:
			fmt.Println("The Page moved")
			return

		default:
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Life is good")
}

// webCall performs a web operation.
func webCall(b bool) error {
	if b {
		return ErrBadRequest
	}

	return ErrPageMoved
}
