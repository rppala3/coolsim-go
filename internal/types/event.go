package types

type EventFactory func() Event

type Event interface {
	Name() string
	Description() string
	// Effects() []Effect
}
