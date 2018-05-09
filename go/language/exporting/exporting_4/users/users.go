// Package users provides support for user management.
package users

// user represents information about a user.
// Unexported type with 2 exported fields.
type user struct {
	Name string
	ID   int
}

// Manager represents information about a manager.
// Exported type embedded the unexported field user.
type Manager struct {
	Title string

	user
}
