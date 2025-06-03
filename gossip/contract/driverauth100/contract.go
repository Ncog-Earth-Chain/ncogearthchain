// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package driverauth100

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"advanceEpochs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"copyCode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"deactivateValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"diff\",\"type\":\"uint256\"}],\"name\":\"incBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"diff\",\"type\":\"uint256\"}],\"name\":\"incNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sfc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_driver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDriverAuth\",\"type\":\"address\"}],\"name\":\"migrateTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"offlineTimes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"offlineBlocks\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"uptimes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"originatedTxsFee\",\"type\":\"uint256[]\"}],\"name\":\"sealEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nextValidatorIDs\",\"type\":\"uint256[]\"}],\"name\":\"sealEpochValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupFromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earlyUnlockPenalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"setGenesisDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_auth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"}],\"name\":\"setGenesisValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"diff\",\"type\":\"bytes\"}],\"name\":\"updateNetworkRules\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateNetworkVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"updateValidatorPubkey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"updateValidatorWeight\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}]"

// ContractBin is the compiled bytecode used for deploying new contracts.
var ContractBin = "0x608060405234801561000f575f80fd5b50611c678061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610163575f3560e01c80638da5cb5b116100c7578063d6a0c7af1161007d578063ebdf104c11610063578063ebdf104c146102d3578063f2fde38b146102e6578063fd1b6ec1146102f9575f80fd5b8063d6a0c7af146102ad578063e08d7e66146102c0575f80fd5b8063a4066fbe116100ad578063a4066fbe14610274578063b9cc6b1c14610287578063c0c53b8b1461029a575f80fd5b80638da5cb5b1461021c5780638f32d59b14610249575f80fd5b80634ddaf8f21161011c57806366e7ea0f1161010257806366e7ea0f146101ee578063715018a61461020157806379bead3814610209575f80fd5b80634ddaf8f2146101c85780634feb92f3146101db575f80fd5b80631e702f831161014c5780631e702f831461018f578063242a6e3f146101a2578063267ab446146101b5575f80fd5b80630aeeca001461016757806318f628d41461017c575b5f80fd5b61017a61017536600461166b565b61030c565b005b61017a61018a3660046116aa565b6103fb565b61017a61019d36600461170a565b61054e565b61017a6101b036600461176f565b610666565b61017a6101c336600461166b565b610759565b61017a6101d63660046117b7565b610817565b61017a6101e93660046117d0565b6108d6565b61017a6101fc366004611850565b6109c9565b61017a610b47565b61017a610217366004611850565b610c1c565b60335460405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b60335473ffffffffffffffffffffffffffffffffffffffff1633146040519015158152602001610240565b61017a61028236600461170a565b610ce2565b61017a610295366004611878565b610da7565b61017a6102a83660046118b7565b610e66565b61017a6102bb3660046118f7565b610fc8565b61017a6102ce366004611969565b61108f565b61017a6102e136600461199c565b611174565b61017a6102f43660046117b7565b61129c565b61017a6103073660046118f7565b61130f565b60335473ffffffffffffffffffffffffffffffffffffffff1633146103785760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b6067546040517f0aeeca000000000000000000000000000000000000000000000000000000000081526004810183905273ffffffffffffffffffffffffffffffffffffffff90911690630aeeca00906024015b5f604051808303815f87803b1580156103e2575f80fd5b505af11580156103f4573d5f803e3d5ffd5b5050505050565b60675473ffffffffffffffffffffffffffffffffffffffff1633146104885760405162461bcd60e51b815260206004820152602560248201527f63616c6c6572206973206e6f7420746865204e6f646544726976657220636f6e60448201527f7472616374000000000000000000000000000000000000000000000000000000606482015260840161036f565b6066546040517f18f628d400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8b81166004830152602482018b9052604482018a9052606482018990526084820188905260a4820187905260c4820186905260e482018590526101048201849052909116906318f628d490610124015b5f604051808303815f87803b15801561052d575f80fd5b505af115801561053f573d5f803e3d5ffd5b50505050505050505050505050565b60675473ffffffffffffffffffffffffffffffffffffffff1633146105db5760405162461bcd60e51b815260206004820152602560248201527f63616c6c6572206973206e6f7420746865204e6f646544726976657220636f6e60448201527f7472616374000000000000000000000000000000000000000000000000000000606482015260840161036f565b6066546040517f1e702f83000000000000000000000000000000000000000000000000000000008152600481018490526024810183905273ffffffffffffffffffffffffffffffffffffffff90911690631e702f83906044015b5f604051808303815f87803b15801561064c575f80fd5b505af115801561065e573d5f803e3d5ffd5b505050505050565b60665473ffffffffffffffffffffffffffffffffffffffff1633146106cd5760405162461bcd60e51b815260206004820152601e60248201527f63616c6c6572206973206e6f74207468652053464320636f6e74726163740000604482015260640161036f565b6067546040517f242a6e3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063242a6e3f9061072790869086908690600401611a9e565b5f604051808303815f87803b15801561073e575f80fd5b505af1158015610750573d5f803e3d5ffd5b50505050505050565b60335473ffffffffffffffffffffffffffffffffffffffff1633146107c05760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161036f565b6067546040517f267ab4460000000000000000000000000000000000000000000000000000000081526004810183905273ffffffffffffffffffffffffffffffffffffffff9091169063267ab446906024016103cb565b60335473ffffffffffffffffffffffffffffffffffffffff16331461087e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161036f565b6067546040517fda7fc24f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301529091169063da7fc24f906024016103cb565b60675473ffffffffffffffffffffffffffffffffffffffff1633146109635760405162461bcd60e51b815260206004820152602560248201527f63616c6c6572206973206e6f7420746865204e6f646544726976657220636f6e60448201527f7472616374000000000000000000000000000000000000000000000000000000606482015260840161036f565b6066546040517f4feb92f300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690634feb92f390610516908c908c908c908c908c908c908c908c908c90600401611ac0565b60665473ffffffffffffffffffffffffffffffffffffffff163314610a305760405162461bcd60e51b815260206004820152601e60248201527f63616c6c6572206973206e6f74207468652053464320636f6e74726163740000604482015260640161036f565b60665473ffffffffffffffffffffffffffffffffffffffff838116911614610ac05760405162461bcd60e51b815260206004820152602160248201527f726563697069656e74206973206e6f74207468652053464320636f6e7472616360448201527f7400000000000000000000000000000000000000000000000000000000000000606482015260840161036f565b60675473ffffffffffffffffffffffffffffffffffffffff9081169063e30443bc908490610af190821631856113d2565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff90921660048301526024820152604401610635565b60335473ffffffffffffffffffffffffffffffffffffffff163314610bae5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161036f565b6033546040515f9173ffffffffffffffffffffffffffffffffffffffff16907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055565b60335473ffffffffffffffffffffffffffffffffffffffff163314610c835760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161036f565b6067546040517f79bead3800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116600483015260248201849052909116906379bead3890604401610635565b60665473ffffffffffffffffffffffffffffffffffffffff163314610d495760405162461bcd60e51b815260206004820152601e60248201527f63616c6c6572206973206e6f74207468652053464320636f6e74726163740000604482015260640161036f565b6067546040517fa4066fbe000000000000000000000000000000000000000000000000000000008152600481018490526024810183905273ffffffffffffffffffffffffffffffffffffffff9091169063a4066fbe90604401610635565b60335473ffffffffffffffffffffffffffffffffffffffff163314610e0e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161036f565b6067546040517fb9cc6b1c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063b9cc6b1c906106359085908590600401611b23565b5f54610100900460ff1680610e7a5750303b155b80610e8757505f5460ff16155b610ef95760405162461bcd60e51b815260206004820152602e60248201527f436f6e747261637420696e7374616e63652068617320616c726561647920626560448201527f656e20696e697469616c697a6564000000000000000000000000000000000000606482015260840161036f565b5f54610100900460ff16158015610f36575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000166101011790555b610f3f826113e6565b6067805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560668054928716929091169190911790558015610fc2575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b50505050565b60335473ffffffffffffffffffffffffffffffffffffffff16331461102f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161036f565b6067546040517fd6a0c7af00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116600483015283811660248301529091169063d6a0c7af90604401610635565b60675473ffffffffffffffffffffffffffffffffffffffff16331461111c5760405162461bcd60e51b815260206004820152602560248201527f63616c6c6572206973206e6f7420746865204e6f646544726976657220636f6e60448201527f7472616374000000000000000000000000000000000000000000000000000000606482015260840161036f565b6066546040517fe08d7e6600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063e08d7e66906106359085908590600401611b87565b60675473ffffffffffffffffffffffffffffffffffffffff1633146112015760405162461bcd60e51b815260206004820152602560248201527f63616c6c6572206973206e6f7420746865204e6f646544726976657220636f6e60448201527f7472616374000000000000000000000000000000000000000000000000000000606482015260840161036f565b6066546040517febdf104c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063ebdf104c90611265908b908b908b908b908b908b908b908b90600401611b9a565b5f604051808303815f87803b15801561127c575f80fd5b505af115801561128e573d5f803e3d5ffd5b505050505050505050505050565b60335473ffffffffffffffffffffffffffffffffffffffff1633146113035760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161036f565b61130c81611555565b50565b60335473ffffffffffffffffffffffffffffffffffffffff1633146113765760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161036f565b813b151580156113865750803b15155b61102f5760405162461bcd60e51b815260206004820152600e60248201527f6e6f74206120636f6e7472616374000000000000000000000000000000000000604482015260640161036f565b5f6113dd8284611bf9565b90505b92915050565b5f54610100900460ff16806113fa5750303b155b8061140757505f5460ff16155b6114795760405162461bcd60e51b815260206004820152602e60248201527f436f6e747261637420696e7374616e63652068617320616c726561647920626560448201527f656e20696e697469616c697a6564000000000000000000000000000000000000606482015260840161036f565b5f54610100900460ff161580156114b6575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000166101011790555b603380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84169081179091556040515f907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a38015611551575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b5050565b73ffffffffffffffffffffffffffffffffffffffff81166115de5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161036f565b60335460405173ffffffffffffffffffffffffffffffffffffffff8084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3603380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b5f6020828403121561167b575f80fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146116a5575f80fd5b919050565b5f805f805f805f805f6101208a8c0312156116c3575f80fd5b6116cc8a611682565b9b60208b01359b5060408b01359a60608101359a506080810135995060a0810135985060c0810135975060e081013596506101000135945092505050565b5f806040838503121561171b575f80fd5b50508035926020909101359150565b5f8083601f84011261173a575f80fd5b50813567ffffffffffffffff811115611751575f80fd5b602083019150836020828501011115611768575f80fd5b9250929050565b5f805f60408486031215611781575f80fd5b83359250602084013567ffffffffffffffff81111561179e575f80fd5b6117aa8682870161172a565b9497909650939450505050565b5f602082840312156117c7575f80fd5b6113dd82611682565b5f805f805f805f805f6101008a8c0312156117e9575f80fd5b6117f28a611682565b985060208a0135975060408a013567ffffffffffffffff811115611814575f80fd5b6118208c828d0161172a565b9a9d999c509a6060810135996080820135995060a0820135985060c0820135975060e09091013595509350505050565b5f8060408385031215611861575f80fd5b61186a83611682565b946020939093013593505050565b5f8060208385031215611889575f80fd5b823567ffffffffffffffff81111561189f575f80fd5b6118ab8582860161172a565b90969095509350505050565b5f805f606084860312156118c9575f80fd5b6118d284611682565b92506118e060208501611682565b91506118ee60408501611682565b90509250925092565b5f8060408385031215611908575f80fd5b61191183611682565b915061191f60208401611682565b90509250929050565b5f8083601f840112611938575f80fd5b50813567ffffffffffffffff81111561194f575f80fd5b6020830191508360208260051b8501011115611768575f80fd5b5f806020838503121561197a575f80fd5b823567ffffffffffffffff811115611990575f80fd5b6118ab85828601611928565b5f805f805f805f806080898b0312156119b3575f80fd5b883567ffffffffffffffff808211156119ca575f80fd5b6119d68c838d01611928565b909a50985060208b01359150808211156119ee575f80fd5b6119fa8c838d01611928565b909850965060408b0135915080821115611a12575f80fd5b611a1e8c838d01611928565b909650945060608b0135915080821115611a36575f80fd5b50611a438b828c01611928565b999c989b5096995094979396929594505050565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b838152604060208201525f611ab7604083018486611a57565b95945050505050565b5f61010073ffffffffffffffffffffffffffffffffffffffff8c1683528a6020840152806040840152611af68184018a8c611a57565b60608401989098525050608081019490945260a084019290925260c083015260e090910152949350505050565b602081525f611b36602083018486611a57565b949350505050565b8183525f7f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831115611b6e575f80fd5b8260051b80836020870137939093016020019392505050565b602081525f611b36602083018486611b3e565b608081525f611bad608083018a8c611b3e565b8281036020840152611bc081898b611b3e565b90508281036040840152611bd5818789611b3e565b90508281036060840152611bea818587611b3e565b9b9a5050505050505050505050565b808201808211156113e0577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffdfea2646970667358221220b080ab110d9a7015d56ae606945f6a0345c31d1f2070b4194506f7458c0c635864736f6c63430008140033"

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Contract *ContractCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Contract *ContractSession) IsOwner() (bool, error) {
	return _Contract.Contract.IsOwner(&_Contract.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Contract *ContractCallerSession) IsOwner() (bool, error) {
	return _Contract.Contract.IsOwner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// AdvanceEpochs is a paid mutator transaction binding the contract method 0x0aeeca00.
//
// Solidity: function advanceEpochs(uint256 num) returns()
func (_Contract *ContractTransactor) AdvanceEpochs(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "advanceEpochs", num)
}

// AdvanceEpochs is a paid mutator transaction binding the contract method 0x0aeeca00.
//
// Solidity: function advanceEpochs(uint256 num) returns()
func (_Contract *ContractSession) AdvanceEpochs(num *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AdvanceEpochs(&_Contract.TransactOpts, num)
}

// AdvanceEpochs is a paid mutator transaction binding the contract method 0x0aeeca00.
//
// Solidity: function advanceEpochs(uint256 num) returns()
func (_Contract *ContractTransactorSession) AdvanceEpochs(num *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AdvanceEpochs(&_Contract.TransactOpts, num)
}

// CopyCode is a paid mutator transaction binding the contract method 0xd6a0c7af.
//
// Solidity: function copyCode(address acc, address from) returns()
func (_Contract *ContractTransactor) CopyCode(opts *bind.TransactOpts, acc common.Address, from common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "copyCode", acc, from)
}

