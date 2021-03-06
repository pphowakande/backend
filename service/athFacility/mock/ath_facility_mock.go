// Automatically generated by MockGen. DO NOT EDIT!
// Source: backend/service/athFacility (interfaces: AthFacilityService)

package mock

import (
	io "backend/pkg/io"
	context "context"
	gomock "github.com/golang/mock/gomock"
)

// Mock of AthFacilityService interface
type MockAthFacilityService struct {
	ctrl     *gomock.Controller
	recorder *_MockAthFacilityServiceRecorder
}

// Recorder for MockAthFacilityService (not exported)
type _MockAthFacilityServiceRecorder struct {
	mock *MockAthFacilityService
}

func NewMockAthFacilityService(ctrl *gomock.Controller) *MockAthFacilityService {
	mock := &MockAthFacilityService{ctrl: ctrl}
	mock.recorder = &_MockAthFacilityServiceRecorder{mock}
	return mock
}

func (_m *MockAthFacilityService) EXPECT() *_MockAthFacilityServiceRecorder {
	return _m.recorder
}

func (_m *MockAthFacilityService) AddFacilityCustomRates(_param0 context.Context, _param1 io.AthFacilityCustomRates) io.Response {
	ret := _m.ctrl.Call(_m, "AddFacilityCustomRates", _param0, _param1)
	ret0, _ := ret[0].(io.Response)
	return ret0
}

func (_mr *_MockAthFacilityServiceRecorder) AddFacilityCustomRates(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddFacilityCustomRates", arg0, arg1)
}

func (_m *MockAthFacilityService) BookFacility(_param0 context.Context, _param1 io.AthFacilityBookings) io.Response {
	ret := _m.ctrl.Call(_m, "BookFacility", _param0, _param1)
	ret0, _ := ret[0].(io.Response)
	return ret0
}

func (_mr *_MockAthFacilityServiceRecorder) BookFacility(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BookFacility", arg0, arg1)
}

func (_m *MockAthFacilityService) BookFacilitySlots(_param0 context.Context, _param1 io.AthFacilityBookingSlots) io.Response {
	ret := _m.ctrl.Call(_m, "BookFacilitySlots", _param0, _param1)
	ret0, _ := ret[0].(io.Response)
	return ret0
}

func (_mr *_MockAthFacilityServiceRecorder) BookFacilitySlots(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BookFacilitySlots", arg0, arg1)
}

func (_m *MockAthFacilityService) CreateFacility(_param0 context.Context, _param1 io.AthFacilities) io.Response {
	ret := _m.ctrl.Call(_m, "CreateFacility", _param0, _param1)
	ret0, _ := ret[0].(io.Response)
	return ret0
}

func (_mr *_MockAthFacilityServiceRecorder) CreateFacility(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateFacility", arg0, arg1)
}

func (_m *MockAthFacilityService) CreateFacilitySlots(_param0 context.Context, _param1 io.AthFacilitySlots) io.Response {
	ret := _m.ctrl.Call(_m, "CreateFacilitySlots", _param0, _param1)
	ret0, _ := ret[0].(io.Response)
	return ret0
}

func (_mr *_MockAthFacilityServiceRecorder) CreateFacilitySlots(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateFacilitySlots", arg0, arg1)
}

func (_m *MockAthFacilityService) EditFacility(_param0 context.Context, _param1 io.AthFacilities) io.Response {
	ret := _m.ctrl.Call(_m, "EditFacility", _param0, _param1)
	ret0, _ := ret[0].(io.Response)
	return ret0
}

func (_mr *_MockAthFacilityServiceRecorder) EditFacility(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EditFacility", arg0, arg1)
}

func (_m *MockAthFacilityService) EditFacilitySlots(_param0 context.Context, _param1 io.AthFacilitySlots) io.Response {
	ret := _m.ctrl.Call(_m, "EditFacilitySlots", _param0, _param1)
	ret0, _ := ret[0].(io.Response)
	return ret0
}

func (_mr *_MockAthFacilityServiceRecorder) EditFacilitySlots(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EditFacilitySlots", arg0, arg1)
}
