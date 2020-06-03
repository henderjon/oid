package main

import (
	"encoding/base32"
	"encoding/base64"
	"fmt"

	"github.com/henderjon/oid"
)

func main() {
	fmt.Println("OID\t", oid.OID())
	fmt.Println("UID\t", oid.UID())

	hex := oid.NewGenerator(oid.HexEncoder, oid.MathSource, 8)
	fmt.Println("hex OID\t", hex.OID())
	fmt.Println("hex UID\t", hex.UID())

	b32 := oid.NewGenerator(base32.StdEncoding, oid.MathSource, 5)
	fmt.Println("base32 OID\t", b32.OID())
	fmt.Println("base32 UID\t", b32.UID())

	b64 := oid.NewGenerator(base64.StdEncoding.WithPadding(base64.NoPadding), oid.MathSource, 6)
	fmt.Println("base64 OID\t", b64.OID())
	fmt.Println("base64 UID\t", b64.UID())
}
