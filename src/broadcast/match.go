package broadcast

type Match struct {
	Live bool
}

func NewMatch(live bool) *Match {
	return &Match{
		Live: live,
	}
}
