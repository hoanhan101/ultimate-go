// -------------------
// Interface Pollution
// -------------------

// It comes from the fact that people are designing software from the interface first down instead
// of concrete type up.

// Why are we using an interface here?
// Myth #1: We are using interfaces because we have to use interfaces.
// No. We don't have to use interfaces. We use it when it is practical and reasonable to do so.
// Even though they are wonderful, there is a cost of using interfaces: a level of indirection and
// potential allocation, when we store concrete type inside of them. Unless the cost of that is
// worth whatever decoupling we are getting, then we shouldn't be taking the cost.

// Myth #2: We need to be able to test our code so we need to use interfaces.
// No. We must design our API that are usable for user application developer first, not our test.

// Below is an example that creates interface pollution by improperly using an interface
// when one is not needed.

package main

// Server defines a contract for tcp servers.
// This is a little bit of smell because this is some sort of APIs that going to be exposed to user
// and already that is a lot of behaviors brought in a generic interface.
type Server interface {
	Start() error
	Stop() error
	Wait() error
}

// server is our Server implementation.
// They match the name. However, that is not necessarily bad.
type server struct {
	host string
	// PRETEND THERE ARE MORE FIELDS.
}

// NewServer returns an interface value of type Server with a server implementation.
// Here is the factory function. It immediately starts to smell even worse. It is returning the
// interface value.
// It is not that functions and interfaces cannot return interface values. They can. But normally,
// that should raise a flag. The concrete type is the data that has the behavior and the interface
// normally should be used as accepting the input to the data, not necessary going out.
func NewServer(host string) Server {
	// SMELL - Storing an unexported type pointer in the interface.
	return &server{host}
}

// Start allows the server to begin to accept requests.
func (s *server) Start() error {
	// PRETEND THERE IS A SPECIFIC IMPLEMENTATION.
	return nil
}

// Stop shuts the server down.
func (s *server) Stop() error {
	// PRETEND THERE IS A SPECIFIC IMPLEMENTATION.
	return nil
}

// Wait prevents the server from accepting new connections.
func (s *server) Wait() error {
	// PRETEND THERE IS A SPECIFIC IMPLEMENTATION.
	return nil
}

func main() {
	// Create a new Server.
	srv := NewServer("localhost")

	// Use the API.
	srv.Start()
	srv.Stop()
	srv.Wait()

	// This code here couldn't care less nor would it change if srv was the concrete type, not the
	// interface. The interface is not providing any level of support whatsoever. There is no
	// decoupling here that is happening. It is not giving us anything special here. All is doing
	// is causing us another level of indirection.
}

// It smells because:
// ------------------
// - The package declares an interface that matches the entire API of its own concrete type.
// - The interface is exported but the concrete type is unexported.
// - The factory function returns the interface value with the unexported concrete type value inside.
// - The interface can be removed and nothing changes for the user of the API.
// - The interface is not decoupling the API from change.
