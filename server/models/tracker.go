package models

import "time"

type TrackerID int

type Tracker struct {
	ID TrackerID `json:"id" gorm:"primary_key;auto_increment"`
	// Label is a human-readable name for the tracker
	Label string `json:"label" codec:"-"`
	// LastPoll indicates when the tracker has last polled for a command
	LastPoll time.Time `json:"last_poll" codec:"-"`
	// LastSleepTimeSec saves the last sleep time the tracker received
	LastSleepTimeSec int `json:"-" codec:"-"`
	// LastBatteryUpdate indicates when the tracker's battery status has last been updated
	LastBatteryUpdate time.Time `json:"last_battery_update" codec:"-"`
	// LastRoom is the last known room in which the tracker was located
	LastRoom RoomID `json:"last_room" codec:"-"`
	// Status indicates what mode of operation the tracker is in
	Status             TrackerStatus `json:"status" codec:"-"`
	BatteryPercentage  int           `json:"battery_percentage" codec:"-"`
	IsCharging         bool          `json:"is_charging" codec:"-"`
	LowBatteryNotified bool          `json:"low_battery_notified" codec:"-"`
}

func (tracker *Tracker) GetSecondsToNextPoll() int {
	pollSecs := tracker.LastPoll.Add(time.Second * time.Duration(tracker.LastSleepTimeSec)).Sub(time.Now()).Seconds()
	if pollSecs < 0 {
		return 0
	}
	return int(pollSecs)
}

type TrackerStatus int8

const (
	TrackerStatusIdle             TrackerStatus = 1
	TrackerStatusLearning         TrackerStatus = 2
	TrackerStatusLearningFinished TrackerStatus = 3
	TrackerStatusTracking         TrackerStatus = 4
)

func (s TrackerStatus) String() string {
	switch s {
	case TrackerStatusIdle:
		return "StatusIdle"
	case TrackerStatusLearning:
		return "StatusLearning"
	case TrackerStatusLearningFinished:
		return "StatusLearningFinished"
	case TrackerStatusTracking:
		return "StatusTracking"
	}
	return "Unknown status"
}
