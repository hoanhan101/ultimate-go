// ---------------------------------
// Unbuffered channel (Tennis match)
// ---------------------------------

// This program will put 2 Goroutines in a tennis match.
// We use an unbuffered channel because we need to guarantee that the ball is hit on both side or
// missed.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Create an unbuffered channel.
	court := make(chan int)

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(2)

	// Launch two players.
	// Both are gonna start out in a receive mode. We are not really sure who is gonna get the ball
	// first. Image the main Goroutine is the judge. It depends on the judge to choose.
	go func() {
		player("Hoanh", court)
		wg.Done()
	}()

	go func() {
		player("Andrew", court)
		wg.Done()
	}()

	// Start the set.
	// The main Goroutine here is performing a send. Since both players are in receive mode, we
	// cannot predict which one will go first.
	court <- 1

	// Wait for the game to finish.
	wg.Wait()
}

// player simulates a person playing the game of tennis.
// We are asking for a channel value using value semantic.
func player(name string, court chan int) {
	for {
		// Wait for the ball to be hit back to us.
		// Notice that this is another form of receive. Instead of getting just the value, we can
		// get a flag indicating how the receive is returned. If the signal happens because of the
		// data, ok will be true. If the signal happens without data, in other word, the channel is
		// closed, ok will be false. In this case, we are gonna use that to determine who won.
		ball, ok := <-court
		if !ok {
			// If the channel was closed we won.
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// Pick a random number and see if we miss the ball (or we lose).
		// If we lose the game, we are gonna close the channel. It then causes the other player to
		// know that he is receiving the signal but without data. The channel is closed so he won.
		// They both return.
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			// Close the channel to signal we lost.
			close(court)
			return
		}

		// Display and then increment the hit count by one.
		// If the 2 cases above doesn't happen, we still have the ball. Increase the value of the
		// ball by one and perform a send. We know that the other player is still in receive mode,
		// therefore, the send and receive will eventually come together.
		// Again, in an unbuffered channel, the receive happens first because it gives us the
		// guarantee.
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// Hit the ball back to the opposing player.
		court <- ball
	}
}
