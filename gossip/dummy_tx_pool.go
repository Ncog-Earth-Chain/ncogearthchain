package gossip

import (
	"math/rand"
	"sort"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	notify "github.com/ethereum/go-ethereum/event"

	"github.com/Ncog-Earth-Chain/ncogearthchain/evmcore"
)

// dummyTxPool is a fake, helper transaction pool for testing purposes
type dummyTxPool struct {
	txFeed notify.Feed
	pool   []*types.Transaction        // Collection of all transactions
	added  chan<- []*types.Transaction // Notification channel for new transactions

	lock sync.RWMutex // Protects the transaction pool
}

// AddRemotes appends a batch of transactions to the pool, and notifies any
// listeners if the addition channel is non nil
func (p *dummyTxPool) AddRemotes(txs []*types.Transaction) []error {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.pool = append(p.pool, txs...)
	if p.added != nil {
		p.added <- txs
	}
	return make([]error, len(txs))
}

// Pending returns all the transactions known to the pool
func (p *dummyTxPool) Pending() (map[common.Address]types.Transactions, error) {
	p.lock.RLock()
	defer p.lock.RUnlock()

	batches := make(map[common.Address]types.Transactions)
	for _, tx := range p.pool {
		from, _ := types.Sender(types.HomesteadSigner{}, tx)
		batches[from] = append(batches[from], tx)
	}
	for _, batch := range batches {
		sort.Sort(types.TxByNonce(batch))
	}
	return batches, nil
}

func (p *dummyTxPool) SubscribeNewTxsNotify(ch chan<- evmcore.NewTxsNotify) notify.Subscription {
	return p.txFeed.Subscribe(ch)
}

func (p *dummyTxPool) Map() map[common.Hash]*types.Transaction {
	p.lock.RLock()
	defer p.lock.RUnlock()
	res := make(map[common.Hash]*types.Transaction, len(p.pool))
	for _, tx := range p.pool {
		res[tx.Hash()] = tx
	}
	return nil
}

func (p *dummyTxPool) Get(txid common.Hash) *types.Transaction {
	return p.Map()[txid]
}

func (p *dummyTxPool) OnlyNotExisting(txids []common.Hash) []common.Hash {
	m := p.Map()
	notExisting := make([]common.Hash, 0, len(txids))
	for _, txid := range txids {
		if m[txid] == nil {
			notExisting = append(notExisting, txid)
		}
	}
	return notExisting
}

func (p *dummyTxPool) SampleHashes(max int) []common.Hash {
	p.lock.RLock()
	defer p.lock.RUnlock()
	res := make([]common.Hash, 0, max)
	skip := 0
	if len(p.pool) > max {
		skip = rand.Intn(len(p.pool) - max)
	}
	for _, tx := range p.pool {
		if len(res) >= max {
			break
		}
		if skip > 0 {
			skip--
			continue
		}
		res = append(res, tx.Hash())
	}
	return res
}
