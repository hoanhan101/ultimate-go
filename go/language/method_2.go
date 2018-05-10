// ---------------------------
// Value and Pointer semantics
// ---------------------------

// When it comes to use built-in type (numeric, string, bool), we should always be using value
// semantics. When a piece of code that take an address of an integer or a bool, this raises a
// big flag. It's hard to say because it depends on the context. But in general, why should
// these values end up on the heap creating garbage? These should be stay on the stack. There
// is an exception to everything. However, until we know it is okay to take the exception, we
// should follow the guideline.

// The reference type (slice, map, channel, interface) also focuses on using value semantic.
// The only time we want to take the address of a slice is when we are sharing it down the call
// stack to Unmarshal function since it always requires the address of a value.

// Examples below are from standard library. By studying them, we learn how important it is to use
// value or pointer semantics in a consistent way.

// When we declare a type, we must ask ourselves immediately:
// - Does this type require value semantic or pointer semantic?
// - If I need to modify this value, should we create a new value or should we modify the value
// itself so everyone can see it?
// It needs to be consistent. It is okay to guess it wrong the first time and refactor it later.

package main

import (
	"sync/atomic"
	"syscall"
)

// --------------
// Value semantic
// --------------

// These is a named type from the net package called IP and IPMask with a base type that is a
// slice of bytes. Since we use value semantics for reference types, the implementation is
// using value semantics for both.
type IP []byte
type IPMask []byte

// Mask is using a value receiver and returning a value of type IP.
// This method is using value semantics for type IP.
func (ip IP) Mask(mask IPMask) IP {
	if len(mask) == IPv6len && len(ip) == IPv4len && allFF(mask[:12]) {
		mask = mask[12:]
	}
	if len(mask) == IPv4len && len(ip) == IPv6len && bytesEqual(ip[:12], v4InV6Prefix) {
		ip = ip[12:]
	}
	n := len(ip)
	if n != len(mask) {
		return nil
	}
	out := make(IP, n)
	for i := 0; i < n; i++ {
		out[i] = ip[i] & mask[i]
	}
	return out
}

// ipEmptyString accepts a value of type IP and returns a value of type string.
// The function is using value semantics for type IP.
func ipEmptyString(ip IP) string {
	if len(ip) == 0 {
		return ""
	}
	return ip.String()
}

// ----------------
// Pointer semantic
// ----------------

// Should time use value or pointer semantics?
// If you need to modify a time value should you mutate the value or create a new one?
type Time struct {
	sec  int64
	nsec int32
	loc  *Location
}

// The best way to understand what semantic is going to be used is to look at the factory function
// for type. It dictates the semantics that will be used. In this example, the Now function
// returns a value of type Time. It is making a copy of it Time value and passing it back up.
// This means Time value can be stayed on the stack. We should be using semantic all the way
// through.
func Now() Time {
	sec, nsec := now()
	return Time{sec + unixToInternal, nsec, Local}
}

// Add is a mutation operation. If we go with the idea that we should be using pointer semantic
// when we mutate something and value semantic when we don't then Add is implemented wrong.
// However, it has not been wrong because it is the type that has to drive the semantic, not the
// implementation of the method. The method must adhere to the semantic that we choose.
// Add is using a value receiver and returning a value of type Time. It is mutating its local copy
// and returning our something new.
func (t Time) Add(d Duration) Time {
	t.sec += int64(d / 1e9)
	nsec := int32(t.nsec) + int32(d%1e9)
	if nsec >= 1e9 {
		t.sec++
		nsec -= 1e9
	} else if nsec < 0 {
		t.sec--
		nsec += 1e9
	}
	t.nsec = nsec
	return t
}

// div accepts a value of type Time and returns values of built-in types.
// The function is using value semantics for type Time.
// func div(t Time, d Duration) (qmod2 int, r Duration) {}

// The only use pointer semantics for the `Time` API are these Unmarshal related functions:
// func (t *Time) UnmarshalBinary(data []byte) error {}
// func (t *Time) GobDecode(data []byte) error {}
// func (t *Time) UnmarshalJSON(data []byte) error {}
// func (t *Time) UnmarshalText(data []byte) error {}

// Observation:
// ------------
// Most struct types are not going to be able to leverage value semantic. Most struct types are
// probably gonna be data that should be shared or more efficient to be shared. For example, an
// User type. Regardless it is possible to copy an User type but it is not a proper thing to do in
// real world.

// Other examples:
// Factory functions dictate the semantics that will be used. The Open function returns a
// pointer of type File. This means we should be using pointer semantics and share File values.
func Open(name string) (file *File, err error) {
	return OpenFile(name, O_RDONLY, 0)
}

// Chdir is using a pointer receiver. This method is using pointer semantics for File.
func (f *File) Chdir() error {
	if f == nil {
		return ErrInvalid
	}
	if e := syscall.Fchdir(f.fd); e != nil {
		return &PathError{"chdir", f.name, e}
	}
	return nil
}

// epipecheck accepts a pointer of type File. The function is using pointer semantics for type File.
func epipecheck(file *File, e error) {
	if e == syscall.EPIPE {
		if atomic.AddInt32(&file.nepipe, 1) >= 10 {
			sigpipe()
		}
	} else {
		atomic.StoreInt32(&file.nepipe, 0)
	}
}
