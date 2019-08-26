// ----------------
// Read/Write Mutex
// ----------------

// There are times when we have a shared resource where we want many Goroutines reading it.
// Occasionally, one Goroutine can come in and make change to the resource. When that happen, every
// body has to stop reading. It doesn't make sense to synchronize read in this type of scenario
// because we are just adding another latency to our software for no reason.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// data is a slice that will be shared.
	data []string

	// rwMutex is used to define a critical section of code.
	// It is a little bit slower than Mutex but we are optimizing the correctness first so we don't
	// care about that for now.
	rwMutex sync.RWMutex

	// Number of reads occurring at any given time.
	// As soon as we see int64 here, we should start thinking about using atomic instruction.
	readCount int64
)

// init is called prior to main.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(1)

	// Create a writer Goroutine that performs 10 different writes.
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			writer(i)
		}
		wg.Done()
	}()

	// Create eight reader Goroutines that runs forever.
	for i := 0; i < 8; i++ {
		go func(i int) {
			for {
				reader(i)
			}
		}(i)
	}

	// Wait for the write Goroutine to finish.
	wg.Wait()
	fmt.Println("Program Complete")
}

// writer adds a new string to the slice in random intervals.
func writer(i int) {
	// Only allow one Goroutine to read/write to the slice at a time.
	rwMutex.Lock()
	{
		// Capture the current read count.
		// Keep this safe though we can due without this call.
		// We want to make sure that no other Goroutines are reading. The value of rc should always
		// be 0 when this code run.
		rc := atomic.LoadInt64(&readCount)

		// Perform some work since we have a full lock.
		fmt.Printf("****> : Performing Write : RCount[%d]\n", rc)
		data = append(data, fmt.Sprintf("String: %d", i))
	}
	rwMutex.Unlock()
}

// reader wakes up and iterates over the data slice.
func reader(id int) {
	// Any Goroutine can read when no write operation is taking place.
	// RLock has the corresponding RUnlock.
	rwMutex.RLock()
	{
		// Increment the read count value by 1.
		rc := atomic.AddInt64(&readCount, 1)

		// Perform some read work and display values.
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		fmt.Printf("%d : Performing Read : Length[%d] RCount[%d]\n", id, len(data), rc)

		// Decrement the read count value by 1.
		atomic.AddInt64(&readCount, -1)
	}
	rwMutex.RUnlock()
}

// Lesson:
// -------
// The atomic functions and mutexes create latency in our software. Latency can be good when we
// have to coordinate orchestrating. However, if we can reduce latency using Read/Write Mutex, life
// is better.

// If we are using mutex, make sure that we need to get in and out of mutex as fast as possible. Do
// not anything extra. Sometime just reading the shared state into a local variable is all we need
// to do, The less operation we can perform on the mutex, the better. We then reduce the latency to
// the bare minimum.
