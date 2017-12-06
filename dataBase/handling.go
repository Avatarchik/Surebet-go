package dataBase

import (
	"surebetSearch/dataBase/types"
	"sync"
	"io/ioutil"
	"encoding/json"
)

type CollectedPairs struct {
	v   []types.EventPair
	mux sync.Mutex
}

func (c *CollectedPairs) AppendUnique(newPairs []types.EventPair) {
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
	c.mux.Lock()
	defer c.mux.Unlock()

	return len(c.v)
}

func (c *CollectedPairs) Load(filename string) error {
	byteData, err := ioutil.ReadFile(filename + ".json")
	if err != nil {
		return err
	}

	c.mux.Lock()
	defer c.mux.Unlock()

	err = json.Unmarshal(byteData, &(c.v))
	if err != nil {
		return err
	}

	return nil
}

func (c *CollectedPairs) Save(filename string) error {
	c.mux.Lock()
	defer c.mux.Unlock()

	byteData, err := json.Marshal(c.v)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filename+".json", byteData, 0644); err != nil {
		return err
	}
	return nil
}
