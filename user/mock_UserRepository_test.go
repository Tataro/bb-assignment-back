// Code generated by mockery v2.44.1. DO NOT EDIT.

package user

import (
	mock "github.com/stretchr/testify/mock"
	domain "github.com/tat-101/bb-assignment-back/domain"
)

// MockUserRepository is an autogenerated mock type for the UserRepository type
type MockUserRepository struct {
	mock.Mock
}

type MockUserRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserRepository) EXPECT() *MockUserRepository_Expecter {
	return &MockUserRepository_Expecter{mock: &_m.Mock}
}

// CreateUser provides a mock function with given fields: user
func (_m *MockUserRepository) CreateUser(user *domain.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserRepository_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type MockUserRepository_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - user *domain.User
func (_e *MockUserRepository_Expecter) CreateUser(user interface{}) *MockUserRepository_CreateUser_Call {
	return &MockUserRepository_CreateUser_Call{Call: _e.mock.On("CreateUser", user)}
}

func (_c *MockUserRepository_CreateUser_Call) Run(run func(user *domain.User)) *MockUserRepository_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*domain.User))
	})
	return _c
}

func (_c *MockUserRepository_CreateUser_Call) Return(_a0 error) *MockUserRepository_CreateUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_CreateUser_Call) RunAndReturn(run func(*domain.User) error) *MockUserRepository_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteUserByID provides a mock function with given fields: id
func (_m *MockUserRepository) DeleteUserByID(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUserByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserRepository_DeleteUserByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteUserByID'
type MockUserRepository_DeleteUserByID_Call struct {
	*mock.Call
}

// DeleteUserByID is a helper method to define mock.On call
//   - id string
func (_e *MockUserRepository_Expecter) DeleteUserByID(id interface{}) *MockUserRepository_DeleteUserByID_Call {
	return &MockUserRepository_DeleteUserByID_Call{Call: _e.mock.On("DeleteUserByID", id)}
}

func (_c *MockUserRepository_DeleteUserByID_Call) Run(run func(id string)) *MockUserRepository_DeleteUserByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockUserRepository_DeleteUserByID_Call) Return(_a0 error) *MockUserRepository_DeleteUserByID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_DeleteUserByID_Call) RunAndReturn(run func(string) error) *MockUserRepository_DeleteUserByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllUsers provides a mock function with given fields:
func (_m *MockUserRepository) GetAllUsers() ([]domain.User, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllUsers")
	}

	var r0 []domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]domain.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_GetAllUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllUsers'
type MockUserRepository_GetAllUsers_Call struct {
	*mock.Call
}

// GetAllUsers is a helper method to define mock.On call
func (_e *MockUserRepository_Expecter) GetAllUsers() *MockUserRepository_GetAllUsers_Call {
	return &MockUserRepository_GetAllUsers_Call{Call: _e.mock.On("GetAllUsers")}
}

func (_c *MockUserRepository_GetAllUsers_Call) Run(run func()) *MockUserRepository_GetAllUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockUserRepository_GetAllUsers_Call) Return(_a0 []domain.User, _a1 error) *MockUserRepository_GetAllUsers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_GetAllUsers_Call) RunAndReturn(run func() ([]domain.User, error)) *MockUserRepository_GetAllUsers_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByEmail provides a mock function with given fields: email
func (_m *MockUserRepository) GetUserByEmail(email string) (*domain.User, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_GetUserByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByEmail'
type MockUserRepository_GetUserByEmail_Call struct {
	*mock.Call
}

// GetUserByEmail is a helper method to define mock.On call
//   - email string
func (_e *MockUserRepository_Expecter) GetUserByEmail(email interface{}) *MockUserRepository_GetUserByEmail_Call {
	return &MockUserRepository_GetUserByEmail_Call{Call: _e.mock.On("GetUserByEmail", email)}
}

func (_c *MockUserRepository_GetUserByEmail_Call) Run(run func(email string)) *MockUserRepository_GetUserByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockUserRepository_GetUserByEmail_Call) Return(_a0 *domain.User, _a1 error) *MockUserRepository_GetUserByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_GetUserByEmail_Call) RunAndReturn(run func(string) (*domain.User, error)) *MockUserRepository_GetUserByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByID provides a mock function with given fields: id
func (_m *MockUserRepository) GetUserByID(id uint) (*domain.User, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByID")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*domain.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *domain.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_GetUserByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByID'
type MockUserRepository_GetUserByID_Call struct {
	*mock.Call
}

// GetUserByID is a helper method to define mock.On call
//   - id uint
func (_e *MockUserRepository_Expecter) GetUserByID(id interface{}) *MockUserRepository_GetUserByID_Call {
	return &MockUserRepository_GetUserByID_Call{Call: _e.mock.On("GetUserByID", id)}
}

func (_c *MockUserRepository_GetUserByID_Call) Run(run func(id uint)) *MockUserRepository_GetUserByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *MockUserRepository_GetUserByID_Call) Return(_a0 *domain.User, _a1 error) *MockUserRepository_GetUserByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_GetUserByID_Call) RunAndReturn(run func(uint) (*domain.User, error)) *MockUserRepository_GetUserByID_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUserByID provides a mock function with given fields: id, updatedUser
func (_m *MockUserRepository) UpdateUserByID(id string, updatedUser domain.User) (*domain.User, error) {
	ret := _m.Called(id, updatedUser)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUserByID")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string, domain.User) (*domain.User, error)); ok {
		return rf(id, updatedUser)
	}
	if rf, ok := ret.Get(0).(func(string, domain.User) *domain.User); ok {
		r0 = rf(id, updatedUser)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string, domain.User) error); ok {
		r1 = rf(id, updatedUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_UpdateUserByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUserByID'
type MockUserRepository_UpdateUserByID_Call struct {
	*mock.Call
}

// UpdateUserByID is a helper method to define mock.On call
//   - id string
//   - updatedUser domain.User
func (_e *MockUserRepository_Expecter) UpdateUserByID(id interface{}, updatedUser interface{}) *MockUserRepository_UpdateUserByID_Call {
	return &MockUserRepository_UpdateUserByID_Call{Call: _e.mock.On("UpdateUserByID", id, updatedUser)}
}

func (_c *MockUserRepository_UpdateUserByID_Call) Run(run func(id string, updatedUser domain.User)) *MockUserRepository_UpdateUserByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(domain.User))
	})
	return _c
}

func (_c *MockUserRepository_UpdateUserByID_Call) Return(_a0 *domain.User, _a1 error) *MockUserRepository_UpdateUserByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_UpdateUserByID_Call) RunAndReturn(run func(string, domain.User) (*domain.User, error)) *MockUserRepository_UpdateUserByID_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUserRepository creates a new instance of MockUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserRepository {
	mock := &MockUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
