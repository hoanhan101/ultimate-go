// -------------------------------------
// Decoupling With Interface Composition
// -------------------------------------

// We change our concrete type System. Instead of using two concrete types Xenia and Pillar, we
// use 2 interface types Puller and Storer. Our concrete type System where we can have concrete
// behaviors is now based on the embedding of 2 interface types. It means that we can inject any
// data, not based on the common DNA but on the data that providing the capability, the behavior
// that we need.
// Now we can be fully decouple because any value that implements the Puller interface can be store
// inside the System (same with Storer interface). We can create multiple Systems and that data can
// be passed in Copy.
// We don't need method here. We just need one function that accept data and its behavior will
// change based on the data we put in.

// Now System is not based on Xenia and Pillar anymore. It is based on 2 interfaces, one that
// stores Xenia and one that stores Pillar. We get the extra layer of decoupling.
// If the system change, no big deal. We replace the system as we need to during the program
// startup.

// We solve this problem. We put this in production. Every single refactoring that we did went into
// production before we did the next one. We keep minimizing technical debt.

//        System                                ps
//  --------------------                      ---------
// |  _________         |-pull               |         |-pull
// | |         |        |-store              | *System |-store
// | | *Xenia  |-pull   |                    |         |
// | |         |        | <------------------ ---------
// |  ---------         |      p             |         |
// | |         |        |    -----           |    *    |
// | |    *    |------- |-> |     |-pull     |         |
// | |         |        |    -----            ---------
// |  ---------         |
// |
// |  __________        |
// | |          |       |
// | | * Pillar |-store |
// | |          |       |
// |  ----------        |      s
// | |          |       |    -----                         p                   s
// | |    *     |------ |-> |     |-store               ---------           ---------
// | |          |       |    -----                     |         |-pull    |         |-store
// |  ----------        |                              | *System |         | *System |
//  --------------------                               |         |         |         |
//          A                                           ---------           ---------
//          |                                          |         |         |         |
//           ------------------------------------------|    *    | ------- |    *    |
//                                                     |         |         |         |
//                                                      ---------           ---------

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

// System wraps Pullers and Stores together into a single system.
type System struct {
	Puller
	Storer
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
		Puller: &Xenia{
			Host:    "localhost:8000",
			Timeout: time.Second,
		},
		Storer: &Pillar{
			Host:    "localhost:9000",
			Timeout: time.Second,
		},
	}

	if err := Copy(&sys, 3); err != io.EOF {
		fmt.Println(err)
	}
}
