package factory

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const FACTORY_ADDRESS = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"

func getTradingPair(ctx context.Context, client *ethclient.Client, tokenA, tokenB common.Address) common.Address {
	caller, err := NewFactoryCaller(common.HexToAddress(FACTORY_ADDRESS), client)
	if err != nil {
		log.Fatal(err)
	}
	callOpts := &bind.CallOpts{
		Context: ctx,
		Pending: false,
	}
	pairAddress, err := caller.GetPair(callOpts, tokenA, tokenB)
	if err != nil {
		log.Fatal(err)
	}
	return pairAddress
}
