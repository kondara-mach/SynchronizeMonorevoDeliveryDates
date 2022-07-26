// Code generated by MockGen. DO NOT EDIT.
// Source: .\jobbook_fetcher.go

// Package mock_orderdb is a generated GoMock package.
package mock_orderdb

import (
	orderdb "SynchronizeMonorevoDeliveryDates/domain/orderdb"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockJobBookFetcher is a mock of JobBookFetcher interface.
type MockJobBookFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockJobBookFetcherMockRecorder
}

// MockJobBookFetcherMockRecorder is the mock recorder for MockJobBookFetcher.
type MockJobBookFetcherMockRecorder struct {
	mock *MockJobBookFetcher
}

// NewMockJobBookFetcher creates a new mock instance.
func NewMockJobBookFetcher(ctrl *gomock.Controller) *MockJobBookFetcher {
	mock := &MockJobBookFetcher{ctrl: ctrl}
	mock.recorder = &MockJobBookFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJobBookFetcher) EXPECT() *MockJobBookFetcherMockRecorder {
	return m.recorder
}

// FetchAll mocks base method.
func (m *MockJobBookFetcher) FetchAll() ([]orderdb.JobBook, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAll")
	ret0, _ := ret[0].([]orderdb.JobBook)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAll indicates an expected call of FetchAll.
func (mr *MockJobBookFetcherMockRecorder) FetchAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAll", reflect.TypeOf((*MockJobBookFetcher)(nil).FetchAll))
}
