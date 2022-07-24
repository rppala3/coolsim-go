package environment

import (
	t_ "coolsim/internal/types"
)

type Office struct {
	*Place
	// state map[string]EnvState // env state map
}

func NewOffice() t_.Environment {
	// func NewOffice() *Office {
	return &Office{
		NewPlace(),
	}
}
