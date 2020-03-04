// Code generated by MockGen. DO NOT EDIT.
// Source: repositories/workflow_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	models "paper-tracker/models"
	reflect "reflect"
)

// MockWorkflowRepository is a mock of WorkflowRepository interface
type MockWorkflowRepository struct {
	ctrl     *gomock.Controller
	recorder *MockWorkflowRepositoryMockRecorder
}

// MockWorkflowRepositoryMockRecorder is the mock recorder for MockWorkflowRepository
type MockWorkflowRepositoryMockRecorder struct {
	mock *MockWorkflowRepository
}

// NewMockWorkflowRepository creates a new mock instance
func NewMockWorkflowRepository(ctrl *gomock.Controller) *MockWorkflowRepository {
	mock := &MockWorkflowRepository{ctrl: ctrl}
	mock.recorder = &MockWorkflowRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWorkflowRepository) EXPECT() *MockWorkflowRepositoryMockRecorder {
	return m.recorder
}

// CreateTemplate mocks base method
func (m *MockWorkflowRepository) CreateTemplate(template *models.WorkflowTemplate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTemplate", template)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTemplate indicates an expected call of CreateTemplate
func (mr *MockWorkflowRepositoryMockRecorder) CreateTemplate(template interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTemplate", reflect.TypeOf((*MockWorkflowRepository)(nil).CreateTemplate), template)
}

// GetAllTemplates mocks base method
func (m *MockWorkflowRepository) GetAllTemplates() ([]*models.WorkflowTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTemplates")
	ret0, _ := ret[0].([]*models.WorkflowTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTemplates indicates an expected call of GetAllTemplates
func (mr *MockWorkflowRepositoryMockRecorder) GetAllTemplates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTemplates", reflect.TypeOf((*MockWorkflowRepository)(nil).GetAllTemplates))
}

// GetTemplateByID mocks base method
func (m *MockWorkflowRepository) GetTemplateByID(templateID models.WorkflowTemplateID) (*models.WorkflowTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTemplateByID", templateID)
	ret0, _ := ret[0].(*models.WorkflowTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTemplateByID indicates an expected call of GetTemplateByID
func (mr *MockWorkflowRepositoryMockRecorder) GetTemplateByID(templateID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTemplateByID", reflect.TypeOf((*MockWorkflowRepository)(nil).GetTemplateByID), templateID)
}

// UpdateTemplate mocks base method
func (m *MockWorkflowRepository) UpdateTemplate(template *models.WorkflowTemplate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTemplate", template)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTemplate indicates an expected call of UpdateTemplate
func (mr *MockWorkflowRepositoryMockRecorder) UpdateTemplate(template interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTemplate", reflect.TypeOf((*MockWorkflowRepository)(nil).UpdateTemplate), template)
}

// DeleteTemplate mocks base method
func (m *MockWorkflowRepository) DeleteTemplate(templateID models.WorkflowTemplateID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTemplate", templateID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTemplate indicates an expected call of DeleteTemplate
func (mr *MockWorkflowRepositoryMockRecorder) DeleteTemplate(templateID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTemplate", reflect.TypeOf((*MockWorkflowRepository)(nil).DeleteTemplate), templateID)
}

// CreateStep mocks base method
func (m *MockWorkflowRepository) CreateStep(step *models.Step) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStep", step)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateStep indicates an expected call of CreateStep
func (mr *MockWorkflowRepositoryMockRecorder) CreateStep(step interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStep", reflect.TypeOf((*MockWorkflowRepository)(nil).CreateStep), step)
}

// GetStepByID mocks base method
func (m *MockWorkflowRepository) GetStepByID(stepID models.StepID) (*models.Step, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStepByID", stepID)
	ret0, _ := ret[0].(*models.Step)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStepByID indicates an expected call of GetStepByID
func (mr *MockWorkflowRepositoryMockRecorder) GetStepByID(stepID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStepByID", reflect.TypeOf((*MockWorkflowRepository)(nil).GetStepByID), stepID)
}

// UpdateStep mocks base method
func (m *MockWorkflowRepository) UpdateStep(step *models.Step) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStep", step)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStep indicates an expected call of UpdateStep
func (mr *MockWorkflowRepositoryMockRecorder) UpdateStep(step interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStep", reflect.TypeOf((*MockWorkflowRepository)(nil).UpdateStep), step)
}

// DeleteStep mocks base method
func (m *MockWorkflowRepository) DeleteStep(stepID models.StepID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStep", stepID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStep indicates an expected call of DeleteStep
func (mr *MockWorkflowRepositoryMockRecorder) DeleteStep(stepID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStep", reflect.TypeOf((*MockWorkflowRepository)(nil).DeleteStep), stepID)
}

// CreateNextStep mocks base method
func (m *MockWorkflowRepository) CreateNextStep(nextStep *models.NextStep) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNextStep", nextStep)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNextStep indicates an expected call of CreateNextStep
func (mr *MockWorkflowRepositoryMockRecorder) CreateNextStep(nextStep interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNextStep", reflect.TypeOf((*MockWorkflowRepository)(nil).CreateNextStep), nextStep)
}

// UpdateNextStep mocks base method
func (m *MockWorkflowRepository) UpdateNextStep(nextStep *models.NextStep) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNextStep", nextStep)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateNextStep indicates an expected call of UpdateNextStep
func (mr *MockWorkflowRepositoryMockRecorder) UpdateNextStep(nextStep interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNextStep", reflect.TypeOf((*MockWorkflowRepository)(nil).UpdateNextStep), nextStep)
}

// DeleteNextStep mocks base method
func (m *MockWorkflowRepository) DeleteNextStep(prevStepID, nextStepID models.StepID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNextStep", prevStepID, nextStepID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNextStep indicates an expected call of DeleteNextStep
func (mr *MockWorkflowRepositoryMockRecorder) DeleteNextStep(prevStepID, nextStepID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNextStep", reflect.TypeOf((*MockWorkflowRepository)(nil).DeleteNextStep), prevStepID, nextStepID)
}

// GetLinearNextStepID mocks base method
func (m *MockWorkflowRepository) GetLinearNextStepID(stepID models.StepID) (models.StepID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLinearNextStepID", stepID)
	ret0, _ := ret[0].(models.StepID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLinearNextStepID indicates an expected call of GetLinearNextStepID
func (mr *MockWorkflowRepositoryMockRecorder) GetLinearNextStepID(stepID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLinearNextStepID", reflect.TypeOf((*MockWorkflowRepository)(nil).GetLinearNextStepID), stepID)
}

// GetNextStepByNextID mocks base method
func (m *MockWorkflowRepository) GetNextStepByNextID(stepID models.StepID) (*models.NextStep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextStepByNextID", stepID)
	ret0, _ := ret[0].(*models.NextStep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextStepByNextID indicates an expected call of GetNextStepByNextID
func (mr *MockWorkflowRepositoryMockRecorder) GetNextStepByNextID(stepID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextStepByNextID", reflect.TypeOf((*MockWorkflowRepository)(nil).GetNextStepByNextID), stepID)
}

// GetNextStepByDecison mocks base method
func (m *MockWorkflowRepository) GetNextStepByDecison(stepID models.StepID, decision string) (models.StepID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextStepByDecison", stepID, decision)
	ret0, _ := ret[0].(models.StepID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextStepByDecison indicates an expected call of GetNextStepByDecison
func (mr *MockWorkflowRepositoryMockRecorder) GetNextStepByDecison(stepID, decision interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextStepByDecison", reflect.TypeOf((*MockWorkflowRepository)(nil).GetNextStepByDecison), stepID, decision)
}

// GetDecisions mocks base method
func (m *MockWorkflowRepository) GetDecisions(stepID models.StepID) ([]*models.NextStep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDecisions", stepID)
	ret0, _ := ret[0].([]*models.NextStep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDecisions indicates an expected call of GetDecisions
func (mr *MockWorkflowRepositoryMockRecorder) GetDecisions(stepID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDecisions", reflect.TypeOf((*MockWorkflowRepository)(nil).GetDecisions), stepID)
}

// CreateExec mocks base method
func (m *MockWorkflowRepository) CreateExec(exec *models.WorkflowExec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExec", exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateExec indicates an expected call of CreateExec
func (mr *MockWorkflowRepositoryMockRecorder) CreateExec(exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExec", reflect.TypeOf((*MockWorkflowRepository)(nil).CreateExec), exec)
}

// GetAllExec mocks base method
func (m *MockWorkflowRepository) GetAllExec() ([]*models.WorkflowExec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllExec")
	ret0, _ := ret[0].([]*models.WorkflowExec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllExec indicates an expected call of GetAllExec
func (mr *MockWorkflowRepositoryMockRecorder) GetAllExec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllExec", reflect.TypeOf((*MockWorkflowRepository)(nil).GetAllExec))
}

// GetExecByID mocks base method
func (m *MockWorkflowRepository) GetExecByID(execID models.WorkflowExecID) (*models.WorkflowExec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExecByID", execID)
	ret0, _ := ret[0].(*models.WorkflowExec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExecByID indicates an expected call of GetExecByID
func (mr *MockWorkflowRepositoryMockRecorder) GetExecByID(execID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExecByID", reflect.TypeOf((*MockWorkflowRepository)(nil).GetExecByID), execID)
}

// GetRunningExecByTrackerID mocks base method
func (m *MockWorkflowRepository) GetRunningExecByTrackerID(trackerID models.TrackerID) (*models.WorkflowExec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunningExecByTrackerID", trackerID)
	ret0, _ := ret[0].(*models.WorkflowExec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRunningExecByTrackerID indicates an expected call of GetRunningExecByTrackerID
func (mr *MockWorkflowRepositoryMockRecorder) GetRunningExecByTrackerID(trackerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunningExecByTrackerID", reflect.TypeOf((*MockWorkflowRepository)(nil).GetRunningExecByTrackerID), trackerID)
}

// GetExecsByTemplateID mocks base method
func (m *MockWorkflowRepository) GetExecsByTemplateID(templateID models.WorkflowTemplateID) ([]*models.WorkflowExec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExecsByTemplateID", templateID)
	ret0, _ := ret[0].([]*models.WorkflowExec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExecsByTemplateID indicates an expected call of GetExecsByTemplateID
func (mr *MockWorkflowRepositoryMockRecorder) GetExecsByTemplateID(templateID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExecsByTemplateID", reflect.TypeOf((*MockWorkflowRepository)(nil).GetExecsByTemplateID), templateID)
}

// UpdateExec mocks base method
func (m *MockWorkflowRepository) UpdateExec(exec *models.WorkflowExec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExec", exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExec indicates an expected call of UpdateExec
func (mr *MockWorkflowRepositoryMockRecorder) UpdateExec(exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExec", reflect.TypeOf((*MockWorkflowRepository)(nil).UpdateExec), exec)
}

// DeleteExec mocks base method
func (m *MockWorkflowRepository) DeleteExec(execID models.WorkflowExecID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExec", execID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExec indicates an expected call of DeleteExec
func (mr *MockWorkflowRepositoryMockRecorder) DeleteExec(execID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExec", reflect.TypeOf((*MockWorkflowRepository)(nil).DeleteExec), execID)
}

// CreateExecStepInfo mocks base method
func (m *MockWorkflowRepository) CreateExecStepInfo(execStepInfo *models.ExecStepInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExecStepInfo", execStepInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateExecStepInfo indicates an expected call of CreateExecStepInfo
func (mr *MockWorkflowRepositoryMockRecorder) CreateExecStepInfo(execStepInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExecStepInfo", reflect.TypeOf((*MockWorkflowRepository)(nil).CreateExecStepInfo), execStepInfo)
}

// GetExecStepInfoByID mocks base method
func (m *MockWorkflowRepository) GetExecStepInfoByID(execID models.WorkflowExecID, stepID models.StepID) (*models.ExecStepInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExecStepInfoByID", execID, stepID)
	ret0, _ := ret[0].(*models.ExecStepInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExecStepInfoByID indicates an expected call of GetExecStepInfoByID
func (mr *MockWorkflowRepositoryMockRecorder) GetExecStepInfoByID(execID, stepID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExecStepInfoByID", reflect.TypeOf((*MockWorkflowRepository)(nil).GetExecStepInfoByID), execID, stepID)
}

// GetExecStepInfoForExecID mocks base method
func (m *MockWorkflowRepository) GetExecStepInfoForExecID(execID models.WorkflowExecID) ([]*models.ExecStepInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExecStepInfoForExecID", execID)
	ret0, _ := ret[0].([]*models.ExecStepInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExecStepInfoForExecID indicates an expected call of GetExecStepInfoForExecID
func (mr *MockWorkflowRepositoryMockRecorder) GetExecStepInfoForExecID(execID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExecStepInfoForExecID", reflect.TypeOf((*MockWorkflowRepository)(nil).GetExecStepInfoForExecID), execID)
}

// UpdateExecStepInfo mocks base method
func (m *MockWorkflowRepository) UpdateExecStepInfo(execStepInfo *models.ExecStepInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExecStepInfo", execStepInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExecStepInfo indicates an expected call of UpdateExecStepInfo
func (mr *MockWorkflowRepositoryMockRecorder) UpdateExecStepInfo(execStepInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExecStepInfo", reflect.TypeOf((*MockWorkflowRepository)(nil).UpdateExecStepInfo), execStepInfo)
}

// DeleteExecStepInfo mocks base method
func (m *MockWorkflowRepository) DeleteExecStepInfo(execID models.WorkflowExecID, stepID models.StepID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExecStepInfo", execID, stepID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExecStepInfo indicates an expected call of DeleteExecStepInfo
func (mr *MockWorkflowRepositoryMockRecorder) DeleteExecStepInfo(execID, stepID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExecStepInfo", reflect.TypeOf((*MockWorkflowRepository)(nil).DeleteExecStepInfo), execID, stepID)
}

// IsRecordNotFoundError mocks base method
func (m *MockWorkflowRepository) IsRecordNotFoundError(err error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRecordNotFoundError", err)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRecordNotFoundError indicates an expected call of IsRecordNotFoundError
func (mr *MockWorkflowRepositoryMockRecorder) IsRecordNotFoundError(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRecordNotFoundError", reflect.TypeOf((*MockWorkflowRepository)(nil).IsRecordNotFoundError), err)
}
