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
		Crockford32Encoder,
		MathSource,
		8,
	)
}

// NewGenerator creates a UID/OID generator based on the given source and the
// given length to be encoded according to the given encoder. There isn't alot
// of error checking. Source should have enough bytes to cover double the
// entropy length (UID reads the entropy length twice). Entropy length must be
// greater than 0.
//
//  OID                  UID
//  +--------+--------+  +--------+--------+
//  |   TS   |   Ent  |  |  Ent   |   Ent  |
//  +--------+--------+  +--------+--------+
//
// TS is the binary encoding of an int64 (8 byte) Unix Timestamp in Nanoseconds
// Ent is the binary encoding of a >=1 byte random number.
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

// OID is the injectable version of `OID` with a configurable number of random
// bytes.
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

// UID is the injectable version of `UID` with a configurable number of random
// bytes. Be mindfull that whatever entropy length is used, the length of UID
// will be double as the random []byte is used internally twice.
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
