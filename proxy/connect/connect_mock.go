// Code generated by MockGen. DO NOT EDIT.
// Source: proxy/connect/connect.go

// Package connect is a generated GoMock package.
package connect

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockConnect is a mock of Connect interface.
type MockConnect struct {
	ctrl     *gomock.Controller
	recorder *MockConnectMockRecorder
}

// MockConnectMockRecorder is the mock recorder for MockConnect.
type MockConnectMockRecorder struct {
	mock *MockConnect
}

// NewMockConnect creates a new mock instance.
func NewMockConnect(ctrl *gomock.Controller) *MockConnect {
	mock := &MockConnect{ctrl: ctrl}
	mock.recorder = &MockConnectMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnect) EXPECT() *MockConnectMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockConnect) Connect(arg0 context.Context, arg1 string, arg2 DestinationHash) (Destination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", arg0, arg1, arg2)
	ret0, _ := ret[0].(Destination)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Connect indicates an expected call of Connect.
func (mr *MockConnectMockRecorder) Connect(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockConnect)(nil).Connect), arg0, arg1, arg2)
}

// MockDestination is a mock of Destination interface.
type MockDestination struct {
	ctrl     *gomock.Controller
	recorder *MockDestinationMockRecorder
}

// MockDestinationMockRecorder is the mock recorder for MockDestination.
type MockDestinationMockRecorder struct {
	mock *MockDestination
}

// NewMockDestination creates a new mock instance.
func NewMockDestination(ctrl *gomock.Controller) *MockDestination {
	mock := &MockDestination{ctrl: ctrl}
	mock.recorder = &MockDestinationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDestination) EXPECT() *MockDestinationMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockDestination) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockDestinationMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDestination)(nil).Close))
}

// SendChannel mocks base method.
func (m *MockDestination) SendChannel() chan<- SendRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendChannel")
	ret0, _ := ret[0].(chan<- SendRequest)
	return ret0
}

// SendChannel indicates an expected call of SendChannel.
func (mr *MockDestinationMockRecorder) SendChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendChannel", reflect.TypeOf((*MockDestination)(nil).SendChannel))
}

// MockDestinationHash is a mock of DestinationHash interface.
type MockDestinationHash struct {
	ctrl     *gomock.Controller
	recorder *MockDestinationHashMockRecorder
}

// MockDestinationHashMockRecorder is the mock recorder for MockDestinationHash.
type MockDestinationHashMockRecorder struct {
	mock *MockDestinationHash
}

// NewMockDestinationHash creates a new mock instance.
func NewMockDestinationHash(ctrl *gomock.Controller) *MockDestinationHash {
	mock := &MockDestinationHash{ctrl: ctrl}
	mock.recorder = &MockDestinationHashMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDestinationHash) EXPECT() *MockDestinationHashMockRecorder {
	return m.recorder
}

// ConnectionClosed mocks base method.
func (m *MockDestinationHash) ConnectionClosed() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ConnectionClosed")
}

// ConnectionClosed indicates an expected call of ConnectionClosed.
func (mr *MockDestinationHashMockRecorder) ConnectionClosed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectionClosed", reflect.TypeOf((*MockDestinationHash)(nil).ConnectionClosed))
}

// RemoveDestination mocks base method.
func (m *MockDestinationHash) RemoveDestination(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveDestination", arg0)
}

// RemoveDestination indicates an expected call of RemoveDestination.
func (mr *MockDestinationHashMockRecorder) RemoveDestination(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDestination", reflect.TypeOf((*MockDestinationHash)(nil).RemoveDestination), arg0)
}
