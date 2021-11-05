package main

import (
	"github.com/minskylab/YOUR-REPO-NAME/configs"
	logger "github.com/sirupsen/logrus"
)

func main() {
	logger.Info("Hello, world!")

	config := configs.MustConfigFromFile("config.yaml")

	if config.Debug {
		logger.SetLevel(logger.DebugLevel)
		logger.Debug("Debug mode enabled")
	}
}
