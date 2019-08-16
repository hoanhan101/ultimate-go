// ------------------
// Language Mechanics
// ------------------

// Channels are for orchestration. They allow us to have 2 Goroutines participate in some sort of
// workflow and give us the ability to orchestrate in a predictable way.

// The one thing that we really need to think about is not that a channel is a queue, even though
// it seems it be implemented like queue, first in first out. We will have a difficult time if we
// think that way. What we want to think about instead is a channel as a way of signaling events to
// another Goroutine. What is nice here is that we can signal an event with data or without data.
// If everything we do have signaling in mind, we are going to use channel in a proper way.

// Go has 2 types of channels: unbuffered and buffered. They both allow us to signal with data.
// The big difference is that, when we use unbuffered channel, we are signaling and getting a
// guarantee the signal was received. We are not gonna be sure if that Goroutine is done whatever
// work we assign it to do but we do have the guarantee. The trade off for the guarantee that the
// signal was received is higher latency because we have to wait to make sure that the Goroutine on
// the other side of that unbuffered channel receive the data.

// This is how the unbuffered channel going to work.
// There is gonna a Goroutine comes to the channel. The channel wants to signal with some piece of
// data. It is gonna put the data right there in the channel. However, the data is locked in and
// cannot move because channel has to know if there is another Goroutine is on the other side to
// receive it. Eventually a Goroutine come and say that it want to receive the data. Both of
// Goroutines are not putting their hands in the channel. The data now can be transferred.
// Here is the key to why that unbuffered channel gives us that guarantee: the receive happens
// first. When the receive happens, we know that the data transfer has occurred and we can walk
// away.

//  G                      G
//  |        Channel       |
//  |      ----------      |
//  |     |   D  D   |     |
//  |-----|--->  <---|-----|
//  |     |          |     |
//  |      ----------      |
//  |                      |

// The unbuffered channel is a very powerful channel. We want to leverage that guarantee as much as
// possible. But again, the cost of the guarantee is higher latency because we have to wait for
// this.

// The buffered channel is a bit different: we do not get the guarantee but we get to reduce the
// amount of latencies on any given send or receive.

// Back to the previous example, we replace the unbuffered channel with a buffered channel. We are
// gonna a buffered channel of just 1. It means there is a space in this channel for 1 piece of
// data that we are using the signal and we don't have to wait for the other side to get it. So
// now a Goroutine comes in, put the data in and then move away immediately. In other word, the
// send is happening before the receive. All the sending Goroutine know is that it issues the
// signal, put that data but has no clue when the signal is going to be received. Now hopefully a
// Goroutine comes in. It see that there is a data there, receive it and move on.

//  G                      G
//  |      Channel (1)     |
//  |      ----------      |
//  |---->|    D     |<----|
//  |      ----------      |
//  |                      |

// We use a buffered of 1 when dealing with these type of latency. We may need buffers that are
// larger but there are some design rules that we are gonna learn later on we use buffers that are
// greater than 1. But if we are in a situation where we can have these sends coming in and they
// could potentially be locked then we have to think again: if the channel of 1 is fast enough to
// reduce the latency that we are dealing with. Because what's gonna happen is the following:
// What we are hoping is, the buffered channel is always empty every time we perform a send.

// Buffered channel ares not for performance. What the buffered channel need to be used for is
// continuity, to keep the wheel moving. One thing we have to understand is that, everybody can
// write a piece of software that works when everything is going well. When things are going
// bad, it's where the architecture and engineer really come in. Our software doesn't enclose and
// it doesn't cost stress. We need to be responsible.

// Back to the example, it's not important that we know exactly the signaling data was received but
// we do have to make sure that it was. The buffered channel of 1 gives us almost guarantee because
// what happen is: it performs a send, puts the data in there, turns around and when it comes back,
// it sees that the buffered is empty. Now we know  that it was received. We don't know immediately
// at the time that we sent but by using a buffer of 1, we do know that is empty when we come back.
// Then it is okay to put another piece of data in there and hopefully when we come back again, it
// is gone. If it's not gone, we have a problem. There is a problem upstream. We cannot move
// forward until the channel is empty. This is something that we want to report immediately because
// we want to know why the data is still there. That's how we can build systems that are reliable.
// We don't take more work at any give time. We identify upstream when there is a problem so we
// don't put more stress on our systems. We don't take more responsibilities for things that we
// shouldn't be.

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("\n=> Basics of a send and receive\n")
	basicSendRecv()

	fmt.Printf("\n=> Close a channel to signal an event\n")
	signalClose()
}

// ---------------------------------------
// Unbuffered channel: Signaling with data
// ---------------------------------------

// basicSendRecv shows the basics of a send and receive.
// We are using make function to create a channel. We have no other way of creating a channel that
// is usable until we use make.
// Channel is also based on type, a type of data that we are gonna do the signaling. In this case,
// we use string. That channel is a reference type. ch is just a pointer variable to larger data
// structure underneath.
func basicSendRecv() {
	// This is an unbuffered channel.
	ch := make(chan string)

	go func() {
		// This is a send: a binary operation with the arrow pointing into the channel.
		// We are signaling with a string "hello".
		ch <- "hello"
	}()

	// This is a receive: also an arrow but it is a unary operation where it is attached to the
	// left hand side of the channel to show that is coming out.
	// We are now have a unbuffered channel where the send and receive have to come together. We
	// also know that the signal has been received because the receive happens first.
	// Both are gonna block until both come together so the exchange can happen.
	fmt.Println(<-ch)
}

// ------------------------------------------
// Unbuffered channel: Signaling without data
// ------------------------------------------

// signalClose shows how to close a channel to signal an event.
func signalClose() {
	// We are making a channel using an empty struct. This is a signal without data.
	ch := make(chan struct{})

	// We are gonna launch a Goroutine to do some work. Suppose that it's gonna take 100
	// millisecond. Then, it wants to signal another Goroutine that it's done.
	// It's gonna close the channel to report that it's done without the need of data.

	// When we create a channel, buffered or unbuffered, that channel can be in 2 different states.
	// All channels start out in open state so we can send and receive data. When we change the
	// state to be closed, it cannot be opened. We also cannot close the channel twice because that
	// is an integrity issue. We cannot signal twice without data twice.

	go func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("signal event")
		close(ch)
	}()

	// When the channel is closed, the receive will immediately return.
	// When we receive on a channel that is open, we cannot return until we receive the data
	// signal. But if we receive on a channel that is closed, we are able to receive the signal
	// without data. We know that event is occurred. Every receive on that channel will immediately
	// return.
	<-ch

	fmt.Println("event received")
}
