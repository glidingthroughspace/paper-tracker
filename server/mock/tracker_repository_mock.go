// Code generated by MockGen. DO NOT EDIT.
// Source: repositories/tracker_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	models "paper-tracker/models"
	reflect "reflect"
)

// MockTrackerRepository is a mock of TrackerRepository interface
type MockTrackerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTrackerRepositoryMockRecorder
}

// MockTrackerRepositoryMockRecorder is the mock recorder for MockTrackerRepository
type MockTrackerRepositoryMockRecorder struct {
	mock *MockTrackerRepository
}

// NewMockTrackerRepository creates a new mock instance
func NewMockTrackerRepository(ctrl *gomock.Controller) *MockTrackerRepository {
	mock := &MockTrackerRepository{ctrl: ctrl}
	mock.recorder = &MockTrackerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTrackerRepository) EXPECT() *MockTrackerRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockTrackerRepository) Create(tracker *models.Tracker) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", tracker)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockTrackerRepositoryMockRecorder) Create(tracker interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTrackerRepository)(nil).Create), tracker)
}

// GetAll mocks base method
func (m *MockTrackerRepository) GetAll() ([]*models.Tracker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*models.Tracker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockTrackerRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockTrackerRepository)(nil).GetAll))
}

// GetByID mocks base method
func (m *MockTrackerRepository) GetByID(trackerID models.TrackerID) (*models.Tracker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", trackerID)
	ret0, _ := ret[0].(*models.Tracker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockTrackerRepositoryMockRecorder) GetByID(trackerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockTrackerRepository)(nil).GetByID), trackerID)
}

// Update mocks base method
func (m *MockTrackerRepository) Update(tracker *models.Tracker) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", tracker)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockTrackerRepositoryMockRecorder) Update(tracker interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTrackerRepository)(nil).Update), tracker)
}

// UpdateLastPoll mocks base method
func (m *MockTrackerRepository) UpdateLastPoll(tracker *models.Tracker) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLastPoll", tracker)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLastPoll indicates an expected call of UpdateLastPoll
func (mr *MockTrackerRepositoryMockRecorder) UpdateLastPoll(tracker interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLastPoll", reflect.TypeOf((*MockTrackerRepository)(nil).UpdateLastPoll), tracker)
}

// SetStatusByID mocks base method
func (m *MockTrackerRepository) SetStatusByID(trackerID models.TrackerID, status models.TrackerStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStatusByID", trackerID, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStatusByID indicates an expected call of SetStatusByID
func (mr *MockTrackerRepositoryMockRecorder) SetStatusByID(trackerID, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatusByID", reflect.TypeOf((*MockTrackerRepository)(nil).SetStatusByID), trackerID, status)
}

// Delete mocks base method
func (m *MockTrackerRepository) Delete(trackerID models.TrackerID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", trackerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockTrackerRepositoryMockRecorder) Delete(trackerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTrackerRepository)(nil).Delete), trackerID)
}

// IsRecordNotFoundError mocks base method
func (m *MockTrackerRepository) IsRecordNotFoundError(err error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRecordNotFoundError", err)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRecordNotFoundError indicates an expected call of IsRecordNotFoundError
func (mr *MockTrackerRepositoryMockRecorder) IsRecordNotFoundError(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRecordNotFoundError", reflect.TypeOf((*MockTrackerRepository)(nil).IsRecordNotFoundError), err)
}
