package repositories

import "paper-tracker/models"

type WorkflowRepository interface {
	CreateTemplate(workflow *models.WorkflowTemplate) error
	GetAllTemplates() ([]*models.WorkflowTemplate, error)
	GetTemplateByID(workflowID models.WorkflowTemplateID) (*models.WorkflowTemplate, error)
	UpdateTemplate(workflow *models.WorkflowTemplate) error
	DeleteTemplate(workflowID models.WorkflowTemplateID) error
	CreateStep(step *models.Step) error
	GetStepByID(stepID models.StepID) (*models.Step, error)
	UpdateStep(step *models.Step) error
	DeleteStep(stepID models.StepID) error
	CreateNextStep(nextStep *models.NextStep) error
	GetLinearNextStepID(stepID models.StepID) (models.StepID, error)
	GetDecisions(stepID models.StepID) ([]*models.NextStep, error)
	IsRecordNotFoundError(err error) bool
}
