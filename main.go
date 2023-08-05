package main

import (
	"github.com/andrioc/job-opening-api/config"
	"github.com/andrioc/job-opening-api/router"
)

func main() {
	logger := config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.ErrorF("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
