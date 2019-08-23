package main

import "fmt"

// notifier is an interface that defines notification type behavior.
type notifier interface {
	notify()
}

// printer displays information.
type printer interface {
	print()
}

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// print displays user's name and email.
func (u user) print() {
	fmt.Printf("My name is %s and my email is %s\n", u.name, u.email)
}

// ------------------------------
// Interface via pointer receiver
// ------------------------------

// notify implements the notifier interface with a pointer receiver.
func (u *user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

// String implements the fmt.Stringer interface.
// The fmt package that we've been using to display things on the screen, if it receive a piece of
// data that implement this behavior, it will use this behavior and overwrite its default.
// Since we are using pointer semantic, only pointer satisfies the interface.
func (u *user) String() string {
	return fmt.Sprintf("My name is %q and my email is %q", u.name, u.email)
}

func main() {
	// Create a value of type User
	u := user{"Hoanh", "hoanhan@email.com"}

	// Call polymorphic function but passing u using value semantic: sendNotification(u).
	// However, the compiler doesn't allow it:
	// "cannot use u (type user) as type notifier in argument to sendNotification:
	// user does not implement notifier (notify method has pointer receiver)"
	// This is setting up for an integrity issue.

	// ----------
	// Method set
	// ----------

	// In the specification, there are sets of rules around the concepts of method sets. What we
	// are doing is against these rules.

	// What are the rules are?
	// For any value of a given type T, only those methods implemented with a value receiver
	// belong to the method sets of that type.
	// For any value of a given type *T (pointer of a given type), both value receiver and pointer
	// receiver methods belong to the method sets.
	// In other word, if we are working with a pointer of some types, all the methods that has been
	// declared are associated with that pointer. But if we are working with a value of some types,
	// only those methods that operated on value semantic can be applied.

	// In the previous lesson about method, we are calling them before without any problem. That is
	// true. When we are dealing with method, method call against the concrete values themselves,
	// Go can adjust to make the call.
	// However, we are not trying to call a method here. We are trying to store a concrete type
	// value inside the interface. For that to happen, that value must satisfy the contract.

	// The question now becomes: Why can't pointer receiver be associated with the method sets for
	// value? What is the integrity issue here that doesn't allow us to use pointer semantic for
	// value of type T?

	// It is not 100% guarantee that any value that can satisfy the interface has an address.
	// We can never call a pointer receiver because if that value doesn't have an address, it is
	// not shareable. For example:
	//      Declare a type name duration that is based on an integer
	//      type duration int

	//      Declare a method name notify using a pointer receiver.
	//      This type now implements the notifier interface using a pointer receiver.
	//      func (d *duration) notify() {
	//          fmt.Println("Sending Notification in", *d)
	//      }

	//      Take a value 42, convert it to type duration and try to call the notify method.
	//      Here are what the compiler says:
	//      - "cannot call pointer method on duration(42)"
	//      - "cannot take the address of duration(42)"
	//      func main() {
	//          duration(42).notify()
	//      }

	//      Why can't we get the address? Because 42 is not stored in a variable. It is still literal
	//      value that we don't know ahead the type. Yet it still does implement the notifier interface.

	// Come back to our example, when we get the error, we know that we are mixing semantics. u
	// implements the interface using a pointer receiver and now we are trying to work with copy of
	// that value, instead of trying to share it. It is not consistent.

	// The lesson:
	// -----------
	// If we implement interface using pointer receiver, we must use pointer semantic.
	// If we implement interface using value receiver, we then have the ability to use value
	// semantic and pointer semantic. However, for consistency, we want to use value semantic most
	// of the time, unless we are doing something like Unmarshal function.

	// To fix the issue, instead of passing value u, we must pass the address of u (&u).
	// We create a user value and pass the address of that, which means the interface now has a
	// pointer of type user and we get to point to the original value.
	//  -------
	// | *User |
	//  -------
	// |   *   | --> original user value
	//  -------
	sendNotification(&u)

	// Similarly, when we pass a value of u to Println, in the output we only see the default
	// formatting. When we pass the address through, it now can overwrite it.
	fmt.Println(u)
	fmt.Println(&u)

	// ------------------
	// Slice of interface
	// ------------------

	// Create a slice of interface value.
	// It means that I can store in this dataset any value or pointer that implement the printer
	// interface.

	//   index 0   index 1
	//  -------------------
	// |   User  |  *User  |
	//  -------------------
	// |    *    |    *    |
	//  -------------------
	//      A         A
	//      |         |
	//     copy    original

	entities := []printer{
		// When we store a value, the interface value has its own copy of the value.
		// Changes to the original value will not be seen.
		u,

		// When we store a pointer, the interface value has its own copy of the address.
		// Changes to the original value will be seen.
		&u,
	}

	// Change the name and email on the user value.
	u.name = "Hoanh An"
	u.email = "hoanhan@bennington.edu"

	// Iterate over the slice of entities and call print against the copied interface value.
	for _, e := range entities {
		e.print()
	}
}

// This is our polymorphic function.
// sendNotification accepts values that implement the notifier interface and sends notifications.
// This is again saying: I will accept any value or pointer that implement the notifier interface.
// I will call that behavior against the interface itself .
func sendNotification(n notifier) {
	n.notify()
}
