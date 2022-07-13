package hospitality

import i_ "coolsim/internal/commons/identity"

type Guest interface {
	i_.Identifiable
	// @note
	// Guest is an subset of Agent, BUT Environment implements the same interface
	// how can we prevent this?
	//
	// an Agent is     a Guest
	// an   Env is NOT a Guest
}
