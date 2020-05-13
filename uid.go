// Package uid provides the ability to generate Un/Sortable Identifiers. These
// are also most likely very unique. At the very least, they are unique enough
// for my purposes, and potentially others' as well.
//
// For my purposes simplicity > speed.
//
// OIDs are sortable, UIDs are not
//
// OID returns a 26 char base 32 encoded string based on timestamp and a random
// number. The base32( binary( XY ) ) where X is the timestamp (8 bytes) and Y
// is a random number (8 bytes).
//
// If (by any chance) OID is called in the same nanosecond, the random number is
// incremented instead of a new one being generated. This makes sure that two
// consecutive Ids generated in the same goroutine also ensure those Ids are
// also sortable.
//
// UID is the same as OID accept that the timestamp is replaced with an 8 byte
// random number. These IDs are not sortable.
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

// OID returns a base 32 encoded string based on timestamp and a random number.
// The `base32( binary( XY ) )` where X is an int64 timestamp (8 bytes) and Y is
// a random number (8 bytes).
//
// If (by any chance) OID is called in the same nanosecond, the random number is
// incremented instead of a new one being generated. This makes sure that two
// consecutive IDs generated in the same goroutine are different and sortable.
//
// It is safe for concurrent use as it provides its own locking.
func OID() string {
	return defaultGenerator.OID()
}

// UID is the same as OID accept that the 8 byte timestamp is replaced with
// an 8 byte random number. These IDs are not sortable.
func UID() string {
	return defaultGenerator.UID()
}
