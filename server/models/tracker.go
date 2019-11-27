package models

import "time"

type Tracker struct {
	ID            int
	Label         string
	LastPoll      time.Time
	LastSleepTime time.Time
	LastLocations []Room
	LastLocation  Room
	Status        TrackerStatus
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
