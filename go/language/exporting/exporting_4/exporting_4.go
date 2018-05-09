// ---------------------------------------------
// Exported types with embedded unexported types
// ---------------------------------------------

package main

import (
	"fmt"

	"github.com/hoanhan101/ultimate-go/go/language/exporting/exporting_4/users"
)

func main() {
	// Create a value of type Manager from the users package.
	// During construction, we are only able to initialize the exported field Title. We cannot
	// access the embedded type directly.
	u := users.Manager{
		Title: "Dev Manager",
	}

	// However, once we have the manager value the exported fields from that unexported type are
	// accessible.
	u.Name = "Hoanh"
	u.ID = 101

	fmt.Printf("User: %#v\n", u)
}

// Again, we don't do this. A better way is to make user exported.
