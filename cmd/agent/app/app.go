package app

import (
	"github.com/dennishilgert/apollo/cmd/agent/config"
	"github.com/dennishilgert/apollo/internal/app/agent"
	"github.com/dennishilgert/apollo/pkg/concurrency/runner"
	"github.com/dennishilgert/apollo/pkg/logger"
	"github.com/dennishilgert/apollo/pkg/signals"
	"github.com/joho/godotenv"
)

var log = logger.NewLogger("apollo.agent")

func Run() {
	// load environment variables from .env file for local development
	godotenv.Load()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("starting apollo agent -- version %s", "TO_BE_IMPLEMENTED")
	log.Infof("log level set to: %s", log.LogLevel())

	ctx := signals.Context()
	agent, err := agent.NewAgent(ctx, agent.Options{
		ApiPort: cfg.ApiPort,
	})
	if err != nil {
		log.Fatalf("error while creating agent: %v", err)
	}

	err = runner.NewRunnerManager(
		agent.Run,
	).Run(ctx)
	if err != nil {
		log.Fatalf("error while running agent: %v", err)
	}

	log.Info("agent shut down gracefully")
}
