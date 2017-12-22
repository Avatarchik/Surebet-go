package types

type Account struct {
	Login    string `json:"login"`
	Password string `json:"pass"`
}

type Range struct {
	From int `json:"from"`
	To   int `json:"to"`
}

type Accounts struct {
	v   []Account
	all []Account
}

func (a *Accounts) SetRange(rng *Range) {
	a.v = a.all[rng.From:rng.To]
}

func (a *Accounts) Set(accounts []Account) {
	a.all = accounts
	a.v = a.all
}

func (a *Accounts) Size() int {
	return len(a.v)
}

func (a *Accounts) Values() []Account {
	return a.v
}
