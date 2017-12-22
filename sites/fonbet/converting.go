package fonbet

import (
	"github.com/korovkinand/surebetSearch/common"
	"github.com/korovkinand/surebetSearch/common/types"
)

func ConvertName(teams types.Teams) types.Teams {
	teams.Team1 = common.ReformatSpaces(teams.Team1)
	teams.Team2 = common.ReformatSpaces(teams.Team2)
	return teams
}
