// Code generated by MockGen. DO NOT EDIT.
// Source: .\proposition_post_case.go

// Package mock_proposition_post_case is a generated GoMock package.
package mock_proposition_post_case

import (
	proposition_post_case "SynchronizeMonorevoDeliveryDates/usecase/proposition_post_case"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPostingExecutor is a mock of PostingExecutor interface.
type MockPostingExecutor struct {
	ctrl     *gomock.Controller
	recorder *MockPostingExecutorMockRecorder
}

// MockPostingExecutorMockRecorder is the mock recorder for MockPostingExecutor.
type MockPostingExecutorMockRecorder struct {
	mock *MockPostingExecutor
}

// NewMockPostingExecutor creates a new mock instance.
func NewMockPostingExecutor(ctrl *gomock.Controller) *MockPostingExecutor {
	mock := &MockPostingExecutor{ctrl: ctrl}
	mock.recorder = &MockPostingExecutorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostingExecutor) EXPECT() *MockPostingExecutorMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockPostingExecutor) Execute(arg0 []proposition_post_case.PostingPropositionPram) ([]proposition_post_case.PostedPropositionDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arg0)
	ret0, _ := ret[0].([]proposition_post_case.PostedPropositionDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockPostingExecutorMockRecorder) Execute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockPostingExecutor)(nil).Execute), arg0)
}
