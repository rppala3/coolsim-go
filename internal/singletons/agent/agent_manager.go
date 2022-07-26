package agent

import (
	t_ "coolsim/internal/types"
	"sync"
)

type AgentManager struct {
	agents map[string]t_.Agent
}

var once sync.Once
var waitGroup sync.WaitGroup
var singleInstance *AgentManager

func GetInstance() *AgentManager {
	if singleInstance == nil {
		once.Do(func() {
			singleInstance = newAgentManager()
		})
	}
	return singleInstance
}

func (manager *AgentManager) GetAgent(uid string) t_.Agent {
	return manager.agents[uid]
}

func (manager *AgentManager) SpawnAnAgent(agentBuilder t_.AgentBuilder) t_.Agent {
	// envMgr := environment.GetInstance() // @todo should be decoupled
	// agent := a_.NewPerson(
	// 	envMgr.SpawnAnHome(),
	// 	envMgr.SpawnAnOffice(),
	// )
	// manager.add(agent)
	// return agent

	// @todo
	return nil
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
