// Package uid provides the ability to generate Un/Sortable Identifiers. These
// are also most likely very unique. At the very least, they are unique
// enough for my purposes, and hopefully yours as well.
//
// This package is simple and only provides two functions. For my purposes
// simplicity > speed.
//
// OIDs are sortable, UIDs are not
//
// OID returns a 26 char base 32 encoded string based on timestamp and a random
// number. The base32( binary( XY ) ) where X is the timestamp ([8]byte) and Y
// is the random number ([8]byte).
//
// If (by any chance) OID is called in the same nanosecond, the random number
// is incremented instead of a new one being generated. This makes sure that two
// consecutive Ids generated in the same goroutine also ensure those Ids are
// also sortable.
//
// UID is the same as OID accept that the [8]byte timestamp is replaced with
// [8]byte random data. These IDs are not sortable.
//
// The use of math/rand means these are not "cryptographically secure".
//
package uid

import (
	"bytes"
	"encoding/base32"
	"encoding/binary"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	// Remember the lastTime so that if (by chance) we get the same NanoSecond,
	// we just incrememt the last random number.
	lastTime int64
	// use base32 to make ascii URL safe strings; using Crockfords
	// dict: https://www.crockford.com/base32.html
	encoder = base32.NewEncoding("0123456789abcdefghjkmnpqrstvwxyz").WithPadding(base32.NoPadding)
	// hold 8 bytes of random data
	lastRand = make([]byte, 8)
	// handle our own concurrency
	mu = &sync.Mutex{}
	// avoid unnecessarily consuming memory
	buf bytes.Buffer
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
	// lock for lastTime, lastRand, and chars
	mu.Lock()
	defer mu.Unlock()

	now := time.Now().UnixNano()

	// if we have the same time, just inc lastRand, else create a new one
	if now == lastTime {
		lastRand[7]++
	} else {
		rand.Read(lastRand)
	}

	// remember this for next time
	lastTime = now

	buf.Reset() // clean our buffer before use
	binary.Write(&buf, binary.BigEndian, now)
	binary.Write(&buf, binary.BigEndian, lastRand)
	return encoder.EncodeToString(buf.Bytes())
}

// UID is the same as OID accept that the 8 byte timestamp is replaced with
// an 8 byte random number. These IDs are not sortable.
func UID() string {
	// lock for lastTime, lastRand, and chars
	mu.Lock()
	defer mu.Unlock()

	buf.Reset() // clean our buffer before use
	rand.Read(lastRand)
	binary.Write(&buf, binary.BigEndian, lastRand)
	rand.Read(lastRand)
	binary.Write(&buf, binary.BigEndian, lastRand)
	return encoder.EncodeToString(buf.Bytes())
}
