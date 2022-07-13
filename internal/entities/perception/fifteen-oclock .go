package perception

import (
	a_ "coolsim/internal/entities/action"
	t_ "coolsim/internal/types"
)

type FifteenOClock struct {
	actions []t_.Action
}

// @todo receive envState
func NewFifteenOClock() *FifteenOClock {
	// @todo get actions from config file ?
	// @todo filter envState (immutable)
	ac01 := a_.NewBackHome()
	return &FifteenOClock{
		actions: [](t_.Action){ac01},
	}
}

func (perception *FifteenOClock) Description() string {
	return "It's 15:00"
}

func (perception *FifteenOClock) GetActions() []t_.Action {
	return perception.actions
}
