// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ledgerwatch/erigon/polygon/sync (interfaces: DB)

// Package sync is a generated GoMock package.
package sync

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/ledgerwatch/erigon/core/types"
)

// MockDB is a mock of DB interface.
type MockDB struct {
	ctrl     *gomock.Controller
	recorder *MockDBMockRecorder
}

// MockDBMockRecorder is the mock recorder for MockDB.
type MockDBMockRecorder struct {
	mock *MockDB
}

// NewMockDB creates a new mock instance.
func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &MockDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDB) EXPECT() *MockDBMockRecorder {
	return m.recorder
}

// WriteHeaders mocks base method.
func (m *MockDB) WriteHeaders(arg0 []*types.Header) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteHeaders", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteHeaders indicates an expected call of WriteHeaders.
func (mr *MockDBMockRecorder) WriteHeaders(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteHeaders", reflect.TypeOf((*MockDB)(nil).WriteHeaders), arg0)
}
