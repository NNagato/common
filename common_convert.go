package common

import "github.com/shopspring/decimal"

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
