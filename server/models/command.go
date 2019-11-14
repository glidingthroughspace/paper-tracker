package models

type Command struct {
	ID         int
	Command    string `json:"command,omitempty"`
	WaitTimeMS int    `gorm:"-"`
}

type CommandType int8

const (
	SendTrackingInformation CommandType = 0
	SignalLocation          CommandType = 1
	Sleep                   CommandType = 2
)
