package types

import (
	h_ "coolsim/internal/components/hospitality"
	i_ "coolsim/internal/components/identity"
)

type EnvironmentBuilder func() Environment

type Environment interface {
	i_.Identifiable
	h_.Hospitable
}
