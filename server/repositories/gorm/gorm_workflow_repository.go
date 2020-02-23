package gorm

import "paper-tracker/models"

func init() {
	databaseModels = append(databaseModels, &models.WorkflowTemplate{})
	databaseModels = append(databaseModels, &models.Step{})
	databaseModels = append(databaseModels, &models.NextStep{})
	databaseModels = append(databaseModels, &models.WorkflowExec{})
	databaseModels = append(databaseModels, &models.ExecStepInfo{})
}

type GormWorkflowRepository struct{}

func CreateGormWorkflowRepository() (*GormWorkflowRepository, error) {
	if databaseConnection == nil {
		return nil, ErrGormNotInitialized
	}
	return &GormWorkflowRepository{}, nil
}

func (rep *GormWorkflowRepository) IsRecordNotFoundError(err error) bool {
	return IsRecordNotFoundError(err)
}

func (rep *GormWorkflowRepository) CreateTemplate(workflow *models.WorkflowTemplate) (err error) {
	err = databaseConnection.Create(workflow).Error
	return
}

func (rep *GormWorkflowRepository) GetAllTemplates() (workflows []*models.WorkflowTemplate, err error) {
	err = databaseConnection.Find(&workflows).Error
	return
}

func (rep *GormWorkflowRepository) GetTemplateByID(workflowID models.WorkflowTemplateID) (workflow *models.WorkflowTemplate, err error) {
	workflow = &models.WorkflowTemplate{}
	err = databaseConnection.First(workflow, &models.WorkflowTemplate{ID: workflowID}).Error
	return
}

func (rep *GormWorkflowRepository) UpdateTemplate(workflow *models.WorkflowTemplate) (err error) {
	err = databaseConnection.Save(workflow).Error
	return
}

func (rep *GormWorkflowRepository) DeleteTemplate(workflowID models.WorkflowTemplateID) (err error) {
	err = databaseConnection.Delete(&models.WorkflowTemplate{ID: workflowID}).Error
	return
}

func (rep *GormWorkflowRepository) CreateStep(step *models.Step) (err error) {
	err = databaseConnection.Create(step).Error
	return
}

func (rep *GormWorkflowRepository) GetStepByID(stepID models.StepID) (step *models.Step, err error) {
	step = &models.Step{}
	err = databaseConnection.Where("id = ?", stepID).First(step).Error
	return
}

func (rep *GormWorkflowRepository) UpdateStep(step *models.Step) (err error) {
	err = databaseConnection.Save(step).Error
	return
}

func (rep *GormWorkflowRepository) DeleteStep(stepID models.StepID) (err error) {
	err = databaseConnection.Delete(&models.Step{ID: stepID}).Error
	return
}

func (rep *GormWorkflowRepository) CreateNextStep(nextStep *models.NextStep) (err error) {
	err = databaseConnection.Create(nextStep).Error
	return
}

func (rep *GormWorkflowRepository) GetLinearNextStepID(stepID models.StepID) (nextStepID models.StepID, err error) {
	nextStep := &models.NextStep{}
	err = databaseConnection.Where("prev_id = ? AND decision_label = \"\"", stepID).First(nextStep).Error
	nextStepID = nextStep.NextID
	return
}

func (rep *GormWorkflowRepository) GetDecisions(stepID models.StepID) (decisions []*models.NextStep, err error) {
	err = databaseConnection.Where("prev_id = ? AND decision_label <> \"\"", stepID).Find(&decisions).Error
	return
}

func (rep *GormWorkflowRepository) CreateExec(exec *models.WorkflowExec) (err error) {
	err = databaseConnection.Create(exec).Error
	return
}

func (rep *GormWorkflowRepository) GetAllExec() (execs []*models.WorkflowExec, err error) {
	err = databaseConnection.Find(&execs).Error
	return
}

func (rep *GormWorkflowRepository) GetExecByID(execID models.WorkflowExecID) (exec *models.WorkflowExec, err error) {
	exec = &models.WorkflowExec{}
	err = databaseConnection.First(exec, &models.WorkflowExec{ID: execID}).Error
	return
}

func (rep *GormWorkflowRepository) UpdateExec(exec *models.WorkflowExec) (err error) {
	err = databaseConnection.Update(exec).Error
	return
}

func (rep *GormWorkflowRepository) DeleteExec(execID models.WorkflowExecID) (err error) {
	err = databaseConnection.Delete(&models.WorkflowExec{ID: execID}).Error
	return
}

func (rep *GormWorkflowRepository) CreateExecStepInfo(execStepInfo *models.ExecStepInfo) (err error) {
	err = databaseConnection.Create(execStepInfo).Error
	return
}

func (rep *GormWorkflowRepository) GetExecStepInfoByID(execID models.WorkflowExecID, stepID models.StepID) (execStepInfo *models.ExecStepInfo, err error) {
	execStepInfo = &models.ExecStepInfo{}
	err = databaseConnection.First(execStepInfo, &models.ExecStepInfo{ExecID: execID, StepID: stepID}).Error
	return
}

func (rep *GormWorkflowRepository) GetExecStepInfoForExecID(execID models.WorkflowExecID) (infos []*models.ExecStepInfo, err error) {
	err = databaseConnection.Find(&infos, &models.ExecStepInfo{ExecID: execID}).Error
	return
}

func (rep *GormWorkflowRepository) UpdateExecStepInfo(execStepInfo *models.ExecStepInfo) (err error) {
	err = databaseConnection.Update(execStepInfo).Error
	return
}

func (rep *GormWorkflowRepository) DeleteExecStepInfo(execID models.WorkflowExecID, stepID models.StepID) (err error) {
	err = databaseConnection.Delete(&models.ExecStepInfo{ExecID: execID, StepID: stepID}).Error
	return
}
