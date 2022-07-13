package environment

import (
	e_ "coolsim/internal/entities/environment"
	t_ "coolsim/internal/types"
	"sync"
)

type EnvironmentManager struct {
	environments map[string]t_.Environment
}

var once sync.Once
var singleInstance *EnvironmentManager

func GetInstance() *EnvironmentManager {
	if singleInstance == nil {
		once.Do(func() {
			singleInstance = newEnvironmentManager()
		})
	}
	return singleInstance
}

func (manager *EnvironmentManager) SpawnAnHome() *e_.Home {
	location := e_.NewHome()
	manager.add(location)
	return location
}

func (manager *EnvironmentManager) SpawnAnOffice() *e_.Office {
	location := e_.NewOffice()
	manager.add(location)
	return location
}

func (manager *EnvironmentManager) Get(uid string) t_.Environment {
	return manager.environments[uid]
}

//
// Private methods
//

func newEnvironmentManager() *EnvironmentManager {
	return &EnvironmentManager{
		make(map[string]t_.Environment),
	}
}

func (manager *EnvironmentManager) add(env t_.Environment) {
	manager.environments[env.GetUID()] = env
}
