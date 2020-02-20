package gorm

import "paper-tracker/models"

func init() {
	databaseModels = append(databaseModels, &models.Workflow{})
	databaseModels = append(databaseModels, &models.Step{})
	databaseModels = append(databaseModels, &models.NextStep{})
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

func (rep *GormWorkflowRepository) CreateWorkflow(workflow *models.Workflow) (err error) {
	err = databaseConnection.Create(workflow).Error
	return
}

func (rep *GormWorkflowRepository) GetAllWorkflows() (workflows []*models.Workflow, err error) {
	err = databaseConnection.Find(&workflows).Error
	return
}

func (rep *GormWorkflowRepository) GetWorkflowByID(workflowID models.WorkflowID) (workflow *models.Workflow, err error) {
	workflow = &models.Workflow{}
	err = databaseConnection.First(workflow, &models.Workflow{ID: workflowID}).Error
	return
}

func (rep *GormWorkflowRepository) UpdateWorkflow(workflow *models.Workflow) (err error) {
	err = databaseConnection.Save(workflow).Error
	return
}

func (rep *GormWorkflowRepository) DeleteWorkflow(workflowID models.WorkflowID) (err error) {
	err = databaseConnection.Delete(&models.Workflow{ID: workflowID}).Error
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