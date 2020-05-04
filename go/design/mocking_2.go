// ------
// Client
// ------

// Run: `go run ./go/design/mocking_1.go ./go/design/mocking_2.go`

// Sample program to show how we can personally mock concrete types when we need to for
// our own packages or tests.
package main

import (
	"fmt"
)

// publisher is an interface to allow this package to mock the pubsub package.
// When we are writing our applications, declare our own interface that map out all the APIs call
// we need for the APIs. The concrete types APIs in the previous files satisfy it out of the box.
// We can write the entire application with mocking decoupling from concrete implementations.
type publisher interface {
	Publish(key string, v interface{}) error
	Subscribe(key string) error
}

// mock is a concrete type to help support the mocking of the pubsub package.
type mock struct{}

// Publish implements the publisher interface for the mock.
func (m *mock) Publish(key string, v interface{}) error {
	// ADD YOUR MOCK FOR THE PUBLISH CALL.
	fmt.Println("Mock PubSub: Publish")
	return nil
}

// Subscribe implements the publisher interface for the mock.
func (m *mock) Subscribe(key string) error {
	// ADD YOUR MOCK FOR THE SUBSCRIBE CALL.
	fmt.Println("Mock PubSub: Subscribe")
	return nil
}

func main() {
	// Create a slice of publisher interface values. Assign the address of a pubsub.
	// PubSub value and the address of a mock value.
	pubs := []publisher{
		New("localhost"),
		&mock{},
	}

	// Range over the interface value to see how the publisher interface provides
	// the level of decoupling the user needs. The pubsub package did not need
	// to provide the interface type.
	for _, p := range pubs {
		p.Publish("key", "value")
		p.Subscribe("key")
	}
}
