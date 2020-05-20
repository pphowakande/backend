// Automatically generated by MockGen. DO NOT EDIT!
// Source: backend/service/athAdmin (interfaces: AthAdminUserService)

package mock

import (
	io "backend/pkg/io"
	context "context"
	gomock "github.com/golang/mock/gomock"
)

// Mock of AthAdminUserService interface
type MockAthAdminUserService struct {
	ctrl     *gomock.Controller
	recorder *_MockAthAdminUserServiceRecorder
}

// Recorder for MockAthAdminUserService (not exported)
type _MockAthAdminUserServiceRecorder struct {
	mock *MockAthAdminUserService
}

func NewMockAthAdminUserService(ctrl *gomock.Controller) *MockAthAdminUserService {
	mock := &MockAthAdminUserService{ctrl: ctrl}
	mock.recorder = &_MockAthAdminUserServiceRecorder{mock}
	return mock
}

func (_m *MockAthAdminUserService) EXPECT() *_MockAthAdminUserServiceRecorder {
	return _m.recorder
}

func (_m *MockAthAdminUserService) LoginAdmin(_param0 context.Context, _param1 io.LoginRequest) io.Response {
	ret := _m.ctrl.Call(_m, "LoginAdmin", _param0, _param1)
	ret0, _ := ret[0].(io.Response)
	return ret0
}

func (_mr *_MockAthAdminUserServiceRecorder) LoginAdmin(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LoginAdmin", arg0, arg1)
}

func (_m *MockAthAdminUserService) SignUpAdmin(_param0 context.Context, _param1 io.AthAdminUser) io.Response {
	ret := _m.ctrl.Call(_m, "SignUpAdmin", _param0, _param1)
	ret0, _ := ret[0].(io.Response)
	return ret0
}

func (_mr *_MockAthAdminUserServiceRecorder) SignUpAdmin(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SignUpAdmin", arg0, arg1)
}
