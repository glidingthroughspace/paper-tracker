package main

import (
	"paper-tracker/repositories"
	"paper-tracker/managers"
	"paper-tracker/router"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func main() {
	err := repositories.InitDatabaseConnection("paper-tracker.db")
	if err != nil {
		log.Fatal("Abort: Failed to initialize database")
	}

	trackerRep, err := repositories.CreateTrackerRepository()
	if err != nil {
		log.Fatal("Abort: Failed to create tracker repository")
	}

	_ = managers.CreateTrackerManager(trackerRep)

	router := router.NewRouter()

	// Start
	log.WithField("port", 8080).Info("Listening on specified port")
	log.Info(router.Router.Run(":" + strconv.Itoa(8080)))
}
