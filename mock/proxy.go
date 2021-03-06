// Automatically generated by MockGen. DO NOT EDIT!
// Source: partition.go

package mock

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of kafkaProxy interface
type MockkafkaProxy struct {
	ctrl     *gomock.Controller
	recorder *_MockkafkaProxyRecorder
}

// Recorder for MockkafkaProxy (not exported)
type _MockkafkaProxyRecorder struct {
	mock *MockkafkaProxy
}

func NewMockkafkaProxy(ctrl *gomock.Controller) *MockkafkaProxy {
	mock := &MockkafkaProxy{ctrl: ctrl}
	mock.recorder = &_MockkafkaProxyRecorder{mock}
	return mock
}

func (_m *MockkafkaProxy) EXPECT() *_MockkafkaProxyRecorder {
	return _m.recorder
}

func (_m *MockkafkaProxy) Add(_param0 string, _param1 int64) error {
	ret := _m.ctrl.Call(_m, "Add", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockkafkaProxyRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Add", arg0, arg1)
}

func (_m *MockkafkaProxy) Remove(_param0 string) error {
	ret := _m.ctrl.Call(_m, "Remove", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockkafkaProxyRecorder) Remove(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Remove", arg0)
}

func (_m *MockkafkaProxy) AddGroup() {
	_m.ctrl.Call(_m, "AddGroup")
}

func (_mr *_MockkafkaProxyRecorder) AddGroup() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddGroup")
}

func (_m *MockkafkaProxy) Stop() {
	_m.ctrl.Call(_m, "Stop")
}

func (_mr *_MockkafkaProxyRecorder) Stop() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stop")
}
