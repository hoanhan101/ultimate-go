// ---------------------
// Interface Composition
// ---------------------

// Let's just add another interface. Let's use interface composition to do this.
// PullStorer has both behaviors: Puller and Storer. Any concrete type that implement both pull and
// store is a PullStorer. System is a PullStorer because it is embedded of these 2 types Xenia and
// Pillar. Now we just need to go into Copy, replace the system pointer with PullStorer and no
// other code need to change. When we call Copy passing the address of our System value in main,
// that already implement the PullStorer interface.

// Looking closely at Copy, there is something that could potentially confuse us. We are passing
// the PullStorer interface value into pull and store respectively.
// If we look into pull and store, they don't want a PullStorer. One want a Puller and one want a
// Storer. Why does the compiler allow us to pass a value of different type value while it didn't
// allow us to do that before?
// This is because Go has what is called: implicit interface conversion.
// This is possible because:
// - All interface values have the exact same model (implementation details).
// - If the type information is clear, the concrete type that exists in one interface has enough
// behaviors for another interface. It is true that any concrete type that is stored inside of a
// PullStorer must also implement the Storer and Puller.

// Let's walkthrough the code.
// In the main function, we are creating a value of our type System. As we know, our type System
// value is based on the embedding of two concrete types: Xenia and Pillar, where Xenia knows how
// to pull and Pillar knows how to store. Because of inner promotion, System knows also how to pull
// and store.
// We are passing the address of our System to Copy. Copy then creates the PullStorer interface.
// The first word is a System pointer and the second word point to the original value. This
// interface now knows how to pull and store. When we call pull off of ps, we call pull off of
// System, which call pull off of Xenia.
// Here is the kicker: the implicit interface conversion.
// We can pass the interface value ps to pull because the compiler knows that any concrete type
// stored inside the PullStorer must also implement Puller. We end up with another interface called
// Puller. Because the memory models are the same for all interfaces, we just copy those 2 words so
// they are all sharing the same interface type. Now when we call pull off of Puller, we call pull
// off of System. Similar to Storer.
// All using value semantic for the interface value and pointer semantic to share.

//        System                       ps
//  ------------------              ---------
// |  _______         |-pull       |         |-pull
// | |        |       |-store      | *System |-store
// | | Xenia  |-pull  |            |         |
// | |        |       |             ---------
// |  -------         |            |         |
// |  _______         |<-----------|    *    |
// | |        |       |            |         |
// | | Pillar |-store |             ---------               p                   s
// | |        |       |                                 ---------           ---------
// |  -------         |                                |         |-pull    |         |-store
// |                  |                                | *System |         | *System |
//  ------------------                                 |         |         |         |
//          A                                           ---------           ---------
//          |                                          |         |         |         |
//           ------------------------------------------|    *    | ------- |    *    |
//                                                     |         |         |         |
//                                                      ---------           ---------

// Next step:
// ----------
// Our system type is still concrete system type because it is still based on two concrete types,
// Xenial and Pillar. If we have another system, say Alice, we have to change in type System
// struct. This it not good. We will solve the last piece in the next file.

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

// PullStorer declares behavior for both pulling and storing.
type PullStorer interface {
	Puller
	Storer
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

// Copy knows how to pull and store data from any System.
func Copy(ps PullStorer, batch int) error {
	data := make([]Data, batch)

	for {
		i, err := pull(ps, data)
		if i > 0 {
			if _, err := store(ps, data[:i]); err != nil {
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
