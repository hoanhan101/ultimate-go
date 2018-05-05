// Constant are not variables.
// Contants have a parallel type system all to themselves. The minimum precision for constant is
// 256 bit. They are considered to be mathematically exact.
// Constants only exist at complied time.

package main

import "fmt"

func main() {
	// ----------------------
	// Declare and initialize
	// ----------------------

	// Constant can be typed or untyped.
	// When it is untyped, we consider it as a kind.
	// They are implicitly converted by the complier.

	// Untyped Constants.
	const ui = 12345    // kind: integer
	const uf = 3.141592 // kind: floating-point

	fmt.Println(ui)
	fmt.Println(uf)

	// Typed Constants still use the constant type system but their precision is restricted.
	const ti int = 12345        // type: int
	const tf float64 = 3.141592 // type: float64

	fmt.Println(ti)
	fmt.Println(tf)

	// This doesn't work because constant 1000 overflows uint8.
	// const myUint8 uint8 = 1000

	// Constant arithmetic supports different kinds.
	// Kind Promotion is used to determine kind in these scenarios.
	// All of this happen implicitly.

	// Variable answer will of type float64.
	var answer = 3 * 0.333 // KindFloat(3) * KindFloat(0.333)

	fmt.Println(answer)

	// Constant third will be of kind floating point.
	const third = 1 / 3.0 // KindFloat(1) / KindFloat(3.0)

	fmt.Println(third)

	// Constant zero will be of kind integer.
	const zero = 1 / 3 // KindInt(1) / KindInt(3)

	fmt.Println(zero)

	// This is an example of constant arithmetic between typed and
	// untyped constants. Must have like types to perform math.
	const one int8 = 1
	const two = 2 * one // int8(2) * int8(1)

	fmt.Println(one)
	fmt.Println(two)

	// Max integer value on 64 bit architecture.
	const maxInt = 9223372036854775807

	fmt.Println(maxInt)

	// Much larger value than int64 but still compile because of untyped system.
	// 256 is a lot of space (depending on the architecture)
	// const bigger = 9223372036854775808543522345

	// Will NOT compile because it exceeds 64 bit
	// const biggerInt int64 = 9223372036854775808543522345

	// ----
	// iota
	// ----

	const (
		A1 = iota // 0 : Start at 0
		B1 = iota // 1 : Increment by 1
		C1 = iota // 2 : Increment by 1
	)

	fmt.Println("1:", A1, B1, C1)

	const (
		A2 = iota // 0 : Start at 0
		B2        // 1 : Increment by 1
		C2        // 2 : Increment by 1
	)

	fmt.Println("2:", A2, B2, C2)

	const (
		A3 = iota + 1 // 1 : Start at 0 + 1
		B3            // 2 : Increment by 1
		C3            // 3 : Increment by 1
	)

	fmt.Println("3:", A3, B3, C3)

	const (
		Ldate         = 1 << iota //  1 : Shift 1 to the left 0.  0000 0001
		Ltime                     //  2 : Shift 1 to the left 1.  0000 0010
		Lmicroseconds             //  4 : Shift 1 to the left 2.  0000 0100
		Llongfile                 //  8 : Shift 1 to the left 3.  0000 1000
		Lshortfile                // 16 : Shift 1 to the left 4.  0001 0000
		LUTC                      // 32 : Shift 1 to the left 5.  0010 0000
	)

	fmt.Println("Log:", Ldate, Ltime, Lmicroseconds, Llongfile, Lshortfile, LUTC)
}
