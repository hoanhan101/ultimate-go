// Methods are really just made up. They are not real. They are just syntactic sugar.
// They give us a believe system that some pieces of data exposed some capabilities.
// Object-oriented programming has driven design and capabilities. However, there is no OOP in Go.
// There is data and behavior. At some time, data can expose some capabilities, but for specific
// purposes, not to really design API around. Methods are really just functions.

package main

import "fmt"

// data is a struct to bind methods to.
type data struct {
	name string
	age  int
}

// displayName provides a pretty print view of the name.
// It uses data as a value receiver.
func (d data) displayName() {
	fmt.Println("My Name Is", d.name)
}

// setAge sets the age and displays the value.
// It uses data as a pointer receiver.
func (d *data) setAge(age int) {
	d.age = age
	fmt.Println(d.name, "Is Age", d.age)
}

func main() {
	// --------------------------
	// Methods are just functions
	// --------------------------

	// Declare a variable of type data.
	d := data{
		name: "Hoanh",
	}

	fmt.Println("Proper Calls to Methods:")

	// How we actually call methods in Go.
	d.displayName()
	d.setAge(21)

	fmt.Println("\nWhat the Compiler is Doing:")

	// This is what Go is doing underneath.
	// When we call d.displayName(), the compiler will call data.displayName, showing that we are
	// using a value receiver of type data, and pass the data in as the first parameter.
	// Taking a look at the function again: "func (d data) displayName()", that receiver is the
	// parameter because it is truly a parameter. It is the first parameter to a function that call
	// displayName.
	// Similar to d.setAge(45). Go is calling a function that based on the pointer receiver and
	// passing data to its parameters. We are adjusting to make the call by taking the address of d.
	data.displayName(d)
	(*data).setAge(&d, 21)

	// -----------------
	// Function variable
	// -----------------

	fmt.Println("\nCall Value Receiver Methods with Variable:")

	// Declare a function variable for the method bound to the d variable.
	// The function variable will get its own copy of d because the method is using a value receiver.
	// f1 is now a reference type: a pointer variable.
	// We don't call the method here. There is no () at the end of displayName.
	f1 := d.displayName

	// Call the method via the variable.
	// f1 is pointer and it points to a special 2 word data structure. The first word points to the
	// code for that method we want to execute, which is displayName in this case. We cannot call
	// displayName unless we have a value of type data. So the second word is a pointer to the
	// copy of data. displayName uses a value receiver so it works on its own copy. When we make
	// an assignment to f1, we are having a copy of d.
	//  -----
	// |  *  | --> code
	//  -----
	// |  *  | --> copy of d
	//  -----
	f1()

	// When we change the value of d to "Hoanh An", f1 is not going to see the change.
	d.name = "Hoanh An"

	// Call the method via the variable. We don't see the change.
	f1()

	// However, if we do this again if f2, then we will see the change.

	fmt.Println("\nCall Pointer Receiver Method with Variable:")

	// Declare a function variable for the method bound to the d variable.
	// The function variable will get the address of d because the method is using a pointer receiver.
	f2 := d.setAge

	// Call the method via the variable.
	// f2 is also a pointer that has 2 word data structure. The first word points to setAge, but
	// the second words doesn't point to its copy any more, but to its original.
	//  -----
	// |  *  | --> code
	//  -----
	// |  *  | --> original d
	//  -----

	// Change the value of d.
	d.name = "Hoanh An Dinh"

	// Call the method via the variable. We see the change.
	f2(21)
}
