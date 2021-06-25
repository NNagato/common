package common

import "time"

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

// PoiterFloat64 ...
func PoiterFloat64(f float64) *float64 {
	return &f
}

func PointerTime(t time.Time) *time.Time {
	return &t
}

// PoiterInt ...
func PoiterInt(i int) *int {
	return &i
}
