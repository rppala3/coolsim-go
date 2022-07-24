package environment

import (
	// p_ "coolsim/internal/entities/perception"
	t_ "coolsim/internal/types"
)

type Gym struct {
	*Place
	// state map[string]EnvState // env state map
}

func NewGym() t_.Environment {
	// func NewGym() *Gym {
	return &Gym{
		NewPlace(),
	}
}
