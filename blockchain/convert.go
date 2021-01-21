package blockchain

import (
	"math/big"

	bigc "github.com/NNagato/common/big"
)

func decimalsBig(decimals int) *big.Int { // return 10^decimals
	return big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
}

// TokenWeiToAmount ...
func TokenWeiToAmount(tokenWei *big.Int, decimals int) float64 {
	var (
		twFloat          = big.NewFloat(0).SetInt(tokenWei)
		decimalsBigFloat = big.NewFloat(0).SetInt(decimalsBig(decimals))
	)
	amount, _ := bigc.QuoFloat(twFloat, decimalsBigFloat).Float64()
	return amount
}

// AmountToTokenWei ...
func AmountToTokenWei(amount float64, decimals int) *big.Int {
	var (
		amountBig        = big.NewFloat(amount)
		decimalsBigFloat = big.NewFloat(0).SetInt(decimalsBig(decimals))
	)
	twei, _ := bigc.MulFloat(amountBig, decimalsBigFloat).Int(nil)
	return twei
}

// GweiToWei convert gas from gwei to wei
func GweiToWei(gwei float64) *big.Int {
	if gwei == 0 {
		return nil
	}
	// 1 gwei = 10^9 wei
	return AmountToTokenWei(gwei, 9)
}

// WeiToGWei convert gwei to wei
func WeiToGWei(wei *big.Int) float64 {
	if wei.Cmp(big.NewInt(0)) == 0 {
		return 0
	}
	// 1 gwei = 10^9 wei
	return TokenWeiToAmount(wei, 9)
}

// GasToEther return gas to eth
func GasToEther(gasUsed float64) float64 {
	const (
		gwei float64 = 0.000000001 // eth
	)
	return gasUsed * gwei
}

// EtherToGas return gwei correspond to given ether
func EtherToGas(ethAmount float64) float64 {
	const (
		gwei float64 = 0.000000001 // eth
	)
	return ethAmount / gwei
}
