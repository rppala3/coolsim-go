package environment

import (
	s_ "coolsim/internal/commons/spawner"
	t_ "coolsim/internal/types"

	"sync"

	b_ "github.com/rppala3/broadcast-channel"
)

type EnvironmentManager struct {
	environments   map[string]t_.Environment
	eventBroadcast *b_.Broadcaster
}

var once sync.Once
var waitGroup sync.WaitGroup
var singleInstance *EnvironmentManager

func GetInstance() *EnvironmentManager {
	if singleInstance == nil {
		once.Do(func() {
			eventBroadcast := b_.NewBroadcaster(0)
			singleInstance = newEnvironmentManager(eventBroadcast)
		})
	}
	return singleInstance
}

func (manager *EnvironmentManager) GetPlace(uid string) t_.Environment {
	return manager.environments[uid]
}

func (manager *EnvironmentManager) SpawnAPlace(envBuilder t_.EnvironmentBuilder, ready s_.AckChannel) t_.Environment {
	location := envBuilder()
	location.Spawn(manager.eventBroadcast.Listen(), ready, &waitGroup)
	waitGroup.Add(1)
	manager.add(location)
	return location
}

func (manager *EnvironmentManager) GetBroadcaster() *b_.Broadcaster {
	return manager.eventBroadcast
}

func (manager *EnvironmentManager) GetWaitGroup() *sync.WaitGroup {
	return &waitGroup
}

//
// Private methods
//

func newEnvironmentManager(eventBroad *b_.Broadcaster) *EnvironmentManager {
	return &EnvironmentManager{
		make(map[string]t_.Environment),
		eventBroad,
	}
}

func (manager *EnvironmentManager) add(env t_.Environment) {
	manager.environments[env.GetUID()] = env
}
