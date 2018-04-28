// Everything is about pass by value.
// Pointer serves only 1 purpose: sharing.
// Pointer shares values across the program boundary.
// There are several types of program boundary. The most common one is between function calls.
// We can also have a boundary between Go routines when we will discuss it later.

// When this program starts up, the runtime creates a Go routine.
// Every Go routine is a separate path of execution that contains instructions that needed to be executed by
// the machine. Can also think of it as thread.
// This program has only 1 Go routine: the main Go routine.

// Every Go routine is given a block of memory, called the stack.
// The stack memory in Go starts out at 2K. It is very small. It can change over time.
// Every time a function is called, a piece of stack is used to help that function run.
// The direction of the stack is downward.
// Every function is given a stack frame, memory execution of a function.
// The size of every stack frame is known at compiled time. No value can be placed on a stack
// unless the compiler knows its size ahead of time.
// If we don't know the size of something at compiled time, it has to be on the heap.

// Zero value enables us to initialize every stack frame that we take.
// Stacks are self cleaning. We clean our stack on the way down.
// Every time we make a function, zero value initialization cleaning stack frame.
// We leave that memory on the way up because we don't know if we would need that again.

package main

// user represents an user in the system.
type user struct {
	name  string
	email string
}

func main() {
	// Declare variable of type int with a value of 10.
	// This value is put on a stack with a value of 10.
	count := 10

	// To get the address of a value, we use &.
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// Pass the "value of" count.
	increment1(count)

	// Printing out the result of count. Nothing is change.
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// Pass the "address of" count.
	// This is still considered pass by value, not by reference because the address itself is a value.
	increment2(&count)

	// Printing out the result of count. count is updated.
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// Value semantic vs pointer semantic.
	stayOnStack()
	escapeToHeap()
}

func increment1(inc int) {
	// Increment the "value of" inc.
	inc++
	println("inc1:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
}

// increment2 declares count as a pointer variable whose value is always an address and points to
// values of type int.
// The * here is not an operator. It is part of the type name.
// Every type that is declared, whether you declare or it is predeclared, you get for free a pointer.
func increment2(inc *int) {
	// Increment the "value of" count that the "pointer points to".
	// The * is an operator. It tells us the value of the pointer points to.
	*inc++
	println("inc2:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tValue Points To[", *inc, "]")
}

// stayOnStack shows how the variable does not escape.
// Since we know the size of the user value at compiled time, the complier will put this on a stack
// frame.
func stayOnStack() user {
	// In the stayOnStack stack frame, create a value and initialize it.
	u := user{
		name:  "Hoanh An",
		email: "hoanhan@bennington.edu",
	}

	// Take the value and return it, pass back up to main stack frame.
	return u
}

// escapeToHeap shows how the variable escape.
func escapeToHeap() *user {
	// In the escapeToHeap stack frame, create a value and initialize it.
	u := user{
		name:  "Hoanh An",
		email: "hoanhan@bennington.edu",
	}

	// Return the address, not the value. Want to share it up the call stack.
	// Because of this, this value cannot be put on the stack frame but out in the heap.
	// In the stayOnStack, we are passing a copy of the value itself, it is safe to
	// keep on the stack.
	// But when we share something up the call stack like this, this memory is no longer gonna be
	// valid when it get back to main. It must be put on the heap. What end up happen is that main
	// will have a pointer to that memory on the heap.

	// In fact, this allocation happens immediately on the heap.
	// escapeToHeap has a pointer but u is based on value semantic.
	return &u
}
