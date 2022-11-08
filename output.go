package mcmf

type Output struct {
	Solutions []Solution `json:"solution"`
}

type Solution struct {
	Matches []Match `json:"match"`
}

type Match struct {
	From   int   `json:"from"`
	To     int   `json:"to"`
	Amount int64 `json:"amount"`
	Cost   int64 `json:"cost"`
}
