package ticker

import (
	e_ "coolsim/internal/entities/event"
	t_ "coolsim/internal/types"
	"coolsim/internal/utils/random"
	"time"

	"fmt"

	b_ "github.com/rppala3/broadcasting"
)

type Ticker struct {
	ticks          int
	delay          time.Duration
	eventFactories []t_.EventFactory // @todo replace with a PriorityQueue
	// https://github.com/emirpasic/gods#priorityqueue
	eventBroadcast *b_.Broadcaster[t_.Event]
}

func NewTicker(ticks int, delay time.Duration, eventBroadcast *b_.Broadcaster[t_.Event]) *Ticker {
	if ticks < 0 {
		ticks = 0
	}
	// @todo load all events from config file
	events := GetEventFactories()
	return &Ticker{
		ticks,
		delay,
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
		index := random.Int(0, eventLen)

		// Create randomized event
		eventFactory := ticker.eventFactories[index]
		event := eventFactory()

		// - DEBUG ---------------------------------------------------------
		// fmt.Printf("Extracted [%d] '%s'\n", index, event.Name())
		fmt.Printf("SEND N.tick: %d Event: %s\n", tick, event.Description())
		// - DEBUG ---------------------------------------------------------

		// Send event
		ticker.eventBroadcast.Send(event)

		time.Sleep(ticker.delay) // slow down the execution
	}
}

//
// Private methods
//
