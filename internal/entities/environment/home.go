package environment

import (
	p_ "coolsim/internal/entities/perception"
	t_ "coolsim/internal/types"
)

type Home struct {
	*Place
	// state map[string]EnvState // env state map
}

func NewHome() *Home {
	return &Home{
		NewPlace(),
	}
}

func (location *Home) GetAPerception() t_.Perception {
	return p_.NewEightOClock(
	// @todo pass envState
	)
}
