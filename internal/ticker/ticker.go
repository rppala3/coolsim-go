package ticker

import (
	e_ "coolsim/internal/entities/event"
	t_ "coolsim/internal/types"
	"coolsim/internal/utils/random"

	"fmt"

	b_ "github.com/rppala3/broadcast-channel"
)

type Ticker struct {
	ticks          int
	eventFactories []t_.EventFactory // @todo replace with a PriorityQueue
	// https://github.com/emirpasic/gods#priorityqueue
	eventBroadcast *b_.Broadcaster
}

func NewTicker(ticks int, eventBroadcast *b_.Broadcaster) *Ticker {
	if ticks < 0 {
		ticks = 0
	}
	// @todo load all events from config file
	events := GetEventFactories()
	return &Ticker{
		ticks,
		events,
		eventBroadcast,
	}
}

// @todo pass events file?
func GetEventFactories() []t_.EventFactory {
	return []t_.EventFactory{
		e_.NewSandStorm,
	}
}

func (ticker *Ticker) Start() {
	eventLen := len(ticker.eventFactories)
	for tick := 1; tick <= ticker.ticks; tick++ {
		// Extract a random index
		index := random.PullOutInt(0, eventLen)

		// Create randomized event
		eventFactory := ticker.eventFactories[index]
		event := eventFactory()

		// - DEBUG ---------------------------------------------------------
		fmt.Printf("SEND N.tick: %d Event: %s\n", tick, event.Description())

		// Send event
		ticker.eventBroadcast.Send(event)

		// time.Sleep(time.Second) // slow down the execution
	}
}

//
// Private methods
//
