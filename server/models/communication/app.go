package communication

type LearningStartResponse struct {
	LearnTimeSec int
}

type LearningStatusResponse struct {
	Done  bool
	SSIDs []string
}
