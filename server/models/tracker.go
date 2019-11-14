package models

import "time"

type Tracker struct {
	ID            int    `json:"id,omitempty"`
	Label         string `json:"label,omitempty"`
	LastPoll      time.Time
	LastSleepTime time.Time
	LastLocations []Room
	NextCommands  []Command
}
