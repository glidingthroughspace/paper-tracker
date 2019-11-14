package models

import "time"

type Command struct {
	ID           int `codec:"-"`
	TrackerID    int `codec:"-"`
	Command      CommandType
	SleepTimeSec int
	CreatedAt    time.Time `codec:"-"`
}

type CommandType int8

const (
	SendTrackingInformation CommandType = 0
	SignalLocation          CommandType = 1
	Sleep                   CommandType = 2
)
