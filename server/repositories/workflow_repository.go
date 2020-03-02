package repositories

import "paper-tracker/models"

type WorkflowRepository interface {
	CreateTemplate(template *models.WorkflowTemplate) error
	GetAllTemplates() ([]*models.WorkflowTemplate, error)
	GetTemplateByID(templateID models.WorkflowTemplateID) (*models.WorkflowTemplate, error)
	UpdateTemplate(template *models.WorkflowTemplate) error
	DeleteTemplate(templateID models.WorkflowTemplateID) error
	CreateStep(step *models.Step) error
	GetStepByID(stepID models.StepID) (*models.Step, error)
	UpdateStep(step *models.Step) error
	DeleteStep(stepID models.StepID) error
	CreateNextStep(nextStep *models.NextStep) error
	UpdateNextStep(nextStep *models.NextStep) error
	DeleteNextStep(prevStepID models.StepID, nextStepID models.StepID) error
	GetLinearNextStepID(stepID models.StepID) (models.StepID, error)
	GetNextStepByNextID(stepID models.StepID) (*models.NextStep, error)
	GetNextStepByDecison(stepID models.StepID, decision string) (models.StepID, error)
	GetDecisions(stepID models.StepID) ([]*models.NextStep, error)
	CreateExec(exec *models.WorkflowExec) error
	GetAllExec() ([]*models.WorkflowExec, error)
	GetExecByID(execID models.WorkflowExecID) (*models.WorkflowExec, error)
	GetExecsByTemplateID(templateID models.WorkflowTemplateID) ([]*models.WorkflowExec, error)
	UpdateExec(exec *models.WorkflowExec) error
	DeleteExec(execID models.WorkflowExecID) error
	CreateExecStepInfo(execStepInfo *models.ExecStepInfo) error
	GetExecStepInfoByID(execID models.WorkflowExecID, stepID models.StepID) (*models.ExecStepInfo, error)
	GetExecStepInfoForExecID(execID models.WorkflowExecID) ([]*models.ExecStepInfo, error)
	UpdateExecStepInfo(execStepInfo *models.ExecStepInfo) error
	DeleteExecStepInfo(execID models.WorkflowExecID, stepID models.StepID) error
	IsRecordNotFoundError(err error) bool
}
