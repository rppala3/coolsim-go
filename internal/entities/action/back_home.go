package action

import t_ "coolsim/internal/types"

type BackHome struct {
}

func NewBackHome() *BackHome {
	return &BackHome{}
}

func (action *BackHome) Name() string {
	return "BackHome"
}

func (action *BackHome) Description() string {
	return "comes back to home"
}

func (action *BackHome) Do(subject t_.Agent) {
	destination := subject.GetHome()
	subject.MoveIn(destination)
}
