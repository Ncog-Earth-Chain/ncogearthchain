// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package ethapi implements the general Ethereum API functions.
package ethapi

import (
	"context"
	"math/big"

	"github.com/Ncog-Earth-Chain/forest-base/hash"
	"github.com/Ncog-Earth-Chain/forest-base/inter/idx"
	"github.com/Ncog-Earth-Chain/forest-base/inter/pos"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethdb"
	notify "github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/Ncog-Earth-Chain/ncogearthchain/evmcore"
	"github.com/Ncog-Earth-Chain/ncogearthchain/gossip/sfcapi"
	"github.com/Ncog-Earth-Chain/ncogearthchain/inter"
)

// PeerProgress is synchronization status of a peer
type PeerProgress struct {
	CurrentEpoch     idx.Epoch
	CurrentBlock     idx.Block
	CurrentBlockHash hash.Event
	CurrentBlockTime inter.Timestamp
	HighestBlock     idx.Block
	HighestEpoch     idx.Epoch
}

// Backend interface provides the common API services (that are provided by
// both full and light clients) with access to necessary functions.
type Backend interface {
	// General Ethereum API
	Progress() PeerProgress
	SuggestPrice(ctx context.Context) (*big.Int, error)
	ChainDb() ethdb.Database
	AccountManager() *accounts.Manager
	ExtRPCEnabled() bool
	RPCGasCap() uint64        // global gas cap for eth_call over rpc: DoS protection
	RPCTxFeeCap() float64     // global tx fee cap for all transaction related APIs
	UnprotectedAllowed() bool // allows only for EIP155 transactions.
	CalcLogsBloom() bool

	// Blockchain API
	HeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*evmcore.EvmHeader, error)
	HeaderByHash(ctx context.Context, hash common.Hash) (*evmcore.EvmHeader, error)
	BlockByNumber(ctx context.Context, number rpc.BlockNumber) (*evmcore.EvmBlock, error)
	StateAndHeaderByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*state.StateDB, *evmcore.EvmHeader, error)
	//GetHeader(ctx context.Context, hash common.Hash) *evmcore.EvmHeader
	BlockByHash(ctx context.Context, hash common.Hash) (*evmcore.EvmBlock, error)
	GetReceiptsByNumber(ctx context.Context, number rpc.BlockNumber) (types.Receipts, error)
	GetTd(hash common.Hash) *big.Int
	GetEVM(ctx context.Context, msg evmcore.Message, state *state.StateDB, header *evmcore.EvmHeader, vmConfig *vm.Config) (*vm.EVM, func() error, error)
	MinGasPrice() *big.Int
	MaxGasLimit() uint64

	// Transaction pool API
	SendTx(ctx context.Context, signedTx *types.Transaction) error
	GetTransaction(ctx context.Context, txHash common.Hash) (*types.Transaction, uint64, uint64, error)
	GetPoolTransactions() (types.Transactions, error)
	GetPoolTransaction(txHash common.Hash) *types.Transaction
	GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error)
	Stats() (pending int, queued int)
	TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions)
	SubscribeNewTxsNotify(chan<- evmcore.NewTxsNotify) notify.Subscription

	ChainConfig() *params.ChainConfig
	CurrentBlock() *evmcore.EvmBlock

	// Forest DAG API
	GetEventPayload(ctx context.Context, shortEventID string) (*inter.EventPayload, error)
	GetEvent(ctx context.Context, shortEventID string) (*inter.Event, error)
	GetHeads(ctx context.Context, epoch rpc.BlockNumber) (hash.Events, error)
	CurrentEpoch(ctx context.Context) idx.Epoch
	SealedEpochTiming(ctx context.Context) (start inter.Timestamp, end inter.Timestamp)

	// Forest SFC API
	GetValidators(ctx context.Context) *pos.Validators
	GetUptime(ctx context.Context, stakerID idx.ValidatorID) (*big.Int, error)
	GetOriginatedFee(ctx context.Context, stakerID idx.ValidatorID) (*big.Int, error)
	GetRewardWeights(ctx context.Context, stakerID idx.ValidatorID) (*big.Int, *big.Int, error)
	GetStakerPoI(ctx context.Context, stakerID idx.ValidatorID) (*big.Int, error)
	GetDowntime(ctx context.Context, stakerID idx.ValidatorID) (idx.Block, inter.Timestamp, error)
	GetDelegationClaimedRewards(ctx context.Context, id sfcapi.DelegationID) (*big.Int, error)
	GetStakerClaimedRewards(ctx context.Context, stakerID idx.ValidatorID) (*big.Int, error)
	GetStakerDelegationsClaimedRewards(ctx context.Context, stakerID idx.ValidatorID) (*big.Int, error)
	GetStaker(ctx context.Context, stakerID idx.ValidatorID) (*sfcapi.SfcStaker, error)
	GetStakerID(ctx context.Context, addr common.Address) (idx.ValidatorID, error)
	GetStakers(ctx context.Context) ([]sfcapi.SfcStakerAndID, error)
	GetDelegationsOf(ctx context.Context, stakerID idx.ValidatorID) ([]sfcapi.SfcDelegationAndID, error)
	GetDelegationsByAddress(ctx context.Context, addr common.Address) ([]sfcapi.SfcDelegationAndID, error)
	GetDelegation(ctx context.Context, id sfcapi.DelegationID) (*sfcapi.SfcDelegation, error)
}

func GetAPIs(apiBackend Backend) []rpc.API {
	nonceLock := new(AddrLocker)
	orig := []rpc.API{
		{
			Namespace: "eth",
			Version:   "1.0",
			Service:   NewPublicEthereumAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "eth",
			Version:   "1.0",
			Service:   NewPublicBlockChainAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "dag",
			Version:   "1.0",
			Service:   NewPublicDAGChainAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "eth",
			Version:   "1.0",
			Service:   NewPublicTransactionPoolAPI(apiBackend, nonceLock),
			Public:    true,
		}, {
			Namespace: "txpool",
			Version:   "1.0",
			Service:   NewPublicTxPoolAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPublicDebugAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPrivateDebugAPI(apiBackend),
		}, {
			Namespace: "eth",
			Version:   "1.0",
			Service:   NewPublicAccountAPI(apiBackend.AccountManager()),
			Public:    true,
		}, {
			Namespace: "personal",
			Version:   "1.0",
			Service:   NewPrivateAccountAPI(apiBackend, nonceLock),
			Public:    false,
		}, {
			Namespace: "sfc",
			Version:   "1.0",
			Service:   NewPublicSfcAPI(apiBackend),
			Public:    false,
		}, {
			Namespace: "abft",
			Version:   "1.0",
			Service:   NewPublicAbftAPI(apiBackend),
			Public:    false,
		},
	}

	// NOTE: eth-namespace is doubled as nec-namespace for branding purpose
	double := []rpc.API{
		{
			Namespace: "nec",
			Version:   "1.0",
			Service:   NewPublicEthereumAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "nec",
			Version:   "1.0",
			Service:   NewPublicBlockChainAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "nec",
			Version:   "1.0",
			Service:   NewPublicTransactionPoolAPI(apiBackend, nonceLock),
			Public:    true,
		}, {
			Namespace: "nec",
			Version:   "1.0",
			Service:   NewPublicAccountAPI(apiBackend.AccountManager()),
			Public:    true,
		},
	}

	return append(orig, double...)
}
