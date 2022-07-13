package environment

import (
	// p_ "coolsim/internal/entities/perception"
	t_ "coolsim/internal/types"
)

type Gym struct {
	*Place
	// state map[string]EnvState // env state map
}

func NewGym() *Gym {
	return &Gym{
		NewPlace(),
	}
}

func (location *Gym) GetAPerception() t_.Perception {
	return nil // @todo
}
