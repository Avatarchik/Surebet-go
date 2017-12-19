package types

import (
	"github.com/emirpasic/gods/sets/hashset"
	"sync"
)

type SafeSet struct {
	v   *hashset.Set
	mux sync.RWMutex
}

func NewSafeSet() *SafeSet {
	return &SafeSet{v: hashset.New()}
}

func (s *SafeSet) Add(elements ...interface{}) {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.v.Add(elements...)
}

func (s *SafeSet) Size() int {
	s.mux.RLock()
	defer s.mux.RUnlock()

	return s.v.Size()
}

func (s *SafeSet) Values() []interface{} {
	s.mux.RLock()
	defer s.mux.RUnlock()

	return s.v.Values()
}

func (s *SafeSet) AddGen(slice interface{}) {
	s.Add(ToInterface(slice)...)
}

func (s *SafeSet) ValuesGen(nilEl interface{}) interface{} {
	return ToType(s.Values(), nilEl)
}
