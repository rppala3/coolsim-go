package types

import i_ "coolsim/internal/components/identity"

type AgentBuilder func() Agent

type Agent interface {
	i_.Identifiable
	GetHome() Environment
	GetWorkPlace() Environment
	MoveIn(destination Environment)
	Perceive()
	Perform(actions []Action)
}
