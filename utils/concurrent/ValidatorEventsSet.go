package concurrent

import (
	"sync"

	"github.com/Ncog-Earth-Chain/forest-base/hash"
	"github.com/Ncog-Earth-Chain/forest-base/inter/idx"
)

type ValidatorEventsSet struct {
	sync.RWMutex
	Val map[idx.ValidatorID]hash.Event
}

func WrapValidatorEventsSet(v map[idx.ValidatorID]hash.Event) *ValidatorEventsSet {
	return &ValidatorEventsSet{
		RWMutex: sync.RWMutex{},
		Val:     v,
	}
}
