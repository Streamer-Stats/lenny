package observers

import (
	"encoding/json"

	"leagueapi.com.br/brain/src/broadcast"
	"leagueapi.com.br/brain/src/enum"
	"leagueapi.com.br/brain/src/infrastructure"
)

// CreatePlayerObserver create player observer
type PlayerObserver struct {
	ID       string
	EventKey enum.EventsEnum
	Stats    *broadcast.Player
	Rabbitmq *infrastructure.RabbitMQ
}

func (c *PlayerObserver) UpdateStats(newStats *broadcast.Player) {
	c.Stats = newStats
}

// Update chanel
func (c *PlayerObserver) Update(observerUpdate chan broadcast.Player, player *broadcast.Player) {
	observerUpdate <- *player
}

func (c *PlayerObserver) GetID() string {
	return c.ID
}

func (c *PlayerObserver) CheckIfPlayerExist() {
	jsonStringfy, _ := json.Marshal(c.Stats)
	c.Rabbitmq.Publish(string(jsonStringfy), "newPlayer")
}

func NewPlayerObserver(id string) *PlayerObserver {
	return &PlayerObserver{
		ID:       id,
		EventKey: enum.CreateNewPlayer,
		Stats:    broadcast.NewPlayer(id, 0.0, false),
		Rabbitmq: infrastructure.NewRabbitMQ(),
	}
}
