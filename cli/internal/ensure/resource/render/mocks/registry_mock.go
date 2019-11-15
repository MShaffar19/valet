// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/valet/cli/internal/ensure/resource/render (interfaces: Registry)

// Package mock_render is a generated GoMock package.
package mock_render

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRegistry is a mock of Registry interface
type MockRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryMockRecorder
}

// MockRegistryMockRecorder is the mock recorder for MockRegistry
type MockRegistryMockRecorder struct {
	mock *MockRegistry
}

// NewMockRegistry creates a new mock instance
func NewMockRegistry(ctrl *gomock.Controller) *MockRegistry {
	mock := &MockRegistry{ctrl: ctrl}
	mock.recorder = &MockRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegistry) EXPECT() *MockRegistryMockRecorder {
	return m.recorder
}

// LoadFile mocks base method
func (m *MockRegistry) LoadFile(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadFile", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadFile indicates an expected call of LoadFile
func (mr *MockRegistryMockRecorder) LoadFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadFile", reflect.TypeOf((*MockRegistry)(nil).LoadFile), arg0)
}