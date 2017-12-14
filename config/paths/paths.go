package paths

import "os"

var projectDir = os.ExpandEnv("$GOPATH/src/github.com/korovkinand/surebetSearch/")

var CollectedPairs = projectDir + "dataBase/collectedPairs"
var PositiveAccounts = projectDir + "config/accounts/positive"
