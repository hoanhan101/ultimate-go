// --------------------------
// Remove Interface Pollution
// --------------------------

// This is basically just removing the improper interface usage from last file.

package main

// Server is our Server implementation.
type Server struct {
	host string
	// PRETEND THERE ARE MORE FIELDS.
}

// NewServer returns an interface value of type Server with a server implementation.
func NewServer(host string) *Server {
	// SMELL - Storing an unexported type pointer in the interface.
	return &Server{host}
}

// Start allows the server to begin to accept requests.
func (s *Server) Start() error {
	// PRETEND THERE IS A SPECIFIC IMPLEMENTATION.
	return nil
}

// Stop shuts the server down.
func (s *Server) Stop() error {
	// PRETEND THERE IS A SPECIFIC IMPLEMENTATION.
	return nil
}

// Wait prevents the server from accepting new connections.
func (s *Server) Wait() error {
	// PRETEND THERE IS A SPECIFIC IMPLEMENTATION.
	return nil
}

func main() {
	// Create a new Server.
	srv := NewServer("localhost")

	// Use the API.
	srv.Start()
	srv.Stop()
	srv.Wait()
}

// Guidelines around interface pollution:
// --------------------------------------
// Use an interface:
// - When users of the API need to provide an implementation detail.
// - When API’s have multiple implementations that need to be maintained.
// - When parts of the API that can change have been identified and require decoupling.
// Question an interface:
// - When its only purpose is for writing testable API’s (write usable API’s first).
// - When it’s not providing support for the API to decouple from change.
// - When it's not clear how the interface makes the code better.
