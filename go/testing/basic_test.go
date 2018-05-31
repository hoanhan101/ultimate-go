// ---------------
// Basic Unit Test
// ---------------

// All of our tests must have the format <filename>_test.go.
// Otherwise, the testing tool is not gonna fine the tests.

// Test files are not compiled into our final binary.

// Test files should be in the same package as your code. We might also want to have a folder
// called test for more than unit test, say integration test.

// The package name can be the name only or name_test.
// If we go with name_test, it allows us to make sure these tests work with the package. The only
// reason that we don't want to do this is when we have a function or method that is unexported.
// However, if we don't use name_test, it will raise a red flag because if we cannot test the
// exported API to get the coverage for unexported API then we know are missing something.
// Therefore, 9/10 this is what we want.

package main

import (
	"net/http"
	"testing" // This is Go testing package.
)

// These constant gives us checkboxes for visualization.
const (
	succeed = "\u2713"
	failed  = "\u2717"
)

// TestBasic validates the http.Get function can download content.
// Every test will be associated with test function. It starts with the word Test and the first
// word after Test must be capitalized. It uses a testing.T pointer as its parameter.

// When writing test, we want to focus on usability first. We must write it the same way as we
// would write it in production.
// We also want the verbosity of tests so we are 3 different methods of t: Log or Logf, Fatal or
// Fatalf, Error or Error f. That is the core APIs for testing.
// Log: Write documentation out into the log output.
// Error: Write documentation and also say that this test is failed but we are continue
// moving forward to execute code in this test.
// Fatal: Similarly, document that this test is failed but we are done. We move on to the next
// test function.

// Given, When, Should format.
// Given: Why are we writing this test?
// When: What data are we using for this test?
// Should: When are we expected to see it happen or not happen?

// We are also using the artificial block between a long Log function.
// They help with readability.
func TestBasic(t *testing.T) {
	url := "https://www.goinggo.net/post/index.xml"
	statusCode := 200

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tTest 0:\tWhen checking %q for status code %d", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to make the Get call : %v", failed, err)
			}
			t.Logf("\t%s\tShould be able to make the Get call.", succeed)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t%s\tShould receive a %d status code.", succeed, statusCode)
			} else {
				t.Errorf("\t%s\tShould receive a %d status code : %d", failed, statusCode, resp.StatusCode)
			}
		}
	}
}

// Output:
// -------
// We can just say "go test" and the testing tool will find that function.
// We can also say "go test -v" for verbosity, we will get a full output of the logging.
// Suppose that we have a lot of test functions, "go test -run TestBasic" will only run the
// TestBasic function.
