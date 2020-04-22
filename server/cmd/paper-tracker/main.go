package main

import (
	"os"
	"paper-tracker/managers"
	"paper-tracker/repositories/gorm"
	"paper-tracker/router"
	"strconv"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	router.SetMode(router.ReleaseMode)
	if debug := os.Getenv("DEBUG"); debug != "" {
		log.SetLevel(log.DebugLevel)
		router.SetMode(router.DebugMode)
	}
}

func main() {
	pflag.String("db.name", "paper-tracker.db", "Path of the database file")
	pflag.String("coap.network", "udp", "Network which should be used for coap requests; 'udp' or 'tcp'")
	pflag.Int("coap.port", 5688, "Port on which the application will listen for coap requests")
	pflag.Int("http.port", 8080, "Port on which the application will listen for http requests")

	pflag.Int("cmd.idle.sleep", 5, "Sleep duration for the tracker before polling for new command in idle")
	pflag.Int("cmd.info.sleep", 5, "Sleep duration for the tracker before sending battery stats when idling")
	pflag.Int("cmd.info.interval", 60, "Interval for the tracker to send battery stats when idling")
	pflag.Int("cmd.track.sleep", 5, "Sleep duration for the tracker before polling for new command in tracking")
	pflag.Int("cmd.learn.sleep", 5, "Sleep duration for the tracker before polling for new command in learning")
	pflag.Int("cmd.learn.count", 5, "Total times the WiFi is scanned when learning a room")
	pflag.Int("cmd.maxSleep", 1800, "Maximum possible sleep time")
	pflag.Int("lowBatteryThreshold", 10, "Threshold that specifies under which threshold a low battery notification will be sent")

	pflag.Int("work.startHour", -1, "Hour of the day the tracker should become active. In 24-Hour format. Set this or end value to -1 to disable.")
	pflag.Int("work.endHour", -1, "Hour of the day the tracker should become inactive. In 24-Hour format. Set this or start value to -1 to disable.")
	pflag.Bool("work.onWeekend", false, "Whether the tracker should be active on weekends")

	pflag.String("mail.username", "", "Username used for connecting to the SMTP server for email notifications. Leave empty for no authorization.")
	pflag.String("mail.password", "", "Password used for connecting to the SMTP server for email notifications.")
	pflag.String("mail.host", "", "SMTP Host used for sending notification emails")
	pflag.Int("mail.port", 25, "Port used for SMTP Host. Defaults to 25")
	pflag.String("mail.sender", "", "Email address to send email from. Leave empty to use 'mail.username'.")
	pflag.StringSlice("mail.recipients", []string{}, "List of email addresses to recevive email notifications")

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.WithError(err).Fatal("Failed to read config file")
		}
	}

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

	coapPort := viper.GetInt("coap.port")
	coapNetwork := viper.GetString("coap.network")
	log.WithField("port", coapPort).Info("Listening for coap on specified port")
	wg.Add(1)
	go coapRouter.Serve(coapNetwork, ":"+strconv.Itoa(coapPort), &wg)

	httpPort := viper.GetInt("http.port")
	log.WithField("port", httpPort).Info("Listening for http on specified port")
	wg.Add(1)
	go httpRouter.Serve(":"+strconv.Itoa(httpPort), &wg)

	wg.Wait()
}
