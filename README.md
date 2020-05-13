# oid/uid

This package provides the ability to generate Un/Sortable Identifiers. These
are also most likely very unique. At the very least, they are unique
enough for my purposes, and potentially yours as well.

For my purposes simplicity > speed.

OIDs are sortable, UIDs are not

OID returns a base 32 encoded string based on timestamp and a random
number. The base32( binary( XY ) ) where X is the timestamp ([]byte len(8)) and Y
is the random number ([8]byte).

If (by any chance) OID is called in the same nanosecond, the random number
is incremented instead of a new one being generated. This makes sure that two
consecutive IDs generated in the same goroutine are different and sortable.

UID is the same as OID accept that the [8]byte timestamp is replaced with
[8]byte random data. These IDs are not sortable.

The use of math/rand means these are not "cryptographically secure".

If there is a desire for greater flexibility you can use NewGenerator which allows you to customize
the final encoding (base32, base64, hex, etc. (see encoders.go)) and the entropy
source and length (math/rand, crypto/rand, etc).

Based *very* heavily on:
- [chilts/sid](https://github.com/chilts/sid)
- [oklog/ulid](https://github.com/oklog/ulid)
- https://blog.kowalczyk.info/article/JyRZ/generating-good-unique-ids-in-go.html
