package config

import (
	"reflect"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	// Define keys used to identify settings
	KeyDBName      = "db.name"
	KeyCoapNetwork = "coap.network"
	KeyCoapPort    = "coap.port"
	KeyHTTPPort    = "http.port"

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

	KeyTrackingScoreInMinMaxRange = "tracking.score.inminmax"
	KeyTrackingScoreInQuartiles   = "tracking.score.inquartiles"
	KeyTrackingScoreMeanFactor    = "tracking.score.mean.factor"
	KeyTrackingScoreMedianFactor  = "tracking.score.median.factor"
	KeyTrackingRangeForMean       = "tracking.range.mean"
	KeyTrackingRangeForMedian     = "tracking.range.median"
	KeyTrackingScoreThreshold     = "tracking.score.threshold"
)

type EditableConfigs struct {
	CmdIdleSleep        int      `json:"cmd_idle_sleep" ptc_key:"cmd.idle.sleep"`
	CmdInfoSleep        int      `json:"cmd_info_sleep" ptc_key:"cmd.info.sleep"`
	CmdInfoInterval     int      `json:"cmd_info_interval" ptc_key:"cmd.info.interval"`
	CmdTrackSleep       int      `json:"cmd_track_sleep" ptc_key:"cmd.track.sleep"`
	CmdLearnSleep       int      `json:"cmd_learn_sleep" ptc_key:"cmd.learn.sleep"`
	CmdLearnCount       int      `json:"cmd_learn_count" ptc_key:"cmd.learn.count"`
	CmdMaxSleep         int      `json:"cmd_max_sleep" ptc_key:"cmd.maxSleep"`
	LowBatteryThreshold int      `json:"low_battery_threshold" ptc_key:"lowBatteryThreshold"`
	WorkStartHour       int      `json:"work_start_hour" ptc_key:"work.startHour"`
	WorkEndHour         int      `json:"work_end_hour" ptc_key:"work.endHour"`
	WorkOnWeekend       bool     `json:"work_on_weekend" ptc_key:"work.onWeekend"`
	MailRecipients      []string `json:"mail_recipients" ptc_key:"mail.recipients"`
	ScoreThreshold      int      `json:"score_threshold" ptc_key:"tracking.score.threshold"`
}

// Initialize sets up the cmd line args and parses them with the config file
func Initialize() {
	pflag.String(KeyDBName, "paper-tracker.db", "Path of the database file")
	pflag.String(KeyCoapNetwork, "udp", "Network which should be used for coap requests; 'udp' or 'tcp'")
	pflag.Int(KeyCoapPort, 5688, "Port on which the application will listen for coap requests")
	pflag.Int(KeyHTTPPort, 8080, "Port on which the application will listen for http requests")

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
	pflag.Float64(KeyTrackingScoreInMinMaxRange, 1.0, "Score a room gets awarded, when a scan result is in the range of the minimum and maximum learned scans")
	pflag.Float64(KeyTrackingScoreInQuartiles, 5.0, "Score a room gets awarded, when a scan result is in the quartile range of the learned scans")
	pflag.Float64(KeyTrackingScoreMeanFactor, 5.0, "tracking.score.mean.factor")
	pflag.Float64(KeyTrackingScoreMedianFactor, 5.0, "tracking.score.median.factor")
	pflag.Float64(KeyTrackingRangeForMean, 5.0, "tracking.range.mean")
	pflag.Float64(KeyTrackingRangeForMedian, 5.0, "tracking.range.median")
	pflag.Int(KeyTrackingScoreThreshold, 750.0, "threshold from which to consider a room matchable")

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.WithError(err).Fatal("Failed to read config file")
		}
	}
}

func GetEditableConfig() (config *EditableConfigs) {
	config = &EditableConfigs{}
	val := reflect.ValueOf(config).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		tag := val.Type().Field(i).Tag.Get("ptc_key")

		val := viper.Get(tag)
		if intVal, ok := val.(int64); ok {
			val = int(intVal)
		} else if arr, ok := val.([]interface{}); ok {
			strArr := make([]string, len(arr))
			for it, arrVal := range arr {
				if arrStr, ok := arrVal.(string); ok {
					strArr[it] = arrStr
				}
			}
			val = strArr
		}

		field.Set(reflect.ValueOf(val))
	}

	return
}

func UpdateEditableConfig(config *EditableConfigs) (err error) {
	val := reflect.ValueOf(config).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		tag := val.Type().Field(i).Tag.Get("ptc_key")

		viper.Set(tag, field.Interface())
	}

	err = viper.WriteConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		err = viper.SafeWriteConfig()
	}
	if err != nil {
		log.WithError(err).Error("Failed to write config file")
		return
	}

	return
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
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
