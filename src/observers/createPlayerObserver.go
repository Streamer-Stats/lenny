package observers

import (
	"fmt"

	"leagueapi.com.br/brain/src/enum"
)

type CreatePlayerObserver struct {
	ID       string
	EventKey enum.EventsEnum
}

func (c *CreatePlayerObserver) Update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.ID, itemName)
}

func (c *CreatePlayerObserver) GetID() string {
	return c.ID
}

func NewCreatePlayerObserver(id string) *CreatePlayerObserver {
	return &CreatePlayerObserver{
		ID:       id,
		EventKey: enum.CreateNewPlayer,
	}
}
