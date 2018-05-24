// ------
// Select
// ------

// This sample program demonstrates how to use a channel to monitor the amount of time
// the program is running and terminate the program if it runs too long.

package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

// Give the program 3 seconds to complete the work.
const timeoutSeconds = 3 * time.Second

// There are 4 channels that we are gonna use: 3 unbuffered and 1 buffered of 1.
var (
	// sigChan receives operating signals.
	// This will allow us to send a Ctrl-C to shut down our program cleanly.
	sigChan = make(chan os.Signal, 1)

	// timeout limits the amount of time the program has.
	// We really don't want to receive on this channel because if we do, that means something bad
	// happens, we are timing out and we need to kill the program.
	timeout = time.After(timeoutSeconds)

	// complete is used to report processing is done.
	// This is the channel we want to receive on. When the Goroutine finish the job, it will signal
	// to us on this complete channel and tell us any error that occurred.
	complete = make(chan error)

	// shutdown provides system wide notification.
	shutdown = make(chan struct{})
)

func main() {
	log.Println("Starting Process")

	// We want to receive all interrupt based signals.
	// We are using a Notify function from the signal package, passing sigChan telling the channel
	// to look for anything that is os.Interrupt related and sending us a data signal on this
	// channel.
	// One important thing about this API is that, it won't wait for us to be ready to receive the
	// signal. If we are not there, it will drop it on the floor. That's why we are using a
	// buffered channel of 1. This way we guarantee to get at least 1 signal. When we are ready to
	// act on that signal, we can come over there and do it.
	signal.Notify(sigChan, os.Interrupt)

	// Launch the process.
	log.Println("Launching Processors")

	// This Goroutine will do the processing job, for example image processing.
	go processor(complete)

	// The main Goroutine here is in this event loop and it's gonna loop forever until the program
	// is terminated.
	// There are 3 cases in select, meaning that there are 3 channels we are trying to receive on
	// at the same time: sigChan, timeout, and complete.

ControlLoop:
	for {
		select {
		case <-sigChan:
			// Interrupt event signaled by the operation system.
			log.Println("OS INTERRUPT")

			// Close the channel to signal to the processor it needs to shutdown.
			close(shutdown)

			// Set the channel to nil so we no longer process any more of these events.
			// If we try to send on a closed channel, we are gonna panic. If we receive on a closed
			// channel, that's gonna immediately return a signal without data. If we receive on a
			// nil channel, we are blocked forever. Similar with send.
			// Why do we want to do that?
			// We don't want user to hold down Ctrl C or hit Ctrl C multiple times. If they do that
			// and we process the signal, we have to call close multiple time. When we call close
			// on a channel that is already closed, the code will panic. Therefore, we cannot have
			// that.
			sigChan = nil

		case <-timeout:
			// We have taken too much time. Kill the app hard.
			log.Println("Timeout - Killing Program")

			// os.Exit will terminate the program immediately.
			os.Exit(1)

		case err := <-complete:
			// Everything completed within the time given.
			log.Printf("Task Completed: Error[%s]", err)

			// We are using a label break here.
			// We put one at the top of the for loop so the case has a break and the for has a
			// break.
			break ControlLoop
		}
	}

	// Program finished.
	log.Println("Process Ended")
}

// processor provides the main program logic for the program.
// There is something interesting in the parameter. We put the arrow on the right hand side of the
// chan keyword. It means this channel is a send-only channel. If we try to receive on this
// channel, the compiler will give us an error.
func processor(complete chan<- error) {
	log.Println("Processor - Starting")

	// Variable to store any error that occurs.
	// Passed into the defer function via closures.
	var err error

	// Defer the send on the channel so it happens regardless of how this function terminates.
	// This is an anonymous function call like we saw with Goroutine. However, we are using the
	// keyword defer here.
	// We want to execute this function but after the processor function returns. This gives us an
	// guarantee that we can have certain things happen before control go back to the caller.
	// Also, defer is the only way to stop a panic. If something bad happens, say the image library
	// is blowing up, that can cause a panic situation throughout the code. In this case, we want
	// to recover from that panic, stop it and then control the shutdown.
	defer func() {
		// Capture any potential panic.
		if r := recover(); r != nil {
			log.Println("Processor - Panic", r)
		}

		// Signal the Goroutine we have shutdown.
		complete <- err
	}()

	// Perform the work.
	err = doWork()

	log.Println("Processor - Completed")
}

// doWork simulates task work.
// Between every single call, we call checkShutdown. After complete every tasks, we are asking:
// Have we been asked to shutdown? The only way we know is that shutdown channel is closed. The
// only way to know if the shutdown channel is closed is to try to receive. If we try to receive on
// a channel that is not closed, it's gonna block. However, the default case is gonna save us here.
func doWork() error {
	log.Println("Processor - Task 1")
	time.Sleep(2 * time.Second)

	if checkShutdown() {
		return errors.New("Early Shutdown")
	}

	log.Println("Processor - Task 2")
	time.Sleep(1 * time.Second)

	if checkShutdown() {
		return errors.New("Early Shutdown")
	}

	log.Println("Processor - Task 3")
	time.Sleep(1 * time.Second)

	return nil
}

// checkShutdown checks the shutdown flag to determine if we have been asked to interrupt processing.
func checkShutdown() bool {
	select {
	case <-shutdown:
		// We have been asked to shutdown cleanly.
		log.Println("checkShutdown - Shutdown Early")
		return true

	default:
		// If the shutdown channel was not closed, presume with normal processing.
		return false
	}
}

// Output:
// -------
// - When we let the program run, since we configure the timeout to be 3 seconds, it will
// then timeout and be terminated.
// - When we hit Ctrl C while the program is running, we will see the OS INTERRUPT and the program
// is being shutdown early.
// - When we send a signal quit by hitting Ctrt \, we will get a full stack trace of all the
// Goroutines.
