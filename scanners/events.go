package scanners

type MatchEvent struct {
	Site  string `json:"BookMaker"`
	Team1 string `json:"FirstTeam"`
	Team2 string `json:"SecondTeam"`
}

type EventPair struct {
	Event1 MatchEvent `json:"FirstEvent"`
	Event2 MatchEvent `json:"SecondEvent"`
}

//new format
//type MatchEvent struct {
//	Site  string `json:"site"`
//	Team1 string `json:"t1"`
//	Team2 string `json:"t2"`
//}
//
//type EventPair struct {
//	Event1 MatchEvent `json:"ev1"`
//	Event2 MatchEvent `json:"ev2"`
//}
