package db

import (
	"github.com/korovkinand/surebetSearch/common"
	"github.com/korovkinand/surebetSearch/config"
	"testing"
	"time"
)

func TestCollect(t *testing.T) {
	s.Rng = &common.Range{11, 14}
	accounts.SetRange(s.Rng)

	s.Time.Sleep = 10 * time.Second
	s.Time.Limit = (s.Time.Sleep * 3) / 2

	if err := Collect(config.DbPath); err != nil {
		t.Error(err)
	}
}