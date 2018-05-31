// -------------
// Internal test
// -------------

// Below is how to test the execution of an internal endpoint without having to stand up the
// server.

// Run test using "go test -run TestSendJSON"

// We are using handlers_test for package name because we want to make sure we only touch the
// exported API.
package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hoanhan101/ultimate-go/go/testing/web_server/handlers"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

// This is very critical. If we forget to do this then nothing will work.
func init() {
	handlers.Routes()
}

// TestSendJSON testing the sendjson internal endpoint.
// In order to mock this call, we don't need the network. What we need to do is create a request
// and run it through the Mux so we are gonna bypass the network call together, run the request
// directly through the Mux to test the route and the handler.
func TestSendJSON(t *testing.T) {
	url := "/sendjson"
	statusCode := 200

	t.Log("Given the need to test the SendJSON endpoint.")
	{
		// Create a nil request GET for the URL.
		r := httptest.NewRequest("GET", url, nil)

		// NewRecorder gives us a pointer to its concrete type called ResponseRecorder that already
		// implemented the ResponseWriter interface.
		w := httptest.NewRecorder()

		// ServerHTTP asks for a ResonseWriter and a Request. This call will perform the Mux and
		// call that handler to test it without network. When his call comes back, the recorder
		// value w has the result of the entire execution. Now we can use that to validate.
		http.DefaultServeMux.ServeHTTP(w, r)

		t.Logf("\tTest 0:\tWhen checking %q for status code %d", url, statusCode)
		{
			if w.Code != 200 {
				t.Fatalf("\t%s\tShould receive a status code of %d for the response. Received[%d].", failed, statusCode, w.Code)
			}
			t.Logf("\t%s\tShould receive a status code of %d for the response.", succeed, statusCode)

			// If we got the 200, we try to unmarshal and validate it.
			var u struct {
				Name  string
				Email string
			}

			if err := json.NewDecoder(w.Body).Decode(&u); err != nil {
				t.Fatalf("\t%s\tShould be able to decode the response.", failed)
			}
			t.Logf("\t%s\tShould be able to decode the response.", succeed)

			if u.Name == "Hoanh An" {
				t.Logf("\t%s\tShould have \"Hoanh An\" for Name in the response.", succeed)
			} else {
				t.Errorf("\t%s\tShould have \"Hoanh An\" for Name in the response : %q", failed, u.Name)
			}

			if u.Email == "hoanhan@bennington.edu" {
				t.Logf("\t%s\tShould have \"hoanhan@bennington.edu\" for Email in the response.", succeed)
			} else {
				t.Errorf("\t%s\tShould have \"hoanhan@bennington.edu\" for Email in the response : %q", failed, u.Email)
			}
		}
	}
}
