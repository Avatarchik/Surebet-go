package types

import (
	"github.com/korovkinand/surebetSearch/common"
	"sync"
)

type CollectedPairs struct {
	v   []EventPair
	mux sync.RWMutex
}

type CollectedPairsItem struct {
	Idx int
	V   EventPair
}

func (c *CollectedPairs) Iter() <-chan CollectedPairsItem {
	ch := make(chan CollectedPairsItem)

	go func() {
		defer close(ch)
		c.mux.RLock()
		defer c.mux.RUnlock()

		for idx, value := range c.v {
			ch <- CollectedPairsItem{idx, value}
		}
	}()

	return ch
}

func (c *CollectedPairs) AppendUnique(newPairs []EventPair) {
	c.mux.Lock()
	defer c.mux.Unlock()

loop:
	for _, newPair := range newPairs {
		for _, collectedPair := range c.v {
			if collectedPair == newPair {
				continue loop
			}
		}
		c.v = append(c.v, newPair)
	}
}

func (c *CollectedPairs) Length() int {
	c.mux.RLock()
	defer c.mux.RUnlock()

	return len(c.v)
}

func (c *CollectedPairs) Load(filename string) error {
	c.mux.Lock()
	defer c.mux.Unlock()

	return common.LoadJson(filename, &(c.v))
}

func (c *CollectedPairs) Save(filename string) error {
	c.mux.RLock()
	defer c.mux.RUnlock()

	return common.SaveJson(filename, c.v)
}
