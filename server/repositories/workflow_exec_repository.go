package repositories

import "paper-tracker/models"

type WorkflowExecRepository interface {
	CreateExec(exec *models.WorkflowExec) error
	GetAllExec() ([]*models.WorkflowExec, error)
	GetExecByID(execID models.WorkflowExecID) (*models.WorkflowExec, error)
	GetRunningExecByTrackerID(trackerID models.TrackerID) (*models.WorkflowExec, error)
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