// CopyCode is a paid mutator transaction binding the contract method 0xd6a0c7af.
//
// Solidity: function copyCode(address acc, address from) returns()
func (_Contract *ContractSession) CopyCode(acc common.Address, from common.Address) (*types.Transaction, error) {
	return _Contract.Contract.CopyCode(&_Contract.TransactOpts, acc, from)
}

// CopyCode is a paid mutator transaction binding the contract method 0xd6a0c7af.
//
// Solidity: function copyCode(address acc, address from) returns()
func (_Contract *ContractTransactorSession) CopyCode(acc common.Address, from common.Address) (*types.Transaction, error) {
	return _Contract.Contract.CopyCode(&_Contract.TransactOpts, acc, from)
}

// DeactivateValidator is a paid mutator transaction binding the contract method 0x1e702f83.
//
// Solidity: function deactivateValidator(uint256 validatorID, uint256 status) returns()
func (_Contract *ContractTransactor) DeactivateValidator(opts *bind.TransactOpts, validatorID *big.Int, status *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "deactivateValidator", validatorID, status)
}

// DeactivateValidator is a paid mutator transaction binding the contract method 0x1e702f83.
//
// Solidity: function deactivateValidator(uint256 validatorID, uint256 status) returns()
func (_Contract *ContractSession) DeactivateValidator(validatorID *big.Int, status *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DeactivateValidator(&_Contract.TransactOpts, validatorID, status)
}

