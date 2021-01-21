package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
)

// RoundingNumber ...
func RoundingNumber(n float64, precision int) float64 {
	rounding := math.Pow10(precision)
	return math.Floor(n*rounding) / rounding
}

// RoundingNumberUp ...
func RoundingNumberUp(n float64, precision int) float64 {
	rounding := math.Pow10(precision)
	return math.Ceil(n*rounding) / rounding
}

// ReadFile ...
func ReadFile(filePath string, result interface{}) error {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(f, &result)
}

// LinkTxEtherescan return link to transaction on etherscan
func LinkTxEtherescan(tx string) string {
	return fmt.Sprintf("https://etherscan.io/tx/%s", tx)
}

// WithPercent ...
func WithPercent(num, percent float64) float64 {
	return num * (100 + percent) / 100
}

// PercentOf ...
func PercentOf(num, percent float64) float64 {
	return num * percent / 100
}

// PercentChange ...
func PercentChange(numA, numB float64) float64 {
	return numB/numA - 1
}
