// Code generated by MockGen. DO NOT EDIT.
// Source: managers/tracker_manager.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	models "paper-tracker/models"
	communication "paper-tracker/models/communication"
	reflect "reflect"
	time "time"
)

// MockTrackerManager is a mock of TrackerManager interface
type MockTrackerManager struct {
	ctrl     *gomock.Controller
	recorder *MockTrackerManagerMockRecorder
}

// MockTrackerManagerMockRecorder is the mock recorder for MockTrackerManager
type MockTrackerManagerMockRecorder struct {
	mock *MockTrackerManager
}

// NewMockTrackerManager creates a new mock instance
func NewMockTrackerManager(ctrl *gomock.Controller) *MockTrackerManager {
	mock := &MockTrackerManager{ctrl: ctrl}
	mock.recorder = &MockTrackerManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTrackerManager) EXPECT() *MockTrackerManagerMockRecorder {
	return m.recorder
}

// GetTrackerByID mocks base method
func (m *MockTrackerManager) GetTrackerByID(trackerID models.TrackerID) (*models.Tracker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrackerByID", trackerID)
	ret0, _ := ret[0].(*models.Tracker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrackerByID indicates an expected call of GetTrackerByID
func (mr *MockTrackerManagerMockRecorder) GetTrackerByID(trackerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrackerByID", reflect.TypeOf((*MockTrackerManager)(nil).GetTrackerByID), trackerID)
}

// GetAllTrackers mocks base method
func (m *MockTrackerManager) GetAllTrackers() ([]*models.Tracker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTrackers")
	ret0, _ := ret[0].([]*models.Tracker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTrackers indicates an expected call of GetAllTrackers
func (mr *MockTrackerManagerMockRecorder) GetAllTrackers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTrackers", reflect.TypeOf((*MockTrackerManager)(nil).GetAllTrackers))
}

// SetTrackerStatus mocks base method
func (m *MockTrackerManager) SetTrackerStatus(trackerID models.TrackerID, status models.TrackerStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTrackerStatus", trackerID, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTrackerStatus indicates an expected call of SetTrackerStatus
func (mr *MockTrackerManagerMockRecorder) SetTrackerStatus(trackerID, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrackerStatus", reflect.TypeOf((*MockTrackerManager)(nil).SetTrackerStatus), trackerID, status)
}

// UpdateTrackerLabel mocks base method
func (m *MockTrackerManager) UpdateTrackerLabel(trackerID models.TrackerID, label string) (*models.Tracker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTrackerLabel", trackerID, label)
	ret0, _ := ret[0].(*models.Tracker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTrackerLabel indicates an expected call of UpdateTrackerLabel
func (mr *MockTrackerManagerMockRecorder) UpdateTrackerLabel(trackerID, label interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTrackerLabel", reflect.TypeOf((*MockTrackerManager)(nil).UpdateTrackerLabel), trackerID, label)
}

// DeleteTracker mocks base method
func (m *MockTrackerManager) DeleteTracker(trackerID models.TrackerID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTracker", trackerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTracker indicates an expected call of DeleteTracker
func (mr *MockTrackerManagerMockRecorder) DeleteTracker(trackerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTracker", reflect.TypeOf((*MockTrackerManager)(nil).DeleteTracker), trackerID)
}

// NotifyNewTracker mocks base method
func (m *MockTrackerManager) NotifyNewTracker() (*models.Tracker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyNewTracker")
	ret0, _ := ret[0].(*models.Tracker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotifyNewTracker indicates an expected call of NotifyNewTracker
func (mr *MockTrackerManagerMockRecorder) NotifyNewTracker() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyNewTracker", reflect.TypeOf((*MockTrackerManager)(nil).NotifyNewTracker))
}

// PollCommand mocks base method
func (m *MockTrackerManager) PollCommand(trackerID models.TrackerID) (*models.Command, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PollCommand", trackerID)
	ret0, _ := ret[0].(*models.Command)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PollCommand indicates an expected call of PollCommand
func (mr *MockTrackerManagerMockRecorder) PollCommand(trackerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PollCommand", reflect.TypeOf((*MockTrackerManager)(nil).PollCommand), trackerID)
}

// InWorkingHours mocks base method
func (m *MockTrackerManager) InWorkingHours(currentTime time.Time) (bool, time.Duration) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InWorkingHours", currentTime)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(time.Duration)
	return ret0, ret1
}

// InWorkingHours indicates an expected call of InWorkingHours
func (mr *MockTrackerManagerMockRecorder) InWorkingHours(currentTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InWorkingHours", reflect.TypeOf((*MockTrackerManager)(nil).InWorkingHours), currentTime)
}

// UpdateFromResponse mocks base method
func (m *MockTrackerManager) UpdateFromResponse(trackerID models.TrackerID, resp communication.TrackerCmdResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFromResponse", trackerID, resp)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFromResponse indicates an expected call of UpdateFromResponse
func (mr *MockTrackerManagerMockRecorder) UpdateFromResponse(trackerID, resp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFromResponse", reflect.TypeOf((*MockTrackerManager)(nil).UpdateFromResponse), trackerID, resp)
}

// UpdateRoom mocks base method
func (m *MockTrackerManager) UpdateRoom(tracker *models.Tracker, roomID models.RoomID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRoom", tracker, roomID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRoom indicates an expected call of UpdateRoom
func (mr *MockTrackerManagerMockRecorder) UpdateRoom(tracker, roomID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRoom", reflect.TypeOf((*MockTrackerManager)(nil).UpdateRoom), tracker, roomID)
}

// NewTrackingData mocks base method
func (m *MockTrackerManager) NewTrackingData(trackerID models.TrackerID, resultID uint64, batchCount uint8, scanRes []*models.ScanResult) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewTrackingData", trackerID, resultID, batchCount, scanRes)
	ret0, _ := ret[0].(error)
	return ret0
}

// NewTrackingData indicates an expected call of NewTrackingData
func (mr *MockTrackerManagerMockRecorder) NewTrackingData(trackerID, resultID, batchCount, scanRes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTrackingData", reflect.TypeOf((*MockTrackerManager)(nil).NewTrackingData), trackerID, resultID, batchCount, scanRes)
}
