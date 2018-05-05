// Idea of reference type: slice, map, channel, interface, function.
// Zero value of a reference type is nil.

package main

import "fmt"

func main() {
	// ------------------
	// Declare and length
	// ------------------

	// Create a slice with a length of 5 elements.
	// make is a special built-in function that only works with slice, map and channel.
	// make creates a slice that has an array of 5 strings behind it. We are getting back a 3 word
	// data structure: the first word points to the backing array, second word is length and third
	// one is capacity.
	//  -----
	// |  *  | --> | nil | nil | nil | nil | nil |
	//  -----      |  0  |  0  |  0  |  0  |  0  |
	// |  5  |
	//  -----
	// |  5  |
	//  -----

	// Difference between length and capacity
	// Length is the number of elements from this pointer position we have access to (read and write).
	// Capacity is the total number of elements from this pointer position that exist in the
	// backing array.

	// Syntactic sugar -> looks like array
	// It also have the same cost that we've seen in array.
	// One thing to be mindful about: there is no value in the bracket []string inside the make
	// function. With that in mind, we can constantly notice that we are dealing with a slice, not
	// array.
	slice1 := make([]string, 5)
	slice1[0] = "Apple"
	slice1[1] = "Orange"
	slice1[2] = "Banana"
	slice1[3] = "Grape"
	slice1[4] = "Plum"

	// We can't access an index of a slice beyond its length.
	// Error: panic: runtime error: index out of range
	// slice1[5] = "Runtime error"

	// We are passing the value of slice, not its address. So the Println function will have its
	// own copy of the slice.
	fmt.Println(slice1)

	// --------------
	// Reference type
	// --------------

	// Create a slice with a length of 5 elements and a capacity of 8.
	// make allows us to adjust the capacity directly on construction of this initialization.
	// What we end up having now is a 3 word data structure where the first word points to an array
	// of 8 elements, length is 5 and capacity is 8.
	//  -----
	// |  *  | --> | nil | nil | nil | nil | nil | nil | nil | nil |
	//  -----      |  0  |  0  |  0  |  0  |  0  |  0  |  0  |  0  |
	// |  5  |
	//  -----
	// |  8  |
	//  -----
	// It means that I can read and write to the first 5 elements and I have 3 elements of capacity
	// that I can leverage later.
	slice2 := make([]string, 5, 8)
	slice2[0] = "Apple"
	slice2[1] = "Orange"
	slice2[2] = "Banana"
	slice2[3] = "Grape"
	slice2[4] = "Plum"

	inspectSlice(slice2)

	// --------------------------------------------------------
	// Idea of appending: making slice a dynamic data structure
	// --------------------------------------------------------

	// Declare a nil slice of strings, set to its zero value.
	// 3 word data structure: first one points to nil, second and last are zero.
	var data []string

	// What if I do data := string{}? Is it the same?
	// No because data in this case is not set to its zero value.
	// This is why we always use var for zero value because not every type when we create an empty
	// literal we have its zero value in return.
	// What actually happen here is that we have a slice but it has a pointer (as opposed to nil).
	// This is consider an empty slice, not a nil slice.
	// There is a semantic between a nil slice and an empty slice. Any reference type that set to
	// its zero value can be considered nil. If we pass a nil slice to a marshal function, we get
	// back a string that said null but when we pass an empty slice, we get an empty JSON document.
	// But where does that pointer point to? It is an empty struct, which we will review later.

	// Capture the capacity of the slice.
	lastCap := cap(data)

	// Append ~100k strings to the slice.
	for record := 1; record <= 102400; record++ {
		// Use the built-in function append to add to the slice.
		// It allows us to add value to a slice, making the data structure dynamic, yet still
		// allows us to use that contiguous block of memory that gives us the predictable access
		// pattern from mechanical sympathy.
		// The append call is working with value semantic. We are not sharing this slice but
		// appending to it and returning a new copy of it. The slice gets to stay on the stack, not
		// heap.
		data = append(data, fmt.Sprintf("Rec: %d", record))

		// Every time append runs, it checks the length and capacity.
		// If it the same, it means that we have no room. append creates a new backing array,
		// double it size, copy the old value back in and append the new value. It mutates its copy
		// on its stack frame and return us a copy. We replace our slice with the new copy.
		// It it not the same, it means that we have extra elements of capacity we can use. Now we
		// can bring these extra capacity into the length and no copy is being made. This is very
		// efficient.

		// Looking at the last column in the output, when the backing array is 1000 elements or
		// less, it doubles the size of the backing array for growth. Once we pass 1000 elements,
		// growth rate moves to 25%.

		// When the capacity of the slice changes, display the changes.
		if lastCap != cap(data) {
			// Calculate the percent of change.
			capChg := float64(cap(data)-lastCap) / float64(lastCap) * 100

			// Save the new values for capacity.
			lastCap = cap(data)

			// Display the results.
			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n", &data[0], record, cap(data), capChg)
		}
	}

	// --------------
	// Slice of slice
	// --------------

	// Take a slice of slice2. We want just indexes 2 and 3.
	// Parameters are [starting_index : (starting_index + length)]
	// By looking at the output, we can see that they are sharing the same backing array.
	// Thes slice headers get to stay on the stack when we use these value semantics. Only the
	// backing array that needed to be on the heap.
	slice3 := slice2[2:4]
	inspectSlice(slice3)

	// When we change the value of the index 0 of slice3, who are going to see this change?
	slice3[0] = "CHANGED"

	// The answer is both.
	// We have to always to aware that we are modifying an existing slice. We have to be aware who
	// are using it, who is sharing that backing array.
	inspectSlice(slice2)
	inspectSlice(slice3)

	// Similar problem will occur with append iff the length and capacity is not the same.
	// Instead of changing slice3 at index 0, we call append on slice3. Since the length of slice3
	// is 2, capacity is 6 at the moment, we have extra rooms for modification. We go and change
	// the element at index 3 of slice3, which is index 4 of slice2. That is very dangerous.

	// So, what if the length and capacity is the same?
	// When append look at this slice and see that the length and capacity is the same, it wouldn't
	// bring in the element at index 4 of slice2. It would detach.
	// When we add another parameter to the slicing syntax that set the capacity to be the same
	// like this: slice3 := slice2[2:4:4]
	// slice3 will have a length of 2 and capacity of 2, still share the same backing array.
	// On the call to append, length and capacity will be different. The addresses are also different.
	// This is called 3 index slice. This new slice will get its own backing array and we don't
	// affect anything at all to out original slice.
}

// inspectSlice exposes the slice header for review.
// Parameter: again, there is no value in side the []string so we want a slice.
// Range over a slice, just like we did with array.
// While len tells us the length, cap tells us the capacity
// In the output, we can see the addresses are aligning as expected.
func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i := range slice {
		fmt.Printf("[%d] %p %s\n", i, &slice[i], slice[i])
	}
}
