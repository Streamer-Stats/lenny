package entities

// Brain is a watcher
type Brain struct {
	events *Events
}

// NewBrain constructor
func NewBrain() *Brain {
	return &Brain{
		events: NewEvents().RegisterEvents(),
	}
}
