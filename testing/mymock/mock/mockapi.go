// Code generated by MockGen. DO NOT EDIT.
// Source: mymock (interfaces: API)

// Package mock_mymock is a generated GoMock package.
package mock_mymock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAPI is a mock of API interface.
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI.
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance.
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// ConsumeMessage mocks base method.
func (m *MockAPI) ConsumeMessage() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsumeMessage")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConsumeMessage indicates an expected call of ConsumeMessage.
func (mr *MockAPIMockRecorder) ConsumeMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumeMessage", reflect.TypeOf((*MockAPI)(nil).ConsumeMessage))
}

// SendMessage mocks base method.
func (m *MockAPI) SendMessage(msg string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockAPIMockRecorder) SendMessage(msg string) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockAPI)(nil).SendMessage), msg)
}
