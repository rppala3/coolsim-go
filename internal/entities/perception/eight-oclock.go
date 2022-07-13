package perception

import (
	a_ "coolsim/internal/entities/action"
	t_ "coolsim/internal/types"
)

type EightOClock struct {
	actions []t_.Action
}

// @todo receive envState
func NewEightOClock() *EightOClock {
	// @todo get actions from config file ?
	// @todo filter envState (immutable)
	ac01 := a_.NewGoOffice()
	return &EightOClock{
		actions: [](t_.Action){ac01},
	}
}

func (perception *EightOClock) Description() string {
	return "It's 8:00"
}

func (perception *EightOClock) GetActions() []t_.Action {
	return perception.actions
}
