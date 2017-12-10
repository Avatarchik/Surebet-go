package common

type Account struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Accounts []Account

type Range struct {
	From, To int
}

func (a *Accounts) LoadRange(filename string, rng Range) error {
	if err := LoadJson(filename, a); err != nil {
		return err
	}
	*a = (*a)[rng.From:rng.To]
	return nil
}
