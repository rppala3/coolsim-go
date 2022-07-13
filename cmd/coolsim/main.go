package main

import (
	agentPkg "coolsim/internal/singletons/agent"

	"go.uber.org/zap"
)

var logger *zap.Logger

func main() {
	thicks := 2

	InitLogger()
	defer logger.Sync()

	agentMgr := agentPkg.GetInstance()
	p01 := agentMgr.SpawnAPerson()

	for i := 0; i < thicks; i++ {
		p01.Perceive()
	}
}

// @see https://developpaper.com/using-zap-log-library-in-go-language-project-translation/
func InitLogger() {
	// @todo logger configuration
	logger, _ = zap.NewProduction()
	zap.ReplaceGlobals(logger)
}
