package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/henderjon/oid"
)

const doc = `
%s is a simple utility for generating un/ordered IDs.

By default, it is the base 32 encoding of a binary encoded string
comprising an 8 byte nanosecond precision unix timestamp and an 8 byte random
number, in that order. The timestamp prefix allows these IDs to be ordered.

Using '-uid' makes it a base 32 encoding of a binary encoded string
comprising two 8 byte random numbers. These are not ordered

version:  %s
compiled: %s
built:    %s

Usage: %s [option [option]...]

Options:
`

var (
	length    int
	unordered bool
	secure    bool
	hex       bool
	help      bool
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(
			os.Stderr,
			doc,
			os.Args[0],
			buildVersion,
			compiledBy,
			buildTimestamp,
			os.Args[0],
		)
		flag.PrintDefaults()
	}

}

func main() {
	flag.IntVar(&length, "len", 8, "The number of bytes of randomness to use. Keep in mind that UIDs use double this value for the overall length, while OIDs are this value +8 Bytes for the timestamp.")
	flag.BoolVar(&unordered, "uid", false, "Generate an 'Unordered' ID (UID) vs the default 'Ordered' ID (OID).")
	flag.BoolVar(&secure, "secure", false, "Use a cryptographically secure randomness generator.")
	flag.BoolVar(&hex, "hex", false, "Use a standard HEX (a-f0-9) dictionary vs the default base32 (Crockford) dictionary.")
	flag.BoolVar(&help, "help", false, "display this message")
	flag.BoolVar(&help, "h", false, "display this message")
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	source := oid.MathSource
	if secure {
		source = oid.CryptoSource
	}

	var dict oid.EncoderToString // assert the interface not the struct pointer
	dict = oid.Crockford32Encoder
	if hex {
		dict = oid.HexEncoder
	}

	gen := oid.NewGenerator(dict, source, length)
	if unordered {
		fmt.Println(gen.UID())
	} else {
		fmt.Println(gen.OID())
	}

}
