// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/valet/cli/internal/ensure/client (interfaces: IngressClient)

// Package mock_client is a generated GoMock package.
package mock_client

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIngressClient is a mock of IngressClient interface
type MockIngressClient struct {
	ctrl     *gomock.Controller
	recorder *MockIngressClientMockRecorder
}

// MockIngressClientMockRecorder is the mock recorder for MockIngressClient
type MockIngressClientMockRecorder struct {
	mock *MockIngressClient
}

// NewMockIngressClient creates a new mock instance
func NewMockIngressClient(ctrl *gomock.Controller) *MockIngressClient {
	mock := &MockIngressClient{ctrl: ctrl}
	mock.recorder = &MockIngressClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIngressClient) EXPECT() *MockIngressClientMockRecorder {
	return m.recorder
}

// GetIngressAddress mocks base method
func (m *MockIngressClient) GetIngressHost(arg0, arg1, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIngressAddress", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIngressAddress indicates an expected call of GetIngressAddress
func (mr *MockIngressClientMockRecorder) GetIngressHost(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIngressAddress", reflect.TypeOf((*MockIngressClient)(nil).GetIngressHost), arg0, arg1, arg2)
}
