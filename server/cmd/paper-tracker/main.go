package main

import (
	"os"
	"paper-tracker/config"
	"paper-tracker/managers"
	"paper-tracker/repositories/gorm"
	"paper-tracker/router"
	"strconv"
	"sync"

	log "github.com/sirupsen/logrus"
)

func init() {
	router.SetMode(router.ReleaseMode)
	if debug := os.Getenv("DEBUG"); debug != "" {
		log.SetLevel(log.DebugLevel)
		router.SetMode(router.DebugMode)
	}
}

func main() {
	config.Initialize()

	err := gorm.InitDatabaseConnection()
	if err != nil {
		log.Fatal("Abort: Failed to initialize database")
	}

	trackerRep, err := gorm.CreateGormTrackerRepository()
	if err != nil {
		log.Fatal("Abort: Failed to create tracker repository")
	}
	scanResultRep, err := gorm.CreateGormScanResultRepository()
	if err != nil {
		log.Fatal("Abort: Failed to create scan result repository")
	}
	roomRep, err := gorm.CreateGormRoomRepository()
	if err != nil {
		log.Fatal("Abort: Failed to create room repository")
	}
	workflowTemplateRep, err := gorm.CreateGormWorkflowTemplateRepository()
	if err != nil {
		log.Fatal("Abort: Failed to create workflow template repository")
	}
	workflowExecRep, err := gorm.CreateGormWorkflowExecRepository()
	if err != nil {
		log.Fatal("Abort: Failed to create workflow repository")
	}

	managers.CreateTrackerManager(trackerRep)
	managers.CreateTrackingManager()
	managers.CreateRoomManager(roomRep)
	managers.CreateLearningManager(scanResultRep)
	managers.CreateWorkflowTemplateManager(workflowTemplateRep)
	managers.CreateWorkflowExecManager(workflowExecRep)
	managers.CreateExportManager()

	coapRouter := router.NewCoapRouter()
	httpRouter := router.NewHttpRouter()

	// Start
	var wg sync.WaitGroup

	coapPort := config.GetInt(config.KeyCoapPort)
	coapNetwork := config.GetString(config.KeyCoapNetwork)
	log.WithField("port", coapPort).Info("Listening for coap on specified port")
	wg.Add(1)
	go coapRouter.Serve(coapNetwork, ":"+strconv.Itoa(coapPort), &wg)

	httpPort := config.GetInt(config.KeyHttpPort)
	log.WithField("port", httpPort).Info("Listening for http on specified port")
	wg.Add(1)
	go httpRouter.Serve(":"+strconv.Itoa(httpPort), &wg)

	wg.Wait()
}
