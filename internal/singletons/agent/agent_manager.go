package agent

import (
	a_ "coolsim/internal/entities/agent"
	"coolsim/internal/singletons/environment"
	t_ "coolsim/internal/types"
	"sync"
)

type AgentManager struct {
	agents map[string]t_.Agent
}

var once sync.Once
var singleInstance *AgentManager

func GetInstance() *AgentManager {
	if singleInstance == nil {
		once.Do(func() {
			singleInstance = newAgentManager()
		})
	}
	return singleInstance
}

func (manager *AgentManager) SpawnAPerson() *a_.Person {
	envMgr := environment.GetInstance() // @todo should be decoupled
	agent := a_.NewPerson(
		envMgr.SpawnAnHome(),
		envMgr.SpawnAnOffice(),
	)
	manager.add(agent)
	return agent
}

func (manager *AgentManager) Get(uid string) t_.Agent {
	return manager.agents[uid]
}

//
// Private methods
//

func newAgentManager() *AgentManager {
	return &AgentManager{
		make(map[string]t_.Agent),
	}
}

func (manager *AgentManager) add(agent t_.Agent) {
	manager.agents[agent.GetUID()] = agent
}
