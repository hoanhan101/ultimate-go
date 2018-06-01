// Sub Test
// --------

// Sub test helps us streamline our test functions, filters out command-line level big tests into
// smaller sub tests.

package main

import (
	"net/http"
	"testing"
)

// TestSub validates the http Get function can download content and
// handles different status conditions properly.
func TestSub(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		statusCode int
	}{
		{"statusok", "https://www.goinggo.net/post/index.xml", http.StatusOK},
		{"statusnotfound", "http://rss.cnn.com/rss/cnn_topstorie.rss", http.StatusNotFound},
	}

	t.Log("Given the need to test downloading different content.")
	{
		// Range over our table but this time, create an anonymous function that takes a testing T
		// parameter. This is a test function inside a test function.
		// What nice about it is that we are gonna have a new function for each set of data that we
		// have in our table. Therefore, we will end up with 2 different functions here.
		for i, tt := range tests {
			tf := func(t *testing.T) {
				t.Logf("\tTest: %d\tWhen checking %q for status code %d", i, tt.url, tt.statusCode)
				{
					resp, err := http.Get(tt.url)
					if err != nil {
						t.Fatalf("\t%s\tShould be able to make the Get call : %v", failed, err)
					}
					t.Logf("\t%s\tShould be able to make the Get call.", succeed)

					defer resp.Body.Close()

					if resp.StatusCode == tt.statusCode {
						t.Logf("\t%s\tShould receive a %d status code.", succeed, tt.statusCode)
					} else {
						t.Errorf("\t%s\tShould receive a %d status code : %v", failed, tt.statusCode, resp.StatusCode)
					}
				}
			}

			// Once we declare this function, we tell the testing tool to register it as a sub test
			// under the test name.
			t.Run(tt.name, tf)
		}
	}
}

// TestParallelize validates the http Get function can download content and
// handles different status conditions properly but runs the tests in parallel.
func TestParallelize(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		statusCode int
	}{
		{"statusok", "https://www.goinggo.net/post/index.xml", http.StatusOK},
		{"statusnotfound", "http://rss.cnn.com/rss/cnn_topstorie.rss", http.StatusNotFound},
	}

	t.Log("Given the need to test downloading different content.")
	{
		for i, tt := range tests {
			tf := func(t *testing.T) {
				// The only difference here is that we call Parallel function inside each of these
				// individual sub test function.
				t.Parallel()

				t.Logf("\tTest: %d\tWhen checking %q for status code %d", i, tt.url, tt.statusCode)
				{
					resp, err := http.Get(tt.url)
					if err != nil {
						t.Fatalf("\t%s\tShould be able to make the Get call : %v", failed, err)
					}
					t.Logf("\t%s\tShould be able to make the Get call.", succeed)

					defer resp.Body.Close()

					if resp.StatusCode == tt.statusCode {
						t.Logf("\t%s\tShould receive a %d status code.", succeed, tt.statusCode)
					} else {
						t.Errorf("\t%s\tShould receive a %d status code : %v", failed, tt.statusCode, resp.StatusCode)
					}
				}
			}

			t.Run(tt.name, tf)
		}
	}
}

// Output:
// -------
// Because we have sub tests, we can run the following to separate them:
// "go test -run TestSub -v"
// "go test -run TestSub/statusok -v"
// "go test -run TestSub/statusnotfound -v"
// "go test -run TestParallelize -v"
