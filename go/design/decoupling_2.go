// -------------------------
// Decoupling With Interface
// -------------------------

// By looking at the API (function), we need to decouple the API from the concrete. The decoupling
// that we do must get all the way down into initialization. To do this right, the only piece of
// code that we need to change is initialization. Everything else should be able to act on the
// behavior that these types are gonna provide.

// pull is based on the concrete. It only knows how to work on Xenia. However, if we are able to
// decouple pull to use any system that know how to pull data, we can get the highest level of
// decoupling. Since the algorithm we have is already efficient, we don't need to add another level
// of generalization and destroy the work we did in the concrete. Same thing with store.

// It is nice to work from the concrete up. When we do this, not only we are solving problem
// efficiently and reducing technical debt but the contracts, they come to us. We already know what
// the contract is for pulling/storing data. We already validate that this is what we need.

// Let's just decouple these 2 functions and add 2 interfaces. The Puller interface knows how to
// pull and the Storer knows how to store.
// Xenia already implemented the Puller interface and Pillar already implemented the Storer
// interface. Now we can come into pull/store, decouple this function from the concrete.
// Instead of passing Xenial and Pillar, we pass in the Puller and Storer. The algorithm doesn't
// change. All we doing is now calling pull/store indirectly through the interface value.

// Next step:
// ----------
// Copy also doesn't have to change because Xenia/Pillar already implemented the interfaces.
// However, we are not done because Copy is still bounded to the concrete. Copy can only work with
// pointer of type system. We need to decouple Copy so we can have a decoupled system that knows
// how to pull and store. We will do it in the next file.

package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Data is the structure of the data we are copying.
type Data struct {
	Line string
}

// Puller declares behavior for pulling data.
type Puller interface {
	Pull(d *Data) error
}

// Storer declares behavior for storing data.
type Storer interface {
	Store(d *Data) error
}

// Xenia is a system we need to pull data from.
type Xenia struct {
	Host    string
	Timeout time.Duration
}

// Pull knows how to pull data out of Xenia.
func (*Xenia) Pull(d *Data) error {
	switch rand.Intn(10) {
	case 1, 9:
		return io.EOF

	case 5:
		return errors.New("Error reading data from Xenia")

	default:
		d.Line = "Data"
		fmt.Println("In:", d.Line)
		return nil
	}
}

// Pillar is a system we need to store data into.
type Pillar struct {
	Host    string
	Timeout time.Duration
}

// Store knows how to store data into Pillar.
func (*Pillar) Store(d *Data) error {
	fmt.Println("Out:", d.Line)
	return nil
}

// System wraps Xenia and Pillar together into a single system.
type System struct {
	Xenia
	Pillar
}

// pull knows how to pull bulks of data from any Puller.
func pull(p Puller, data []Data) (int, error) {
	for i := range data {
		if err := p.Pull(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

// store knows how to store bulks of data from any Storer.
func store(s Storer, data []Data) (int, error) {
	for i := range data {
		if err := s.Store(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

// Copy knows how to pull and store data from the System.
func Copy(sys *System, batch int) error {
	data := make([]Data, batch)

	for {
		i, err := pull(&sys.Xenia, data)
		if i > 0 {
			if _, err := store(&sys.Pillar, data[:i]); err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}
	}
}

func main() {
	sys := System{
		Xenia: Xenia{
			Host:    "localhost:8000",
			Timeout: time.Second,
		},
		Pillar: Pillar{
			Host:    "localhost:9000",
			Timeout: time.Second,
		},
	}

	if err := Copy(&sys, 3); err != io.EOF {
		fmt.Println(err)
	}
}
