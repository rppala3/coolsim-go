package environment

import (
	h_ "coolsim/internal/components/hospitality"
	i_ "coolsim/internal/components/identity"
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
