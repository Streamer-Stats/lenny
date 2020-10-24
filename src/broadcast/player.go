package broadcast

// Player i a struct with data
type Player struct {
	Username string
	Damage   float64
	Match    *Match
	Account  *Account
}

func NewPlayer(username string, damage float64, live bool) *Player {
	return &Player{
		Username: username,
		Damage:   damage,
		Match:    NewMatch(live),
		Account:  NewAccount(),
	}
}
