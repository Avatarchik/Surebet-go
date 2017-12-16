package common

type factor float32

type condFactor struct {
	Cond  float32
	Value factor
}

type factors []factor
type condFactors []condFactor

type Bets struct {
	Part  int
	O1    factors
	OX    factors
	O2    factors
	O1X   factors
	O12   factors
	OX2   factors
	TO    condFactors
	TU    condFactors
	IndTO condFactors
	IndTU condFactors
	Hand1 condFactors
	Hand2 condFactors
}

type Event struct {
	Team1 string
	Team2 string
	Parts []Bets
}

type Events []Event

type Sports struct {
	Soccer Events
	Tennis Events
	Hockey Events
	Basket Events
	Volley Events
}

type Bookmakers struct {
	Fonbet   Sports
	Olimp    Sports
	Marathon Sports
}
