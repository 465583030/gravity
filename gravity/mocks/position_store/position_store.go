// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/moiot/gravity/gravity/position_store (interfaces: MySQLPositionStore,MongoPositionStore)

// Package mock_position_store is a generated GoMock package.
package mock_position_store

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	config "github.com/moiot/gravity/gravity/config"
	utils "github.com/moiot/gravity/pkg/utils"
)

// MockMySQLPositionStore is a mock of MySQLPositionStore interface
type MockMySQLPositionStore struct {
	ctrl     *gomock.Controller
	recorder *MockMySQLPositionStoreMockRecorder
}

// MockMySQLPositionStoreMockRecorder is the mock recorder for MockMySQLPositionStore
type MockMySQLPositionStoreMockRecorder struct {
	mock *MockMySQLPositionStore
}

// NewMockMySQLPositionStore creates a new mock instance
func NewMockMySQLPositionStore(ctrl *gomock.Controller) *MockMySQLPositionStore {
	mock := &MockMySQLPositionStore{ctrl: ctrl}
	mock.recorder = &MockMySQLPositionStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMySQLPositionStore) EXPECT() *MockMySQLPositionStoreMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockMySQLPositionStore) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockMySQLPositionStoreMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMySQLPositionStore)(nil).Close))
}

// Get mocks base method
func (m *MockMySQLPositionStore) Get() utils.MySQLBinlogPosition {
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(utils.MySQLBinlogPosition)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockMySQLPositionStoreMockRecorder) Get() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMySQLPositionStore)(nil).Get))
}

// Put mocks base method
func (m *MockMySQLPositionStore) Put(arg0 utils.MySQLBinlogPosition, arg1 chan struct{}) {
	m.ctrl.Call(m, "Put", arg0, arg1)
}

// Put indicates an expected call of Put
func (mr *MockMySQLPositionStoreMockRecorder) Put(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockMySQLPositionStore)(nil).Put), arg0, arg1)
}

// Start mocks base method
func (m *MockMySQLPositionStore) Start() error {
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockMySQLPositionStoreMockRecorder) Start() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockMySQLPositionStore)(nil).Start))
}

// MockMongoPositionStore is a mock of MongoPositionStore interface
type MockMongoPositionStore struct {
	ctrl     *gomock.Controller
	recorder *MockMongoPositionStoreMockRecorder
}

// MockMongoPositionStoreMockRecorder is the mock recorder for MockMongoPositionStore
type MockMongoPositionStoreMockRecorder struct {
	mock *MockMongoPositionStore
}

// NewMockMongoPositionStore creates a new mock instance
func NewMockMongoPositionStore(ctrl *gomock.Controller) *MockMongoPositionStore {
	mock := &MockMongoPositionStore{ctrl: ctrl}
	mock.recorder = &MockMongoPositionStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMongoPositionStore) EXPECT() *MockMongoPositionStoreMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockMongoPositionStore) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockMongoPositionStoreMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMongoPositionStore)(nil).Close))
}

// Get mocks base method
func (m *MockMongoPositionStore) Get(arg0 string) config.MongoPosition {
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(config.MongoPosition)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockMongoPositionStoreMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMongoPositionStore)(nil).Get), arg0)
}

// Put mocks base method
func (m *MockMongoPositionStore) Put(arg0 string, arg1 config.MongoPosition) {
	m.ctrl.Call(m, "Put", arg0, arg1)
}

// Put indicates an expected call of Put
func (mr *MockMongoPositionStoreMockRecorder) Put(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockMongoPositionStore)(nil).Put), arg0, arg1)
}

// Start mocks base method
func (m *MockMongoPositionStore) Start() error {
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockMongoPositionStoreMockRecorder) Start() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockMongoPositionStore)(nil).Start))
}
