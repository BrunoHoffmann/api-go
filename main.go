package main

import (
	"github.com/BrunoHoffmann/api-go/config"
	"github.com/BrunoHoffmann/api-go/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
