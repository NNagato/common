package blockchain

import (
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

// TransactOptsFromPK ...
func TransactOptsFromPK(pk string, chainID *big.Int) (*bind.TransactOpts, error) {
	ePK, err := crypto.HexToECDSA(pk)
	if err != nil {
		return nil, err
	}
	return bind.NewKeyedTransactorWithChainID(ePK, chainID)
}

// TransactOptsFromKeystore ...
func TransactOptsFromKeystore(keyStorePath, passphrase string, chainID *big.Int) (*bind.TransactOpts, error) {
	keyStore, err := os.Open(keyStorePath)
	if err != nil {
		return nil, err
	}
	return bind.NewTransactorWithChainID(keyStore, passphrase, chainID)
}
