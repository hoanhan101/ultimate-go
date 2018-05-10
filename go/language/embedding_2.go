// ---------------
// Embedding types
// ---------------

package main

import "fmt"

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method notifies users of different events.
func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s>\n", u.name, u.email)
}

// admin represents an admin user with privileges.
// Notice that we don't use the field person here anymore.
// We are now embedding a value of type user inside value of type admin.
// This is an inner type-outer-type relationship where user is an inner type and admin is the
// outer-type.

// --------------------
// Inner type promotion
// --------------------

// What special about embedding in Go is that we have inner type promotion mechanism.
// In other word, anything relates to the inner type can be promoted up to the outer type.
// It will mean more in the construction below.
type admin struct {
	user  // Embedded Type
	level string
}

func main() {
	// We are now constructing outer-type admin and inner type user.
	// This inner type value now looks like a field, but not a field. We can access it through the
	// type name like a field.
	// We are initializing the inner value through the struct literal of user.
	ad := admin{
		user: user{
			name:  "Hoanh An",
			email: "hoanhan@bennington.edu",
		},
		level: "superuser",
	}

	// We can access the inner type's method directly.
	ad.user.notify()

	// Because of inner type promotion, we can access the notify method directly through the outer
	// type. Therefore, the output will be the same.
	ad.notify()
}
