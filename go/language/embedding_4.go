// ---------------------------------------------
// OUTER AND INNER TYPE IMPLEMENT SAME INTERFACE
// ---------------------------------------------

package main

import "fmt"

// notifier is an interface that defined notification type behavior.
type notifier interface {
	notify()
}

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
type admin struct {
	user
	level string
}

// notify implements a method notifies admins of different events.
// We now have two different implementations of notifier interface, one for the inner type,
// one for the outer type. Because the outer type now implements that interface, the inner type
// promotion doesn't happen. We have overwritten through the outer type anything that inner type
// provides to us.
func (a *admin) notify() {
	fmt.Printf("Sending admin email To %s<%s>\n", a.name, a.email)
}

func main() {
	// Create an admin user.
	ad := admin{
		user: user{
			name:  "Hoanh An",
			email: "hoanhan@bennington.edu",
		},
		level: "superuser",
	}

	// Send the admin user a notification.
	// The embedded inner type's implementation of the interface is NOT "promoted"
	// to the outer type.
	sendNotification(&ad)

	// We can access the inner type's method directly.
	ad.user.notify()

	// The inner type's method is NOT promoted.
	ad.notify()
}

// sendNotification accepts values that implement the notifier interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}
