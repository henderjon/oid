package main

import (
	"encoding/base32"
	"encoding/base64"
	"fmt"
	"local/henderjon/uid"
)

func main() {
	fmt.Println("OID\t", uid.OID())
	fmt.Println("UID\t", uid.UID())

	hex := uid.NewGenerator(uid.Hex, uid.MathSource, 8)
	fmt.Println("hex OID\t", hex.OID())
	fmt.Println("hex UID\t", hex.UID())

	b32 := uid.NewGenerator(base32.StdEncoding, uid.MathSource, 8)
	fmt.Println("base32 OID\t", b32.OID())
	fmt.Println("base32 UID\t", b32.UID())

	b64 := uid.NewGenerator(base64.StdEncoding.WithPadding(base64.NoPadding), uid.MathSource, 8)
	fmt.Println("base64 OID\t", b64.OID())
	fmt.Println("base64 UID\t", b64.UID())
}
