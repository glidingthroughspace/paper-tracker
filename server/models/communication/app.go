package communication

import (
	"fmt"
	"paper-tracker/models"
)

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

type StepMoveDirection string

const (
	StepMoveUp   = StepMoveDirection("up")
	StepMoveDown = StepMoveDirection("down")
)

func StepMoveDirectionFromString(raw string) (direction StepMoveDirection, err error) {
	if StepMoveDirection(raw) == StepMoveUp {
		direction = StepMoveUp
		return
	} else if raw == "" || StepMoveDirection(raw) == StepMoveDown {
		direction = StepMoveDown
		return
	}
	err = fmt.Errorf("Could not parse '%s' as StepMoveDirection", raw)
	return
}

type TrackerNextPollResponse struct {
	NextPollSec int `json:"next_poll_sec"`
}
