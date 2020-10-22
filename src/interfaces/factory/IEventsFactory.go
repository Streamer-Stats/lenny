package interfacefactory

import (
	"fmt"

	"leagueapi.com.br/brain/src/enum"
	IWatcher "leagueapi.com.br/brain/src/interfaces/watcherInterfaces"

	factoryevents "leagueapi.com.br/brain/src/factory/events"
)

// IEventsFactory interface of abstract factory: events
type IEventsFactory interface {
	CreateEvent() IWatcher.IEvents
}

// GetEventsFactory static method to get new factory
func GetEventsFactory(event enum.EventsEnum) (IEventsFactory, error) {

	switch event {
	case enum.CreateNewPlayer:
		return factoryevents.NewCreateUserEventFactory(), nil
	default:
		return nil, fmt.Errorf("Wrong event type passed")
	}

}
