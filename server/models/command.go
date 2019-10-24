package models

type Command struct {
	Command    string `json:"command,omitempty"`
	WaitTimeMS int    `json:"wait_time_ms,omitempty"`
	PollAgain  bool   `json:"poll_again,omitempty"`
}
