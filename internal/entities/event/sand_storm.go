package event

import (
	t_ "coolsim/internal/types"
	"coolsim/internal/utils/random"
	"fmt"
)

type SandStorm struct {
	intensity int
	// durantion: // ?
}

// @todo receive envState
// func NewSandStorm() *SandStorm {
// 	return &SandStorm{
// 		random.PullOutInt(1, 100),
// 	}
// }

func NewSandStorm() t_.Event {
	return &SandStorm{
		random.PullOutInt(1, 100),
	}
}

func (event *SandStorm) Description() string {
	return fmt.Sprintf("sand storm with %d intensity", event.intensity)
}

func (event *SandStorm) Effects() []t_.Effect {
	return nil
}
