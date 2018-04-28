// Everything is about pass by value

// Every Go routine is a separate path of execution that contains instructions to be executed by
// the machine. Can think of it as thread.
// Every Go routine is given a block of memory, called the stack.
// The stack memory in Go starts out at 2K. It is very small. It can change over time.
// Every time a function is called, a piece of stack is used to help that function run.
// The direction of the stack is downward.
// Every function is given a stack frame, memory execution of a function.
// The size of every stack frame is known at compiled time. No value can be placed on a stack
// unless the compiler knows its size ahead of time.
// If we don't know the size of something at compiled time, it has to be on the heap.

package main

func main() {

	// Declare variable of type int with a value of 10.
	// This value is put on a stack with a value of 10.

	count := 10

	// To get the address of a value, we use &

	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// Pass the "value of" count.

	increment1(count)

	// Pass the "address of" count.
	// This is still considered pass by value, not by reference because the address itself is a value.

	increment2(&count)

	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
}

func increment1(inc int) {
	// Increment the "value of" inc.
	inc++
	println("inc1:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
}

// increment2 declares count as a pointer variable whose value is
// always an address and points to values of type int.
// The * here is not an operator. It is part of the type name.
// Every type that is declared, whether you do it or it is predeclared, you get for free a pointer.
func increment2(inc *int) {
	// Increment the "value of" count that the "pointer points to".
	// The * is an operator. It tells us the value of the pointer points to.
	*inc++
	println("inc2:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tValue Points To[", *inc, "]")
}
