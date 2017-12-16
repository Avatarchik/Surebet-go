package info

import (
	"github.com/korovkinand/surebetSearch/common"
	"os"
	"time"
)

var projDir = os.ExpandEnv("$GOPATH/src/github.com/korovkinand/surebetSearch/")

var Fonbet = common.SiteInfo{
	Url:  "https://www.fonbet.com/live/",
	Node: `#lineTable > tbody`,
}

var Marat = common.SiteInfo{
	Url: "https://www.marathonbet.com/en/live/",
}

var Olimp = common.SiteInfo{
	Url: "https://olimp.com/betting",
}

var Posit = common.SiteInfo{
	Acc: common.Accounts{
		Rng:  common.Range{0, 7},
		Path: projDir + "config/accounts/positive",
	},
	Time: common.Time{
		Sleep: 35 * time.Second,
		Limit: 24 * time.Hour,
	},
	Url:  "https://positivebet.com/en/user/login",
	Node: `.grid-view > table`,
}
