package caller

import (
	"context"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Caller ...
type Caller struct {
	cli      *ethclient.Client
	cABI     abi.ABI
	cAddress common.Address
}

// NewContractCaller ...
func NewContractCaller(contractABI, nodeEndpoint string, contractAddress common.Address) (*Caller, error) {
	cABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return nil, err
	}
	cli, err := ethclient.Dial(nodeEndpoint)
	if err != nil {
		return nil, err
	}
	return &Caller{
		cli:      cli,
		cABI:     cABI,
		cAddress: contractAddress,
	}, nil
}

// ensureContext is a helper method to ensure a context is not nil, even if the
// user specified it as such.
func ensureContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.TODO()
	}
	return ctx
}

// Call ...
func (c *Caller) Call(opts *bind.CallOpts, funcName string, result interface{}, params ...interface{}) error {
	input, err := c.cABI.Pack(funcName, params...)
	if err != nil {
		return err
	}
	msg := ethereum.CallMsg{To: &c.cAddress, Data: input}
	ctx := ensureContext(opts.Context)
	outByte, err := c.cli.CallContract(ctx, msg, opts.BlockNumber)
	if err != nil {
		return err
	}
	return c.cABI.UnpackIntoInterface(&result, funcName, outByte)
}
