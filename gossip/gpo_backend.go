package gossip

import (
	"github.com/Ncog-Earth-Chain/forest-base/hash"
	"github.com/Ncog-Earth-Chain/forest-base/inter/idx"

	"github.com/Ncog-Earth-Chain/ncogearthchain/inter"
	"github.com/Ncog-Earth-Chain/ncogearthchain/ncogearthchain"
	"github.com/Ncog-Earth-Chain/ncogearthchain/utils/concurrent"
)

type GPOBackend struct {
	store *Store
}

func (b *GPOBackend) GetLatestBlockIndex() idx.Block {
	return b.store.GetLatestBlockIndex()
}

func (b *GPOBackend) GetRules() ncogearthchain.Rules {
	return b.store.GetRules()
}

func (b *GPOBackend) GetPendingRules() ncogearthchain.Rules {
	return b.store.GetBlockState().DirtyRules
}

// TotalGasPowerLeft returns a total amount of obtained gas power by the validators, according to the latest events from each validator
func (b *GPOBackend) TotalGasPowerLeft() uint64 {
	es := b.store.GetEpochState()
	set := b.store.GetLastEvents(es.Epoch)
	if set == nil {
		set = concurrent.WrapValidatorEventsSet(map[idx.ValidatorID]hash.Event{})
	}
	set.RLock()
	defer set.RUnlock()
	metValidators := map[idx.ValidatorID]bool{}
	total := uint64(0)
	// count GasPowerLeft from latest events of this epoch
	for _, tip := range set.Val {
		e := b.store.GetEvent(tip)
		total += e.GasPowerLeft().Gas[inter.LongTermGas]
		metValidators[e.Creator()] = true
	}
	// count GasPowerLeft from last events of prev epoch if no event in current epoch is present
	for i := idx.Validator(0); i < es.Validators.Len(); i++ {
		vid := es.Validators.GetID(i)
		if !metValidators[vid] {
			prevEvent := b.store.GetEvent(es.ValidatorStates[i].PrevEpochEvent)
			if prevEvent != nil {
				total += prevEvent.GasPowerLeft().Gas[inter.LongTermGas]
			}
		}
	}

	return total
}
