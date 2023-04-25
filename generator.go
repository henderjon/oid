package oid

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"sync"
	"time"
)

type GeneratorInterface interface {
	OID() string // generate a random, orderable ID
	UID() string // generate a random, Un-orderable ID
	SID() string // generate a shorter random, Un-orderable ID
}

// Generator is an injectable way of creating UIDs and OIDs.
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

// NewGenerator creates a custom UID/OID generator that reads `len` number of
// bytes from `src` and uses `enc` to encode it as a string. `src` should have
// enough bytes to cover double the `len` as UID reads `len` twice). `len`
// should be greater than 0.
func NewGenerator(enc EncoderToString, src io.Reader, len int) *Generator {
	if len < 1 {
		len = 1
		log.Println("illegal value; entropy length coerced to 1")
	}
	return &Generator{
		encoder:  enc,
		lastRand: make([]byte, len),
		source:   src,
	}
}

// OID is a configurable and injectable version of `OID()`.
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

// UID is a configurable and injectable version of `OID()`. Be mindful that
// whatever `len` was used, the length of the resulting UID will be double as
// `len` is used internally twice.
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

// SID is a configurable and injectable version of `UID()` but does not double
// the number of bytes used. This simply encodes N number of bytes.
func (gen *Generator) SID() string {
	// lock for lastTime, lastRand, and chars
	gen.mu.Lock()
	defer gen.mu.Unlock()

	gen.buf.Reset() // clean our buffer before use
	gen.source.Read(gen.lastRand)
	binary.Write(&gen.buf, binary.BigEndian, gen.lastRand)
	return gen.encoder.EncodeToString(gen.buf.Bytes())
}
