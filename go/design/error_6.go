// ---------------
// Wrapping Errors
// ---------------

// Error handling has to be part of our code and usually it is bounded to logging.
// The main goal of logging is to debug.

// We only log things that are actionable. Only log the contexts that are allowed us to identify
// that is going on. Anything else ideally is noise and would be better suited up on the dashboard
// through metrics. For example, socket connection and disconnection, we can log these but these
// are not actionable because we don't necessary lookup the log for that.

// There is a package that is written by Dave Cheney called errors that let us simplify error
// handling and logging that the same time. Below is a demonstration on how to leverage the package
// to simplify our code. By reducing logging, we also reduce a large amount of pressure on the heap
// (garbage collection).

package main

import (
	"fmt"

	// This is Dave Cheney's errors package that have all the wrapping functions.
	"github.com/pkg/errors"
)

// AppError represents a custom error type.
type AppError struct {
	State int
}

// Error implements the error interface.
func (c *AppError) Error() string {
	return fmt.Sprintf("App Error, State: %d", c.State)
}

func main() {
	// Make the function call and validate the error.

	// firtCall calls secondCall calls thirdCall then results in AppError.
	// Start down the call stack, in thirdCall, where the error occurs. The is the root of the
	// error. We return it up the call stack in our traditional error interface value.
	// Back to secondCall, we get the interface value and there is a concrete type stored inside
	// the value. secondCall has to make a decision whether to handle the error and push up the
	// call stack if it cannot handle. If secondCall decides to handle the error, it has the
	// responsibility of logging it. If not, its responsibility is to move it up. However, if we
	// are going to push it up the call stack, we cannot lose context. This is where the error
	// package comes in. We create a new interface value that wraps this error, add a context
	// around it and push it up. This maintains the call stack of where we are in the code.
	// Similarly, firstCall doesn't handle the error but wraps and pushes it up.

	// In main, we are handling the call, which means the bug stops here and I have to log it.
	// In order to properly handle this error, we need to know that the root cause of this error
	// was. It is the original error that is not wrapped. Cause will bubble up this error out of
	// these wrapping and allow us to be able to use all the language mechanics we have.

	// We are not only be able to access the State even though we've done this assertion back to
	// concrete, we can log out the entire stack trace by using %+v for this call.

	if err := firstCall(10); err != nil {
		// Use type as context to determine cause.
		switch v := errors.Cause(err).(type) {
		case *AppError:
			// We got our custom error type.
			fmt.Println("Custom App Error:", v.State)

		default:
			// We did not get any specific error type.
			fmt.Println("Default Error")
		}

		// Display the stack trace for the error.
		fmt.Println("\nStack Trace\n********************************")
		fmt.Printf("%+v\n", err)
		fmt.Println("\nNo Trace\n********************************")
		fmt.Printf("%v\n", err)
	}
}

// firstCall makes a call to a second function and wraps any error.
func firstCall(i int) error {
	if err := secondCall(i); err != nil {
		return errors.Wrapf(err, "firstCall->secondCall(%d)", i)
	}
	return nil
}

// secondCall makes a call to a third function and wraps any error.
func secondCall(i int) error {
	if err := thirdCall(); err != nil {
		return errors.Wrap(err, "secondCall->thirdCall()")
	}
	return nil
}

// thirdCall create an error value we will validate.
func thirdCall() error {
	return &AppError{99}
}
