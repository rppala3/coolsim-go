package types

import (
	h_ "coolsim/internal/commons/hospitality"
	i_ "coolsim/internal/commons/identity"
	s_ "coolsim/internal/commons/spawner"
)

type EnvironmentBuilder func() Environment

type Environment interface {
	i_.Identifiable
	h_.Hospitable
	s_.Spawnable
	// s_.Spawnable[Event]
}
