package communication

type LearningStartResponse struct {
	LearnTimeSec int `json:"learn_time_sec,omitempty"`
}

type LearningStatusResponse struct {
	Done  bool
	SSIDs []string
}
