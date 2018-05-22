// -------------------------
// Buffered channel: Fan Out
// -------------------------

// This is a classic use of a buffered channel that is greater than 1.
// It is called a Fan Out Pattern.

// Idea: A Goroutine is doing its thing and decides to run a bunch of database operation. It is
// gonna create a bunch of Gouroutines, say 10, to do that. Each Goroutine will perform 2 database
// operations. We end up having 20 database operations across 10 Goroutines. In other word, the
// original Goroutine will fan 10 Goroutines out, wait for them all to report back.

// The buffered channel is fantastic here because we know ahead of time that there are 10
// Goroutines performing 20 operations, so the size of the buffer is 20. There is no reason for any
// of these operation signal to block because we know that we have to receive this at the end of
// the day.

package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// result is what is sent back from each operation.
type result struct {
	id  int
	op  string
	err error
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Set the number of Goroutines and insert operations.
	const routines = 10
	const inserts = routines * 2

	// Buffered channel to receive information about any possible insert.
	ch := make(chan result, inserts)

	// Number of responses we need to handle.
	// Instead of using a WaitGroup, since this Goroutine can maintain its stack space, we are
	// gonna use a local variable as our WaitGroup. We will decrement that as we go.
	// Therefore, we set it to 20 inserts right out the box.
	waitInserts := inserts

	// Perform all the inserts. This is the fan out.
	// We are gonna have 10 Goroutines. Each Goroutine performs 2 inserts. The result of the insert
	// is used in a ch channel. Because this is a buffered channel, none of these send blocks.
	for i := 0; i < routines; i++ {
		go func(id int) {
			ch <- insertUser(id)

			// We don't need to wait to start the second insert thanks to the buffered channel.
			// The first send will happen immediately.
			ch <- insertTrans(id)
		}(i)
	}

	// Process the insert results as they complete.
	for waitInserts > 0 {
		// Wait for a response from a Goroutine.
		// This is a receive. We are receiving one result at a time and decrement the waitInserts
		// until it gets down to 0.
		r := <-ch

		// Display the result.
		log.Printf("N: %d ID: %d OP: %s ERR: %v", waitInserts, r.id, r.op, r.err)

		// Decrement the wait count and determine if we are done.
		waitInserts--
	}

	log.Println("Inserts Complete")
}

// insertUser simulates a database operation.
func insertUser(id int) result {
	r := result{
		id: id,
		op: fmt.Sprintf("insert USERS value (%d)", id),
	}

	// Randomize if the insert fails or not.
	if rand.Intn(10) == 0 {
		r.err = fmt.Errorf("Unable to insert %d into USER table", id)
	}

	return r
}

// insertTrans simulates a database operation.
func insertTrans(id int) result {
	r := result{
		id: id,
		op: fmt.Sprintf("insert TRANS value (%d)", id),
	}

	// Randomize if the insert fails or not.
	if rand.Intn(10) == 0 {
		r.err = fmt.Errorf("Unable to insert %d into USER table", id)
	}

	return r
}
