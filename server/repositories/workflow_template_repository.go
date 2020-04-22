package repositories

import "paper-tracker/models"

type WorkflowTemplateRepository interface {
	CreateTemplate(template *models.WorkflowTemplate) error
	GetAllTemplates() ([]*models.WorkflowTemplate, error)
	GetTemplateByID(templateID models.WorkflowTemplateID) (*models.WorkflowTemplate, error)
	UpdateTemplate(template *models.WorkflowTemplate) error
	DeleteTemplate(templateID models.WorkflowTemplateID) error
	CreateStep(step *models.Step) error
	CreateStepRoom(stepRoom *models.StepRoom) error
	GetStepByID(stepID models.StepID) (*models.Step, error)
	GetRoomsByStepID(stepID models.StepID) ([]*models.StepRoom, error)
	GetStepsByRoomID(roomID models.RoomID) ([]*models.StepRoom, error)
	UpdateStep(step *models.Step) error
	DeleteStep(stepID models.StepID) error
	ClearStepRooms(stepID models.StepID) error
	CreateNextStep(nextStep *models.NextStep) error
	UpdateNextStep(nextStep *models.NextStep) error
	DeleteNextStep(prevStepID models.StepID, nextStepID models.StepID) error
	GetLinearNextStepID(stepID models.StepID) (models.StepID, error)
	GetNextStepByNextID(stepID models.StepID) (*models.NextStep, error)
	GetNextStepByDecison(stepID models.StepID, decision string) (models.StepID, error)
	GetDecisions(stepID models.StepID) ([]*models.NextStep, error)
	IsRecordNotFoundError(err error) bool
}
