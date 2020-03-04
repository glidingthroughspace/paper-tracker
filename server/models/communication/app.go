package communication

import "paper-tracker/models"

type LearningStartResponse struct {
	LearnTimeSec int `json:"learn_time_sec"`
}

type LearningStatusResponse struct {
	Done  bool     `json:"done"`
	SSIDs []string `json:"ssids"`
}

type LearningFinishRequest struct {
	RoomID models.RoomID `json:"room_id"`
	SSIDs  []string      `json:"ssids"`
}

type CreateStepRequest struct {
	PrevStepID    models.StepID `json:"prev_step_id"`
	DecisionLabel string        `json:"decision_label"`
	Step          *models.Step  `json:"step"`
}

type CreateRevisionRequest struct {
	RevisionLabel string `json:"revision_label"`
}
