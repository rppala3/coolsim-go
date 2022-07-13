package environment

import (
	p_ "coolsim/internal/entities/perception"
	t_ "coolsim/internal/types"
)

type Office struct {
	*Place
	// state map[string]EnvState // env state map
}

func NewOffice() *Office {
	return &Office{
		NewPlace(),
	}
}

func (location *Office) GetAPerception() t_.Perception {
	return p_.NewFifteenOClock(
	// @todo pass envState
	)
}
