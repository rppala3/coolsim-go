package action

import t_ "coolsim/internal/types"

type Rest struct {
}

func NewRest() *Rest {
	return &Rest{}
}

func (action *Rest) Description() string {
	return "is resting"
}

func (action *Rest) Do(subject t_.Agent) {
}
