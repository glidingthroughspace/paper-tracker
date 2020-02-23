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
	CreateExec(exec *models.WorkflowExec) error
	GetAllExec() ([]*models.WorkflowExec, error)
	GetExecByID(execID models.WorkflowExecID) (*models.WorkflowExec, error)
	UpdateExec(exec *models.WorkflowExec) error
	DeleteExec(execID models.WorkflowExecID) error
	CreateExecStepInfo(execStepInfo *models.ExecStepInfo) error
	GetExecStepInfoByID(execID models.WorkflowExecID, stepID models.StepID) (*models.ExecStepInfo, error)
	UpdateExecStepInfo(execStepInfo *models.ExecStepInfo) error
	DeleteExecStepInfo(execID models.WorkflowExecID, stepID models.StepID) error
	IsRecordNotFoundError(err error) bool
}
