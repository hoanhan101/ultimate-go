// -------------------------------
// Declaring Fields, NOT embedding
// -------------------------------

package main

import "fmt"

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method notifies users
// of different events.
func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s>\n", u.name, u.email)
}

// admin represents an admin user with privileges.
// person user is not embedding. All we do here just create a person field based on that other
// concrete type name user.
type admin struct {
	person user // NOT Embedding
	level  string
}

func main() {
	// Create an admin user using struct literal.
	// Since person also has struct type, we use another literal to initialize it.
	ad := admin{
		person: user{
			name:  "Hoanh An",
			email: "hoanhan@bennington.edu",
		},
		level: "superuser",
	}

	// We call notify through the person field through the admin type value.
	ad.person.notify()
}
