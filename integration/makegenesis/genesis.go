package makegenesis

import (
	"math/big"
	"time"

	"github.com/Ncog-Earth-Chain/forest-base/hash"
	"github.com/Ncog-Earth-Chain/forest-base/inter/idx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/cryptod"

	"github.com/Ncog-Earth-Chain/ncogearthchain/inter"
	"github.com/Ncog-Earth-Chain/ncogearthchain/inter/validatorpk"
	"github.com/Ncog-Earth-Chain/ncogearthchain/ncogearthchain"
	"github.com/Ncog-Earth-Chain/ncogearthchain/ncogearthchain/genesis"
	"github.com/Ncog-Earth-Chain/ncogearthchain/ncogearthchain/genesis/driver"
	"github.com/Ncog-Earth-Chain/ncogearthchain/ncogearthchain/genesis/driverauth"
	"github.com/Ncog-Earth-Chain/ncogearthchain/ncogearthchain/genesis/evmwriter"
	"github.com/Ncog-Earth-Chain/ncogearthchain/ncogearthchain/genesis/gpos"
	"github.com/Ncog-Earth-Chain/ncogearthchain/ncogearthchain/genesis/netinit"
	"github.com/Ncog-Earth-Chain/ncogearthchain/ncogearthchain/genesis/sfc"
	"github.com/Ncog-Earth-Chain/ncogearthchain/ncogearthchain/genesisstore"
)

var (
	//FakeGenesisTime = inter.Timestamp(1608600000 * time.Second)
	FakeGenesisTime = inter.Timestamp(1642595720 * time.Second)
)

// FakeKey gets n-th fake private key.
func FakeKey(n int) *cryptod.PrivateKey {

	key, err := cryptod.GenerateMLDsa87Key()
	if err != nil {
		panic(err)
	}

	return key
}

func FakeGenesisStore(num int, balance, stake *big.Int) *genesisstore.Store {
	genStore := genesisstore.NewMemStore()
	genStore.SetRules(ncogearthchain.FakeNetRules())

	validators := GetFakeValidators(num)

	totalSupply := new(big.Int)
	for _, val := range validators {
		genStore.SetEvmAccount(val.Address, genesis.Account{
			Code:    []byte{},
			Balance: balance,
			Nonce:   0,
		})
		genStore.SetDelegation(val.Address, val.ID, genesis.Delegation{
			Stake:              stake,
			Rewards:            new(big.Int),
			LockedStake:        new(big.Int),
			LockupFromEpoch:    0,
			LockupEndTime:      0,
			LockupDuration:     0,
			EarlyUnlockPenalty: new(big.Int),
		})
		totalSupply.Add(totalSupply, balance)
	}

	var owner common.Address
	if num != 0 {
		owner = validators[0].Address
	}

	genStore.SetMetadata(genesisstore.Metadata{
		Validators:    validators,
		FirstEpoch:    2,
		Time:          FakeGenesisTime,
		PrevEpochTime: FakeGenesisTime - inter.Timestamp(time.Hour),
		ExtraData:     []byte("fake"),
		DriverOwner:   owner,
		TotalSupply:   totalSupply,
	})
	genStore.SetBlock(0, genesis.Block{
		Time:        FakeGenesisTime - inter.Timestamp(time.Minute),
		Atropos:     hash.Event{},
		Txs:         types.Transactions{},
		InternalTxs: types.Transactions{},
		Root:        hash.Hash{},
		Receipts:    []*types.ReceiptForStorage{},
	})
	// pre deploy NetworkInitializer
	genStore.SetEvmAccount(netinit.ContractAddress, genesis.Account{
		Code:    netinit.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy NodeDriver
	genStore.SetEvmAccount(driver.ContractAddress, genesis.Account{
		Code:    driver.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy NodeDriverAuth
	genStore.SetEvmAccount(driverauth.ContractAddress, genesis.Account{
		Code:    driverauth.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy SFC
	genStore.SetEvmAccount(sfc.ContractAddress, genesis.Account{
		Code:    sfc.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// set non-zero code for pre-compiled contracts
	genStore.SetEvmAccount(evmwriter.ContractAddress, genesis.Account{
		Code:    []byte{0},
		Balance: new(big.Int),
		Nonce:   0,
	})

	return genStore
}

// GetFakeValidators generates fake validators for testing purposes.
func GetFakeValidators(num int) gpos.Validators {
	validators := make(gpos.Validators, 0, num)

	for i := 1; i <= num; i++ {
		key := FakeKey(i)
		addr := cryptod.PubkeyToAddress(*key.Public().(*cryptod.PublicKey))
		pubkeyraw := cryptod.FromMLDsa87Pub(key.Public().(*cryptod.PublicKey))
		validatorID := idx.ValidatorID(i)
		validators = append(validators, gpos.Validator{
			ID:      validatorID,
			Address: addr,
			PubKey: validatorpk.PubKey{
				Raw:  pubkeyraw,
				Type: validatorpk.Types.MLDsa87,
			},
			CreationTime:     FakeGenesisTime,
			CreationEpoch:    0,
			DeactivatedTime:  0,
			DeactivatedEpoch: 0,
			Status:           0,
		})
	}

	return validators
}
