package main

import (
	"os"

	"procpipe/config"
	"procpipe/logger"
	"procpipe/orchestrator"
)

func main() {
	flags := config.ParseFlags()

	cfg, err := config.Load(flags)
	if err != nil {
		logger.Error("Config error: %s", err)
		os.Exit(1)
	}

	exitCode, err := orchestrator.Run(cfg)
	if err != nil {
		logger.Error("Fatal: %s", err)
		os.Exit(1)
	}

	os.Exit(exitCode)
}
