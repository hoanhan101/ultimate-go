// ---------
// CPU CACHE
// ---------
// Cores DO NOT access main memory directly but their local caches.
// What store in caches are data and instruction.

// Cache speed from fastest to slowest: L1 -> L2 -> L3 -> main memory.
// Scott Meyers: "If performance matter then the total memeory you have is the total amount of
// caches" -> access to main memory is incredibly slow; practically speaking it might not even be there.

// How do we write code that can be sympathetic with the caching system to make sure that
// we don't have a cache miss or at least, we minimalize cache misses to our fullest potential?

// Processor has a Prefetcher. It predicts what data is needed ahead of time.
// There are different granularity depending on where we are on the machine.
// Our programming model uses a byte. We can read and write to a byte at a time. However, from the
// caching system POV, our granularity is not 1 byte. It is 64 bytes, called a cache line. All
// memory us junked up in this 64 bytes cache line.

// Since the caching mechanism is complex, Prefetcher tries to hide all the latency from us.
// It has to be able to pick up on predictable access pattern to data.
// -> We need to write code that create predictable access pattern to data

// One easy way is to create a contiguous allocation of memory and to iterate over them.
// The array data structure gives us ability to do so.
// From the hardware perspective, array is the most important data structure.
// From Go perspective, slice is. Array is the backing data structure for slice (like Vector in C++).
// Once we allocate an array, whatever it size, every element is equal distant from other element.
// As we iterate over that array, we begin to walk cache line by cache line. As the Prefetcher see
// that access pattern, it can pick it up and hide all the latency from us.

// For example, we have a big nxn matrix. We do LinkedList Traverse, Column Traverse, and Row Traverse
// and benchmark against them.
// Unsurprisingly, Row Traverse has the best performance. It walk through the matrix cache line
// by cache line and create a predictable access pattern.
// Column Traverse does not walk through the matrix cache line by cache line. It looks like random
// access memory pattern. That is why is the slowest among those.
// However, that doesn't explain why the LinkedList Traverse's performance is in the middle. We
// just think that it might perform as poorly as the Column Traverse.
// -> This leads us to another cache: TLB - Translation lookaside buffer. Its job is to maintain
// operating system page and offset to where physical memory is.

// Back to the different granularity, the caching system moves data in and out the hardware at 64
// bytes at a time. However, the operating system manages memory by paging its 4K (traditional page
// size for an operating system).
// TLB: For every page that we are managing, let's take our virtual memory addresses because that
// that we use (softwares run virtual addresses, its sandbox, that is how we use/share physical
// memory) and map it to the right page and offset for that physical memory.

// A miss on the TLB can be worse than just the cache miss alone.
// The LinkedList is somewhere in between is because the chance of multiple nodes being on the same
// page is probably pretty good. Even though we can get cache misses because cache lines aren't
// necessary in the distance that is predictable, we probably not have so many TLB cache misses.
// In the Column Traverse, not only we have cache misses, we probably have a TLB cache miss on
// every access as well.

// Data-oriented design matters.
// It is not enough to write the most efficient algorithm, how we access our data can have much
// more lasting effect on the performance than the algorithm itself.

package main

import "fmt"

func main() {
	fmt.Println("ok")
}
