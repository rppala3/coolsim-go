package action

import t_ "coolsim/internal/types"

type Clean struct {
}

func NewClean() *Clean {
	return &Clean{}
}

func (action *Clean) Name() string {
	return "Clean"
}

func (action *Clean) Description() string {
	return "is cleaning"
}

func (action *Clean) Do(subject t_.Agent) {
	// return new env state
}
