package common

import (
	"strconv"

	"github.com/shopspring/decimal"
)

// MustStringToFloat convert string to float if got error panic program
func MustStringToFloat(fs string) float64 {
	f, err := decimal.NewFromString(fs)
	if err != nil {
		panic(err)
	}
	realf, _ := f.Float64()
	return realf
}

// StringToFloat convert string to float if got error panic program
func StringToFloat(fs string) (float64, error) {
	f, err := decimal.NewFromString(fs)
	if err != nil {
		return 0, err
	}
	realf, _ := f.Float64()
	return realf, nil
}

// FloatToString convert float to string
func FloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// Uint64ToString convert uint64 to string
func Uint64ToString(u uint64) string {
	return strconv.FormatUint(u, 10)
}

// Int64ToString convert int64 to string
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}
