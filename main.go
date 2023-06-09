package main

import (
	"github.com/LuisMarchio03/golang-boilerplate-api/config"
	"github.com/LuisMarchio03/golang-boilerplate-api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize the config
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize the router
	router.Initialize()
}
