// Command oid is a simple utility for generating un/ordered IDs.
//
// By default, it is the base 32 encoding of a binary encoded string
// comprising an 8 byte nanosecond precision unix timestamp and an 8 byte random
// number, in that order. The timestamp prefix allows these IDs to be ordered.
//
// Using '-uid' makes it a base 32 encoding of a binary encoded string
// comprising two 8 byte random numbers. These are not ordered
//
// Usage: oid [option [option]...]
//
// Options:
//   -h	display this message
//   -help
//     	display this message
//   -hex
//     	Use a standard HEX (a-f0-9) dictionary vs the default base32 (Crockford) dictionary.
//   -len int
//     	The number of bytes of randomness to use. Keep in mind that UIDs use double this value for the overall length, while OIDs are this value +8 Bytes for the timestamp. (default 8)
//   -secure
//     	Use a cryptographically secure randomness generator.
//   -uid
//     	Generate an 'Unordered' ID (UID) vs the default 'Ordered' ID (OID).
package main
