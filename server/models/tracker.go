package models

import "time"

type TrackerID int

type Tracker struct {
	ID                TrackerID     `json:"id" gorm:"primary_key;auto_increment"`
	Label             string        `json:"label" codec:"-"`
	LastPoll          time.Time     `json:"last_poll" codec:"-"`
	LastSleepTime     time.Time     `json:"last_sleep_time" codec:"-"`
	LastRoom          RoomID        `json:"last_room" codec:"-"`
	Status            TrackerStatus `json:"status" codec:"-"`
	BatteryPercentage int           `json:"battery_percentage" codec:"-"`
	IsCharging        bool          `json:"is_charging" codec:"-"`
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
