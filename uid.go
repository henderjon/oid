// Package uid provides the ability to generate Un/Ordered Identifiers. These
// are also most likely very unique. At the very least, they are unique enough
// for my purposes, and potentially others' as well.
//
// For my purposes simplicity > ∞.
//
// OIDs are ordered, UIDs are not
//
// The use of math/rand means these are not cryptographically secure.
//
// If there is a desire for greater flexibility use a Generator which allows the
// customization of the final encoding (base32, base64, hex, etc. (see
// encoders.go)) and the entropy source and length (math/rand, crypto/rand,
// etc).
package uid

var (
	defaultGenerator = DefaultGenerator()
)

// OID returns a base 32 encoded binary encoded string based on timestamp and a
// random number.
//
//  +--------+--------+
//  |   TS   |   Ent  |
//  +--------+--------+
//
// TS is the binary encoding of an int64 (8 byte) Unix Timestamp in Nanoseconds
// Ent is the binary encoding of an 8 byte random number
//
// The 16 bytes are then base32 encoded for human readability and URL safety.
//
// If (by any chance) OID is called in the same nanosecond, the random number is
// incremented instead of a new one being generated. This makes sure that two
// consecutive IDs generated in the same goroutine are different and ordered.
//
// It is safe for concurrent use as it provides its own locking.
//
// OID will run out of nanoseconds on Fri, 11 Apr 2262 23:47:16 UTC
func OID() string {
	return defaultGenerator.OID()
}

// UID is the same as OID accept that the 8 byte timestamp is replaced with
// an 8 byte random number. These IDs are not ordered.
//
//  +--------+--------+
//  |  Ent   |   Ent  |
//  +--------+--------+
//
// Ent is the binary encoding of an 8 byte random number
//
// The 16 bytes are then base32 encoded for human readability and URL safety.
func UID() string {
	return defaultGenerator.UID()
}
