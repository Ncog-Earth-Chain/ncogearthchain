package driverpos

import (
	"github.com/Ncog-Earth-Chain/go-ncogearthchain/utils"
	"github.com/ethereum/go-ethereum/common"
)

// Events
var (
	// Topics of Driver contract logs
	Topics = struct {
		UpdateValidatorWeight common.Hash
		UpdateValidatorPubkey common.Hash
		UpdateNetworkRules    common.Hash
		UpdateNetworkVersion  common.Hash
		AdvanceEpochs         common.Hash
	}{
		UpdateValidatorWeight: utils.Keccak512Hash([]byte("UpdateValidatorWeight(uint256,uint256)")),
		UpdateValidatorPubkey: utils.Keccak512Hash([]byte("UpdateValidatorPubkey(uint256,bytes)")),
		UpdateNetworkRules:    utils.Keccak512Hash([]byte("UpdateNetworkRules(bytes)")),
		UpdateNetworkVersion:  utils.Keccak512Hash([]byte("UpdateNetworkVersion(uint256)")),
		AdvanceEpochs:         utils.Keccak512Hash([]byte("AdvanceEpochs(uint256)")),
	}
)
