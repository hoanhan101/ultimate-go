// -------------------
// Behavior as context
// -------------------

// Behavior as context allows us to use a custom error type as our context but avoid that type
// assertion back to the concrete. We get to maintain a level of decoupling in our code.

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

// client represents a single connection in the room.
type client struct {
	name   string
	reader *bufio.Reader
}

// TypeAsContext shows how to check multiple types of possible custom error
// types that can be returned from the net package.
func (c *client) TypeAsContext() {
	for {
		// We are using reader interface value to decouple ourselves from the network read.
		line, err := c.reader.ReadString('\n')
		if err != nil {
			// This is using type as context like the previous example.
			// What special here is the method Temporary. If it is, we can keep going but if not,
			// we have to break thing down and build thing back up.
			// Every one of these cases care only about 1 thing: the behavior of Temporary. This is
			// what important. We can switch here, from type as context to type as behavior if we
			// do this type assertion and only ask about the potential behavior of that concrete
			// type itself.
			// We can go ahead and declare our own interface called temporary like below.
			switch e := err.(type) {
			case *net.OpError:
				if !e.Temporary() {
					log.Println("Temporary: Client leaving chat")
					return
				}

			case *net.AddrError:
				if !e.Temporary() {
					log.Println("Temporary: Client leaving chat")
					return
				}

			case *net.DNSConfigError:
				if !e.Temporary() {
					log.Println("Temporary: Client leaving chat")
					return
				}

			default:
				if err == io.EOF {
					log.Println("EOF: Client leaving chat")
					return
				}

				log.Println("read-routine", err)
			}
		}

		fmt.Println(line)
	}
}

// temporary is declared to test for the existence of the method coming from the net package.
// Because Temporary this the only behavior we care about. If the concrete type has the method
// named temporary then this is what we want. We get to stay decoupled and continue to work at the
// interface level.
type temporary interface {
	Temporary() bool
}

// BehaviorAsContext shows how to check for the behavior of an interface
// that can be returned from the net package.
func (c *client) BehaviorAsContext() {
	for {
		line, err := c.reader.ReadString('\n')
		if err != nil {
			switch e := err.(type) {
			// We can reduce 3 cases into 1 by asking in the case here during type assertion: Does
			// the concrete type stored inside the error interface also implement this interface.
			// We can declare that interface for myself and we leverage it ourselves.
			case temporary:
				if !e.Temporary() {
					log.Println("Temporary: Client leaving chat")
					return
				}

			default:
				if err == io.EOF {
					log.Println("EOF: Client leaving chat")
					return
				}

				log.Println("read-routine", err)
			}
		}

		fmt.Println(line)
	}
}

// Lesson:
// If we can one of these methods to our concrete error type, we can maintain a level of decoupling:
// - Temporary
// - Timeout
// - NotFound
// - NotAuthorized
