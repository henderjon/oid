package oid

import (
	"encoding/base32"
	"encoding/hex"
)

// Encoder is an adapter akin to http.HandlerFunc
type Encoder func([]byte) string

// EncoderToString is an interface for encoding bytes to strings
type EncoderToString interface {
	EncodeToString([]byte) string
}

// EncodeToString allows any Encoder type to match the EncoderToString interface
func (e Encoder) EncodeToString(b []byte) string {
	return e(b)
}

var (
	// Crockford32Encoder is a base 32 dictionary from Dougland Crockford
	// (https://www.crockford.com/base32.html)
	Crockford32Encoder = base32.NewEncoding("0123456789abcdefghjkmnpqrstvwxyz").WithPadding(base32.NoPadding)
	// HexEncoder is a typical hex representation of bytes. I'm casting it here
	// to allow for a top level package function to satisfy a proper interface.
	HexEncoder = Encoder(hex.EncodeToString)
)
