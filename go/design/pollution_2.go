// --------------------------
// Remove Interface Pollution
// --------------------------

// This is basically just removing the improper interface usage from previous file pollution_1.go.

package main

// Server implementation.
type Server struct {
	host string
	// PRETEND THERE ARE MORE FIELDS.
}

// NewServer returns just a concrete pointer of type Server
func NewServer(host string) *Server {
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

	// Use the APIs.
	srv.Start()
	srv.Stop()
	srv.Wait()
}

// Guidelines around interface pollution:
// --------------------------------------
// Use an interface:
// - When users of the API need to provide an implementation detail.
// - When APIs have multiple implementations that need to be maintained.
// - When parts of the APIs that can change have been identified and require decoupling.
// Question an interface:
// - When its only purpose is for writing testable API’s (write usable APIs first).
// - When it’s not providing support for the API to decouple from change.
// - When it's not clear how the interface makes the code better.
