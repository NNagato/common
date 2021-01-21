package blockchain

import (
	"github.com/ethereum/go-ethereum/common"
)

// IsZeroAddress ...
func IsZeroAddress(addr common.Address) bool {
	return addr == common.HexToAddress("0x0000000000000000000000000000000000000000")
}

// IsZeroTxHash ...
func IsZeroTxHash(tx common.Hash) bool {
	return tx == common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")
}
