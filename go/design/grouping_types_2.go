// --------------------
// Grouping By Behavior
// --------------------

// This is an example of using composition and interfaces.
// This is something we want to do in Go.
// This pattern does provide a good design principle in a Go program.

// We will group common types by their behavior and not by their state.
// What brilliant about Go is that it doesn't have to be configured ahead of time. The compiler
// identifies interface and behaviours at compiled time. In means that we can write code today that
// compliant with any interface that exists today or tomorrow. It doesn't matter where that is
// declared because the compiler can do this on the fly.

// Stop thinking about a concrete base type. Let's think about what we do instead.

package main

import "fmt"

// Speaker provide a common behavior for all concrete types to follow if they want to be a
// part of this group. This is a contract for these concrete types to follow.
// We get rid of the Animal type.
type Speaker interface {
	Speak()
}

// Dog contains everything a Dog needs.
type Dog struct {
	Name       string
	IsMammal   bool
	PackFactor int
}

// Speak knows how to speak like a dog.
// This makes a Dog now part of a group of concrete types that know how to speak.
func (d Dog) Speak() {
	fmt.Println("Woof!",
		"My name is", d.Name,
		", it is", d.IsMammal,
		"I am a mammal with a pack factor of", d.PackFactor)
}

// Cat contains everything a Cat needs.
// A little copy and paste can go a long way. Decoupling, in many cases, is a much better option
// than reusing the code.
type Cat struct {
	Name        string
	IsMammal    bool
	ClimbFactor int
}

// Speak knows how to speak like a cat.
// This makes a Cat now part of a group of concrete types that know how to speak.
func (c Cat) Speak() {
	fmt.Println("Meow!",
		"My name is", c.Name,
		", it is", c.IsMammal,
		"I am a mammal with a climb factor of", c.ClimbFactor)
}

func main() {
	// Create a list of Animals that know how to speak.
	speakers := []Speaker{
		// Create a Dog by initializing its Animal parts and then its specific Dog attributes.
		Dog{
			Name:       "Fido",
			IsMammal:   true,
			PackFactor: 5,
		},

		// Create a Cat by initializing its Animal parts and then its specific Cat attributes.
		Cat{
			Name:        "Milo",
			IsMammal:    true,
			ClimbFactor: 4,
		},
	}

	// Have the Animals speak.
	for _, spkr := range speakers {
		spkr.Speak()
	}
}

// ---------------------------------
// Guidelines around declaring types
// ---------------------------------

// - Declare types that represent something new or unique. We don't want to create aliases just for readability.
// - Validate that a value of any type is created or used on its own.
// - Embed types not because we need the state but because we need the behavior. If we not thinking
// about behavior, we are really locking ourselves into the design that we cannot grow in the future.
// - Question types that are an alias or abstraction for an existing type.
// - Question types whose sole purpose is to share common state.
