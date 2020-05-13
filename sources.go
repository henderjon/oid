package uid

import (
	crand "crypto/rand"
	"io"
	mrand "math/rand"
	"time"
)

var (
	// MathSource uses math/rand
	MathSource io.Reader = mathSource{}
	// CryptoSource uses crypto/rand
	CryptoSource io.Reader = cryptoSource{}
)

func init() {
	mrand.Seed(time.Now().UnixNano())
}

type mathSource struct{}

func (s mathSource) Read(p []byte) (n int, err error) {
	return mrand.Read(p)
}

type cryptoSource struct{}

func (s cryptoSource) Read(p []byte) (n int, err error) {
	return crand.Read(p)
}
