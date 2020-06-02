package oid

var (
	defaultGenerator = DefaultGenerator()
)

// OID generates an ordered ID
func OID() string {
	return defaultGenerator.OID()
}

// UID generates an unordered ID
func UID() string {
	return defaultGenerator.UID()
}
