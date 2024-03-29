// Code generated by MockGen. DO NOT EDIT.
// Source: managers/learning_manager.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	models "paper-tracker/models"
	reflect "reflect"
)

// MockLearningManager is a mock of LearningManager interface
type MockLearningManager struct {
	ctrl     *gomock.Controller
	recorder *MockLearningManagerMockRecorder
}

// MockLearningManagerMockRecorder is the mock recorder for MockLearningManager
type MockLearningManagerMockRecorder struct {
	mock *MockLearningManager
}

// NewMockLearningManager creates a new mock instance
func NewMockLearningManager(ctrl *gomock.Controller) *MockLearningManager {
	mock := &MockLearningManager{ctrl: ctrl}
	mock.recorder = &MockLearningManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLearningManager) EXPECT() *MockLearningManagerMockRecorder {
	return m.recorder
}

// StartLearning mocks base method
func (m *MockLearningManager) StartLearning(trackerID models.TrackerID) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartLearning", trackerID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartLearning indicates an expected call of StartLearning
func (mr *MockLearningManagerMockRecorder) StartLearning(trackerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartLearning", reflect.TypeOf((*MockLearningManager)(nil).StartLearning), trackerID)
}

// FinishLearning mocks base method
func (m *MockLearningManager) FinishLearning(trackerID models.TrackerID, roomID models.RoomID, ssids []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinishLearning", trackerID, roomID, ssids)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinishLearning indicates an expected call of FinishLearning
func (mr *MockLearningManagerMockRecorder) FinishLearning(trackerID, roomID, ssids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinishLearning", reflect.TypeOf((*MockLearningManager)(nil).FinishLearning), trackerID, roomID, ssids)
}

// GetLearningStatus mocks base method
func (m *MockLearningManager) GetLearningStatus(trackerID models.TrackerID) (bool, []string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLearningStatus", trackerID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetLearningStatus indicates an expected call of GetLearningStatus
func (mr *MockLearningManagerMockRecorder) GetLearningStatus(trackerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLearningStatus", reflect.TypeOf((*MockLearningManager)(nil).GetLearningStatus), trackerID)
}

// CancelLearning mocks base method
func (m *MockLearningManager) CancelLearning(trackerID models.TrackerID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelLearning", trackerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelLearning indicates an expected call of CancelLearning
func (mr *MockLearningManagerMockRecorder) CancelLearning(trackerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelLearning", reflect.TypeOf((*MockLearningManager)(nil).CancelLearning), trackerID)
}

// NewLearningTrackingData mocks base method
func (m *MockLearningManager) NewLearningTrackingData(trackerID models.TrackerID, scanRes []*models.ScanResult) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewLearningTrackingData", trackerID, scanRes)
	ret0, _ := ret[0].(error)
	return ret0
}

// NewLearningTrackingData indicates an expected call of NewLearningTrackingData
func (mr *MockLearningManagerMockRecorder) NewLearningTrackingData(trackerID, scanRes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewLearningTrackingData", reflect.TypeOf((*MockLearningManager)(nil).NewLearningTrackingData), trackerID, scanRes)
}
