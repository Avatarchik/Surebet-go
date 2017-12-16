package common

type Account struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Range struct {
	From, To int
}

type Accounts struct {
	V    []Account
	Rng  Range
	Path string
}

func (a *Accounts) Load() error {
	var accounts []Account
	if err := LoadJson(a.Path, &accounts); err != nil {
		return err
	}
	a.V = accounts[a.Rng.From:a.Rng.To]
	return nil
}
