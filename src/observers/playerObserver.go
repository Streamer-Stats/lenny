package observers

import (
	"leagueapi.com.br/brain/src/broadcast"
	"leagueapi.com.br/brain/src/enum"
)

// CreatePlayerObserver create player observer
type PlayerObserver struct {
	ID       string
	EventKey enum.EventsEnum
	Stats    *broadcast.Player
}

// Update chanel
func (c *PlayerObserver) Update(observerUpdate chan broadcast.Player, player *broadcast.Player) {
	observerUpdate <- *player
}

func (c *PlayerObserver) GetID() string {
	return c.ID
}

func NewPlayerObserver(id string) *PlayerObserver {
	return &PlayerObserver{
		ID:       id,
		EventKey: enum.CreateNewPlayer,
		Stats:    broadcast.NewPlayer(id, 0.0, false),
	}
}
