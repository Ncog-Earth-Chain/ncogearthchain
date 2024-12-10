package sfcapi

import (
	"github.com/Ncog-Earth-Chain/go-ncogearthchain/utils"
	"github.com/ethereum/go-ethereum/common"
)

// Events

//event CreatedValidator(uint256 indexed validatorID, address indexed auth, uint256 createdEpoch, uint256 createdTime);
//event DeactivatedValidator(uint256 indexed validatorID, uint256 deactivatedEpoch, uint256 deactivatedTime);
//event ChangedValidatorStatus(uint256 indexed validatorID, uint256 status);
//event Delegated(address indexed delegator, uint256 indexed toValidatorID, uint256 amount);
//event Undelegated(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount);
//event ClaimedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 rewards);

var (
	// Topics of SFC contract logs
	Topics = struct {
		ClaimedRewards          common.Hash
		RestakedRewards         common.Hash
		ClaimedDelegationReward common.Hash
		ClaimedValidatorReward  common.Hash
		CreatedValidator        common.Hash
		DeactivatedValidator    common.Hash
		ChangedValidatorStatus  common.Hash
		Delegated               common.Hash
		Undelegated             common.Hash
	}{
		ClaimedRewards:          utils.Keccak512Hash([]byte("ClaimedRewards(address,uint256,uint256,uint256,uint256)")),
		RestakedRewards:         utils.Keccak512Hash([]byte("RestakedRewards(address,uint256,uint256,uint256,uint256)")),
		ClaimedDelegationReward: utils.Keccak512Hash([]byte("ClaimedDelegationReward(address,uint256,uint256,uint256,uint256)")),
		ClaimedValidatorReward:  utils.Keccak512Hash([]byte("ClaimedValidatorReward(uint256,uint256,uint256,uint256)")),
		CreatedValidator:        utils.Keccak512Hash([]byte("CreatedValidator(uint256,address,uint256,uint256)")),
		DeactivatedValidator:    utils.Keccak512Hash([]byte("DeactivatedValidator(uint256,uint256,uint256)")),
		ChangedValidatorStatus:  utils.Keccak512Hash([]byte("ChangedValidatorStatus(uint256,uint256)")),
		Delegated:               utils.Keccak512Hash([]byte("Delegated(address,uint256,uint256)")),
		Undelegated:             utils.Keccak512Hash([]byte("Undelegated(address,uint256,uint256,uint256)")),
	}
)
