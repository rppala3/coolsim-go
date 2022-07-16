package types

type EventFactory func() Event

type Event interface {
	Description() string
	// Effects() []Effect
}
