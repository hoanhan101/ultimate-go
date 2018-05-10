// ----------------------------------
// Everything is about pass by value.
// ----------------------------------

// Pointer serves only 1 purpose: sharing.
// Pointer shares values across the program boundary.
// There are several types of program boundary. The most common one is between function calls.
// We can also have a boundary between Goroutines when we will discuss it later.

// When this program starts up, the runtime creates a Goroutine.
// Every Goroutine is a separate path of execution that contains instructions that needed to be executed by
// the machine. Can also think of Goroutine as a lightweight thread.
// This program has only 1 Goroutine: the main Goroutine.

// Every Goroutine is given a block of memory, called the stack.
// The stack memory in Go starts out at 2K. It is very small. It can change over time.
// Every time a function is called, a piece of stack is used to help that function run.
// The growing direction of the stack is downward.

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

	// Printing out the result of count. Nothing is change.
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
// This looks almost identical to the stayOnStack function.
// It creates a value of type user and initialize it. It seems like we are doing the same here.
// However, there is one subtle difference: we do not return the value itself but the address
// of u. That is the value that is being passed back up the call stack. We are using pointer
// semantic.

// You might think about what we have after this call is: main has a pointer to a value that is
// on a stack frame below. It is the case then we are in trouble.
// Once we come back up the call stack, this memory is there but it is reusable again. It is no
// longer valid. Anytime now main makes a function call, we need to allocate the frame and
// initialize it.

// Think about zero value for a second here. It is enable to us to initialize every stack frame that
// we take. Stack are self cleaning. We clean our stack on the way down. Every time we make a
// function call, zero value, initialization, we are cleaning those stack frames. We leave that
// memory on the way up because we don't know if we need that again.

// Back to the example, it is bad because it looks like we take the address of user value, pass it
// back up to the call stack and we now have a pointer which is about to get erased. Therefore, it
// is not what will happen.

// What actually going to happen is the idea of escape analysis.
// Because of line "return &u", this value cannot be put inside the stack frame for this function
// so we have to put it out on the heap.
// Escape analysis decides what stay on stack and what not.
// In the stayOnStack function, because we are passing the copy of the value itself, it is safe to
// keep these things on the stack. But when we SHARE something above the call stack like this,
// escape analysis said this memory is no longer be valid when we get back to main, we must put it
// out there on the heap. main is end up having a pointer to the heap.
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

// What happen next is during that function call, there is a little preamble that asks "Do we have
// enough stack space? for this frame?". If yes then no problem because at complied time we know
// the size of every frame. If no, we have to have bigger frame and these values need to be copy
// over. The memory on that stack move. It is a trade off. We have to take the cost of this copy
// because it doesn't happen a lot. The benefit of using less memory any Goroutine outrace the
// cost.

// Because stack can grow, no Goroutine can have a pointer to some other Goroutine stack.
// There would be too much overhead for complier to keep track of every pointer. The latency will
// be insane.
// -> The stack for a Goroutine is only for that Goroutine only. It cannot be shared between
// Goroutine.

// ------------------
// Garbage collection
// ------------------

// Once something is moved to the heap, Garbage Collection has to get in.
// The most important thing about the Garbage Collector (GC) is the pacinng algorithm.
// It determines the frequency/pace that the GC has to run in order to maintain the smallest t as
// possible.

// Image a program where you have a 4 MB heap. GC is trying to maintain a live heap of 2 MB.
// If the live heap grow pass 4 MB we have allocate a larger heap.
// Depending how fast the heap grow, we determine the pace that the GC has to run. The smaller the
// pace, the less impact it is going to have. The goal is to get the live heap back down.

// When the GC is running, we have to take a performance cost so all Goroutine can keep running
// concurrently. The GC also have a group of Goroutine that perform the garbage collection work.
// It takes 25% of our available CPU capacity to itself.
// More details about GC and pacing algorithm can be find at:
// https://github.com/ardanlabs/gotraining/blob/master/topics/go/language/pointers/README.md
