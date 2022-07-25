package action

import t_ "coolsim/internal/types"

type GoOffice struct {
}

func NewGoOffice() *GoOffice {
	return &GoOffice{}
}

func (action *GoOffice) Name() string {
	return "GoOffice"
}

func (action *GoOffice) Description() string {
	return "goes to office"
}

func (action *GoOffice) Do(subject t_.Agent) {
	destination := subject.GetWorkPlace()
	subject.MoveIn(destination)
}
