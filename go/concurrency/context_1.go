// ----------------------------------------
// Store and retrieve values from a context
// ----------------------------------------

// Context package is the answer to cancellation and deadline in Go.

package main

import (
	"context"
	"fmt"
)

// user is the type of value to store in the context.
type user struct {
	name string
}

// userKey is the type of value to use for the key. The key is type specific and only values of the
// same type will match.
// When we store a value inside a context, what getting stored is not just a value but also a type
// associated with the storage. We can only pull a value out of that context if we know the type of
// value that we are looking for.
// The idea of this userKey type becomes really important when we want to store a value inside the
// context.
type userKey int

func main() {
	// Create a value of type user.
	u := user{
		name: "Hoanh",
	}

	// Declare a key with the value of zero of type userKey.
	const uk userKey = 0

	// Store the pointer to the user value inside the context with a value of zero of type userKey.
	// We are using context.WithValue because a new context value and we want to initialize that
	// with data to begin with. Anytime we work a context, the context has to have a parent
	// context. This is where the Background function comes in. We are gonna store the key uk to
	// its value (which is 0 in this case), and address of user.
	ctx := context.WithValue(context.Background(), uk, &u)

	// Retrieve that user pointer back by user the same key type value.
	// Value allows us to pass the key of the corrected type (in our case is uk of userKey type)
	// and returns an empty interface. Because we are working with an interface, we have to perform
	// a type assertion to pull the value that we store in there out the interface so we can work
	// with the concrete again.
	if u, ok := ctx.Value(uk).(*user); ok {
		fmt.Println("User", u.name)
	}

	// Attempt to retrieve the value again using the same value but of a different type.
	// Even though the key value is 0, if we just pass 0 into this function call, we are not gonna
	// get back that address to the user because 0 is based on integer type, not our userKey type.
	// It's important that when we store the value inside the context to not use the built-in type.
	// Declare our own key type. That way, only us and who understand that type can pull that out.
	// Because what if multiple partial programs want to use that value of 0, we are all being
	// tripped up on each other. That type extends an extra level of protection on being able to
	// store and retrieve value out of context.
	// If we are using this, we want to raise a flag because we have to ask twice why do we want to
	// do that instead of passing down the call stack. Because if we can pass it down the call
	// stack, it would be much better for readability and maintainability for our legacy code in
	// the future.
	if _, ok := ctx.Value(0).(*user); !ok {
		fmt.Println("User Not Found")
	}
}
