package uid

import (
	crand "crypto/rand"
	"io"
	mrand "math/rand"
	"time"
)

var (
	// MathSource uses math/rand
	MathSource io.Reader = readerFunc(mrand.Read)
	// CryptoSource uses crypto/rand
	CryptoSource io.Reader = readerFunc(crand.Read)
)

func init() {
	mrand.Seed(time.Now().UnixNano())
}

type readerFunc func(p []byte) (n int, err error)

func (s readerFunc) Read(p []byte) (n int, err error) {
	return s(p)
}
