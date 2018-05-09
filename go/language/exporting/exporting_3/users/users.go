// Package users provides support for user management.
package users

// Exported type User represents information about a user.
// It has 2 exported fields: Name and ID and 1 unexported field: password.
type User struct {
	Name string
	ID   int

	password string
}
