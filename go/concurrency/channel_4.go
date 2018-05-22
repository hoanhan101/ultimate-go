// --------------------------------
// Unbuffered channel (Replay race)
// --------------------------------

// The program shows how to use an unbuffered channel to simulate a relay race between four Goroutines.
// Imagine we have 4 runners that are on the track. Only 1 can run at a time. We have the second
// runner on the track until the last one. The second one wait to be exchanged.

package main

import (
	"fmt"
	"sync"
	"time"
)

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

func main() {
	// Create an unbuffered channel.
	track := make(chan int)

	// Add a count of one for the last runner.
	// We only add one because all we care about is the last runner in the race telling us that he
	// is done.
	wg.Add(1)

	// Create a first runner to his mark.
	go Runner(track)

	// The main Goroutine start the race (shoot the gun).
	// At this moment, we know that on the other side, a Goroutine is performing a receive.
	track <- 1

	// Wait for the race to finish.
	wg.Wait()
}

// Runner simulates a person running in the relay race.
// This Runner doesn't have a loop because it's gonna do everything from the beginning to end and
// then terminate. We are gonna keep adding Goroutines (Runners) in order to make this pattern
// work.
func Runner(track chan int) {
	// The number of exchanges of the baton.
	const maxExchanges = 4

	var exchange int

	// Wait to receive the baton with data.
	baton := <-track

	// Start running around the track.
	fmt.Printf("Runner %d Running With Baton\n", baton)

	// New runner to the line. Are we the last runner on the race?
	// If not, we increment the data by 1 to keep track which runner we are on.
	// We will create another Goroutine. It will go immediately into a receive. We are now having a
	// second Groutine on the track, in the receive waiting for the baton. (1)
	if baton < maxExchanges {
		exchange = baton + 1
		fmt.Printf("Runner %d To The Line\n", exchange)
		go Runner(track)
	}

	// Running around the track.
	time.Sleep(100 * time.Millisecond)

	// Is the race over.
	if baton == maxExchanges {
		fmt.Printf("Runner %d Finished, Race Over\n", baton)
		wg.Done()
		return
	}

	// Exchange the baton for the next runner.
	fmt.Printf("Runner %d Exchange With Runner %d\n", baton, exchange)

	// Since we are not the last runner, perform a send so (1) can receive it.
	track <- exchange
}
