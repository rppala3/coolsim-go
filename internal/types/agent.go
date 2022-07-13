package types

import i_ "coolsim/internal/commons/identity"

type Agent interface {
	i_.Identifiable
	GetHome() Environment
	GetWorkPlace() Environment
	MoveIn(destination Environment)
	Perceive()
	Perform(actions []Action)
}
