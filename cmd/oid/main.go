package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"text/template"

	"github.com/henderjon/oid"
)

// Tmpl is a basic man page[-ish] looking template
const Tmpl = `
{{define "manual"}}
NAME
  {{.Bin}} - a simple utility for generating un/ordered IDs.

SYNOPSIS
  {{.Bin}} [-len=8] [-hex] [-secure] [-uid]

DESCRIPTION
  By default, it is the base 32 encoding of a binary encoded string comprising an 8 byte nanosecond precision unix timestamp and an 8 byte random number, in that order. The timestamp prefix allows these IDs to be ordered.

  Using '-uid' makes it a base 32 encoding of a binary encoded string comprising two 8 byte random numbers. These are not ordered

  The Crockford base32 dictionary used (https://www.crockford.com/base32.html).

OPTIONS
{{.Options}}
VERSION
  version:  {{.Version}}
  compiled: {{.CompiledBy}}
  built:    {{.BuildTimestamp}}

{{end}}
`

var (
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
	if unordered {
		fmt.Println(gen.UID())
	} else {
		fmt.Println(gen.OID())
	}

}

// Info represents the infomation used in the default Tmpl string
type Info struct {
	Tmpl           string
	Bin            string
	Version        string
	CompiledBy     string
	BuildTimestamp string
	Options        string
}

// Usage wraps a set of `Info` and creates a flag.Usage func
func Usage(info Info) func() {
	if len(info.Tmpl) == 0 {
		info.Tmpl = Tmpl
	}

	t := template.Must(template.New("manual").Parse(info.Tmpl))

	return func() {
		var def bytes.Buffer
		flag.CommandLine.SetOutput(&def)
		flag.PrintDefaults()

		info.Options = def.String()
		t.Execute(os.Stdout, info)
	}
}
