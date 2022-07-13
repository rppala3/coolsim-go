package environment

import (
	h_ "coolsim/internal/commons/hospitality"
	i_ "coolsim/internal/commons/identity"
)

type Place struct {
	*i_.Identity
	*h_.Hospitality
}

func NewPlace() *Place {
	return &Place{
		i_.NewIdentity(),
		h_.NewHospitality(),
	}
}
