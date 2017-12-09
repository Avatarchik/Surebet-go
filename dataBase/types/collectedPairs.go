package types

import (
	"sync"
	"surebetSearch/common"
)

type CollectedPairs struct {
	v   []EventPair
	mux sync.RWMutex
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
