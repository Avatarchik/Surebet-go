package paths

import "os"

var projectDir = os.ExpandEnv("$GOPATH/src/surebetSearch/")

var CollectedPairs = projectDir + "dataBase/collectedPairs"
var PositiveAccounts = projectDir + "config/accounts/positive"
