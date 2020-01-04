package communication

type LearningStartResponse struct {
	LearnTimeSec int `json:"learn_time_sec"`
}

type LearningStatusResponse struct {
	Done  bool     `json:"done"`
	SSIDs []string `json:"ssids"`
}

type LearningFinishRequest struct {
	RoomID int      `json:"room_id"`
	SSIDs  []string `json:"ssids"`
}
