// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/mulukenhailu/Binary/domain"
	mock "github.com/stretchr/testify/mock"
)

// RoleUsecase is an autogenerated mock type for the RoleUsecase type
type RoleUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, role
func (_m *RoleUsecase) Create(c context.Context, role *domain.CreateRoleDto) error {
	ret := _m.Called(c, role)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.CreateRoleDto) error); ok {
		r0 = rf(c, role)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: c, roleId
func (_m *RoleUsecase) Delete(c context.Context, roleId int32) error {
	ret := _m.Called(c, roleId)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) error); ok {
		r0 = rf(c, roleId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchByName provides a mock function with given fields: c, roleName
func (_m *RoleUsecase) FetchByName(c context.Context, roleName string) (domain.Role, error) {
	ret := _m.Called(c, roleName)

	if len(ret) == 0 {
		panic("no return value specified for FetchByName")
	}

	var r0 domain.Role
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.Role, error)); ok {
		return rf(c, roleName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Role); ok {
		r0 = rf(c, roleName)
	} else {
		r0 = ret.Get(0).(domain.Role)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, roleName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchRoles provides a mock function with given fields: c
func (_m *RoleUsecase) FetchRoles(c context.Context) ([]domain.Role, error) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for FetchRoles")
	}

	var r0 []domain.Role
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.Role, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Role); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Role)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: c, updateRoleDto
func (_m *RoleUsecase) Update(c context.Context, updateRoleDto *domain.UpdateRoleDto) error {
	ret := _m.Called(c, updateRoleDto)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.UpdateRoleDto) error); ok {
		r0 = rf(c, updateRoleDto)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRoleUsecase creates a new instance of RoleUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRoleUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *RoleUsecase {
	mock := &RoleUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
