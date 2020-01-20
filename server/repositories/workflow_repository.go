package repositories

import "paper-tracker/models"

type WorkflowRepository interface {
	CreateWorkflow(workflow *models.Workflow) error
	GetWorkflowByID(workflowID models.WorkflowID) (*models.Workflow, error)
	UpdateWorkflow(workflow *models.Workflow) error
	DeleteWorkflow(workflowID models.WorkflowID) error
	CreateStep(step *models.Step) error
	GetStepByID(stepID models.StepID) (*models.Step, error)
	UpdateStep(step *models.Step) error
	DeleteStep(stepID models.StepID) error
	IsRecordNotFoundError(err error) bool
}
