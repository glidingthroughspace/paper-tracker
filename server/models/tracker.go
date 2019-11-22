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
	StatusIdle     TrackerStatus = 1
	StatusLearning TrackerStatus = 2
	StatusTracking TrackerStatus = 3
)
