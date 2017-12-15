package dataBase

import (
	"github.com/korovkinand/surebetSearch/common"
	"github.com/korovkinand/surebetSearch/config/accounts"
	"github.com/korovkinand/surebetSearch/config/intervals"
	"testing"
)

func TestCollect(t *testing.T) {
	accounts.PositiveRange = common.Range{11, 14}
	intervals.PositiveWorkLimit = intervals.PositiveSleep + 10
	if err := Collect(); err != nil {
		t.Error(err)
	}
}
