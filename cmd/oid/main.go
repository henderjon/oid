package main

import (
	"flag"
	"fmt"

	"github.com/henderjon/oid"
)

var (
	num       int
	length    int
	unordered bool
	secure    bool
	hex       bool
	help      bool
)

func main() {
	flag.Usage = Usage(Info{
		Tmpl:           Tmpl,
		Bin:            "oid",
		Version:        buildVersion,
		CompiledBy:     compiledBy,
		BuildTimestamp: buildTimestamp,
	})

	flag.IntVar(&num, "num", 1, "the number of u/oids to generate")
	flag.IntVar(&length, "len", 8, "the number of bytes of randomness to use. Keep in mind that UIDs use double this value for the overall length, while OIDs are this value +8 bytes for the timestamp.")
	flag.BoolVar(&unordered, "uid", false, "generate an unordered ID (UID) vs the default ordered ID (OID).")
	flag.BoolVar(&secure, "secure", false, "use a cryptographically secure randomness generator.")
	flag.BoolVar(&hex, "hex", false, "use a standard HEX (a-f0-9) dictionary vs the default base32 (Crockford) dictionary.")
	flag.Parse()

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

	for ; num > 0; num-- {
		if unordered {
			fmt.Println(gen.UID())
		} else {
			fmt.Println(gen.OID())
		}
	}
}
