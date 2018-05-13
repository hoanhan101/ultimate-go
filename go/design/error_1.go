// --------------------
// Default error values
// --------------------

// Integrity matters. Nothing trumps integrity. Therefore, part of integrity is error handling.
// It is a big part of what we do everyday. It has to be a part of the main code.

// First, let's look at the language mechanic first on how the default error type is implemented.

package main

import "fmt"

// http://golang.org/pkg/builtin/#error
// This is built in the language so it looks like an unexported type. It has one active behavior,
// which is Error returned a string.
// Error handling is decoupled because we are always working with error interface when we are
// testing our code.
// Errors in Go are really just values. We are going to valuate these through the decoupling of
// the interface. Decoupling error handling means that cascading changes will bubble up through the
// user application, causes cascading wide effect through the code base. It's important that we
// leverage the interface here as much as we can.
type error interface {
	Error() string
}

// http://golang.org/src/pkg/errors/errors.go
// This is the default concrete type that comes from the error package. It is an unexported type
// that has an unexported field. This gives us enough context to make us form a decision.
// We have responsibility around error handling to give the caller enough context to make them form
// a decision so they know how to handle this situation.
type errorString struct {
	s string
}

// http://golang.org/src/pkg/errors/errors.go
// This is using a pointer receiver and returning a string.
// If the caller must call this method and parse a string to see what is going on then we fail.
// This method is only for logging information about the error.
func (e *errorString) Error() string {
	return e.s
}

// http://golang.org/src/pkg/errors/errors.go
// New returns an error interface that formats as the given text.
// When we call New, what we are doing is creating errorString value, putting some sort of string
// in there.. Since we are returning the address of a concrete type, the user will get an error
// interface value where the first word is a *errorString and the second word points to the
// original value. We are going to stay decoupled during the error handling.
//       error
//  --------------
// | *errorString |      errorString
//  --------------       -----------
// |      *       | --> |   "Bad"   |
//  --------------       -----------
func New(text string) error {
	return &errorString{text}
}

func main() {
	// This is a very traditional way of error handling in Go.
	// We are calling webCall and return the error interface and store that in a variable.
	// nil is a special value in Go. What "error != nil" actually means is that we are asking if
	// there is a concrete type value that is stored in error type interface. Because if error is
	// not nil, there is a concrete value stored inside. If that's is the case, we've got an error.
	// Now do we handle the error, do we return the error up the call stack for someone else to
	// handle? We will talk about this latter.
	if err := webCall(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Life is good")
}

// webCall performs a web operation.
func webCall() error {
	return New("Bad Request")
}
