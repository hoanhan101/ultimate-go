// ---------------------
// Interface Conversions
// ---------------------

package main

import "fmt"

// Mover provides support for moving things.
type Mover interface {
	Move()
}

// Locker provides support for locking and unlocking things.
type Locker interface {
	Lock()
	Unlock()
}

// MoveLocker provides support for moving and locking things.
type MoveLocker interface {
	Mover
	Locker
}

// bike represents a concrete type for the example.
type bike struct{}

// Move can change the position of a bike.
func (bike) Move() {
	fmt.Println("Moving the bike")
}

// Lock prevents a bike from moving.
func (bike) Lock() {
	fmt.Println("Locking the bike")
}

// Unlock allows a bike to be moved.
func (bike) Unlock() {
	fmt.Println("Unlocking the bike")
}

func main() {
	// Declare variables of the MoveLocker and Mover interfaces set to their zero value.
	var ml MoveLocker
	var m Mover

	// Create a value of type bike and assign the value to the MoveLocker interface value.
	ml = bike{}

	// An interface value of type MoveLocker can be implicitly converted into
	// a value of type Mover. They both declare a method named move.
	m = ml

	//    ml                          m
	//  ------                      ------
	// | bike |        bike        | bike |
	//  ------        ------        ------
	// |  *   | ---> |      | <--- |      |
	//  ------        ------        ------

	// However, we cannot go in the other direction, like so:
	// ml = m

	// The compiler will say:
	// cannot use m (type Mover) as type MoveLocker in assignment: Mover does not
	// implement MoveLocker (missing Lock method).

	// --------------
	// Type assertion
	// --------------

	// Interface type Mover does not declare methods named lock and unlock. Therefore, the compiler
	// can't perform an implicit conversion to assign a value of interface type Mover to an
	// interface value of type MoveLocker. It is irrelevant that the concrete type value of
	// type bike that is stored inside of the Mover interface value implements the MoveLocker interface.

	// We can perform a type assertion at runtime to support the assignment.

	// Perform a type assertion against the Mover interface value to access a COPY of the concrete type
	// value of type bike that was stored inside of it. Then assign the COPY of the concrete type
	// to the MoveLocker interface.

	// This is the syntax for type assertion.
	// We are taking the interface value itself, dot (bike). We are using bike as an parameter.
	// If there is a bike inside of m, we will get a copy of it since we are using value semantic.
	// b is having a copy of bike value.
	b := m.(bike)
	ml = b

	// It's important to note that the type assertion syntax provides a way to state what type
	// of value is stored inside the interface. This is more powerful from a language and readability
	// standpoint, than using a casting syntax, like in other languages.
}
