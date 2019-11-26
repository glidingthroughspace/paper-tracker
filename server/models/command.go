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
	CmdSendTrackingInformation CommandType = 0
	CmdSignalLocation          CommandType = 1
	CmdSleep                   CommandType = 2
)
