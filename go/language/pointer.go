// ----------------------------------
// Everything is about pass by value.
// ----------------------------------

// Pointers serve only 1 purpose: sharing.
// Pointers share values across the program boundaries.
// There are several types of program boundaries. The most common one is between function calls.
// We can also have a boundary between Goroutines which we will discuss later.

// When this program starts up, the runtime creates a Goroutine.
// Every Goroutine is a separate path of execution that contains instructions that needed to be executed by
// the machine. We can also think of Goroutines as lightweight threads.
// This program has only 1 Goroutine: the main Goroutine.

// Every Goroutine is given a block of memory, called the stack.
// The stack memory in Go starts out at 2K. It is very small. It can change over time.
// Every time a function is called, a piece of stack is used to help that function run.
// The growing direction of the stack is downward.

// Every function is given a stack frame, memory execution of a function.
// The size of every stack frame is known at compile time. No value can be placed on a stack
// unless the compiler knows its size ahead of time.
// If we don't know the size of something at compile time, it has to be on the heap.

// Zero value enables us to initialize every stack frame that we take.
// Stacks are self cleaning. We clean our stack on the way down.
// Every time we make a function, zero value initialization cleans the stack frame.
// We leave that memory on the way up because we don't know if we would need that again.

package main

// user represents an user in the system.
type user struct {
	name  string
	email string
}

func main() {
	// -------------
	// Pass by value
	// -------------

	// Declare variable of type int with a value of 10.
	// This value is put on a stack with a value of 10.
	count := 10

	// To get the address of a value, we use &.
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// Pass the "value of" count.
	increment1(count)

	// Printing out the result of count. Nothing has changed.
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// Pass the "address of" count.
	// This is still considered pass by value, not by reference because the address itself is a value.
	increment2(&count)

	// Printing out the result of count. count is updated.
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// ---------------
	// Escape analysis
	// ---------------

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
// Since we know the size of the user value at compile time, the compiler will put this on a stack
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
// This looks almost identical to the stayOnStack function.
// It creates a value of type user and initializes it. It seems like we are doing the same here.
// However, there is one subtle difference: we do not return the value itself but the address
// of u. That is the value that is being passed back up the call stack. We are using pointer
// semantic.

// You might think about what we have after this call is: main has a pointer to a value that is
// on a stack frame below. If this is the case, then we are in trouble.
// Once we come back up the call stack, this memory is there but it is reusable again. It is no
// longer valid. Anytime now main makes a function call, we need to allocate the frame and
// initialize it.

// Think about zero value for a second here. It enables us to initialize every stack frame that
// we take. Stacks are self cleaning. We clean our stack on the way down. Every time we make a
// function call, zero value, initialization, we are cleaning those stack frames. We leave that
// memory on the way up because we don't know if we need that again.

// Back to the example. It is bad because it looks like we take the address of user value, pass it
// back up to the call stack giving us a pointer which is about to get erased.
// However, that is not what will happen.

// What is actually going to happen is escape analysis.
// Because of the line "return &u", this value cannot be put inside the stack frame for this function
// so we have to put it out on the heap.
// Escape analysis decides what stays on the stack and what does not.
// In the stayOnStack function, because we are passing the copy of the value itself, it is safe to
// keep these things on the stack. But when we SHARE something above the call stack like this,
// escape analysis said this memory is no longer valid when we get back to main, we must put it
// out there on the heap. main will end up having a pointer to the heap.
// In fact, this allocation happens immediately on the heap. escapeToHeap is gonna have a pointer
// to the heap. But u is gonna base on value semantic.
func escapeToHeap() *user {
	u := user{
		name:  "Hoanh An",
		email: "hoanhan@bennington.edu",
	}

	return &u
}

// ----------------------------------
// What if we run out of stack space?
// ----------------------------------

// What happens next is during that function call, there is a little preamble that asks "Do we have
// enough stack space for this frame?". If yes then no problem because at complie time we know
// the size of every frame. If not, we have to have bigger frame and these values need to be copied
// over. The memory on that stack moves. It is a trade off. We have to take the cost of this copy
// because it doesn't happen a lot. The benefit of using less memory in any Goroutine outweighs the
// cost.

// Because the stack can grow, no Goroutine can have a pointer to some other Goroutine stack.
// There would be too much overhead for the compiler to keep track of every pointer. The latency will
// be insane.
// -> The stack for a Goroutine is only for that Goroutine. It cannot be shared between Goroutines.

// ------------------
// Garbage collection
// ------------------

// Once something is moved to the heap, Garbage Collection has to get in.
// The most important thing about the Garbage Collector (GC) is the pacing algorithm.
// It determines the frequency/pace that the GC has to run in order to maintain the smallest possible t.

// Imagine a program where you have a 4 MB heap. GC is trying to maintain a live heap of 2 MB.
// If the live heap grows beyond 4 MB, we have to allocate a larger heap.
// The pace the GC runs at depends on how fast the heap grows in size. The slower the
// pace, the less impact it is going to have. The goal is to get the live heap back down.

// When the GC is running, we have to take a performance cost so all Goroutines can keep running
// concurrently. The GC also has a group of Goroutines that perform the garbage collection work.
// It uses 25% of our available CPU capacity for itself.
// More details about GC and pacing algorithm can be find at:
// https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/pointers/README.md
