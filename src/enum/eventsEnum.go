package enum

//EventsEnum enum of events
type EventsEnum int

const (
	// CreateNewPlayer call create player event
	CreateNewPlayer EventsEnum = iota
)

func (e EventsEnum) String() string {
	return [...]string{"CreateNewPlayer"}[e]
}
