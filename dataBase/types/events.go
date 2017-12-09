package types

type MatchEvent struct {
	BookMaker  string `json:"BookMaker"`
	FirstTeam  string `json:"FirstTeam"`
	SecondTeam string `json:"SecondTeam"`
}

type EventPair struct {
	FirstEvent  MatchEvent `json:"FirstEvent"`
	SecondEvent MatchEvent `json:"SecondEvent"`
}
