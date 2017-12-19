package config

import "os"

var projDir = os.ExpandEnv("$GOPATH/src/github.com/korovkinand/surebetSearch/")

var (
	sitesPath    = projDir + "config/sites/"
	accountsPath = projDir + "config/accounts/"
)

var DbPath = projDir + "db/eventPairs"
