package observers

import (
	"leagueapi.com.br/brain/src/enum"
)

type CreatePlayerObserver struct {
	ID       string
	EventKey enum.EventsEnum
}

func (c *CreatePlayerObserver) Update(observerUpdate chan string) {
	observerUpdate <- "PLAYER " + c.ID + "FOI NOTIFICADO DO CADASTRO\n"
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
