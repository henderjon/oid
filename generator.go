package uid

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"sync"
	"time"
)

// Generator wraps all the necessary components for creating UIDs and OIDs in an
// injectable package.
type Generator struct {
	lastTime int64
	encoder  EncoderToString
	lastRand []byte
	source   io.Reader
	mu       sync.Mutex
	buf      bytes.Buffer
}

// DefaultGenerator returns a generator setup to emulate the bare OID/UID funcs.
// This is useful for injecting the default behavior.
func DefaultGenerator() *Generator {
	return NewGenerator(
		Crockford32,
		MathSource,
		8,
	)
}

// NewGenerator creates a UID/OID generator based on the given source and the
// given length to be encoded according to the given encoder. There isn't alot
// of error checking. Source should have enough bytes to cover double the
// entropy length (UID reads the entropy length twice).
func NewGenerator(enc EncoderToString, source io.Reader, entropyLen int) *Generator {
	if entropyLen < 1 {
		entropyLen = 1
		log.Println("illegal value; entropy length coerced to 1")
	}
	return &Generator{
		encoder:  enc,
		lastRand: make([]byte, entropyLen),
		source:   source,
	}
}

// OID returns a base 32 encoded string based on timestamp and a random number.
// The `base32( binary( XY ) )` where X is an int64 timestamp (8 bytes) and Y is
// a random number (8 bytes). As opposed to the `func OID()`, the length of the
// random number is configurable.
//
// If (by any chance) OID is called in the same nanosecond, the random number is
// incremented instead of a new one being generated. This makes sure that two
// consecutive IDs generated in the same goroutine are different and sortable.
//
// It is safe for concurrent use as it provides its own locking.
func (gen *Generator) OID() string {
	// lock for lastTime, lastRand, and chars
	gen.mu.Lock()
	defer gen.mu.Unlock()

	now := time.Now().UnixNano()
	// now = int64(2398476238476) // debugging

	// if we have the same time, just inc lastRand, else create a new one
	if now == gen.lastTime {
		gen.lastRand[len(gen.lastRand)-1]++
	} else {
		gen.source.Read(gen.lastRand)
	}

	// remember this for next time
	gen.lastTime = now

	gen.buf.Reset() // clean our buffer before use
	binary.Write(&gen.buf, binary.BigEndian, now)
	binary.Write(&gen.buf, binary.BigEndian, gen.lastRand)
	return gen.encoder.EncodeToString(gen.buf.Bytes())
}

// UID is the same as OID accept that the 8 byte timestamp is replaced with an 8
// byte random number. These IDs are not sortable. As opposed to the `func
// OID()`, the length of the random number is configurable. keep in mind that
// whatever entropy length is used, this value will be double as the `[]byte` is
// used internally twice.
func (gen *Generator) UID() string {
	// lock for lastTime, lastRand, and chars
	gen.mu.Lock()
	defer gen.mu.Unlock()

	gen.buf.Reset() // clean our buffer before use
	gen.source.Read(gen.lastRand)
	binary.Write(&gen.buf, binary.BigEndian, gen.lastRand)
	gen.source.Read(gen.lastRand)
	binary.Write(&gen.buf, binary.BigEndian, gen.lastRand)
	return gen.encoder.EncodeToString(gen.buf.Bytes())
}
