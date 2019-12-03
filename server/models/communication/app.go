package communication

type LearningStartResponse struct {
	LearnTimeSec int `json:"learn_time_sec"`
}

type LearningStatusResponse struct {
	Done  bool
	SSIDs []string
}
