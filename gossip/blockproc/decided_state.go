package blockproc

import (
	"crypto/sha256"
	"math/big"

	"github.com/Ncog-Earth-Chain/forest-base/forest"
	"github.com/Ncog-Earth-Chain/forest-base/hash"
	"github.com/Ncog-Earth-Chain/forest-base/inter/idx"
	"github.com/Ncog-Earth-Chain/forest-base/inter/pos"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/Ncog-Earth-Chain/ncogearthchain/inter"
	"github.com/Ncog-Earth-Chain/ncogearthchain/ncogearthchain"
)

type ValidatorBlockState struct {
	Cheater          bool
	LastEvent        hash.Event
	Uptime           inter.Timestamp
	LastOnlineTime   inter.Timestamp
	LastGasPowerLeft inter.GasPowerLeft
	LastBlock        idx.Block
	DirtyGasRefund   uint64
	Originated       *big.Int
}

type ValidatorEpochState struct {
	GasRefund      uint64
	PrevEpochEvent hash.Event
}

type BlockCtx struct {
	Idx     idx.Block
	Time    inter.Timestamp
	Atropos hash.Event
}

type BlockState struct {
	LastBlock          BlockCtx
	FinalizedStateRoot hash.Hash

	EpochGas      uint64
	EpochCheaters forest.Cheaters

	ValidatorStates       []ValidatorBlockState
	NextValidatorProfiles ValidatorProfiles

	DirtyRules ncogearthchain.Rules

	AdvanceEpochs idx.Epoch
}

func (bs BlockState) Copy() BlockState {
	cp := bs
	cp.EpochCheaters = make(forest.Cheaters, len(bs.EpochCheaters))
	copy(cp.EpochCheaters, bs.EpochCheaters)
	cp.ValidatorStates = make([]ValidatorBlockState, len(bs.ValidatorStates))
	copy(cp.ValidatorStates, bs.ValidatorStates)
	cp.NextValidatorProfiles = make(ValidatorProfiles, len(bs.NextValidatorProfiles))
	for k, v := range bs.NextValidatorProfiles {
		cp.NextValidatorProfiles[k] = v
	}
	cp.DirtyRules = bs.DirtyRules.Copy()
	return cp
}

func (bs *BlockState) GetValidatorState(id idx.ValidatorID, validators *pos.Validators) *ValidatorBlockState {
	validatorIdx := validators.GetIdx(id)
	return &bs.ValidatorStates[validatorIdx]
}

type EpochState struct {
	Epoch          idx.Epoch
	EpochStart     inter.Timestamp
	PrevEpochStart inter.Timestamp

	EpochStateRoot hash.Hash

	Validators        *pos.Validators
	ValidatorStates   []ValidatorEpochState
	ValidatorProfiles ValidatorProfiles

	Rules ncogearthchain.Rules
}

func (es *EpochState) GetValidatorState(id idx.ValidatorID, validators *pos.Validators) *ValidatorEpochState {
	validatorIdx := validators.GetIdx(id)
	return &es.ValidatorStates[validatorIdx]
}

func (es EpochState) Duration() inter.Timestamp {
	return es.EpochStart - es.PrevEpochStart
}

func (es EpochState) Hash() hash.Hash {
	hasher := sha256.New()
	err := rlp.Encode(hasher, &es)
	if err != nil {
		panic("can't hash: " + err.Error())
	}
	return hash.BytesToHash(hasher.Sum(nil))
}

func (es EpochState) Copy() EpochState {
	cp := es
	cp.ValidatorStates = make([]ValidatorEpochState, len(es.ValidatorStates))
	copy(cp.ValidatorStates, es.ValidatorStates)
	cp.ValidatorProfiles = make(ValidatorProfiles, len(es.ValidatorProfiles))
	for k, v := range es.ValidatorProfiles {
		cp.ValidatorProfiles[k] = v
	}
	cp.Rules = es.Rules.Copy()
	return cp
}
