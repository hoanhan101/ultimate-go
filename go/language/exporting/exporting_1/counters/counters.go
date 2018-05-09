// Package counters provides alert counter support.
package counters

// AlertCounter is an exported named type that contains an integer counter for alerts.
// The first word is in upper-case format so it is considered to be exported.
type AlertCounter int

// alertCounter is an unexported named type that contains an integer counter for alerts.
// The first word is in lower-case format so it is considered to be unexported.
// It is not accessible out the package, unless it is part of the package counter itself.
type alertCounter int
