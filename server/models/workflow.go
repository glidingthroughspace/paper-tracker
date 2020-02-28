package models

import "time"

type WorkflowTemplateID int
type StepID int
type WorkflowExecID int

type WorkflowTemplate struct {
	ID              WorkflowTemplateID `json:"id" gorm:"primary_key;auto_increment"`
	Label           string             `json:"label"`
	StartStep       StepID             `json:"-"`
	FirstRevisionID WorkflowTemplateID `json:"first_revision_id"`
	Steps           []*Step            `json:"steps" gorm:"-"`
	EditingLocked   bool               `json:"editing_locked" gorm:"-"`
}

type Step struct {
	ID      StepID             `json:"id" gorm:"primary_key;auto_increment"`
	Label   string             `json:"label"`
	RoomID  RoomID             `json:"room_id"`
	Options map[string][]*Step `json:"options" gorm:"-"`
}

type NextStep struct {
	PrevID        StepID `json:"prev_id" gorm:"primary_key;auto_increment:false"`
	NextID        StepID `json:"next_id" gorm:"primary_key;auto_increment:false"`
	DecisionLabel string `json:"decision_label"`
}

type WorkflowExec struct {
	ID            WorkflowExecID           `json:"id" gorm:"primary_key;auto_increment"`
	Label         string                   `json:"label"`
	TemplateID    WorkflowTemplateID       `json:"template_id"`
	TrackerID     TrackerID                `json:"tracker_id"`
	Completed     bool                     `json:"completed"`
	StartedOn     time.Time                `json:"started_on"`
	CompletedOn   time.Time                `json:"completed_on"`
	CurrentStepID StepID                   `json:"current_step_id"`
	StepInfos     map[StepID]*ExecStepInfo `json:"step_infos" gorm:"-"`
}

type ExecStepInfo struct {
	ExecID      WorkflowExecID `json:"-" gorm:"primary_key;auto_increment:false"`
	StepID      StepID         `json:"-" gorm:"primary_key;auto_increment:false"`
	Decision    string         `json:"decision"`
	StartedOn   time.Time      `json:"started_on"`
	CompletedOn time.Time      `json:"completed_on"`
}
