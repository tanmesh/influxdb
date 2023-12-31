// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/influxdata/influxdb/v2/v1/authorization (interfaces: UserFinder)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	influxdb "github.com/influxdata/influxdb/v2"
	"github.com/influxdata/influxdb/v2/kit/platform"
)

// MockUserFinder is a mock of UserFinder interface
type MockUserFinder struct {
	ctrl     *gomock.Controller
	recorder *MockUserFinderMockRecorder
}

// MockUserFinderMockRecorder is the mock recorder for MockUserFinder
type MockUserFinderMockRecorder struct {
	mock *MockUserFinder
}

// NewMockUserFinder creates a new mock instance
func NewMockUserFinder(ctrl *gomock.Controller) *MockUserFinder {
	mock := &MockUserFinder{ctrl: ctrl}
	mock.recorder = &MockUserFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserFinder) EXPECT() *MockUserFinderMockRecorder {
	return m.recorder
}

// FindUserByID mocks base method
func (m *MockUserFinder) FindUserByID(arg0 context.Context, arg1 platform.ID) (*influxdb.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByID", arg0, arg1)
	ret0, _ := ret[0].(*influxdb.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByID indicates an expected call of FindUserByID
func (mr *MockUserFinderMockRecorder) FindUserByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByID", reflect.TypeOf((*MockUserFinder)(nil).FindUserByID), arg0, arg1)
}
