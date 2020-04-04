package main

import (
	"flag"
	"os"
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
	dbNamePtr := flag.String("db-name", "paper-tracker.db", "Path of the database file")
	coapNetworkPtr := flag.String("coap-network", "udp", "Network which should be used for coap requests; 'udp' or 'tcp'")
	coapPortPtr := flag.Int("coap-port", 5688, "Port on which the application will listen for coap requests")
	httpPortPtr := flag.Int("http-port", 8080, "Port on which the application will listen for http requests")

	idleSleepSecPtr := flag.Int("idle-sleep", 5, "Sleep duration for the tracker before polling for new command in idle")
	sendInfoSleepSecPtr := flag.Int("info-sleep", 5, "Sleep duration for the tracker before sending battery stats when idling")
	sendInfoIntervalSecPtr := flag.Int("info-interval", 60, "Interval for the tracker to send battery stats when idling")
	trackingSleepSecPtr := flag.Int("tracking-sleep", 5, "Sleep duration for the tracker before polling for new command in tracking")
	learnSleepSecPtr := flag.Int("learn-sleep", 5, "Sleep duration for the tracker before polling for new command in learning")
	learnCountPtr := flag.Int("learn-count", 5, "Total times the WiFi is scanned when learning a room")

	workStartHourPtr := flag.Int("work-start-hour", -1, "Hour of the day the tracker should become active. In 24-Hour format. Set this or end value to -1 to disable.")
	workEndHourPtr := flag.Int("work-end-hour", -1, "Hour of the day the tracker should become inactive. In 24-Hour format. Set this or start value to -1 to disable.")
	workOnWeekend := flag.Bool("work-on-weekend", false, "Whether the tracker should be inactive on weekends")

	err := gorm.InitDatabaseConnection(*dbNamePtr)
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

	managers.CreateTrackerManager(trackerRep, *idleSleepSecPtr, *trackingSleepSecPtr, *learnSleepSecPtr,
		*sendInfoSleepSecPtr, *sendInfoIntervalSecPtr, *workStartHourPtr, *workEndHourPtr, *workOnWeekend)
	managers.CreateRoomManager(roomRep)
	managers.CreateLearningManager(scanResultRep, *learnCountPtr, *learnSleepSecPtr)
	managers.CreateWorkflowTemplateManager(workflowTemplateRep)
	managers.CreateWorkflowExecManager(workflowExecRep)

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
