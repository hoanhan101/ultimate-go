// ----------
// Table test
// ----------

// Set up a data structure of input to expected output.
// This way we don't need a separate function for each one of these. We just have 1 test function.
// As we go along, we just add more to the table.

package main

import (
	"net/http"
	"testing"
)

// TestTable validates the http Get function can download content and
// handles different status conditions properly.
func TestTable(t *testing.T) {
	// This table is a slice of anonymous struct type. It is the URL we are gonna call and
	// statusCode are what we expect.
	tests := []struct {
		url        string
		statusCode int
	}{
		{"https://www.goinggo.net/post/index.xml", http.StatusOK},
		{"http://rss.cnn.com/rss/cnn_topstorie.rss", http.StatusNotFound},
	}

	t.Log("Given the need to test downloading different content.")
	{
		for i, tt := range tests {
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
	}
}
