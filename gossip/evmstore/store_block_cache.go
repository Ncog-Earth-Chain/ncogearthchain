package evmstore

import (
	"github.com/Ncog-Earth-Chain/forest-base/inter/idx"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Ncog-Earth-Chain/ncogearthchain/evmcore"
)

func (s *Store) GetCachedEvmBlock(n idx.Block) *evmcore.EvmBlock {
	c, ok := s.cache.EvmBlocks.Get(n)
	if !ok {
		return nil
	}

	return c.(*evmcore.EvmBlock)
}

func (s *Store) SetCachedEvmBlock(n idx.Block, b *evmcore.EvmBlock) {
	var empty = common.Hash{}
	if b.EvmHeader.TxHash == empty {
		panic("You have to cache only completed blocks (with txs)")
	}
	s.cache.EvmBlocks.Add(n, b, uint(b.EstimateSize()))
}
