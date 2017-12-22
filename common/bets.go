package common

type CondBet struct {
	Cond Factor
	V1   Factor
	V2   Factor
}

type Factor float64

type CondBets []CondBet

type Bets struct {
	Part      int
	O1        Factor
	OX        Factor
	O2        Factor
	O1X       Factor
	O12       Factor
	OX2       Factor
	Total     CondBets
	IndTotal1 CondBets
	IndTotal2 CondBets
	Hand      CondBets
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
	Fonbet   *Sports
	Olimp    *Sports
	Marathon *Sports
}

type Teams struct {
	Team1 string
	Team2 string
}

func (c *CondBet) isNotEmpty() bool {
	return c.V1 != 0 && c.V2 != 0
}

func (e *Event) isNotEmpty() bool {
	return len(e.Parts) != 0
}

func (c *CondBets) AppendNotEmpty(elems ...CondBet) {
	for _, elem := range elems {
		if elem.isNotEmpty() {
			*c = append(*c, elem)
		}
	}
}

func (e *Events) AppendNotEmpty(elems ...Event) {
	for _, elem := range elems {
		if elem.isNotEmpty() {
			*e = append(*e, elem)
		}
	}
}
