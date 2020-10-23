package syncronize

import (
	"os"
	"os/signal"
)

type Syncronize struct {
	Interrupt      chan os.Signal
	Done           chan struct{}
	NewMessage     chan string
	ObserverUpdate chan string
}

func (s *Syncronize) InterruptNotify() {
	signal.Notify(s.Interrupt, os.Interrupt)
}
func NewSyncronize() *Syncronize {
	return &Syncronize{
		Interrupt:      make(chan os.Signal, 1),
		Done:           make(chan struct{}),
		NewMessage:     make(chan string),
		ObserverUpdate: make(chan string),
	}
}
