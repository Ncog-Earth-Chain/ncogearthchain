package netinitcall

import (
	"math/big"
	"strings"

	"github.com/Ncog-Earth-Chain/forest-base/inter/idx"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Ncog-Earth-Chain/ncogearthchain/utils"
)

const ContractABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sealedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sfc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_auth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_driver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_evmWriter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initializeAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

var (
	sAbi, _ = abi.JSON(strings.NewReader(ContractABI))
)

// Methods

func InitializeAll(sealedEpoch idx.Epoch, totalSupply *big.Int, sfcAddr common.Address, driverAuthAddr common.Address, driverAddr common.Address, evmWriterAddr common.Address, owner common.Address) []byte {
	data, _ := sAbi.Pack("initializeAll", utils.U64toBig(uint64(sealedEpoch)), totalSupply, sfcAddr, driverAuthAddr, driverAddr, evmWriterAddr, owner)
	return data
}
