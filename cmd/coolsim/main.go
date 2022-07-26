package main

import (
	a_ "coolsim/internal/entities/agent"
	agentMgr "coolsim/internal/singletons/agent"
	envMgr "coolsim/internal/singletons/environment"
	"coolsim/internal/ticker"
	"fmt"
	"time"

	"go.uber.org/zap"
)

var logger *zap.Logger

func main() {
	// delay := time.Nanosecond
	delay := time.Second
	ticks := 10
	numOfAgents := 1
	numOfEnvs := 5

	fmt.Println("Init")
	fmt.Println(" - logger")
	initLogger()
	defer logger.Sync()

	fmt.Println(" - agents")
	initAgents(numOfAgents)

	fmt.Println(" - environments")
	initEnvironments(numOfEnvs)

	fmt.Println("Start simulation")

	eventBroadcast := envMgr.GetInstance().GetBroadcaster()
	tckr := ticker.NewTicker(ticks, delay, eventBroadcast)
	tckr.Start()
	eventBroadcast.Close()

	envMgr.GetInstance().GetWaitGroup().Wait()
	fmt.Println("End simulation")
}

//
// Private methods
//

// @see https://developpaper.com/using-zap-log-library-in-go-language-project-translation/
func initLogger() {
	// @todo logger configuration
	logger, _ = zap.NewProduction()
	zap.ReplaceGlobals(logger)
}

func initAgents(numOfAgents int) {
	agentMgr := agentMgr.GetInstance()
	for i := 1; i <= numOfAgents; i++ {
		agent := agentMgr.SpawnAnAgent(a_.NewPerson)
		fmt.Printf(" * Agent %d spawned %s\n", i, agent.GetUID())
	}
}

func initEnvironments(numOfEnvs int) {
	// ready := make(m_.AckChannel)
	// envMgr := envMgr.GetInstance()
	// for i := 1; i <= numOfEnvs; i++ {
	// 	office := envMgr.SpawnAPlace(e_.NewOffice, ready)
	// 	<-ready
	// 	fmt.Printf(" * Office %d spawned %s\n", i, office.GetUID())
	// }
}
