// -----------
// Web testing
// -----------

// Those basic tests that we just went through were cool but had a flaw: they require the use of
// Internet. We cannot assume that we always have access to the resources we need. Therefore,
// mocking becomes an important part of testing in many cases. (Mocking databases if not the case
// here because it is hard to do so but other networking related thing, we surely can do that).

// The standard library already has the http test package that let us mock different http stuff
// right out of the box. Below is how to mock an http GET call internally.

package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// feed is mocking the XML document we expect to receive.
// Notice that we are using ` instead of " so we can reserve special characters.
var feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
<channel>
    <title>Going Go Programming</title>
    <description>Golang : https://github.com/goinggo</description>
    <link>http://www.goinggo.net/</link>
    <item>
        <pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
        <title>Object Oriented Programming Mechanics</title>
        <description>Go is an object oriented language.</description>
        <link>http://www.goinggo.net/2015/03/object-oriented</link>
    </item>
</channel>
</rss>`

// Item defines the fields associated with the item tag in the mock RSS document.
type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
}

// Channel defines the fields associated with the channel tag in the mock RSS document.
type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Items       []Item   `xml:"item"`
}

// Document defines the fields associated with the mock RSS document.
type Document struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
	URI     string
}

// mockServer returns a pointer of type httptest.Server to handle the mock get call.
// This mock function calls NewServer function that gonna stand up a web server for us
// automatically. All we have to give NewServer is a function of the Handler type, which is f.

// f creates an anonymous function with the signature of ResponseWriter and Request. This is the
// core signature of everything related to http in Go.
// ResponseWriter is an interface that allows us to write the response out. Normally when we get
// this interface value, there is already a concrete type value stored inside of it that support
// what we are doing.
// Request is a concrete type that we are gonna get with the request.

// This is how it's gonna work.
// We are gonna get a mock server started by making NewServer call. When the request comes into
// it, execute f. Therefore, f is doing the entire mock.
// We are gonna send 200 down the line, set the header to XML and use Fprintln to take the
// ResponseWriter interface value and feeding with the raw string we defined above.
func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, feed)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

// TestWeb validates the http Get function can download content and
// the content can be unmarshaled and clean.
func TestWeb(t *testing.T) {
	statusCode := http.StatusOK

	// Call the mock sever and defer close to shut it down cleanly.
	server := mockServer()
	defer server.Close()

	// Now, it's just the matter of using server value to know what URL we need to use to run this
	// mock. From the http.Get point of view, it is making an URL call. It has no idea that it's
	// hitting the mock server. We have mocked out a perfect response.
	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tTest 0:\tWhen checking %q for status code %d", server.URL, statusCode)
		{
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to make the Get call : %v", failed, err)
			}
			t.Logf("\t%s\tShould be able to make the Get call.", succeed)

			defer resp.Body.Close()

			if resp.StatusCode != statusCode {
				t.Fatalf("\t%s\tShould receive a %d status code : %v", failed, statusCode, resp.StatusCode)
			}
			t.Logf("\t%s\tShould receive a %d status code.", succeed, statusCode)

			// When we get the response back, we are unmarshaling it from XML to our struct type
			// and do some extra validation with that as we go.
			var d Document
			if err := xml.NewDecoder(resp.Body).Decode(&d); err != nil {
				t.Fatalf("\t%s\tShould be able to unmarshal the response : %v", failed, err)
			}
			t.Logf("\t%s\tShould be able to unmarshal the response.", succeed)

			if len(d.Channel.Items) == 1 {
				t.Logf("\t%s\tShould have 1 item in the feed.", succeed)
			} else {
				t.Errorf("\t%s\tShould have 1 item in the feed : %d", failed, len(d.Channel.Items))
			}
		}
	}
}

// Output:
// -------
// First of all, it runs much faster because we know that it is not leaving the building.
// We also see that we have the localhost IP address on a port.
