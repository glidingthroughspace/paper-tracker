package repositories

import "paper-tracker/models"

type WorkflowRepository interface {
	CreateWorkflow(workflow *models.Workflow) error
	GetWorkflowByID(workflowID models.WorkflowID) (*models.Workflow, error)
	CreateStep(step *models.Step) error
	GetStepByID(stepID models.StepID) (*models.Step, error)
	IsRecordNotFoundError(err error) bool
}
