package models

import "time"

type WorkflowID int
type StepID int

type Workflow struct {
	ID          WorkflowID
	Label       string
	IsTemplate  bool
	StartStep   StepID
	CurrentStep StepID
}

type Step struct {
	ID    StepID
	Label string
	StartedOn time.Time
	CompletedOn time.Time
	RoomID RoomID
}

type NextStep struct {
	PrevID StepID
	NextID StepID
	DecisionLabel string
}
