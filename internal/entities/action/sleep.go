package action

import t_ "coolsim/internal/types"

type Sleep struct {
}

func NewSleep() *Sleep {
	return &Sleep{}
}

func (action *Sleep) Description() string {
	return "is sleeping"
}

func (action *Sleep) Do(subject t_.Agent) {
}
