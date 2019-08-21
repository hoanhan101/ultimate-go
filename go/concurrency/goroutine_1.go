// ----------------------
// Go Scheduler Internals
// ----------------------

// Every time our Go's program starts up, it looks to see how many cores are available. Then it
// creates a logical processor.

// The operating system scheduler is considered a preemptive scheduler. It runs down there in the
// kernel. Its job is to look at all the threads that are in runnable states and gives them the
// opportunity to run on some cores. These algorithms are fairly complex: waiting, bouncing
// threads, keeping memory of threads, caching,... The operating system is doing all of that for
// us. The algorithm is really smart when it comes to multicore processor. Go doesn't want to
// reinvent the wheel. It wants to sit on top of the operating system and leverage it.

// The operating system is still responsible for operating system threads, scheduling operating
// system threads efficiently. If we have a 2 core machine and a thousands threads that the
// operating system has to schedule, that's a lot of work. A context switch on on some operating
// system thread is expensive when the operating system have no clues of what that thread is doing.
// It have to save all the possible state in order to be able to restore that to exactly the way it
// was. If there is less threads, each thread can get more time to be reschedule. If the is more
// thread, each thread have less time over a long period of time.

// "Less is more" is a really big concept here when we start to write concurrent software. We want to
// leverage the preemptive scheduler. So the Go's scheduler, the logical processor actually runs in
// user mode, the mode our application is running at. Because of that, we have to call the Go's
// scheduler a cooperating scheduler. What brilliant here is the runtime that coordinating the
// operation. It sill look and feel as a preemptive scheduler up in user land. We will see how "less
// is more" concept gets to present itself and we get to do a lot more work with less. Our goal needs
// to be how much work we get done with the less number of threads.

// Think about this in a simple way because processors are complex: hyperthreading, multiple threads
// per core, clock cycle. We can only execute one operating system thread at a time on any given
// core. If we only have 1 core, only 1 thread can be executed at a time. Anytime we have more
// threads in runnable states than we have cores, we are creating load, latency and we are getting
// more work done as we want. There needs to be this balance because not every thread is
// necessarily gonna be active at the same time. It all comes down to determining, understanding
// the workload for the software that we are writing.

// Back to the first idea, when our Go program comes up, it has to see how many cores that
// available. Let's say it found 1. It is going to create a logical processor P for that core.
// Again, the operating system is scheduling things around operating system threads. What this
// processor P will get is an m, where m stands for machine. It represents an operating system
// thread that the operating system is going to schedule and allows out code to run.

// The Linux scheduler has a run queue. Threads are placed in run queue in certain cores ore
// family of cores and those are constantly bounded as threads are running. Go is gonna do the same
// thing. Go has its run queue as well. It has Global Run Queue (GRQ) and every P has a Local Run
// Queue (LRQ).

// Goroutine
// ---------
// What is a Goroutine? It is a path of execution. Threads are paths of execution. That path of
// execution needs to be scheduled. In Go, every function or method can be created to be a
// Goroutine, can become an independent path of execution that can be scheduled to run on some
// operating system threads against some cores.

// When we start our Go program, the first thing runtime gonna do is creating that a Go routine
// and putting that in some main LRQ for some P. In our case, we only have 1 P here so we can
// imagine that Goroutine is attached to P.

// A Goroutine, just like thread, can be in one of three major states: sleeping, executing or in
// runnable state asking to wait for some time to execute on the hardware. When the runtime creates
// a Goroutine, it is gonna placed in P and multiplex on this thread. Remember that it's the
// operating system that taking the thread, scheduling it, placing it on some core and doing
// execution. So Go's scheduler is gonna take all the code related to that Goroutine's path of
// execution, place it on a thread, tell the operating system that this thread is in runnable state
// and can we execute it. If the answer is yes, the operating system starts to execute on some
// cores there in the hardware.

// As the main Goroutine runs, it might want to creates more path of execution, more Goroutines.
// When that happens, those Goroutines might find themselves initially in the GRQ. These would be
// Goroutines that are in runnable states but haven't been assigned to some P yet. Eventually, they
// would end up in the LRQ where they're saying they would like some time to execute.

// This queue does not necessarily follow First-In-First-Out protocol. We have to understand that
// everything here is non-deterministic, just like the operating system scheduler is. We cannot
// predict what the scheduler is gonna do when all things are equal. It is gonna make sure there is
// a balance. Until we get into orchestration, till we learn how to coordinate these execution of
// these Goroutines, there is no predictability.

// Here is the mental model of our example.
// GRQ
//      m
//      |
//    -----          LRQ
//   |  P  | ----------
//    -----           |
//      |             G1
//     Gm             |
//                    G2

// We have Gm executing on for this thread for this P, and we are creating 2 more Goroutines G1 and
// G2. Because this is a cooperating scheduler, that means that these Goroutines have to cooperate
// to be scheduled, to be swapped context switch on this operating system thread m.

// There are 4 major places in our code where the scheduler has the opportunity to make a
// scheduling decision.
// - The keyword Go that we are going to create Goroutines. That is also an opportunity for the
// scheduler to rebalance when it has multiple P.
// - A system call. These system calls tend to happen all the time already.
// - A channel operation because there is mutex (blocking call) that we will learn later.
// - Garbage collection.

