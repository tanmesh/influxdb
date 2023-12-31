// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/influxdata/influxdb/v2/kv (interfaces: Bucket)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	kv "github.com/influxdata/influxdb/v2/kv"
)

// MockBucket is a mock of Bucket interface
type MockBucket struct {
	ctrl     *gomock.Controller
	recorder *MockBucketMockRecorder
}

// MockBucketMockRecorder is the mock recorder for MockBucket
type MockBucketMockRecorder struct {
	mock *MockBucket
}

// NewMockBucket creates a new mock instance
func NewMockBucket(ctrl *gomock.Controller) *MockBucket {
	mock := &MockBucket{ctrl: ctrl}
	mock.recorder = &MockBucketMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBucket) EXPECT() *MockBucketMockRecorder {
	return m.recorder
}

// Cursor mocks base method
func (m *MockBucket) Cursor(arg0 ...kv.CursorHint) (kv.Cursor, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Cursor", varargs...)
	ret0, _ := ret[0].(kv.Cursor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cursor indicates an expected call of Cursor
func (mr *MockBucketMockRecorder) Cursor(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cursor", reflect.TypeOf((*MockBucket)(nil).Cursor), arg0...)
}

// Delete mocks base method
func (m *MockBucket) Delete(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockBucketMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBucket)(nil).Delete), arg0)
}

// ForwardCursor mocks base method
func (m *MockBucket) ForwardCursor(arg0 []byte, arg1 ...kv.CursorOption) (kv.ForwardCursor, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ForwardCursor", varargs...)
	ret0, _ := ret[0].(kv.ForwardCursor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ForwardCursor indicates an expected call of ForwardCursor
func (mr *MockBucketMockRecorder) ForwardCursor(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForwardCursor", reflect.TypeOf((*MockBucket)(nil).ForwardCursor), varargs...)
}

// Get mocks base method
func (m *MockBucket) Get(arg0 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockBucketMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBucket)(nil).Get), arg0)
}

// GetBatch mocks base method
func (m *MockBucket) GetBatch(arg0 ...[]byte) ([][]byte, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBatch", varargs...)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBatch indicates an expected call of GetBatch
func (mr *MockBucketMockRecorder) GetBatch(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBatch", reflect.TypeOf((*MockBucket)(nil).GetBatch), arg0...)
}

// Put mocks base method
func (m *MockBucket) Put(arg0, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put
func (mr *MockBucketMockRecorder) Put(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockBucket)(nil).Put), arg0, arg1)
}
