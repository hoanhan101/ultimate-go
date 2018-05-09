// -----------------------------------------
// Unexported fields from an exported struct
// -----------------------------------------

package main

import (
	"fmt"

	"github.com/hoanhan101/ultimate-go/go/language/exporting/exporting_3/users"
)

func main() {
	// Create a value of type User from the users package using struct literal.
	// However, since password is unexported, it cannot be compiled:
	// - unknown users.User field 'password' in struct literal
	u := users.User{
		Name: "Hoanh",
		ID:   101,

		password: "xxxx",
	}

	fmt.Printf("User: %#v\n", u)
}
