// Code generated by MockGen. DO NOT EDIT.
// Source: repositories/room_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	models "paper-tracker/models"
	reflect "reflect"
)

// MockRoomRepository is a mock of RoomRepository interface
type MockRoomRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRoomRepositoryMockRecorder
}

// MockRoomRepositoryMockRecorder is the mock recorder for MockRoomRepository
type MockRoomRepositoryMockRecorder struct {
	mock *MockRoomRepository
}

// NewMockRoomRepository creates a new mock instance
func NewMockRoomRepository(ctrl *gomock.Controller) *MockRoomRepository {
	mock := &MockRoomRepository{ctrl: ctrl}
	mock.recorder = &MockRoomRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRoomRepository) EXPECT() *MockRoomRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockRoomRepository) Create(room *models.Room) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", room)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockRoomRepositoryMockRecorder) Create(room interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRoomRepository)(nil).Create), room)
}

// GetByID mocks base method
func (m *MockRoomRepository) GetByID(roomID models.RoomID) (*models.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", roomID)
	ret0, _ := ret[0].(*models.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockRoomRepositoryMockRecorder) GetByID(roomID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockRoomRepository)(nil).GetByID), roomID)
}

// GetAll mocks base method
func (m *MockRoomRepository) GetAll() ([]*models.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*models.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockRoomRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockRoomRepository)(nil).GetAll))
}

// Delete mocks base method
func (m *MockRoomRepository) Delete(roomID models.RoomID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", roomID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockRoomRepositoryMockRecorder) Delete(roomID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRoomRepository)(nil).Delete), roomID)
}

// SetLearnedByID mocks base method
func (m *MockRoomRepository) SetLearnedByID(roomID models.RoomID, learned bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLearnedByID", roomID, learned)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLearnedByID indicates an expected call of SetLearnedByID
func (mr *MockRoomRepositoryMockRecorder) SetLearnedByID(roomID, learned interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLearnedByID", reflect.TypeOf((*MockRoomRepository)(nil).SetLearnedByID), roomID, learned)
}

// Update mocks base method
func (m *MockRoomRepository) Update(room *models.Room) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", room)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockRoomRepositoryMockRecorder) Update(room interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRoomRepository)(nil).Update), room)
}

// IsRecordNotFoundError mocks base method
func (m *MockRoomRepository) IsRecordNotFoundError(err error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRecordNotFoundError", err)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRecordNotFoundError indicates an expected call of IsRecordNotFoundError
func (mr *MockRoomRepositoryMockRecorder) IsRecordNotFoundError(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRecordNotFoundError", reflect.TypeOf((*MockRoomRepository)(nil).IsRecordNotFoundError), err)
}
