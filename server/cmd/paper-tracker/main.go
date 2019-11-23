package main

import (
	"flag"
	"paper-tracker/managers"
	"paper-tracker/repositories/gorm"
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
	defaultSleepSecPtr := flag.Int("default-sleep", 5, "Default sleep duration for the tracker before polling for new commands")
	learnCountPtr := flag.Int("learn-count", 5, "Total times the WiFi is scanned when learning a room")
	sleepBetweenLearnSecPtr := flag.Int("sleep-between-learn", 5, "Sleep duration between two scans during learning")

	err := gorm.InitDatabaseConnection(*dbNamePtr)
	if err != nil {
		log.Fatal("Abort: Failed to initialize database")
	}

	trackerRep, err := gorm.CreateGormTrackerRepository()
	if err != nil {
		log.Fatal("Abort: Failed to create tracker repository")
	}
	cmdRep, err := gorm.CreateGormCommandRepository()
	if err != nil {
		log.Fatal("Abort: Failed to create command repository")
	}

	trackerMgr := managers.CreateTrackerManager(trackerRep, cmdRep, *defaultSleepSecPtr, *learnCountPtr, *sleepBetweenLearnSecPtr)

	coapRouter := router.NewCoapRouter(trackerMgr)
	httpRouter := router.NewHttpRouter(trackerMgr)

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
