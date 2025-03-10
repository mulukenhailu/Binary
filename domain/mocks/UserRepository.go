// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/mulukenhailu/Binary/domain"
	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, createUserDto
func (_m *UserRepository) Create(c context.Context, createUserDto *domain.CreateUserDto) error {
	ret := _m.Called(c, createUserDto)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.CreateUserDto) error); ok {
		r0 = rf(c, createUserDto)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: c, userId
func (_m *UserRepository) Delete(c context.Context, userId int32) error {
	ret := _m.Called(c, userId)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) error); ok {
		r0 = rf(c, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchByRoleId provides a mock function with given fields: c, roleId
func (_m *UserRepository) FetchByRoleId(c context.Context, roleId int32) ([]domain.User, error) {
	ret := _m.Called(c, roleId)

	if len(ret) == 0 {
		panic("no return value specified for FetchByRoleId")
	}

	var r0 []domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) ([]domain.User, error)); ok {
		return rf(c, roleId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32) []domain.User); ok {
		r0 = rf(c, roleId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(c, roleId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchByUserName provides a mock function with given fields: c, userName
func (_m *UserRepository) FetchByUserName(c context.Context, userName string) (domain.User, error) {
	ret := _m.Called(c, userName)

	if len(ret) == 0 {
		panic("no return value specified for FetchByUserName")
	}

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.User, error)); ok {
		return rf(c, userName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.User); ok {
		r0 = rf(c, userName)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, userName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchUsers provides a mock function with given fields: c
func (_m *UserRepository) FetchUsers(c context.Context) ([]domain.User, error) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for FetchUsers")
	}

	var r0 []domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.User, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.User); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: c, updateUserDto
func (_m *UserRepository) Update(c context.Context, updateUserDto *domain.UpdateUserDto) error {
	ret := _m.Called(c, updateUserDto)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.UpdateUserDto) error); ok {
		r0 = rf(c, updateUserDto)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
