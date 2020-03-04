package gorm

import "paper-tracker/models"

func init() {
	databaseModels = append(databaseModels, &models.WorkflowTemplate{})
	databaseModels = append(databaseModels, &models.Step{})
	databaseModels = append(databaseModels, &models.NextStep{})
	databaseModels = append(databaseModels, &models.WorkflowExec{})
	databaseModels = append(databaseModels, &models.ExecStepInfo{})
}

type GormWorkflowTemplateRepository struct{}

func CreateGormWorkflowTemplateRepository() (*GormWorkflowTemplateRepository, error) {
	if databaseConnection == nil {
		return nil, ErrGormNotInitialized
	}
	return &GormWorkflowTemplateRepository{}, nil
}

func (rep *GormWorkflowTemplateRepository) IsRecordNotFoundError(err error) bool {
	return IsRecordNotFoundError(err)
}

func (rep *GormWorkflowTemplateRepository) CreateTemplate(workflow *models.WorkflowTemplate) (err error) {
	err = databaseConnection.Create(workflow).Error
	return
}

func (rep *GormWorkflowTemplateRepository) GetAllTemplates() (workflows []*models.WorkflowTemplate, err error) {
	err = databaseConnection.Find(&workflows).Error
	return
}

func (rep *GormWorkflowTemplateRepository) GetTemplateByID(workflowID models.WorkflowTemplateID) (workflow *models.WorkflowTemplate, err error) {
	workflow = &models.WorkflowTemplate{}
	err = databaseConnection.First(workflow, &models.WorkflowTemplate{ID: workflowID}).Error
	return
}

func (rep *GormWorkflowTemplateRepository) UpdateTemplate(workflow *models.WorkflowTemplate) (err error) {
	err = databaseConnection.Save(workflow).Error
	return
}

func (rep *GormWorkflowTemplateRepository) DeleteTemplate(workflowID models.WorkflowTemplateID) (err error) {
	err = databaseConnection.Delete(&models.WorkflowTemplate{ID: workflowID}).Error
	return
}

func (rep *GormWorkflowTemplateRepository) CreateStep(step *models.Step) (err error) {
	err = databaseConnection.Create(step).Error
	return
}

func (rep *GormWorkflowTemplateRepository) GetStepByID(stepID models.StepID) (step *models.Step, err error) {
	step = &models.Step{}
	err = databaseConnection.First(step, &models.Step{ID: stepID}).Error
	return
}

func (rep *GormWorkflowTemplateRepository) UpdateStep(step *models.Step) (err error) {
	err = databaseConnection.Save(step).Error
	return
}

func (rep *GormWorkflowTemplateRepository) DeleteStep(stepID models.StepID) (err error) {
	err = databaseConnection.Delete(&models.Step{ID: stepID}).Error
	return
}

func (rep *GormWorkflowTemplateRepository) CreateNextStep(nextStep *models.NextStep) (err error) {
	err = databaseConnection.Create(nextStep).Error
	return
}

func (rep *GormWorkflowTemplateRepository) UpdateNextStep(nextStep *models.NextStep) (err error) {
	err = databaseConnection.Save(nextStep).Error
	return
}

func (rep *GormWorkflowTemplateRepository) DeleteNextStep(prevStepID models.StepID, nextStepID models.StepID) (err error) {
	err = databaseConnection.Delete(&models.NextStep{PrevID: prevStepID, NextID: nextStepID}).Error
	return
}

func (rep *GormWorkflowTemplateRepository) GetLinearNextStepID(stepID models.StepID) (nextStepID models.StepID, err error) {
	nextStep := &models.NextStep{}
	err = databaseConnection.Where("prev_id = ? AND decision_label = \"\"", stepID).First(nextStep).Error
	nextStepID = nextStep.NextID
	return
}

func (rep *GormWorkflowTemplateRepository) GetNextStepByNextID(stepID models.StepID) (nextStep *models.NextStep, err error) {
	nextStep = &models.NextStep{}
	err = databaseConnection.First(nextStep, &models.NextStep{NextID: stepID}).Error
	return
}

func (rep *GormWorkflowTemplateRepository) GetNextStepByDecison(stepID models.StepID, decision string) (nextStepID models.StepID, err error) {
	nextStep := &models.NextStep{}
	err = databaseConnection.Where("prev_id = ? AND decision_label = ?", stepID, decision).First(nextStep).Error
	nextStepID = nextStep.NextID
	return
}

func (rep *GormWorkflowTemplateRepository) GetDecisions(stepID models.StepID) (decisions []*models.NextStep, err error) {
	err = databaseConnection.Where("prev_id = ? AND decision_label <> \"\"", stepID).Find(&decisions).Error
	return
}
