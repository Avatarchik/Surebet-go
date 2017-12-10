package common

type Bet struct {
	Bookmaker map[string]map[string]int `json:"bookmaker"`
}