// DeactivateValidator is a paid mutator transaction binding the contract method 0x1e702f83.
//
// Solidity: function deactivateValidator(uint256 validatorID, uint256 status) returns()
func (_Contract *ContractTransactorSession) DeactivateValidator(validatorID *big.Int, status *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DeactivateValidator(&_Contract.TransactOpts, validatorID, status)
}

// IncBalance is a paid mutator transaction binding the contract method 0x66e7ea0f.
//
// Solidity: function incBalance(address acc, uint256 diff) returns()
func (_Contract *ContractTransactor) IncBalance(opts *bind.TransactOpts, acc common.Address, diff *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "incBalance", acc, diff)
}

// IncBalance is a paid mutator transaction binding the contract method 0x66e7ea0f.
//
// Solidity: function incBalance(address acc, uint256 diff) returns()
func (_Contract *ContractSession) IncBalance(acc common.Address, diff *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.IncBalance(&_Contract.TransactOpts, acc, diff)
}

// IncBalance is a paid mutator transaction binding the contract method 0x66e7ea0f.
//
// Solidity: function incBalance(address acc, uint256 diff) returns()
func (_Contract *ContractTransactorSession) IncBalance(acc common.Address, diff *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.IncBalance(&_Contract.TransactOpts, acc, diff)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _sfc, address _driver, address _owner) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, _sfc common.Address, _driver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", _sfc, _driver, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _sfc, address _driver, address _owner) returns()
func (_Contract *ContractSession) Initialize(_sfc common.Address, _driver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _sfc, _driver, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _sfc, address _driver, address _owner) returns()
func (_Contract *ContractTransactorSession) Initialize(_sfc common.Address, _driver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _sfc, _driver, _owner)
}

