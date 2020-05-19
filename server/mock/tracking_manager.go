// Code generated by MockGen. DO NOT EDIT.
// Source: managers/tracking_manager.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	models "paper-tracker/models"
	reflect "reflect"
)

// MockTrackingManager is a mock of TrackingManager interface
type MockTrackingManager struct {
	ctrl     *gomock.Controller
	recorder *MockTrackingManagerMockRecorder
}

// MockTrackingManagerMockRecorder is the mock recorder for MockTrackingManager
type MockTrackingManagerMockRecorder struct {
	mock *MockTrackingManager
}

// NewMockTrackingManager creates a new mock instance
func NewMockTrackingManager(ctrl *gomock.Controller) *MockTrackingManager {
	mock := &MockTrackingManager{ctrl: ctrl}
	mock.recorder = &MockTrackingManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTrackingManager) EXPECT() *MockTrackingManagerMockRecorder {
	return m.recorder
}

// GetRoomMatchingBest mocks base method
func (m *MockTrackingManager) GetRoomMatchingBest(scoredRooms []map[models.RoomID]float64) models.RoomID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoomMatchingBest", scoredRooms)
	ret0, _ := ret[0].(models.RoomID)
	return ret0
}

// GetRoomMatchingBest indicates an expected call of GetRoomMatchingBest
func (mr *MockTrackingManagerMockRecorder) GetRoomMatchingBest(scoredRooms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoomMatchingBest", reflect.TypeOf((*MockTrackingManager)(nil).GetRoomMatchingBest), scoredRooms)
}

// ConsolidateScanResults mocks base method
func (m *MockTrackingManager) ConsolidateScanResults(scanResults []*models.ScanResult) []models.BSSIDTrackingData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsolidateScanResults", scanResults)
	ret0, _ := ret[0].([]models.BSSIDTrackingData)
	return ret0
}

// ConsolidateScanResults indicates an expected call of ConsolidateScanResults
func (mr *MockTrackingManagerMockRecorder) ConsolidateScanResults(scanResults interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsolidateScanResults", reflect.TypeOf((*MockTrackingManager)(nil).ConsolidateScanResults), scanResults)
}

// ScoreRoomsForScanResults mocks base method
func (m *MockTrackingManager) ScoreRoomsForScanResults(rooms []*models.Room, scanResults []*models.ScanResult) map[models.RoomID]float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScoreRoomsForScanResults", rooms, scanResults)
	ret0, _ := ret[0].(map[models.RoomID]float64)
	return ret0
}

// ScoreRoomsForScanResults indicates an expected call of ScoreRoomsForScanResults
func (mr *MockTrackingManagerMockRecorder) ScoreRoomsForScanResults(rooms, scanResults interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScoreRoomsForScanResults", reflect.TypeOf((*MockTrackingManager)(nil).ScoreRoomsForScanResults), rooms, scanResults)
}
