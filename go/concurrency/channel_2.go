// ------------------
// Language Mechanics
// ------------------

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("\n=> Double signal\n")
	signalAck()

	fmt.Printf("\n=> Select and receive\n")
	selectRecv()

	fmt.Printf("\n=> Select and send\n")
	selectRecv()

	fmt.Printf("\n=> Select and drop\n")
	selectDrop()
}

// ---------------------------------
// Unbuffered channel: Double signal
// ---------------------------------

// signalAck shows how to signal an event and wait for an acknowledgement it is done
// It does not only want to guarantee that a signal is received but also want to know when that
// work is done. This is gonna like a double signal.
func signalAck() {
	ch := make(chan string)

	go func() {
		fmt.Println(<-ch)
		ch <- "ok done"
	}()

	// It blocks on the receive. This Goroutine can no longer move on until we receive a signal.
	ch <- "do this"
	fmt.Println(<-ch)
}

// ---------------------------------
// Buffered channel: Close and range
// ---------------------------------

// closeRange shows how to use range to receive value and using close to terminate the loop.
func closeRange() {
	// This is a buffered channel of 5.
	ch := make(chan int, 5)

	// Populate with value
	for i := 0; i < 5; i++ {
		ch <- i
	}

	// Close the channel.
	close(ch)

	// Every iteration on the range is a receive.
	// When the range notices that the channel is closed, the loop will terminate.
	for v := range ch {
		fmt.Println(v)
	}
}

// --------------------------------------
// Unbuffered channel: select and receive
// --------------------------------------

// Select allows a Goroutine to work with multiple channel at a time, including send and receive.
// This can be great when creating an event loop but not good for serializing shared state.

// selectRecv shows how to use the select statement to wait for a specific amount of time to
// receive a value.
func selectRecv() {
	ch := make(chan string)

	// Wait for some amount of time and perform a send.
	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		ch <- "work"
	}()

	// Perform 2 different receives on 2 different channels: one above and one for time.
	// time.After returns a channel that will send the current time after that duration.
	// We want to receive the signal from the work sent but we are not willing to wait forever. We
	// only wait 100 milliseconds then we will move on.
	select {
	case v := <-ch:
		fmt.Println(v)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timed out")
	}

	// However, there is a very common bug in this code.
	// One of the biggest bug we are going to have and potential memory is when we write code like
	// this and we don't give the Goroutine an opportunity to terminate.
	// We are using an unbuffered channel and this Goroutine at some point, its duration will
	// finish and it will want to perform a send. But this is an unbuffered channel. This send
	// cannot be completed unless there is a corresponding receive. What if this Goroutine times
	// out and moves on? There is no more corresponding receive. Therefore, we will have a Goroutine
	// leak, which means it will never be terminated.

	// The cleanest way to fix this bug is to use the buffered channel of 1. If this send happens,
	// we don't necessarily have the guarantee. We don't need it. We just need to perform the
	// signal then we can walk away. Therefore, either we get the signal on the other side or we
	// walk away. Even if we walk away, this send can still be completed because there is room in
	// the buffer for that send to happen.
}

// -----------------------------------
// Unbuffered channel: select and send
// -----------------------------------

// selectSend shows how to use the select statement to attempt a send on a channel for a specific
// amount of time.
func selectSend() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		fmt.Println(<-ch)
	}()

	select {
	case ch <- "work":
		fmt.Println("send work")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timed out")
	}

	// Similar to the above function, Goroutine leak will occur.
	// Once again, a buffered channel of 1 will save us here.
}

// ---------------------------------
// Buffered channel: Select and drop
// ---------------------------------

// selectDrop shows how to use the select to walk away from a channel operation if it will
// immediately block.
// This is a really important pattern. Imagine a situation where our service is flushed with work
// to do or work is gonna coming. Something upstream is not functioning properly. We can't just
// back up the work. We have to throw it away so we can keep moving on.
// A Denial-of-service attack is a great example. We get a bunch of requests coming to our server.
// If we try to handle every single request, we are gonna implode. We have to handle what we can
// and drop other requests.
// Using this type of pattern (fanout), we are willing to drop some data. We can use buffer that
// are larger than 1. We have to measure what the buffer should be. It cannot be random.
func selectDrop() {
	ch := make(chan int, 5)

	go func() {
		// We are in the receive loop waiting for data to work on.
		for v := range ch {
			fmt.Println("recv", v)
		}
	}()

	// This will send the work to the channel.
	// If the buffer fills up, which means it blocks, the default case comes in and drop things.
	for i := 0; i < 20; i++ {
		select {
		case ch <- i:
			fmt.Println("send work", i)
		default:
			fmt.Println("drop", i)
		}
	}

	close(ch)
}
