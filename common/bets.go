package common

type factor float32

type condFactor struct {
	Cond  float32
	Value factor
}

type Bets struct {
	Part  int
	O1    []factor
	OX    []factor
	O2    []factor
	O1X   []factor
	O12   []factor
	OX2   []factor
	TO    []condFactor
	TU    []condFactor
	IndTO []condFactor
	IndTU []condFactor
	Hand1 []condFactor
	Hand2 []condFactor
}

type Event struct {
	Team1 string
	Team2 string
	Parts []Bets
}

type Sport map[string][]Event

type Bookmaker map[string]Sport

type AllEvents map[string]Bookmaker
