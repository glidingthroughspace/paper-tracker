package models

type Command struct {
	Type         CommandType `json:"type"`
	SleepTimeSec int         `json:"sleep_time_sec"`
}

type CommandType int8

const (
	CmdSendTrackingInformation CommandType = 0
	CmdUnused                  CommandType = 1
	CmdSleep                   CommandType = 2
	CmdSendInformation         CommandType = 3
)
