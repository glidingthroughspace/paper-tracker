// Code generated by MockGen. DO NOT EDIT.
// Source: repositories/workflow_template_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	models "paper-tracker/models"
	reflect "reflect"
)

// MockWorkflowTemplateRepository is a mock of WorkflowTemplateRepository interface
type MockWorkflowTemplateRepository struct {
	ctrl     *gomock.Controller
	recorder *MockWorkflowTemplateRepositoryMockRecorder
}

// MockWorkflowTemplateRepositoryMockRecorder is the mock recorder for MockWorkflowTemplateRepository
type MockWorkflowTemplateRepositoryMockRecorder struct {
	mock *MockWorkflowTemplateRepository
}

// NewMockWorkflowTemplateRepository creates a new mock instance
func NewMockWorkflowTemplateRepository(ctrl *gomock.Controller) *MockWorkflowTemplateRepository {
	mock := &MockWorkflowTemplateRepository{ctrl: ctrl}
	mock.recorder = &MockWorkflowTemplateRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWorkflowTemplateRepository) EXPECT() *MockWorkflowTemplateRepositoryMockRecorder {
	return m.recorder
}

// CreateTemplate mocks base method
func (m *MockWorkflowTemplateRepository) CreateTemplate(template *models.WorkflowTemplate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTemplate", template)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTemplate indicates an expected call of CreateTemplate
func (mr *MockWorkflowTemplateRepositoryMockRecorder) CreateTemplate(template interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTemplate", reflect.TypeOf((*MockWorkflowTemplateRepository)(nil).CreateTemplate), template)
}

// GetAllTemplates mocks base method
func (m *MockWorkflowTemplateRepository) GetAllTemplates() ([]*models.WorkflowTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTemplates")
	ret0, _ := ret[0].([]*models.WorkflowTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTemplates indicates an expected call of GetAllTemplates
func (mr *MockWorkflowTemplateRepositoryMockRecorder) GetAllTemplates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTemplates", reflect.TypeOf((*MockWorkflowTemplateRepository)(nil).GetAllTemplates))
}

// GetTemplateByID mocks base method
func (m *MockWorkflowTemplateRepository) GetTemplateByID(templateID models.WorkflowTemplateID) (*models.WorkflowTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTemplateByID", templateID)
	ret0, _ := ret[0].(*models.WorkflowTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTemplateByID indicates an expected call of GetTemplateByID
func (mr *MockWorkflowTemplateRepositoryMockRecorder) GetTemplateByID(templateID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTemplateByID", reflect.TypeOf((*MockWorkflowTemplateRepository)(nil).GetTemplateByID), templateID)
}

// UpdateTemplate mocks base method
func (m *MockWorkflowTemplateRepository) UpdateTemplate(template *models.WorkflowTemplate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTemplate", template)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTemplate indicates an expected call of UpdateTemplate
func (mr *MockWorkflowTemplateRepositoryMockRecorder) UpdateTemplate(template interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTemplate", reflect.TypeOf((*MockWorkflowTemplateRepository)(nil).UpdateTemplate), template)
}

// DeleteTemplate mocks base method
func (m *MockWorkflowTemplateRepository) DeleteTemplate(templateID models.WorkflowTemplateID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTemplate", templateID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTemplate indicates an expected call of DeleteTemplate
func (mr *MockWorkflowTemplateRepositoryMockRecorder) DeleteTemplate(templateID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTemplate", reflect.TypeOf((*MockWorkflowTemplateRepository)(nil).DeleteTemplate), templateID)
}

// CreateStep mocks base method
func (m *MockWorkflowTemplateRepository) CreateStep(step *models.Step) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStep", step)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateStep indicates an expected call of CreateStep
func (mr *MockWorkflowTemplateRepositoryMockRecorder) CreateStep(step interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStep", reflect.TypeOf((*MockWorkflowTemplateRepository)(nil).CreateStep), step)
}

// GetStepByID mocks base method
func (m *MockWorkflowTemplateRepository) GetStepByID(stepID models.StepID) (*models.Step, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStepByID", stepID)
	ret0, _ := ret[0].(*models.Step)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStepByID indicates an expected call of GetStepByID
func (mr *MockWorkflowTemplateRepositoryMockRecorder) GetStepByID(stepID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStepByID", reflect.TypeOf((*MockWorkflowTemplateRepository)(nil).GetStepByID), stepID)
}

// GetStepsByRoomID mocks base method
func (m *MockWorkflowTemplateRepository) GetStepsByRoomID(roomID models.RoomID) ([]*models.Step, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStepsByRoomID", roomID)
	ret0, _ := ret[0].([]*models.Step)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStepsByRoomID indicates an expected call of GetStepsByRoomID
func (mr *MockWorkflowTemplateRepositoryMockRecorder) GetStepsByRoomID(roomID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStepsByRoomID", reflect.TypeOf((*MockWorkflowTemplateRepository)(nil).GetStepsByRoomID), roomID)
}

// UpdateStep mocks base method
func (m *MockWorkflowTemplateRepository) UpdateStep(step *models.Step) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStep", step)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStep indicates an expected call of UpdateStep
func (mr *MockWorkflowTemplateRepositoryMockRecorder) UpdateStep(step interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStep", reflect.TypeOf((*MockWorkflowTemplateRepository)(nil).UpdateStep), step)
}

// DeleteStep mocks base method
func (m *MockWorkflowTemplateRepository) DeleteStep(stepID models.StepID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStep", stepID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStep indicates an expected call of DeleteStep
func (mr *MockWorkflowTemplateRepositoryMockRecorder) DeleteStep(stepID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStep", reflect.TypeOf((*MockWorkflowTemplateRepository)(nil).DeleteStep), stepID)
}

// CreateNextStep mocks base method
func (m *MockWorkflowTemplateRepository) CreateNextStep(nextStep *models.NextStep) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNextStep", nextStep)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNextStep indicates an expected call of CreateNextStep
func (mr *MockWorkflowTemplateRepositoryMockRecorder) CreateNextStep(nextStep interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNextStep", reflect.TypeOf((*MockWorkflowTemplateRepository)(nil).CreateNextStep), nextStep)
}

// UpdateNextStep mocks base method
func (m *MockWorkflowTemplateRepository) UpdateNextStep(nextStep *models.NextStep) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNextStep", nextStep)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateNextStep indicates an expected call of UpdateNextStep
func (mr *MockWorkflowTemplateRepositoryMockRecorder) UpdateNextStep(nextStep interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNextStep", reflect.TypeOf((*MockWorkflowTemplateRepository)(nil).UpdateNextStep), nextStep)
}

// DeleteNextStep mocks base method
func (m *MockWorkflowTemplateRepository) DeleteNextStep(prevStepID, nextStepID models.StepID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNextStep", prevStepID, nextStepID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNextStep indicates an expected call of DeleteNextStep
func (mr *MockWorkflowTemplateRepositoryMockRecorder) DeleteNextStep(prevStepID, nextStepID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNextStep", reflect.TypeOf((*MockWorkflowTemplateRepository)(nil).DeleteNextStep), prevStepID, nextStepID)
}

// GetLinearNextStepID mocks base method
func (m *MockWorkflowTemplateRepository) GetLinearNextStepID(stepID models.StepID) (models.StepID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLinearNextStepID", stepID)
	ret0, _ := ret[0].(models.StepID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLinearNextStepID indicates an expected call of GetLinearNextStepID
func (mr *MockWorkflowTemplateRepositoryMockRecorder) GetLinearNextStepID(stepID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLinearNextStepID", reflect.TypeOf((*MockWorkflowTemplateRepository)(nil).GetLinearNextStepID), stepID)
}

// GetNextStepByNextID mocks base method
func (m *MockWorkflowTemplateRepository) GetNextStepByNextID(stepID models.StepID) (*models.NextStep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextStepByNextID", stepID)
	ret0, _ := ret[0].(*models.NextStep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextStepByNextID indicates an expected call of GetNextStepByNextID
func (mr *MockWorkflowTemplateRepositoryMockRecorder) GetNextStepByNextID(stepID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextStepByNextID", reflect.TypeOf((*MockWorkflowTemplateRepository)(nil).GetNextStepByNextID), stepID)
}

// GetNextStepByDecison mocks base method
func (m *MockWorkflowTemplateRepository) GetNextStepByDecison(stepID models.StepID, decision string) (models.StepID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextStepByDecison", stepID, decision)
	ret0, _ := ret[0].(models.StepID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextStepByDecison indicates an expected call of GetNextStepByDecison
func (mr *MockWorkflowTemplateRepositoryMockRecorder) GetNextStepByDecison(stepID, decision interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextStepByDecison", reflect.TypeOf((*MockWorkflowTemplateRepository)(nil).GetNextStepByDecison), stepID, decision)
}

// GetDecisions mocks base method
func (m *MockWorkflowTemplateRepository) GetDecisions(stepID models.StepID) ([]*models.NextStep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDecisions", stepID)
	ret0, _ := ret[0].([]*models.NextStep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDecisions indicates an expected call of GetDecisions
func (mr *MockWorkflowTemplateRepositoryMockRecorder) GetDecisions(stepID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDecisions", reflect.TypeOf((*MockWorkflowTemplateRepository)(nil).GetDecisions), stepID)
}

// IsRecordNotFoundError mocks base method
func (m *MockWorkflowTemplateRepository) IsRecordNotFoundError(err error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRecordNotFoundError", err)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRecordNotFoundError indicates an expected call of IsRecordNotFoundError
func (mr *MockWorkflowTemplateRepositoryMockRecorder) IsRecordNotFoundError(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRecordNotFoundError", reflect.TypeOf((*MockWorkflowTemplateRepository)(nil).IsRecordNotFoundError), err)
}
