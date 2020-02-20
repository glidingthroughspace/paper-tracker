package models

import "time"

type WorkflowID int
type StepID int

type Workflow struct {
	ID          WorkflowID `json:"id" gorm:"primary_key;auto_increment"`
	Label       string     `json:"label"`
	IsTemplate  bool       `json:"is_template"`
	StartStep   StepID     `json:"-"`
	CurrentStep StepID     `json:"-"`
	Steps       []*Step    `json:"steps" gorm:"-"`
}

type Step struct {
	ID          StepID             `json:"id" gorm:"primary_key;auto_increment"`
	Label       string             `json:"label"`
	StartedOn   time.Time          `json:"started_on"`
	CompletedOn time.Time          `json:"completed_on"`
	RoomID      RoomID             `json:"room_id"`
	Options     map[string][]*Step `json:"options" gorm:"-"`
}

type NextStep struct {
	PrevID        StepID `json:"prev_id" gorm:"primary_key;auto_increment:false"`
	NextID        StepID `json:"next_id" gorm:"primary_key;auto_increment:false"`
	DecisionLabel string `json:"decision_label"`
}