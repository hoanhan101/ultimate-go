// -----------------------------
// Embedded types and interfaces
// -----------------------------

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

// notify implements a method notifies users of different events using a pointer receiver.
func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s>\n", u.name, u.email)
}

// admin represents an admin user with privileges.
type admin struct {
	user
	level string
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
	// We are passing the address of outer type value. Because of inner type promotion, the outer
	// type now implements all the same contract as the inner type.
	sendNotification(&ad)

	// Embedding does not create a sub typing relationship. user is still user and admin is still
	// admin. The behavior that inner type value uses, the outer type exposes it as well.
	// It means that outer type value can implement the same interface/same contract as the inner
	// type.

	// We are getting type reuse. We are not mixing or sharing state but extending the behavior up
	// to the outer type.
}

// We have our polymorphic function here.
// sendNotification accepts values that implement the notifier interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}
