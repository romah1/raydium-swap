package raydium

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
)

type ParsedPool struct {
	ID solana.PublicKey `json:"id"`
	BaseMint solana.PublicKey `json:"baseMint"`
	QuoteMint solana.PublicKey `json:"quoteMint"`
	LpMint solana.PublicKey `json:"lpMint"`
	BaseDecimals int `json:"baseDecimals"`
	QuoteDecimals int `json:"quoteDecimals"`
	LpDecimals int `json:"lpDecimals"`
	Version int `json:"version"`
	ProgramId solana.PublicKey `json:"programId"`
	Authority solana.PublicKey `json:"authority"`
	OpenOrders solana.PublicKey `json:"openOrders"`
	TargetOrders solana.PublicKey `json:"targetOrders"`
	BaseVault solana.PublicKey `json:"baseVault"`
	QuoteVault solana.PublicKey `json:"quoteVault"`
	WithdrawQueue solana.PublicKey `json:"withdrawQueue"`
	LpVault solana.PublicKey `json:"lpVault"`
	MarketVersion int `json:"marketVersion"`
	MarketProgramId solana.PublicKey `json:"marketProgramId"`
	MarketId solana.PublicKey `json:"marketId"`
	MarketAuthority solana.PublicKey `json:"marketAuthority"`
	MarketBaseVault solana.PublicKey `json:"marketBaseVault"`
	MarketQuoteVault solana.PublicKey `json:"marketQuoteVault"`
	MarketBids solana.PublicKey `json:"marketBids"`
	MarketAsks solana.PublicKey `json:"marketAsks"`
	MarketEventQueue solana.PublicKey `json:"marketEventQueue"`
	LookupTableAccount solana.PublicKey `json:"lookupTableAccount"`
}

func RetrievePoolsFromPoolInfo(poolInfo *PoolInfo) ([]ParsedPool, error){
	pools := append(poolInfo.Official, poolInfo.UnOfficial...)
	var result []ParsedPool
	for _, pool := range pools {
		parsedPool, err := ParsePool(&pool)
		if err != nil {
			return nil, err
		}
		result = append(result, *parsedPool)
	}
	return result, nil
}

func ParsePool(pool *Pool) (*ParsedPool, error) {
	id, err := solana.PublicKeyFromBase58(pool.ID)
	if err != nil {
		return nil, err
	}
	baseMint, err := solana.PublicKeyFromBase58(pool.BaseMint)
	if err != nil {
		return nil, err
	}
	quoteMint, err := solana.PublicKeyFromBase58(pool.QuoteMint)
	if err != nil {
		return nil, err
	}
	lpMint, err := solana.PublicKeyFromBase58(pool.LpMint)
	if err != nil {
		return nil, err
	}
	programId, err := solana.PublicKeyFromBase58(pool.ProgramId)
	if err != nil {
		return nil, err
	}
	authority, err := solana.PublicKeyFromBase58(pool.Authority)
	if err != nil {
		return nil, err
	}
	openOrders, err := solana.PublicKeyFromBase58(pool.OpenOrders)
	if err != nil {
		return nil, err
	}
	targetOrders, err := solana.PublicKeyFromBase58(pool.TargetOrders)
	if err != nil {
		return nil, err
	}
	baseVault, err := solana.PublicKeyFromBase58(pool.BaseVault)
	if err != nil {
		return nil, err
	}
	quoteVault, err := solana.PublicKeyFromBase58(pool.QuoteVault)
	if err != nil {
		return nil, err
	}
	withdrawQueue, err := solana.PublicKeyFromBase58(pool.WithdrawQueue)
	if err != nil {
		return nil, err
	}
	lpVault, err := solana.PublicKeyFromBase58(pool.LpVault)
	if err != nil {
		return nil, err
	}
	marketProgramId, err := solana.PublicKeyFromBase58(pool.MarketProgramId)
	if err != nil {
		return nil, err
	}
	marketId, err := solana.PublicKeyFromBase58(pool.MarketId)
	if err != nil {
		return nil, err
	}
	marketAuthority, err := solana.PublicKeyFromBase58(pool.MarketAuthority)
	if err != nil {
		return nil, err
	}
	marketBaseVault, err := solana.PublicKeyFromBase58(pool.MarketBaseVault)
	if err != nil {
		return nil, err
	}
	marketQuoteVault, err := solana.PublicKeyFromBase58(pool.MarketQuoteVault)
	if err != nil {
		return nil, err
	}
	marketBids, err := solana.PublicKeyFromBase58(pool.MarketBids)
	if err != nil {
		return nil, err
	}
	marketAsks, err := solana.PublicKeyFromBase58(pool.MarketAsks)
	if err != nil {
		return nil, err
	}
	marketEventQueue, err := solana.PublicKeyFromBase58(pool.MarketEventQueue)
	if err != nil {
		return nil, err
	}
	lookupTableAccount, err := solana.PublicKeyFromBase58(pool.LookupTableAccount)
	if err != nil {
		return nil, err
	}
	
	return &ParsedPool{
		ID: id,
		BaseMint: baseMint,
		QuoteMint: quoteMint,
		LpMint: lpMint,
		BaseDecimals: pool.BaseDecimals,
		QuoteDecimals: pool.QuoteDecimals,
		LpDecimals: pool.LpDecimals,
		Version: pool.Version,
		ProgramId: programId,
		Authority: authority,
		OpenOrders: openOrders,
		TargetOrders: targetOrders,
		BaseVault: baseVault,
		QuoteVault: quoteVault,
		WithdrawQueue: withdrawQueue,
		LpVault: lpVault,
		MarketVersion: pool.MarketVersion,
		MarketProgramId: marketProgramId,
		MarketId: marketId,
		MarketAuthority: marketAuthority,
		MarketBaseVault: marketBaseVault,
		MarketQuoteVault: marketQuoteVault,
		MarketBids: marketBids,
		MarketAsks: marketAsks,
		MarketEventQueue: marketEventQueue,
		LookupTableAccount: lookupTableAccount,
	}, nil
}

func FindPoolInfoByID(id string) (*Pool, error) {
	poolsInfo, err := GetMainnetPoolsInfo()
	if err != nil {
		return nil, err
	}
	pools := append(poolsInfo.Official, poolsInfo.UnOfficial...)
	for _, pool := range pools {
		if pool.ID == id {
			return &pool, nil
		}
	}
	return nil, fmt.Errorf("pool was not found by id [%s]", id)
}
