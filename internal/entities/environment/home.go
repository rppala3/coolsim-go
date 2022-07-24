package environment

import (
	t_ "coolsim/internal/types"
)

type Home struct {
	*Place
	// state map[string]EnvState // env state map
}

func NewHome() t_.Environment {
	// func NewHome() *Home {
	return &Home{
		NewPlace(),
	}
}
