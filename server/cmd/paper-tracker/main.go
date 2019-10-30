package main

import (
	"paper-tracker/managers"
	"paper-tracker/repositories"
	"paper-tracker/router"

	"github.com/go-ocf/go-coap"
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
	log.WithField("port", 5688).Info("Listening on specified port")
	log.Info(coap.ListenAndServe("udp", ":5688", router.Mux))
}
