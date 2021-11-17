package main

import (
	"bytes"
	"flag"
	"os"
	"text/template"
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
SEE ALSO
  https://github.com/chilts/sid, https://github.com/oklog/ulid, https://blog.kowalczyk.info/article/JyRZ/generating-good-unique-ids-in-go.html

VERSION
  version:  {{.Version}}
  compiled: {{.CompiledBy}}
  built:    {{.BuildTimestamp}}

{{end}}
`

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
