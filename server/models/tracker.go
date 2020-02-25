package models

import "time"

type TrackerID int

type Tracker struct {
	ID            TrackerID     `json:"id" gorm:"primary_key;auto_increment"`
	Label         string        `json:"label" codec:"-"`
	LastPoll      time.Time     `json:"last_poll" codec:"-"`
	LastSleepTime time.Time     `json:"last_sleep_time" codec:"-"`
	LastLocations []Room        `json:"last_locations" codec:"-"`
	LastLocation  *Room         `json:"last_location" codec:"-"`
	Status        TrackerStatus `json:"status" codec:"-"`
}

type TrackerStatus int8

const (
	StatusIdle             TrackerStatus = 1
	StatusLearning         TrackerStatus = 2
	StatusLearningFinished TrackerStatus = 3
	StatusTracking         TrackerStatus = 4
)

func (s TrackerStatus) String() string {
	switch s {
	case StatusIdle:
		return "StatusIdle"
	case StatusLearning:
		return "StatusLearning"
	case StatusLearningFinished:
		return "StatusLearningFinished"
	case StatusTracking:
		return "StatusTracking"
	}
	return "Unknown status"
}
