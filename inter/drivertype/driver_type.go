package drivertype

import (
	"math/big"

	"github.com/Ncog-Earth-Chain/forest-base/inter/idx"

	"github.com/Ncog-Earth-Chain/ncogearthchain/inter/validatorpk"
)

var (
	// DoublesignBit is set if validator has a confirmed pair of fork events
	DoublesignBit = uint64(1 << 7)
	OkStatus      = uint64(0)
)

// Validator is the node-side representation of Driver validator
type Validator struct {
	Weight *big.Int
	PubKey validatorpk.PubKey
}

// ValidatorAndID is pair Validator + ValidatorID
type ValidatorAndID struct {
	ValidatorID idx.ValidatorID
	Validator   Validator
}
