package environment

import (
	h_ "coolsim/internal/commons/hospitality"
	i_ "coolsim/internal/commons/identity"
	s_ "coolsim/internal/commons/spawner"
	t_ "coolsim/internal/types"
	"fmt"
	"sync"

	b_ "github.com/rppala3/broadcast-channel"
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

func (place *Place) Spawn(
	eventListener *b_.Listener,
	ready s_.AckChannel,
	waitGroup *sync.WaitGroup,
) {
	go func() {
		defer (*waitGroup).Done()
		ready <- true
		for {
			msg, open := <-eventListener.Channel()
			if !open {
				fmt.Println("eventListener closed") // DEBUG
				fmt.Println("Place closed")         // DEBUG
				break
			}
			event := msg.(t_.Event)
			// DEBUG
			fmt.Printf("RECEIVED _%s_ Event: '%s' \n",
				place.GetUID(),
				event.Description(),
			)
			// DEBUG
		}
	}()
}
