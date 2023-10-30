package main

import (
	"github.com/vinisioux/go-api/config"
	"github.com/vinisioux/go-api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}

	router.Init()
}