// MigrateTo is a paid mutator transaction binding the contract method 0x4ddaf8f2.
//
// Solidity: function migrateTo(address newDriverAuth) returns()
func (_Contract *ContractTransactor) MigrateTo(opts *bind.TransactOpts, newDriverAuth common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "migrateTo", newDriverAuth)
}

// MigrateTo is a paid mutator transaction binding the contract method 0x4ddaf8f2.
//
// Solidity: function migrateTo(address newDriverAuth) returns()
func (_Contract *ContractSession) MigrateTo(newDriverAuth common.Address) (*types.Transaction, error) {
	return _Contract.Contract.MigrateTo(&_Contract.TransactOpts, newDriverAuth)
}

// MigrateTo is a paid mutator transaction binding the contract method 0x4ddaf8f2.
//
// Solidity: function migrateTo(address newDriverAuth) returns()
func (_Contract *ContractTransactorSession) MigrateTo(newDriverAuth common.Address) (*types.Transaction, error) {
	return _Contract.Contract.MigrateTo(&_Contract.TransactOpts, newDriverAuth)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// SealEpoch is a paid mutator transaction binding the contract method 0xebdf104c.
//
// Solidity: function sealEpoch(uint256[] offlineTimes, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee) returns()
func (_Contract *ContractTransactor) SealEpoch(opts *bind.TransactOpts, offlineTimes []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "sealEpoch", offlineTimes, offlineBlocks, uptimes, originatedTxsFee)
}

// SealEpoch is a paid mutator transaction binding the contract method 0xebdf104c.
//
// Solidity: function sealEpoch(uint256[] offlineTimes, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee) returns()
func (_Contract *ContractSession) SealEpoch(offlineTimes []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SealEpoch(&_Contract.TransactOpts, offlineTimes, offlineBlocks, uptimes, originatedTxsFee)
}

// SealEpoch is a paid mutator transaction binding the contract method 0xebdf104c.
//
// Solidity: function sealEpoch(uint256[] offlineTimes, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee) returns()
func (_Contract *ContractTransactorSession) SealEpoch(offlineTimes []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SealEpoch(&_Contract.TransactOpts, offlineTimes, offlineBlocks, uptimes, originatedTxsFee)
}

// SealEpochValidators is a paid mutator transaction binding the contract method 0xe08d7e66.
//
// Solidity: function sealEpochValidators(uint256[] nextValidatorIDs) returns()
func (_Contract *ContractTransactor) SealEpochValidators(opts *bind.TransactOpts, nextValidatorIDs []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "sealEpochValidators", nextValidatorIDs)
}

