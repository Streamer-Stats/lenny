package events

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"leagueapi.com.br/brain/src/broadcast"
	"leagueapi.com.br/brain/src/enum"
	"leagueapi.com.br/brain/src/infrastructure"
	IWatcher "leagueapi.com.br/brain/src/interfaces/watcherInterfaces"
	"leagueapi.com.br/brain/src/observers"
)

// CreatePlayerEvent event of player criation
type CreatePlayerEvent struct {
	ID           enum.EventsEnum
	ObserverList []*observers.PlayerObserver
	RabbitMQ     *infrastructure.RabbitMQ
}

// Register register a new type event player
func (e *CreatePlayerEvent) Register(observer IWatcher.IObserverEvents) {
	e.ObserverList = append(e.ObserverList, observer.(*observers.PlayerObserver))
}

// Deregister deregister a new type event player
func (e *CreatePlayerEvent) Deregister(observer IWatcher.IObserverEvents) {}

// NotifyAll notify all watcher
func (e *CreatePlayerEvent) NotifyAll(observerUpdate chan broadcast.Player) {
	for _, observers := range e.ObserverList {
		observers.Update(observerUpdate, observers.Stats)
	}
}

func (e *CreatePlayerEvent) notifyOnce(observerUpdate chan broadcast.Player, playerID string, newStats *broadcast.Player) {
	for _, observers := range e.ObserverList {
		if observers.GetID() == playerID {
			observers.UpdateStats(newStats)
			observers.Update(observerUpdate, observers.Stats)
		}

	}
}

// ReturnID return eventID
func (e *CreatePlayerEvent) ReturnID() enum.EventsEnum {
	return e.ID
}

func (e *CreatePlayerEvent) CheckChanges(observerUpdate chan broadcast.Player) {
	for range time.Tick(time.Second * 5) {
		log.Println("event new player launch.... ")
		for playerData := range e.RabbitMQ.Consume("brainPlayerData") {
			var newPlayerStats *broadcast.Player
			json.Unmarshal(playerData.Body, newPlayerStats)
			// e.notifyOnce(observerUpdate, newPlayerStats.Username, newPlayerStats)
			fmt.Println(newPlayerStats)
		}
	}
}

// NewCreatePlayerEvent constructor
func NewCreatePlayerEvent() *CreatePlayerEvent {
	return &CreatePlayerEvent{
		ID:       enum.CreateNewPlayer,
		RabbitMQ: infrastructure.NewRabbitMQ(),
	}
}
