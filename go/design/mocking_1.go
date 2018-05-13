// ---------------
// Package To Mock
// ---------------

// It is important to mock thing.
// Most things over the network can be mocked in our test. However, mocking our database is a
// different story because it is too complex. This is where Docker can come in and simplify our
// code by allowing us to launch our database while running our tests and have that clean database
// for everything we do.

// Every API only need to focus on its test. We no longer have to worry about the application user
// or user over API test. We used to worry about: if we don't have that interface, then the user
// who use our API can't write test. That is gone. The example below will demonstrate the reason.

// Imagine we are working at a company that decides to incorporate Go as a part of its stack. They
// have their internal pubsub system that all applications are supposed to used. Maybe they are
// doing event sourcing and there is a single pubsub platform they are using that is not going to
// be replaced. They need the pubsub API for Go that they can start building services that connect
// into this event source.
// So what can change? Can the even source change?
// If the answer is no, then it immediately tells us that we don't need to use interfaces. We can
// built the entire API in the concrete, which we would do it first anyway. We then write tests to
// make sure everything work.

// A couple days later, they come to us with a problem. They have to write tests and they cannot
// hit the pubsub system directly when my test run so they need to mock that out. They want us to
// give them an interface. However, we don't need an interface because our API doesn't need an
// interface. They need an interface, not us. They need to decouple from the pubsub system, not us.
// They can do any decoupling they want because this is Go. The next file will be an example of
// their application.

// Package pubsub simulates a package that provides publication/subscription type services.

package main // should be pubsub, but leave main here for it to compile

// PubSub provides access to a queue system.
type PubSub struct {
	host string
	// PRETEND THERE ARE MORE FIELDS.
}

// New creates a pubsub value for use.
func New(host string) *PubSub {
	ps := PubSub{
		host: host,
	}

	// PRETEND THERE IS A SPECIFIC IMPLEMENTATION.

	return &ps
}

// Publish sends the data for the specified key.
func (ps *PubSub) Publish(key string, v interface{}) error {
	// PRETEND THERE IS A SPECIFIC IMPLEMENTATION.
	return nil
}

// Subscribe sets up an request to receive messages for the specified key.
func (ps *PubSub) Subscribe(key string) error {
	// PRETEND THERE IS A SPECIFIC IMPLEMENTATION.
	return nil
}
