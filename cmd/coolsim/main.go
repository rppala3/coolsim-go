package main

import (
	s_ "coolsim/internal/commons/spawner"
	e_ "coolsim/internal/entities/environment"
	envMgr "coolsim/internal/singletons/environment"
	ticker "coolsim/internal/ticker"
	"fmt"

	"go.uber.org/zap"
)

var logger *zap.Logger

func main() {
	// var routines sync.WaitGroup
	ticks := 5
	envs := 1

	fmt.Println("Init")
	fmt.Println(" - logger")
	initLogger()
	defer logger.Sync()

	fmt.Println(" - environment")
	spawnEnvironments(envs)

	// @todo spawn all agents

	fmt.Println("Start simulation")

	eventBroadcast := envMgr.GetInstance().GetBroadcaster()
	tckr := ticker.NewTicker(ticks, eventBroadcast)
	tckr.Start()
	eventBroadcast.Discard()

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

func spawnEnvironments(nEnvs int) {
	ready := make(s_.AckChannel)
	envMgr := envMgr.GetInstance()
	for i := 1; i <= nEnvs; i++ {
		office := envMgr.SpawnAPlace(e_.NewOffice, ready)
		<-ready
		fmt.Printf(" * Office %d spawned %s\n", i, office.GetUID())
	}
}
