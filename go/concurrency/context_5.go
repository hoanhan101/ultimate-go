// ----------------
// Request/Response
// ----------------

// Sample program that implements a web request with a context that is
// used to timeout the request if it takes too long.

package main

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	// Create a new request.
	req, err := http.NewRequest("GET", "https://www.ardanlabs.com/blog/post/index.xml", nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Create a context with a timeout of 50 milliseconds.
	ctx, cancel := context.WithTimeout(req.Context(), 50*time.Millisecond)
	defer cancel()

	// Declare a new transport and client for the call.
	tr := http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := http.Client{
		Transport: &tr,
	}

	// Make the web call in a separate Goroutine so it can be cancelled.
	ch := make(chan error, 1)
	go func() {
		log.Println("Starting Request")

		// Make the web call and return any error.
		// client.Do is going out and trying to hit the request URL. It's probably blocked right
		// now because it will need to wait for the entire document to comeback.
		resp, err := client.Do(req)

		// It the error occurs, we perform a send on the channel to report that we are done. We are
		// going to use this channel at some point to report back what is happening.
		if err != nil {
			ch <- err
			return
		}

		// If it doesn't fail, we close the response body on the return.
		defer resp.Body.Close()

		// Write the response to stdout.
		io.Copy(os.Stdout, resp.Body)

		// Then send back the nil instead of error.
		ch <- nil
	}()

	// Wait the request or timeout.
	// We perform a receive on ctx.Done saying that we want to wait 50 ms for that whole process
	// above to happen. If it doesn't, we signal back to that Goroutine to cancel the sending
	// request. We don't have to just walk away and let that eat up resources and finish because we
	// are not gonna need it. We are able to call CancelRequest and underneath, we are able to kill
	// that connection.
	select {
	case <-ctx.Done():
		log.Println("timeout, cancel work...")

		// Cancel the request and wait for it to complete.
		tr.CancelRequest(req)
		log.Println(<-ch)
	case err := <-ch:
		if err != nil {
			log.Println(err)
		}
	}
}
