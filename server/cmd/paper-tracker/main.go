package main

import (
	"flag"
	"paper-tracker/managers"
	"paper-tracker/repositories"
	"paper-tracker/router"
	"strconv"
	"sync"

	log "github.com/sirupsen/logrus"
)

func main() {
	dbNamePtr := flag.String("db-name", "paper-tracker.db", "Path of the database file")
	coapNetworkPtr := flag.String("coap-network", "udp", "Network which should be used for coap requests; 'udp' or 'tcp'")
	coapPortPtr := flag.Int("coap-port", 5688, "Port on which the application will listen for coap requests")
	httpPortPtr := flag.Int("http-port", 8080, "Port on which the application will listen for http requests")

	err := repositories.InitDatabaseConnection(*dbNamePtr)
	if err != nil {
		log.Fatal("Abort: Failed to initialize database")
	}

	trackerRep, err := repositories.CreateTrackerRepository()
	if err != nil {
		log.Fatal("Abort: Failed to create tracker repository")
	}

	_ = managers.CreateTrackerManager(trackerRep)

	coapRouter := router.NewCoapRouter()
	httpRouter := router.NewHttpRouter()

	// Start
	var wg sync.WaitGroup

	log.WithField("port", *coapPortPtr).Info("Listening for coap on specified port")
	wg.Add(1)
	go coapRouter.Serve(*coapNetworkPtr, ":"+strconv.Itoa(*coapPortPtr), &wg)

	log.WithField("port", *httpPortPtr).Info("Listening for http on specified port")
	wg.Add(1)
	go httpRouter.Serve(":"+strconv.Itoa(*httpPortPtr), &wg)

	wg.Wait()
}