// SealEpochValidators is a paid mutator transaction binding the contract method 0xe08d7e66.
//
// Solidity: function sealEpochValidators(uint256[] nextValidatorIDs) returns()
func (_Contract *ContractSession) SealEpochValidators(nextValidatorIDs []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SealEpochValidators(&_Contract.TransactOpts, nextValidatorIDs)
}

// SealEpochValidators is a paid mutator transaction binding the contract method 0xe08d7e66.
//
// Solidity: function sealEpochValidators(uint256[] nextValidatorIDs) returns()
func (_Contract *ContractTransactorSession) SealEpochValidators(nextValidatorIDs []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SealEpochValidators(&_Contract.TransactOpts, nextValidatorIDs)
}

// SetGenesisDelegation is a paid mutator transaction binding the contract method 0x18f628d4.
//
// Solidity: function setGenesisDelegation(address delegator, uint256 toValidatorID, uint256 stake, uint256 lockedStake, uint256 lockupFromEpoch, uint256 lockupEndTime, uint256 lockupDuration, uint256 earlyUnlockPenalty, uint256 rewards) returns()
func (_Contract *ContractTransactor) SetGenesisDelegation(opts *bind.TransactOpts, delegator common.Address, toValidatorID *big.Int, stake *big.Int, lockedStake *big.Int, lockupFromEpoch *big.Int, lockupEndTime *big.Int, lockupDuration *big.Int, earlyUnlockPenalty *big.Int, rewards *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setGenesisDelegation", delegator, toValidatorID, stake, lockedStake, lockupFromEpoch, lockupEndTime, lockupDuration, earlyUnlockPenalty, rewards)
}

// SetGenesisDelegation is a paid mutator transaction binding the contract method 0x18f628d4.
//
// Solidity: function setGenesisDelegation(address delegator, uint256 toValidatorID, uint256 stake, uint256 lockedStake, uint256 lockupFromEpoch, uint256 lockupEndTime, uint256 lockupDuration, uint256 earlyUnlockPenalty, uint256 rewards) returns()
func (_Contract *ContractSession) SetGenesisDelegation(delegator common.Address, toValidatorID *big.Int, stake *big.Int, lockedStake *big.Int, lockupFromEpoch *big.Int, lockupEndTime *big.Int, lockupDuration *big.Int, earlyUnlockPenalty *big.Int, rewards *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetGenesisDelegation(&_Contract.TransactOpts, delegator, toValidatorID, stake, lockedStake, lockupFromEpoch, lockupEndTime, lockupDuration, earlyUnlockPenalty, rewards)
}

// SetGenesisDelegation is a paid mutator transaction binding the contract method 0x18f628d4.
//
// Solidity: function setGenesisDelegation(address delegator, uint256 toValidatorID, uint256 stake, uint256 lockedStake, uint256 lockupFromEpoch, uint256 lockupEndTime, uint256 lockupDuration, uint256 earlyUnlockPenalty, uint256 rewards) returns()
func (_Contract *ContractTransactorSession) SetGenesisDelegation(delegator common.Address, toValidatorID *big.Int, stake *big.Int, lockedStake *big.Int, lockupFromEpoch *big.Int, lockupEndTime *big.Int, lockupDuration *big.Int, earlyUnlockPenalty *big.Int, rewards *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetGenesisDelegation(&_Contract.TransactOpts, delegator, toValidatorID, stake, lockedStake, lockupFromEpoch, lockupEndTime, lockupDuration, earlyUnlockPenalty, rewards)
}

// SetGenesisValidator is a paid mutator transaction binding the contract method 0x4feb92f3.
//
// Solidity: function setGenesisValidator(address _auth, uint256 validatorID, bytes pubkey, uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime) returns()
func (_Contract *ContractTransactor) SetGenesisValidator(opts *bind.TransactOpts, _auth common.Address, validatorID *big.Int, pubkey []byte, status *big.Int, createdEpoch *big.Int, createdTime *big.Int, deactivatedEpoch *big.Int, deactivatedTime *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setGenesisValidator", _auth, validatorID, pubkey, status, createdEpoch, createdTime, deactivatedEpoch, deactivatedTime)
}

// SetGenesisValidator is a paid mutator transaction binding the contract method 0x4feb92f3.
//
// Solidity: function setGenesisValidator(address _auth, uint256 validatorID, bytes pubkey, uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime) returns()
func (_Contract *ContractSession) SetGenesisValidator(_auth common.Address, validatorID *big.Int, pubkey []byte, status *big.Int, createdEpoch *big.Int, createdTime *big.Int, deactivatedEpoch *big.Int, deactivatedTime *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetGenesisValidator(&_Contract.TransactOpts, _auth, validatorID, pubkey, status, createdEpoch, createdTime, deactivatedEpoch, deactivatedTime)
}

