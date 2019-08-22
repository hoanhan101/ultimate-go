package main

import "fmt"

// --------------
// Valueless type
// --------------

// reader is an interface that defines the act of reading data.
// interface is technically a valueless type. This interface doesn't declare state.
// It defines a contract of behavior. Through that contract of behavior, we have polymorphism.
// It is a 2 word data structure that has 2 pointers.
// When we say var r reader, we would have a nil value interface because interface is a reference
// type.
type reader interface {
	read(b []byte) (int, error) // (1)

	// We could have written this API a little bit differently.
	// Technically, I could have said: How many bytes do you want me to read and I will return that
	// slice of byte and an error, like so: read(i int) ([]byte, error) (2).
	// Why do we choose the other one instead?
	// Every time we call (2), it will cost an allocation because the method would have to
	// allocate a slice of some unknown type and share it back up the call stack. The method
	// would have to allocate a slice of some unknown type and share it back up the call stack. The
	// backing array for that slice has to be an allocation. But if we stick with (1), the caller
	// is allocating a slice. Even the backing array for that is ended up on a heap, it is just 1
	// allocation. We can call this 10000 times and it is still 1 allocation.
}

// -------------------------------
// Concrete type vs Interface type
// -------------------------------

// A concrete type is any type that can have a method. Only user defined type can have a method.
// Method allows a piece of data to expose capabilities, primarily around interfaces.
// file defines a system file.
// It is a concrete type because the have the method read below. It is identical to the method in
// the reader interface. Because of this, we can say the concrete type file implements the reader
// interface using a value receiver.
// There is no fancy syntax. The complier can recognize the implementation here.

// ------------
// Relationship
// ------------

// We store concrete type values inside interfaces.
type file struct {
	name string
}

// read implements the reader interface for a file.
func (file) read(b []byte) (int, error) {
	s := "<rss><channel><title>Going Go Programming</title></channel></rss>"
	copy(b, s)
	return len(s), nil
}

// pipe defines a named pipe network connection.
// This is the second concrete type that uses a value receiver.
// We now have two different pieces of data, both exposing the reader's contract and implementation
// for this contract.
type pipe struct {
	name string
}

// read implements the reader interface for a network connection.
func (pipe) read(b []byte) (int, error) {
	s := `{name: "hoanh", title: "developer"}`
	copy(b, s)
	return len(s), nil
}

func main() {
	// Create two values one of type file and one of type pipe.
	f := file{"data.json"}
	p := pipe{"cfg_service"}

	// Call the retrieve function for each concrete type.
	// Here we are passing the value itself, which means a copy of f going to pass
	// across the program boundary.
	// The compiler will ask: Does this file value implement the reader interface?
	// The answer is Yes because there is a method there using the value receiver that implements
	// its contract.
	// The second word of the interface value will store its won copy of f. The first word points
	// to a special data structure that we call the iTable.
	// The iTable serves 2 purposes:
	// - The first part describes the type of value being stored. In our case, it is the file value.
	// - The second part gives us a matrix of function pointers so we can actually execute the
	// right method when we call that through the interface.

	//       reader           iTable
	//    -----------        --------
	//   |           |      |  file  |
	//   |     *     | -->   --------
	//   |           |      |   *    | --> code
	//    -----------        --------
	//   |           |       --------
	//   |     *     | -->  | f copy | --> read()
	//   |           |       --------
	//    -----------

	// When we do a read against the interface, we can do an iTable lookup, find that read
	// associated with this type, then call that value against the read method - the concrete
	// implementation of read for this type of value.
	retrieve(f)

	// Similar with p. Now the first word of reader interface point to pipe, not file and the
	// second word points to a copy of pip value.

	//       reader           iTable
	//    -----------        -------
	//   |           |      |  pipe  |
	//   |     *     | -->   -------
	//   |           |      |    *   | --> code
	//    -----------        --------
	//   |           |       --------
	//   |     *     | -->  | p copy | --> read()
	//   |           |       --------
	//    -----------

	// The behavior changes because the data changes.
	retrieve(p)

	// Important note:
	// Later on, for simplicity, instead of drawing the a pointer pointing to iTable, we only draw
	// *pipe, like so:
	//  -------
	// | *pipe |
	//  -------
	// |   *   |  --> p copy
	//  -------
}

// --------------------
// Polymorphic function
// --------------------

// retrieve can read any device and process the data.
// This is called a polymorphic function.
// The parameter is being used here is the reader type. But it is valueless. What does it mean?
// This function will accept any data that implement the reader contract.
// This function knows nothing about the concrete and it is completely decoupled.
// It is the highest level of decoupling we can get. The algorithm is still efficient and compact.
// All we have is a level of indirection to the concrete type data values in order to be able to
// execute the algorithm.
func retrieve(r reader) error {
	data := make([]byte, 100)

	len, err := r.read(data)
	if err != nil {
		return err
	}

	fmt.Println(string(data[:len]))
	return nil
}
