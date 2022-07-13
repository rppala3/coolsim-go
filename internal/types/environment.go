package types

import (
	h_ "coolsim/internal/commons/hospitality"
	i_ "coolsim/internal/commons/identity"
)

type Environment interface {
	i_.Identifiable
	h_.Hospitable
	GetAPerception() Perception
}
