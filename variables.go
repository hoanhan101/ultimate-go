package main

import "fmt"

func main() {

	// Type provides integrity and readability.
	// - What is the amount of memory that we allocate?
	// - What does that memory represent?

	// Type can be specific such as int32 or int64.
	// When we declare a type without being very specific, such as int, Go will choose it base on
	// our architecture. Go has something called a word size, which matches our address size.
	// For example, in 64-bit architecture, our word size is 64, address size is 64 then our
	// integer should be 64.

	// Zero value concept
	// Every single value we create must be initialized. If we don't specify it, it will be set to
	// the zero value. The entire allocation of memory, we reset that bit to 0.
	// - Boolean false
	// - Integer 0
	// - Floating Point 0
	// - Complex 0i
	// - String "" (empty string)
	// - Pointer nil

	// String type
	// String is a 2 word data structure: first word represent a pointer to a backing array, the
	// second word represent it length.
	// If it is a zero value then the first word is nil, the second word is 0

	// var is the only guarantee to initialize a zero value for a type.

	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n\n", d, d)

	// Using the short variable declaration operator, we can define and initialize at the same time.

	aa := 10
	bb := "hello" // 1st word points to a array of characters, 2nd word is 5 bytes
	cc := 3.14159
	dd := true

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.14159 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n\n", dd, dd)

	// Conversion vs casting
	// Go doesn't have casting, but conversion.
	// Instead of telling a compiler to pretend to have some more bytes, we have to allocate more
	// memory.

	// Specify type and perform a conversion.

	aaa := int32(10)

	fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa)
}
