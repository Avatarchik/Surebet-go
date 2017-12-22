package fonbet

import (
	"github.com/korovkinand/surebetSearch/common"
)

func ReformatName(teams common.Teams) common.Teams {
	teams.Team1 = common.ReformatSpaces(teams.Team1)
	teams.Team2 = common.ReformatSpaces(teams.Team2)
	return teams
}
