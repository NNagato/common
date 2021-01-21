package common

// PoiterString ...
func PoiterString(s string) *string {
	return &s
}

// PoiterUint64 ...
func PoiterUint64(u uint64) *uint64 {
	return &u
}

// PoiterInt64 ...
func PoiterInt64(i int64) *int64 {
	return &i
}
