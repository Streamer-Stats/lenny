package syncronize

import (
	"os"
	"os/signal"

	"leagueapi.com.br/brain/src/broadcast"
)

type Syncronize struct {
	Interrupt  chan os.Signal
	Done       chan struct{}
	NewMessage chan string
	BroadCast  chan broadcast.Player
}

func (s *Syncronize) InterruptNotify() {
	signal.Notify(s.Interrupt, os.Interrupt)
}
func NewSyncronize() *Syncronize {
	return &Syncronize{
		Interrupt:  make(chan os.Signal, 1),
		Done:       make(chan struct{}),
		NewMessage: make(chan string),
		BroadCast:  make(chan broadcast.Player),
	}
}
