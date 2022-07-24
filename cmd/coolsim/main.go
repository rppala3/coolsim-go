package main

import (
	m_ "coolsim/internal/commons/messaging"
	e_ "coolsim/internal/entities/environment"
	envMgr "coolsim/internal/singletons/environment"
	"coolsim/internal/ticker"
	"fmt"
	"time"

	"go.uber.org/zap"
)

var logger *zap.Logger

func main() {
	delay := time.Nanosecond
	// delay := time.Second
	ticks := 10
	envs := 5

	fmt.Println("Init")
	fmt.Println(" - logger")
	initLogger()
	defer logger.Sync()

	fmt.Println(" - environments")
	spawnEnvironments(envs)

	// @todo
	// fmt.Println(" - agents")
	// spawnAgents(envs)

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

func spawnEnvironments(nEnvs int) {
	ready := make(m_.AckChannel)
	envMgr := envMgr.GetInstance()
	for i := 1; i <= nEnvs; i++ {
		office := envMgr.SpawnAPlace(e_.NewOffice, ready)
		<-ready
		fmt.Printf(" * Office %d spawned %s\n", i, office.GetUID())
	}
}
