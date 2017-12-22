package config

import "os"

var projDir = os.ExpandEnv("$GOPATH/src/github.com/korovkinand/surebetSearch/")

var (
	sitesConfigDir   = projDir + "config/sites/"
	accountsDir      = projDir + "config/accounts/"
	FonbetSamplesDir = projDir + "sites/fonbet/testing/"
)

var DbPath = projDir + "db/eventPairs"
