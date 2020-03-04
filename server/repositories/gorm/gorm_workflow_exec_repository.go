package gorm

import "paper-tracker/models"

func init() {
	databaseModels = append(databaseModels, &models.WorkflowTemplate{})
	databaseModels = append(databaseModels, &models.Step{})
	databaseModels = append(databaseModels, &models.NextStep{})
	databaseModels = append(databaseModels, &models.WorkflowExec{})
	databaseModels = append(databaseModels, &models.ExecStepInfo{})
}

type GormWorkflowExecRepository struct{}

func CreateGormWorkflowExecRepository() (*GormWorkflowExecRepository, error) {
	if databaseConnection == nil {
		return nil, ErrGormNotInitialized
	}
	return &GormWorkflowExecRepository{}, nil
}

func (rep *GormWorkflowExecRepository) IsRecordNotFoundError(err error) bool {
	return IsRecordNotFoundError(err)
}

func (rep *GormWorkflowExecRepository) CreateExec(exec *models.WorkflowExec) (err error) {
	err = databaseConnection.Create(exec).Error
	return
}

func (rep *GormWorkflowExecRepository) GetAllExec() (execs []*models.WorkflowExec, err error) {
	err = databaseConnection.Find(&execs).Error
	return
}

func (rep *GormWorkflowExecRepository) GetExecByID(execID models.WorkflowExecID) (exec *models.WorkflowExec, err error) {
	exec = &models.WorkflowExec{}
	err = databaseConnection.First(exec, &models.WorkflowExec{ID: execID}).Error
	return
}

func (rep *GormWorkflowExecRepository) GetRunningExecByTrackerID(trackerID models.TrackerID) (exec *models.WorkflowExec, err error) {
	exec = &models.WorkflowExec{}
	err = databaseConnection.First(exec, &models.WorkflowExec{TrackerID: trackerID, Status: models.ExecStatusRunning}).Error
	return
}

func (rep *GormWorkflowExecRepository) GetExecsByTemplateID(templateID models.WorkflowTemplateID) (execs []*models.WorkflowExec, err error) {
	err = databaseConnection.Where(&models.WorkflowExec{TemplateID: templateID}).Find(&execs).Error
	return
}

func (rep *GormWorkflowExecRepository) UpdateExec(exec *models.WorkflowExec) (err error) {
	err = databaseConnection.Save(exec).Error
	return
}

func (rep *GormWorkflowExecRepository) DeleteExec(execID models.WorkflowExecID) (err error) {
	err = databaseConnection.Delete(&models.WorkflowExec{ID: execID}).Error
	return
}

func (rep *GormWorkflowExecRepository) CreateExecStepInfo(execStepInfo *models.ExecStepInfo) (err error) {
	err = databaseConnection.Create(execStepInfo).Error
	return
}

func (rep *GormWorkflowExecRepository) GetExecStepInfoByID(execID models.WorkflowExecID, stepID models.StepID) (execStepInfo *models.ExecStepInfo, err error) {
	execStepInfo = &models.ExecStepInfo{}
	err = databaseConnection.First(execStepInfo, &models.ExecStepInfo{ExecID: execID, StepID: stepID}).Error
	return
}

func (rep *GormWorkflowExecRepository) GetExecStepInfoForExecID(execID models.WorkflowExecID) (infos []*models.ExecStepInfo, err error) {
	err = databaseConnection.Find(&infos, &models.ExecStepInfo{ExecID: execID}).Error
	return
}

func (rep *GormWorkflowExecRepository) UpdateExecStepInfo(execStepInfo *models.ExecStepInfo) (err error) {
	err = databaseConnection.Save(execStepInfo).Error
	return
}

func (rep *GormWorkflowExecRepository) DeleteExecStepInfo(execID models.WorkflowExecID, stepID models.StepID) (err error) {
	err = databaseConnection.Delete(&models.ExecStepInfo{ExecID: execID, StepID: stepID}).Error
	return
}
