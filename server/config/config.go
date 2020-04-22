package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	KeyDBName      = "db.name"
	KeyCoapNetwork = "coap.network"
	KeyCoapPort    = "coap.port"
	KeyHttpPort    = "http.port"

	KeyCmdIdleSleep        = "cmd.idle.sleep"
	KeyCmdInfoSleep        = "cmd.info.sleep"
	KeyCmdInfoInterval     = "cmd.info.interval"
	KeyCmdTrackSleep       = "cmd.track.sleep"
	KeyCmdLearnSleep       = "cmd.learn.sleep"
	KeyCmdLearnCount       = "cmd.learn.count"
	KeyCmdMaxSleep         = "cmd.maxSleep"
	KeyLowBatteryThreshold = "lowBatteryThreshold"

	KeyWorkStartHour = "work.startHour"
	KeyWorkEndHour   = "work.endHour"
	KeyWorkOnWeekend = "work.onWeekend"

	KeyMailUsername   = "mail.username"
	KeyMailPassword   = "mail.password"
	KeyMailHost       = "mail.host"
	KeyMailPort       = "mail.port"
	KeyMailSender     = "mail.sender"
	KeyMailRecipients = "mail.recipients"
)

// Initialize sets up the cmd line args and parses them with the config file
func Initialize() {
	pflag.String(KeyDBName, "paper-tracker.db", "Path of the database file")
	pflag.String(KeyCoapNetwork, "udp", "Network which should be used for coap requests; 'udp' or 'tcp'")
	pflag.Int(KeyCoapPort, 5688, "Port on which the application will listen for coap requests")
	pflag.Int(KeyHttpPort, 8080, "Port on which the application will listen for http requests")

	pflag.Int(KeyCmdIdleSleep, 5, "Sleep duration for the tracker before polling for new command in idle")
	pflag.Int(KeyCmdInfoSleep, 5, "Sleep duration for the tracker before sending battery stats when idling")
	pflag.Int(KeyCmdInfoInterval, 60, "Interval for the tracker to send battery stats when idling")
	pflag.Int(KeyCmdTrackSleep, 5, "Sleep duration for the tracker before polling for new command in tracking")
	pflag.Int(KeyCmdLearnSleep, 5, "Sleep duration for the tracker before polling for new command in learning")
	pflag.Int(KeyCmdLearnCount, 5, "Total times the WiFi is scanned when learning a room")
	pflag.Int(KeyCmdMaxSleep, 1800, "Maximum possible sleep time")
	pflag.Int(KeyLowBatteryThreshold, 10, "Threshold that specifies under which threshold a low battery notification will be sent")

	pflag.Int(KeyWorkStartHour, -1, "Hour of the day the tracker should become active. In 24-Hour format. Set this or end value to -1 to disable.")
	pflag.Int(KeyWorkEndHour, -1, "Hour of the day the tracker should become inactive. In 24-Hour format. Set this or start value to -1 to disable.")
	pflag.Bool(KeyWorkOnWeekend, false, "Whether the tracker should be active on weekends")

	pflag.String(KeyMailUsername, "", "Username used for connecting to the SMTP server for email notifications. Leave empty for no authorization.")
	pflag.String(KeyMailPassword, "", "Password used for connecting to the SMTP server for email notifications.")
	pflag.String(KeyMailHost, "", "SMTP Host used for sending notification emails")
	pflag.Int(KeyMailPort, 25, "Port used for SMTP Host. Defaults to 25")
	pflag.String(KeyMailSender, "", "Email address to send email from. Leave empty to use 'mail.username'.")
	pflag.StringSlice(KeyMailRecipients, []string{}, "List of email addresses to recevive email notifications")

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.WithError(err).Fatal("Failed to read config file")
		}
	}
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}
