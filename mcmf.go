package mcmf

type Input struct {
	Ids      []int     `json:"ids"`
	Costs    [][]int64 `json:"costs"`
	Problems []Problem `json:"problems"`
}

type Output struct {
	Solutions []Solution `json:"solution"`
}

type Problem struct {
	Sources []Edge `json:"sources"`
	Sinks   []Edge `json:"sinks"`
}

type Edge struct {
	Id       int   `json:"id"`
	Capacity int64 `json:"capacity"`
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

func Solve(input Input) (*Solution, error) {

	return nil, nil
}
