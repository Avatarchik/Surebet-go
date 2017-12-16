package scanners

import (
	"github.com/korovkinand/surebetSearch/common"
	"sync"
)

type EventPairs struct {
	v   []EventPair
	mux sync.RWMutex
}

type EvPairsItem struct {
	Idx int
	V   EventPair
}

func (c *EventPairs) Iter() <-chan EvPairsItem {
	ch := make(chan EvPairsItem)

	go func() {
		defer close(ch)
		c.mux.RLock()
		defer c.mux.RUnlock()

		for idx, value := range c.v {
			ch <- EvPairsItem{idx, value}
		}
	}()

	return ch
}

func (c *EventPairs) AppendUnique(newPairs []EventPair) {
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

func (c *EventPairs) Length() int {
	c.mux.RLock()
	defer c.mux.RUnlock()

	return len(c.v)
}

func (c *EventPairs) Load(filename string) error {
	c.mux.Lock()
	defer c.mux.Unlock()

	return common.LoadJson(filename, &(c.v))
}

func (c *EventPairs) Save(filename string) error {
	c.mux.RLock()
	defer c.mux.RUnlock()

	return common.SaveJson(filename, c.v)
}