// Back to the example, says the scheduler might decide Gm has enough time to run, it will put Gm
// back to the run queue and allow G1 to run on that m. We are now having context switch.
//      m
//      |
//    -----      LRQ
//   |  P  | ------
//    -----       |
//      |         Gm
//      G1        |
//                G2

// Let's say G1 decides to open up a file. Opening up a file can microsecond or 10 milliseconds. We
// don't really know. If we allow this Goroutine to block this operating system thread while we
// open up that file, we are not getting more work done. In this scenario here, having a single P,
// we are single threaded software application. All Goroutines only execute on the m attached to
// this P. What happen is this Goroutine is gonna block this m for potential a long time. We are
// basically be stalled while we still have works that need to get done. So the scheduler is not
// gonna allow that to happen, What actually happen is that the scheduler is gonna detach that m
// and G1. It is gonna bring a new m, say m2, then decide what G from the run queue should run
// next, say G2.
//            m2
//            |
//    m     -----      LRQ
//    |    |  P  | ------
//    G1    -----       |
//            |         Gm
//            G2

// We are now have 2 threads in a single threaded program. From our perspective, we are still
// single threading because the code that we are writing, the code associated with any G can only
// run against this P and this m. However, we don't know at any given time what m we are running on. m
// can get swapped out but we are still single threaded.

// Eventually, G1 will come back, the file will be opened. The scheduler is gonna take this G1 and
// put it back to the run queue so we can be executed against on this P for some m (m2 in this
// case). m is get placed on the side for later use. We are still maintaining these 2 threads. The
// whole process can happen again.
//            m2
//            |
//    m     -----      LRQ
//         |  P  | ------
//          -----       |
//            |         Gm
//            G2        |
//                      G1

// It is a really brilliant system of trying to leverage this thread to its fullest capability by
// doing more on 1 thread. Let's do so much on this thread we don't need another.

// There is something called a Network poller. It is gonna do all the low level networking
// asynchronous networking stuff. Our G, if it is gonna do anything like that, might be moved out
// to the Network poller and then brought back in. From our perspective, here is what we have to
// remember:
// The code that we are writing always run on some P against some m. Depending on how many P we
// have, that's how many threads variables for us to run.

// Concurrency is about managing a lot of thing at once. This is what the scheduler is doing. It
// manages the execution of these 3 Goroutines against this one m for this P. Only 1 Goroutine can
// be executed at a single time.

// If we want to do something in parallel, which means doing a lot of things at once, then we would
// need to create another P that has another m, say m3.
//    m3              m2
//    |               |
//  -----            -----      LRQ
// |  P  | ------   |  P  | ------
//  -----       |    -----       |
//    |         Gx     |         Gm
//    Gx        |      G2        |
//              Gx               G1

// Both are are scheduled by the operating system. So now we can have 2 Goroutines running at the
// same time in parallel.

// Let's try another example.
// --------------------------
// We have a multiple threaded software. The program launched 2 threads. Even if both threads end
// up on the same cord, each want to pass a message to each other. What has to happen from the
// operating system point of view?
// We have to wait for thread 1 to get scheduled and placed on some cores - a context switch (CTX)
// has to happen here. While that's happening, thread is is asleep so it's not running at all. From
// thread 1, we send a message over and want to wait to get a message back. In order to do that,
// there is another to context switch need to be happened because we can put a different thread on
// that core (?). We are waiting for the operating system to schedule thread 2 so we are going to
// get another context switch, waking up and running, processing the message and sending the
// message back. On every single message that we are passing back and forth, thread is gonna from
// executable state to runnable state to asleep state. This is gonna cost a lot of context switches
// to occur.
//  T1                         T2
//  | CTX      message     CTX |
//  | -----------------------> |
//  | CTX                    | |
//  | CTX      message     CTX |
//  | <----------------------- |
//  | CTX                  CTX |

// Let's see what happen when we are using Goroutines, even on a single core.
// G1 wants to send a message to G2 and we perform a context switch. However, the context here
// is user's space switch. G1 can be taken of the thread and G2 can be put on the thread. From the
// operating system point of view, this thread never go to sleep. This thread is always executing
// and never needed to be context switched out. It is the Go's scheduler that keeps the Goroutines
// context switched.
//               m
//               |
//             -----
//            |  P  |
//             -----
//  G1                         G2
//  | CTX      message     CTX |
//  | -----------------------> |
//  | CTX                    | |
//  | CTX      message     CTX |
//  | <----------------------- |
//  | CTX                  CTX |

// If a P for some m here has no work to do, there is no G, the runtime scheduler will try to spin
// that m for a little bit to keep it hot on the core. Because if that thread goes cold, the
// operating system will pull it off the core and put something else on. So it just spin a little
// bit to see if there will be another G comes in to get some work done.

// This is how the scheduler work underneath. We have a P, attached to thread m. The operating
// system will do the scheduling. We don't want any more than we have cores. We don't need any more
// operating system threads than we have cores. If we have more threads then we have cores, all we
// do is putting load on the operating system. We allow the Go's scheduler to make decisions on our
// Goroutines, keeping the least number of threads we need and hot all time if we have work. The
// Go's scheduler is gonna look and feel preemptive even though we are calling a cooperating
// scheduler.

// However, let's not think about how the scheduler work. Think the following way makes it easier
// for future development.
// Every single G, every Goroutine that is in runnable state, is running at the same time.

package main

import "fmt"

func main() {
	fmt.Println("ok")
}
