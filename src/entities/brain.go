package entities

import (
	"fmt"
	"time"

	"leagueapi.com.br/brain/src/events"
)

// Brain is a watcher
type Brain struct {
	events  *Events
	watcher *Watcher
}

func (brain *Brain) createNewPlayer(username string, event *events.CreatePlayerEvent) {
	player := brain.watcher.CreateNewPlayerWatcher(username, event.ID)
	event.Register(player)
	fmt.Println(username, " IS NOW WATCHING CREATE NEW PLAYER EVENT")
}

// Handle handle the decisions
func (brain *Brain) Handle() {
	newPlayerEvent := brain.events.GetCreatePlayerEvent()
	brain.createNewPlayer("BANOFFE", newPlayerEvent)
	brain.createNewPlayer("Frango SafaJhin", newPlayerEvent)
	brain.createNewPlayer("Zoiao", newPlayerEvent)
	time.Sleep(5 * time.Second)
	newPlayerEvent.NotifyAll()
}

// NewBrain constructor
func NewBrain() *Brain {
	return &Brain{
		events:  NewEvents().RegisterEvents(),
		watcher: NewWatcher(),
	}
}
