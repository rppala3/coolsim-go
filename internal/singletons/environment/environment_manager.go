package environment

import (
	m_ "coolsim/internal/commons/messaging"
	t_ "coolsim/internal/types"
	"fmt"

	"sync"

	b_ "github.com/rppala3/broadcasting"
)

type EnvironmentManager struct {
	environments   map[string]t_.Environment
	eventBroadcast *b_.Broadcaster[t_.Event]
}

var once sync.Once
var waitGroup sync.WaitGroup
var singleInstance *EnvironmentManager

func GetInstance() *EnvironmentManager {
	if singleInstance == nil {
		once.Do(func() {
			eventBroadcast := b_.NewBroadcaster[t_.Event](0)
			singleInstance = newEnvironmentManager(eventBroadcast)
		})
	}
	return singleInstance
}

func (manager *EnvironmentManager) GetPlace(uid string) t_.Environment {
	return manager.environments[uid]
}

func (manager *EnvironmentManager) GetBroadcaster() *b_.Broadcaster[t_.Event] {
	return manager.eventBroadcast
}

func (manager *EnvironmentManager) GetWaitGroup() *sync.WaitGroup {
	return &waitGroup
}

func (manager *EnvironmentManager) SpawnAPlace(
	envBuilder t_.EnvironmentBuilder,
	ready m_.AckChannel,
) t_.Environment {
	place := envBuilder()
	manager.add(place)

	eventListener := manager.eventBroadcast.Listen()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		ready <- true
		for {
			event := <-eventListener.Channel()
			if event == nil {
				eventListener.Discard()
				fmt.Printf("Place %s closed\n", place.GetUID()) // DEBUG
				break
			}
			// - DEBUG ---------------------------------------------------------------
			fmt.Printf("RECEIVED _%s_ Event: '%s' \n",
				place.GetUID(),
				event.Description(),
			)
			// - DEBUG ---------------------------------------------------------------
		}
	}()

	return place
}

//
// Private methods
//

func newEnvironmentManager(eventBroad *b_.Broadcaster[t_.Event]) *EnvironmentManager {
	return &EnvironmentManager{
		make(map[string]t_.Environment),
		eventBroad,
	}
}

func (manager *EnvironmentManager) add(env t_.Environment) {
	manager.environments[env.GetUID()] = env
}