// SetGenesisValidator is a paid mutator transaction binding the contract method 0x4feb92f3.
//
// Solidity: function setGenesisValidator(address _auth, uint256 validatorID, bytes pubkey, uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime) returns()
func (_Contract *ContractTransactorSession) SetGenesisValidator(_auth common.Address, validatorID *big.Int, pubkey []byte, status *big.Int, createdEpoch *big.Int, createdTime *big.Int, deactivatedEpoch *big.Int, deactivatedTime *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetGenesisValidator(&_Contract.TransactOpts, _auth, validatorID, pubkey, status, createdEpoch, createdTime, deactivatedEpoch, deactivatedTime)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// UpdateNetworkRules is a paid mutator transaction binding the contract method 0xb9cc6b1c.
//
// Solidity: function updateNetworkRules(bytes diff) returns()
func (_Contract *ContractTransactor) UpdateNetworkRules(opts *bind.TransactOpts, diff []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateNetworkRules", diff)
}

// UpdateNetworkRules is a paid mutator transaction binding the contract method 0xb9cc6b1c.
//
// Solidity: function updateNetworkRules(bytes diff) returns()
func (_Contract *ContractSession) UpdateNetworkRules(diff []byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateNetworkRules(&_Contract.TransactOpts, diff)
}

// UpdateNetworkRules is a paid mutator transaction binding the contract method 0xb9cc6b1c.
//
// Solidity: function updateNetworkRules(bytes diff) returns()
func (_Contract *ContractTransactorSession) UpdateNetworkRules(diff []byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateNetworkRules(&_Contract.TransactOpts, diff)
}

// UpdateNetworkVersion is a paid mutator transaction binding the contract method 0x267ab446.
//
// Solidity: function updateNetworkVersion(uint256 version) returns()
func (_Contract *ContractTransactor) UpdateNetworkVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateNetworkVersion", version)
}

// UpdateNetworkVersion is a paid mutator transaction binding the contract method 0x267ab446.
//
// Solidity: function updateNetworkVersion(uint256 version) returns()
func (_Contract *ContractSession) UpdateNetworkVersion(version *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateNetworkVersion(&_Contract.TransactOpts, version)
}

// UpdateNetworkVersion is a paid mutator transaction binding the contract method 0x267ab446.
//
// Solidity: function updateNetworkVersion(uint256 version) returns()
func (_Contract *ContractTransactorSession) UpdateNetworkVersion(version *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateNetworkVersion(&_Contract.TransactOpts, version)
}

// UpdateValidatorPubkey is a paid mutator transaction binding the contract method 0x242a6e3f.
//
// Solidity: function updateValidatorPubkey(uint256 validatorID, bytes pubkey) returns()
func (_Contract *ContractTransactor) UpdateValidatorPubkey(opts *bind.TransactOpts, validatorID *big.Int, pubkey []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateValidatorPubkey", validatorID, pubkey)
}

// UpdateValidatorPubkey is a paid mutator transaction binding the contract method 0x242a6e3f.
//
// Solidity: function updateValidatorPubkey(uint256 validatorID, bytes pubkey) returns()
func (_Contract *ContractSession) UpdateValidatorPubkey(validatorID *big.Int, pubkey []byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateValidatorPubkey(&_Contract.TransactOpts, validatorID, pubkey)
}

// UpdateValidatorPubkey is a paid mutator transaction binding the contract method 0x242a6e3f.
//
// Solidity: function updateValidatorPubkey(uint256 validatorID, bytes pubkey) returns()
func (_Contract *ContractTransactorSession) UpdateValidatorPubkey(validatorID *big.Int, pubkey []byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateValidatorPubkey(&_Contract.TransactOpts, validatorID, pubkey)
}

// UpdateValidatorWeight is a paid mutator transaction binding the contract method 0xa4066fbe.
//
// Solidity: function updateValidatorWeight(uint256 validatorID, uint256 value) returns()
func (_Contract *ContractTransactor) UpdateValidatorWeight(opts *bind.TransactOpts, validatorID *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateValidatorWeight", validatorID, value)
}

// UpdateValidatorWeight is a paid mutator transaction binding the contract method 0xa4066fbe.
//
// Solidity: function updateValidatorWeight(uint256 validatorID, uint256 value) returns()
func (_Contract *ContractSession) UpdateValidatorWeight(validatorID *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateValidatorWeight(&_Contract.TransactOpts, validatorID, value)
}

