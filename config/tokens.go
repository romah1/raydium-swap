package config

import (
	_ "embed"
	"math"
)

const (
	TokenAccountSize = 165
	NativeSOL        = "11111111111111111111111111111111"
	WrappedSOL       = "So11111111111111111111111111111111111111112"
)

type TokenInfo struct {
	Symbol   string
	Decimals uint8
}

func (s *TokenInfo) Pow() float64 {
	return math.Pow10(int(s.Decimals))
}

func (s *TokenInfo) ToFloat(v uint64) float64 {
	return float64(v) / math.Pow10(int(s.Decimals))
}

func (s *TokenInfo) FromFloat(v float64) uint64 {
	return uint64(v * math.Pow10(int(s.Decimals)))
}