package models

import "time"

type Tracker struct {
	ID            int
	Label         string
	LastPoll      time.Time
	LastSleepTime time.Time
	LastLocations []Room
}
