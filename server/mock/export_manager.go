// Code generated by MockGen. DO NOT EDIT.
// Source: managers/export_manager.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockExportManager is a mock of ExportManager interface
type MockExportManager struct {
	ctrl     *gomock.Controller
	recorder *MockExportManagerMockRecorder
}

// MockExportManagerMockRecorder is the mock recorder for MockExportManager
type MockExportManagerMockRecorder struct {
	mock *MockExportManager
}

// NewMockExportManager creates a new mock instance
func NewMockExportManager(ctrl *gomock.Controller) *MockExportManager {
	mock := &MockExportManager{ctrl: ctrl}
	mock.recorder = &MockExportManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExportManager) EXPECT() *MockExportManagerMockRecorder {
	return m.recorder
}

// GenerateExport mocks base method
func (m *MockExportManager) GenerateExport(writer io.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateExport", writer)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateExport indicates an expected call of GenerateExport
func (mr *MockExportManagerMockRecorder) GenerateExport(writer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateExport", reflect.TypeOf((*MockExportManager)(nil).GenerateExport), writer)
}
