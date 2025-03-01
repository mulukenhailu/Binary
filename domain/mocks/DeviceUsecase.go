// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/mulukenhailu/Binary/domain"
	mock "github.com/stretchr/testify/mock"
)

// DeviceUsecase is an autogenerated mock type for the DeviceUsecase type
type DeviceUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, createDeviceDto
func (_m *DeviceUsecase) Create(c context.Context, createDeviceDto *domain.CreateDeviceDto) error {
	ret := _m.Called(c, createDeviceDto)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.CreateDeviceDto) error); ok {
		r0 = rf(c, createDeviceDto)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: c, deviceId
func (_m *DeviceUsecase) Delete(c context.Context, deviceId int32) error {
	ret := _m.Called(c, deviceId)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) error); ok {
		r0 = rf(c, deviceId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchByCampus provides a mock function with given fields: c, campusName
func (_m *DeviceUsecase) FetchByCampus(c context.Context, campusName string) ([]domain.Device, error) {
	ret := _m.Called(c, campusName)

	if len(ret) == 0 {
		panic("no return value specified for FetchByCampus")
	}

	var r0 []domain.Device
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]domain.Device, error)); ok {
		return rf(c, campusName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []domain.Device); ok {
		r0 = rf(c, campusName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Device)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, campusName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchDevices provides a mock function with given fields: c
func (_m *DeviceUsecase) FetchDevices(c context.Context) ([]domain.Device, error) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for FetchDevices")
	}

	var r0 []domain.Device
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.Device, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Device); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Device)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: c, updateDeviceDto
func (_m *DeviceUsecase) Update(c context.Context, updateDeviceDto *domain.UpdateDeviceDto) error {
	ret := _m.Called(c, updateDeviceDto)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.UpdateDeviceDto) error); ok {
		r0 = rf(c, updateDeviceDto)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDeviceUsecase creates a new instance of DeviceUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDeviceUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *DeviceUsecase {
	mock := &DeviceUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
