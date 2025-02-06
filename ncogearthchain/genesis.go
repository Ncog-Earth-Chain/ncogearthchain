package ncogearthchain

import (
	"math/big"

	"github.com/Ncog-Earth-Chain/forest-base/hash"
	"github.com/Ncog-Earth-Chain/forest-base/inter/idx"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Ncog-Earth-Chain/ncogearthchain/inter"
	"github.com/Ncog-Earth-Chain/ncogearthchain/ncogearthchain/genesis"
	"github.com/Ncog-Earth-Chain/ncogearthchain/ncogearthchain/genesis/gpos"
)

type Genesis struct {
	Accounts    genesis.Accounts
	Storage     genesis.Storage
	Delegations genesis.Delegations
	Blocks      genesis.Blocks
	RawEvmItems genesis.RawEvmItems
	Validators  gpos.Validators

	FirstEpoch    idx.Epoch
	PrevEpochTime inter.Timestamp
	Time          inter.Timestamp
	ExtraData     []byte

	TotalSupply *big.Int

	DriverOwner common.Address

	Rules Rules

	Hash func() hash.Hash
}
