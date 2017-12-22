package db

import (
	"github.com/korovkinand/surebetSearch/common/types"
	"github.com/korovkinand/surebetSearch/config"
	"testing"
	"time"
)

func TestCollect(t *testing.T) {
	s.Rng = &types.Range{From: 11, To: 14}
	accounts.SetRange(s.Rng)

	s.Time.Sleep = 10 * time.Second
	s.Time.Limit = (s.Time.Sleep * 3) / 2

	if err := Collect(config.DbPath); err != nil {
		t.Error(err)
	}
}
