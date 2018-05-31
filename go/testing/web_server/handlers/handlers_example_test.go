// ------------
// Example test
// ------------

// This is another type of test in Go. Examples are both tests and documentations.
// If we execute "godoc -http :3000", Go will generate for us a server that present the
// documentation of our code. The interface will looks like the official golang interface, but then
// inside Packages section are our local packages.

// Example functions are a little bit more concrete in term of showing people how to use
// our API. More interestingly, Examples are not only for documentation but they can also be tests.
// For them to be tests, we need to add a comment at the end of the functions: one is Output and
// one is expected output. If we change the expected output to be something wrong then, the
// complier will tell us when we run the test. Below is an example.

// Example tests are really powerful. They give users examples how to use the API and validate that
// the APIs and examples are working.

// Run test: "go test -run ExampleSendJSON"

package handlers_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

// ExampleSendJSON provides a basic example example.
// Notice that we are binding that Example to our SendJSON function.
func ExampleSendJSON() {
	r := httptest.NewRequest("GET", "/sendjson", nil)
	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, r)

	var u struct {
		Name  string
		Email string
	}

	if err := json.NewDecoder(w.Body).Decode(&u); err != nil {
		log.Println("ERROR:", err)
	}

	fmt.Println(u)
	// Output:
	// {Hoanh An hoanhan@bennington.edu}
}
