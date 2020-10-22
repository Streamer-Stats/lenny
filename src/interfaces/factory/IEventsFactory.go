package interfacefactory

import (
	"fmt"

	factoryevents "leagueapi.com.br/brain/src/factory/events"
	"leagueapi.com.br/brain/src/interfaces"
)

type IEventsFactory interface {
	CreateEvent(id string) interfaces.IEvents
}

func GetEventsFactory(event string) (IEventsFactory, error) {
	if event == "registerUser" {
		return factoryevents.NewCreateUserEventFactory(), nil
	}

	return nil, fmt.Errorf("Wrong event type passed")
}
