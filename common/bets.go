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

type Sports struct {
	Soccer []Event
	Tennis []Event
	Hockey []Event
	Basket []Event
	Volley []Event
}

type Bookmakers struct {
	Fonbet   Sports
	Olimp    Sports
	Marathon Sports
}
