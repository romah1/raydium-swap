package main

import (
	"context"
	"raydium-swap/raydium"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
)

func Swap() error {
	pools, err := raydium.GetMainnetPoolsInfo()
	if err != nil {
		return err
	}
	spew.Dump(pools)
	return nil
}

func main() {
	swap := raydium.NewRaydiumSwap(nil, solana.PrivateKey{})
	ctx := context.Background()
	sig, err := swap.EasySwap(
		ctx,
		"", // targetPool
		0, // amount
		"", // fromToken
		solana.PublicKey{}, // fromAccount
		"", // toToken
		solana.PublicKey{}, // toAccount
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(sig)
}