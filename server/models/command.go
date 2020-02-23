package models

import "time"

type CommandID int

type Command struct {
	ID           CommandID   `codec:"-" gorm:"primary_key;auto_increment" json:"id"`
	TrackerID    TrackerID   `codec:"-" json:"tracker_id"`
	Command      CommandType `json:"command"`
	SleepTimeSec int         `json:"sleep_time_sec"`
	CreatedAt    time.Time   `codec:"-" json:"created_at"`
}

type CommandType int8

const (
	CmdSendTrackingInformation CommandType = 0
	CmdSignalLocation          CommandType = 1
	CmdSleep                   CommandType = 2
)
