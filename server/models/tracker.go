package models

import "time"

type Tracker struct {
	ID            int           `json:"id,omitempty"`
	Label         string        `json:"label,omitempty"`
	LastPoll      time.Time     `json:"last_poll,omitempty"`
	LastSleepTime time.Time     `json:"last_sleep_time,omitempty"`
	LastLocations []Room        `json:"last_locations,omitempty"`
	LastLocation  *Room         `json:"last_location,omitempty"`
	Status        TrackerStatus `json:"status,omitempty"`
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
