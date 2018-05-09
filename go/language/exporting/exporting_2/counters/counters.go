// Package counters provides alert counter support.
package counters

// alertCounter is an unexported named type that contains an integer counter for alerts.
type alertCounter int

// Declare an exported function called New - a factory function that knows how to create and
// initialize the value of an unexported type.
// It returns an unexported value of alertCounter.
func New(value int) alertCounter {
	return alertCounter(value)
}

// The compiler is okay with this because exporting and unexporting is not about the value like
// private and public mechanism, it is about the identifier itself.
// However, we don't do this since there is no encapsulation here. We can just make the type
// exported.
