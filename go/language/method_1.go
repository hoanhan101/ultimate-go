package main

import "fmt"

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method with a value receiver: u of type user
// In Go, a function is called a method if that function has declare within it a receiver.
// It looks and feels like a parameter but it is exactly what it is.
// Using the value receiver, the method operates on its own copy of the value that is used to make
// the call.
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

// changeEmail implements a method with a pointer receiver: u of type pointer user
// Using the pointer reciever, the method operates on share access.
func (u *user) changeEmail(email string) {
	u.email = email
}

// These 2 methods above are just for studying the difference between a value receiver and a
// pointer receiver. In production, we will have to ask ourself why we choose to use inconsistent
// receiver's type. We will talk about this later on.

func main() {
	// -------------------------------
	// Value and pointer receiver call
	// -------------------------------

	// Values of type user can be used to call methods declared with both value and pointer receivers.
	bill := user{"Bill", "bill@email.com"}
	bill.notify()
	bill.changeEmail("bill@hotmail.com")

	// Pointers of type user can also be used to call methods declared with both value and pointer receiver.
	hoanh := &user{"Hoanh", "hoanhan@email.com"}
	hoanh.notify()
	hoanh.changeEmail("hoanhan@bennington.edu")

	// hoanh in this example is a pointer that has the type *user. We are still able to call notify.
	// This is still correct. As long as we deal with the type user, Go can adjust to make the call.

	// Behind the scene, we have something like (*hoanh).notify(). Go will take the value that hoanh
	// points to and make sure that notify leverages its value semantic and works on its own copy.

	// Similarly, bill has the type user but still be able to call changeEmail. Go will take the
	// address of bill and do the rest for you: (*bill).changeEmail().

	// Create a slice of user values with two users.
	users := []user{
		{"bill", "bill@email.com"},
		{"hoanh", "hoanh@email.com"},
	}

	// We are ranging over this slice of values, making a copy of each value and call notify to
	// make another copy.
	for _, u := range users {
		u.notify()
	}

	// Iterate over the slice of users switching semantics. Not good practice.
	for _, u := range users {
		u.changeEmail("it@wontmatter.com")
	}
}
