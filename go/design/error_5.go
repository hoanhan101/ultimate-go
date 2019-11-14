// ------------
// Finding the bug/pitfall of nil value of error interface
// ------------

// Please run this file to see runtime behavior of this example: `go run ./go/design/error_5.go`

package main

import "log"

// customError is just an empty struct.
type customError struct{}

// customError implements the error interface.
func (*customError) Error() string {
	return "Find the bug."
}

// stub function to be placeholder for pretending if something bad happen
func isSomethingBadHappen() bool {
	return false
}

// doSomething1 returns nil values for error
func doSomething1() ([]byte, *customError) {
	var err *customError

	if isSomethingBadHappen() {
		err = &customError{}
		return nil, err
	}
	return []byte{}, nil
}

// doSometing2 is the same as doSomething1, except the error interface return type instead of concrete type
func doSomething2() ([]byte, error) {
	// This variable is redundant for this use case, it is here for comparing to doSomething1 function
	var err *customError

	if isSomethingBadHappen() {
		err = &customError{}
		return nil, err
	}

	// Always use explicit 'nil' literal if nil must be returned for error interface.
	// DO NOT use nil value assigned in concrete type.
	return []byte{}, nil
}

func main() {
	// This err variable is first initialized as nil value of error interface
	var err error
	log.Printf("before              : err == %#v\n\n", err)

	// Reminder: variable of interface type consists of two elements: type T, and value V: (T)(V)
	_, err = doSomething1()
	log.Printf("after doSomething1(): err == %#v\n", err)
	log.Printf("%#v != %#v is evaluated as %#v\n", err, nil, err != nil)

	// This if statement `err != nil` evaluated as true even though we return nil to err!
	if err != nil {
		// Reference: https://golang.org/doc/faq#nil_error

		// This is because when we assign returned value nil, which is of type *customError struct, to err variable, 
		// which is of type error interface,
		// Go implicit conversion converts the nil value of type *customError as (T=*customError)(V=nil), 
		// which is considered non-nil value even though the V is nil, and assign it to err variable.
		// Value that is considered nil for interface type is (T=nil)(V=nil).
		// When we compare (T=*customError)(V=nil) to (T=nil)(V=nil), it is evaluated as true!

		log.Println("### doSomething1 function fail! Read the comment for more explaination ###\n")
	}

	// The solutions are:
	// * Always use error interface as return type, not the concrete
	// * If nil error is needed to be returned, always use `nil` literal explicitly, 
	// instead of nil value that get assigned into custom type

	// The solutions mentioned above are implemented in doSomething2 function

	_, err = doSomething2()
	log.Printf("after doSomething2(): err == %#v\n", err)
	log.Printf("%#v != %#v is evaluated as %#v\n", err, nil, err != nil)

	// This if statement doesn't happens as expected
	if err != nil {
		log.Println("doSomething2 should NOT be fail!")
	}
}