// UpdateValidatorWeight is a paid mutator transaction binding the contract method 0xa4066fbe.
//
// Solidity: function updateValidatorWeight(uint256 validatorID, uint256 value) returns()
func (_Contract *ContractTransactorSession) UpdateValidatorWeight(validatorID *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateValidatorWeight(&_Contract.TransactOpts, validatorID, value)
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var ContractBinRuntime = "0x608060405234801561001057600080fd5b50600436106101365760003560e01c80638da5cb5b116100b2578063c0c53b8b11610081578063e08d7e6611610066578063e08d7e66146104f5578063ebdf104c14610565578063f2fde38b146106cb57610136565b8063c0c53b8b14610475578063d6a0c7af146104ba57610136565b80638da5cb5b146103955780638f32d59b146103c6578063a4066fbe146103e2578063b9cc6b1c1461040557610136565b8063267ab446116101095780634feb92f3116100ee5780634feb92f3146102a957806366e7ea0f14610354578063715018a61461038d57610136565b8063267ab446146102595780634ddaf8f21461027657610136565b80630aeeca001461013b57806318f628d41461015a5780631e702f83146101bf578063242a6e3f146101e2575b600080fd5b6101586004803603602081101561015157600080fd5b50356106fe565b005b610158600480360361012081101561017157600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060208101359060408101359060608101359060808101359060a08101359060c08101359060e08101359061010001356107e5565b610158600480360360408110156101d557600080fd5b508035906020013561090c565b610158600480360360408110156101f857600080fd5b8135919081019060408101602082013564010000000081111561021a57600080fd5b82018360208201111561022c57600080fd5b8035906020019184600183028401116401000000008311171561024e57600080fd5b5090925090506109f8565b6101586004803603602081101561026f57600080fd5b5035610b27565b6101586004803603602081101561028c57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610bf3565b61015860048036036101008110156102c057600080fd5b73ffffffffffffffffffffffffffffffffffffffff823516916020810135918101906060810160408201356401000000008111156102fd57600080fd5b82018360208201111561030f57600080fd5b8035906020019184600183028401116401000000008311171561033157600080fd5b919350915080359060208101359060408101359060608101359060800135610cc0565b6101586004803603604081101561036a57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135610e1b565b610158610f80565b61039d611048565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6103ce611064565b604080519115158252519081900360200190f35b610158600480360360408110156103f857600080fd5b5080359060200135611082565b6101586004803603602081101561041b57600080fd5b81019060208101813564010000000081111561043657600080fd5b82018360208201111561044857600080fd5b8035906020019184600183028401116401000000008311171561046a57600080fd5b509092509050611168565b6101586004803603606081101561048b57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813581169160208101358216916040909101351661125e565b610158600480360360408110156104d057600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813581169160200135166113b9565b6101586004803603602081101561050b57600080fd5b81019060208101813564010000000081111561052657600080fd5b82018360208201111561053857600080fd5b8035906020019184602083028401116401000000008311171561055a57600080fd5b50909250905061151d565b6101586004803603608081101561057b57600080fd5b81019060208101813564010000000081111561059657600080fd5b8201836020820111156105a857600080fd5b803590602001918460208302840111640100000000831117156105ca57600080fd5b9193909290916020810190356401000000008111156105e857600080fd5b8201836020820111156105fa57600080fd5b8035906020019184602083028401116401000000008311171561061c57600080fd5b91939092909160208101903564010000000081111561063a57600080fd5b82018360208201111561064c57600080fd5b8035906020019184602083028401116401000000008311171561066e57600080fd5b91939092909160208101903564010000000081111561068c57600080fd5b82018360208201111561069e57600080fd5b803590602001918460208302840111640100000000831117156106c057600080fd5b509092509050611616565b610158600480360360208110156106e157600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661181c565b610706611064565b610757576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b606754604080517f0aeeca0000000000000000000000000000000000000000000000000000000000815260048101849052905173ffffffffffffffffffffffffffffffffffffffff90921691630aeeca009160248082019260009290919082900301818387803b1580156107ca57600080fd5b505af11580156107de573d6000803e3d6000fd5b5050505050565b60675473ffffffffffffffffffffffffffffffffffffffff16331461083b5760405162461bcd60e51b8152600401808060200182810382526025815260200180611bad6025913960400191505060405180910390fd5b606654604080517f18f628d400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8c81166004830152602482018c9052604482018b9052606482018a90526084820189905260a4820188905260c4820187905260e482018690526101048201859052915191909216916318f628d49161012480830192600092919082900301818387803b1580156108e957600080fd5b505af11580156108fd573d6000803e3d6000fd5b50505050505050505050505050565b60675473ffffffffffffffffffffffffffffffffffffffff1633146109625760405162461bcd60e51b8152600401808060200182810382526025815260200180611bad6025913960400191505060405180910390fd5b606654604080517f1e702f830000000000000000000000000000000000000000000000000000000081526004810185905260248101849052905173ffffffffffffffffffffffffffffffffffffffff90921691631e702f839160448082019260009290919082900301818387803b1580156109dc57600080fd5b505af11580156109f0573d6000803e3d6000fd5b505050505050565b60665473ffffffffffffffffffffffffffffffffffffffff163314610a64576040805162461bcd60e51b815260206004820152601e60248201527f63616c6c6572206973206e6f74207468652053464320636f6e74726163740000604482015290519081900360640190fd5b606754604080517f242a6e3f00000000000000000000000000000000000000000000000000000000815260048101868152602482019283526044820185905273ffffffffffffffffffffffffffffffffffffffff9093169263242a6e3f928792879287929091606401848480828437600081840152601f19601f820116905080830192505050945050505050600060405180830381600087803b158015610b0a57600080fd5b505af1158015610b1e573d6000803e3d6000fd5b50505050505050565b610b2f611064565b610b80576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b606754604080517f267ab44600000000000000000000000000000000000000000000000000000000815260048101849052905173ffffffffffffffffffffffffffffffffffffffff9092169163267ab4469160248082019260009290919082900301818387803b1580156107ca57600080fd5b610bfb611064565b610c4c576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b606754604080517fda7fc24f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301529151919092169163da7fc24f91602480830192600092919082900301818387803b1580156107ca57600080fd5b60675473ffffffffffffffffffffffffffffffffffffffff163314610d165760405162461bcd60e51b8152600401808060200182810382526025815260200180611bad6025913960400191505060405180910390fd5b606660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634feb92f38a8a8a8a8a8a8a8a8a6040518a63ffffffff1660e01b8152600401808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001898152602001806020018781526020018681526020018581526020018481526020018381526020018281038252898982818152602001925080828437600081840152601f19601f8201169050808301925050509a5050505050505050505050600060405180830381600087803b1580156108e957600080fd5b60665473ffffffffffffffffffffffffffffffffffffffff163314610e87576040805162461bcd60e51b815260206004820152601e60248201527f63616c6c6572206973206e6f74207468652053464320636f6e74726163740000604482015290519081900360640190fd5b60665473ffffffffffffffffffffffffffffffffffffffff838116911614610ee05760405162461bcd60e51b8152600401808060200182810382526021815260200180611b8c6021913960400191505060405180910390fd5b60675473ffffffffffffffffffffffffffffffffffffffff9081169063e30443bc908490610f17908216318563ffffffff61188116565b6040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b1580156109dc57600080fd5b610f88611064565b610fd9576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b60335460405160009173ffffffffffffffffffffffffffffffffffffffff16907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055565b60335473ffffffffffffffffffffffffffffffffffffffff1690565b60335473ffffffffffffffffffffffffffffffffffffffff16331490565b60665473ffffffffffffffffffffffffffffffffffffffff1633146110ee576040805162461bcd60e51b815260206004820152601e60248201527f63616c6c6572206973206e6f74207468652053464320636f6e74726163740000604482015290519081900360640190fd5b606754604080517fa4066fbe0000000000000000000000000000000000000000000000000000000081526004810185905260248101849052905173ffffffffffffffffffffffffffffffffffffffff9092169163a4066fbe9160448082019260009290919082900301818387803b1580156109dc57600080fd5b611170611064565b6111c1576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6067546040517fb9cc6b1c0000000000000000000000000000000000000000000000000000000081526020600482019081526024820184905273ffffffffffffffffffffffffffffffffffffffff9092169163b9cc6b1c91859185918190604401848480828437600081840152601f19601f8201169050808301925050509350505050600060405180830381600087803b1580156109dc57600080fd5b600054610100900460ff168061127757506112776118e2565b80611285575060005460ff16155b6112c05760405162461bcd60e51b815260040180806020018281038252602e815260200180611b5e602e913960400191505060405180910390fd5b600054610100900460ff1615801561132657600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b61132f826118e8565b6067805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255606680549287169290911691909117905580156113b357600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b50505050565b6113c1611064565b611412576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b60665473ffffffffffffffffffffffffffffffffffffffff83811691161480611450575073ffffffffffffffffffffffffffffffffffffffff821630145b6114a1576040805162461bcd60e51b815260206004820152601760248201527f6e6f7420534643206f722073656c662061646472657373000000000000000000604482015290519081900360640190fd5b606754604080517fd6a0c7af00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff858116600483015284811660248301529151919092169163d6a0c7af91604480830192600092919082900301818387803b1580156109dc57600080fd5b60675473ffffffffffffffffffffffffffffffffffffffff1633146115735760405162461bcd60e51b8152600401808060200182810382526025815260200180611bad6025913960400191505060405180910390fd5b6066546040517fe08d7e660000000000000000000000000000000000000000000000000000000081526020600482018181526024830185905273ffffffffffffffffffffffffffffffffffffffff9093169263e08d7e6692869286929182916044909101908590850280828437600081840152601f19601f8201169050808301925050509350505050600060405180830381600087803b1580156109dc57600080fd5b60675473ffffffffffffffffffffffffffffffffffffffff16331461166c5760405162461bcd60e51b8152600401808060200182810382526025815260200180611bad6025913960400191505060405180910390fd5b606660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ebdf104c89898989898989896040518963ffffffff1660e01b8152600401808060200180602001806020018060200185810385528d8d82818152602001925060200280828437600083820152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01690910186810385528b8152602090810191508c908c0280828437600083820152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169091018681038452898152602090810191508a908a0280828437600083820152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169091018681038352878152602090810191508890880280828437600081840152601f19601f8201169050808301925050509c50505050505050505050505050600060405180830381600087803b1580156117fa57600080fd5b505af115801561180e573d6000803e3d6000fd5b505050505050505050505050565b611824611064565b611875576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b61187e81611a57565b50565b6000828201838110156118db576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b303b1590565b600054610100900460ff168061190157506119016118e2565b8061190f575060005460ff16155b61194a5760405162461bcd60e51b815260040180806020018281038252602e815260200180611b5e602e913960400191505060405180910390fd5b600054610100900460ff161580156119b057600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b603380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84811691909117918290556040519116906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a38015611a5357600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b5050565b73ffffffffffffffffffffffffffffffffffffffff8116611aa95760405162461bcd60e51b8152600401808060200182810382526026815260200180611b386026913960400191505060405180910390fd5b60335460405173ffffffffffffffffffffffffffffffffffffffff8084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3603380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905556fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a6564726563697069656e74206973206e6f74207468652053464320636f6e747261637463616c6c6572206973206e6f7420746865204e6f646544726976657220636f6e7472616374a265627a7a7231582085f0b421e25c5c41c27b69f82881c0fbe44729bdd97b1b955484e0541fc3768f64736f6c63430005110032"
