// Package uid provides the ability to generate Un/Ordered Identifiers. These
// are also most likely very unique. At the very least, they are unique enough.
//
// simplicity > ∞
//
// OIDs are ordered, UIDs are not and the use of math/rand means these are not
// cryptographically secure.
//
// If there is a desire for greater flexibility use a Generator which allows the
// customization of the final encoding (base32, base64, hex, etc. (see
// encoders.go)) and the entropy source and length (math/rand, crypto/rand,
// etc).
//
// By default, an OID is the base 32 encoding of a binary encoded string
// comprising an 8 byte nanosecond precision unix timestamp and an 8 byte random
// number, in that order. The timestamp prefix allows these IDs to be ordered.
//
// If (by any chance) OID is called in the same nanosecond, the random number is
// incremented instead of a new one being generated. This makes sure that two
// consecutive IDs generated in the same goroutine are different and ordered.
//
// By default a UID is the base 32 encoding of a binary encoded string
// comprising two 8 byte random numbers.
//
// Both are safe for concurrent use as they provide their own locking. OID will
// run out of nanoseconds on Fri, 11 Apr 2262 23:47:16 UTC
package uid
