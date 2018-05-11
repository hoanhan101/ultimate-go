// Reference types: slice, map, channel, interface, function.
// Zero value of a reference type is nil.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// ----------------------
	// Declare and initialize
	// ----------------------

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

	// ------------------
	// Length vs Capacity
	// ------------------

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
	fmt.Printf("\n=> Printing a slice\n")
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

	fmt.Printf("\n=> Length vs Capacity\n")
	inspectSlice(slice2)

	// --------------------------------------------------------
	// Idea of appending: making slice a dynamic data structure
	// --------------------------------------------------------
	fmt.Printf("\n=> Idea of appending\n")

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
	// The length is slice3 is 2 and capacity is 6.
	// Parameters are [starting_index : (starting_index + length)]
	// By looking at the output, we can see that they are sharing the same backing array.
	// Thes slice headers get to stay on the stack when we use these value semantics. Only the
	// backing array that needed to be on the heap.
	slice3 := slice2[2:4]

	fmt.Printf("\n=> Slice of slice (before)\n")
	inspectSlice(slice2)
	inspectSlice(slice3)

	// When we change the value of the index 0 of slice3, who are going to see this change?
	slice3[0] = "CHANGED"

	// The answer is both.
	// We have to always to aware that we are modifying an existing slice. We have to be aware who
	// are using it, who is sharing that backing array.
	fmt.Printf("\n=> Slice of slice (after)\n")
	inspectSlice(slice2)
	inspectSlice(slice3)

	// How about slice3 := append(slice3, "CHANGED")?
	// Similar problem will occur with append iff the length and capacity is not the same.
	// Instead of changing slice3 at index 0, we call append on slice3. Since the length of slice3
	// is 2, capacity is 6 at the moment, we have extra rooms for modification. We go and change
	// the element at index 3 of slice3, which is index 4 of slice2. That is very dangerous.

	// So, what if the length and capacity is the same? Instead of making slice3 capacity 6, we set
	// it to 2 by adding  another parameter to the slicing syntax like this: slice3 := slice2[2:4:4]
	// When append look at this slice and see that the length and capacity is the same, it wouldn't
	// bring in the element at index 4 of slice2. It would detach.
	// slice3 will have a length of 2 and capacity of 2, still share the same backing array.
	// On the call to append, length and capacity will be different. The addresses are also different.
	// This is called 3 index slice. This new slice will get its own backing array and we don't
	// affect anything at all to out original slice.

	// ------------
	// Copy a slice
	// ------------

	// copy only works with string and slice only.
	// Make a new slice big enough to hold elements of slice 1 and copy the values over using
	// the builtin copy function.
	slice4 := make([]string, len(slice2))
	copy(slice4, slice2)

	fmt.Printf("\n=> Copy a slice\n")
	inspectSlice(slice4)

	// -------------------
	// Slice and reference
	// -------------------

	// Declare a slice of integers with 7 values.
	x := make([]int, 7)

	// Random starting counters.
	for i := 0; i < 7; i++ {
		x[i] = i * 100
	}

	// Set a pointer to the second element of the slice.
	twohundred := &x[1]

	// Append a new value to the slice. This line of code raises a red flag.
	// We have x is a slice with length 7, capacity 7. Since the length and capacity is the same,
	// append doubles it size the copy values over. x nows points to diffrent memeory block and
	// has a length of 8, capacity of 14.
	x = append(x, 800)

	// When we change the value of the second element of the slice, twohundred is not gonna change
	// because it points to the old slice. Everytime we read it, we will get the wrong value.
	x[1]++

	// By printing out the output, we can see that we are in trouble.
	fmt.Printf("\n=> Slice and reference\n")
	fmt.Println("twohundred:", *twohundred, "x[1]:", x[1])

	// -----
	// UTF-8
	// -----
	fmt.Printf("\n=> UTF-8\n")

	// Everything in Go is based on UTF-8 character sets.
	// If we use different encoding scheme, we might have a problem.

	// Declare a string with both Chinese and English characters.
	// For each Chinese character, we need 3 byte for each one.
	// The UTF-8 is built on 3 layers: bytes, code point and character. From Go perspective, string
	// are just bytes. That is what we are storing.
	// In our example, the first 3 byte represents a single code point that represents that single
	// character. We can have anywhere from 1 to 4 bytes representing a code point (a code point is
	// a 32 bit value) and anywhere from 1 to multiple code points can actually represent a
	// character. To keep it simple, we only have 3 byte representing 1 code point representing 1
	// character. So we can read s as 3 byte, 3 byte, 1 byte, 1 byte,... (since there are only 2
	// Chinese characters in the first place, the rests are English)
	s := "世界 means world"

	// UTFMax is 4 -- up to 4 bytes per encoded rune -> maximum number of bytes we need to
	// represent any code point is 4.
	// Rune is its own type. It is an alias for int32 type. Similar to type byte we are using, it
	// is just an alias for uint8.
	var buf [utf8.UTFMax]byte

	// When we ranging over a string, are we doing it byte by byte or code point by code point or
	// character by character?
	// The answer is code point by code point.
	// On the first iteration, i is 0. On the next one, i is 3 because we are moving to the next
	// code point. Then i is 6.
	for i, r := range s {
		// Capture the number of bytes for this rune/code point.
		rl := utf8.RuneLen(r)

		// Calculate the slice offset for the bytes associated with this rune.
		si := i + rl

		// Copy of rune from the string to our buffer.
		// We want to go through every code point and copy them into our array buf, and display
		// them in the screen.
		// "Every array is just a slice waiting to happen." - Go saying
		// We are using the slicing syntax, creating our slice header where buf becomes the backing
		// array. All of them are on the stack. There is no allocation here.
		copy(buf[:], s[i:si])

		// Display the details.
		fmt.Printf("%2d: %q; codepoint: %#6x; encoded bytes: %#v\n", i, r, r, buf[:rl])
	}
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
