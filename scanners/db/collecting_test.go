package db

import (
	"github.com/korovkinand/surebetSearch/common"
	"github.com/korovkinand/surebetSearch/config/info"
	"testing"
)

func TestCollect(t *testing.T) {
	info.Posit.Acc.Rng = common.Range{11, 14}
	info.Posit.Time.Limit = (info.Posit.Time.Sleep * 3) / 2
	if err := Collect(); err != nil {
		t.Error(err)
	}
}
