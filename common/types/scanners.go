package types

type MatchEvent struct {
	Site  string `json:"s"`
	Team1 string `json:"t1"`
	Team2 string `json:"t2"`
}

type EventPair struct {
	Event1 MatchEvent `json:"e1"`
	Event2 MatchEvent `json:"e2"`
}
