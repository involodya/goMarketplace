// Code generated by mockery v2.40.3. DO NOT EDIT.

package entity

import (
	entity "fullstack/backend/internal/entity"

	mock "github.com/stretchr/testify/mock"
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

// Create provides a mock function with given fields: _a0
func (_m *MockUserRepository) Create(_a0 *entity.User) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockUserRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 *entity.User
func (_e *MockUserRepository_Expecter) Create(_a0 interface{}) *MockUserRepository_Create_Call {
	return &MockUserRepository_Create_Call{Call: _e.mock.On("Create", _a0)}
}

func (_c *MockUserRepository_Create_Call) Run(run func(_a0 *entity.User)) *MockUserRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.User))
	})
	return _c
}

func (_c *MockUserRepository_Create_Call) Return(_a0 error) *MockUserRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_Create_Call) RunAndReturn(run func(*entity.User) error) *MockUserRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: id
func (_m *MockUserRepository) Delete(id uint) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockUserRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - id uint
func (_e *MockUserRepository_Expecter) Delete(id interface{}) *MockUserRepository_Delete_Call {
	return &MockUserRepository_Delete_Call{Call: _e.mock.On("Delete", id)}
}

func (_c *MockUserRepository_Delete_Call) Run(run func(id uint)) *MockUserRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *MockUserRepository_Delete_Call) Return(_a0 error) *MockUserRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_Delete_Call) RunAndReturn(run func(uint) error) *MockUserRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: id
func (_m *MockUserRepository) Get(id uint) (*entity.User, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*entity.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *entity.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockUserRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - id uint
func (_e *MockUserRepository_Expecter) Get(id interface{}) *MockUserRepository_Get_Call {
	return &MockUserRepository_Get_Call{Call: _e.mock.On("Get", id)}
}

func (_c *MockUserRepository_Get_Call) Run(run func(id uint)) *MockUserRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *MockUserRepository_Get_Call) Return(_a0 *entity.User, _a1 error) *MockUserRepository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_Get_Call) RunAndReturn(run func(uint) (*entity.User, error)) *MockUserRepository_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields:
func (_m *MockUserRepository) GetAll() (*[]entity.User, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 *[]entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func() (*[]entity.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *[]entity.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type MockUserRepository_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
func (_e *MockUserRepository_Expecter) GetAll() *MockUserRepository_GetAll_Call {
	return &MockUserRepository_GetAll_Call{Call: _e.mock.On("GetAll")}
}

func (_c *MockUserRepository_GetAll_Call) Run(run func()) *MockUserRepository_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockUserRepository_GetAll_Call) Return(_a0 *[]entity.User, _a1 error) *MockUserRepository_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_GetAll_Call) RunAndReturn(run func() (*[]entity.User, error)) *MockUserRepository_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// GetByEmail provides a mock function with given fields: email
func (_m *MockUserRepository) GetByEmail(email string) (*entity.User, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for GetByEmail")
	}

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entity.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *entity.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_GetByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByEmail'
type MockUserRepository_GetByEmail_Call struct {
	*mock.Call
}

// GetByEmail is a helper method to define mock.On call
//   - email string
func (_e *MockUserRepository_Expecter) GetByEmail(email interface{}) *MockUserRepository_GetByEmail_Call {
	return &MockUserRepository_GetByEmail_Call{Call: _e.mock.On("GetByEmail", email)}
}

func (_c *MockUserRepository_GetByEmail_Call) Run(run func(email string)) *MockUserRepository_GetByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockUserRepository_GetByEmail_Call) Return(_a0 *entity.User, _a1 error) *MockUserRepository_GetByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_GetByEmail_Call) RunAndReturn(run func(string) (*entity.User, error)) *MockUserRepository_GetByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: _a0
func (_m *MockUserRepository) Update(_a0 *entity.User) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockUserRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - _a0 *entity.User
func (_e *MockUserRepository_Expecter) Update(_a0 interface{}) *MockUserRepository_Update_Call {
	return &MockUserRepository_Update_Call{Call: _e.mock.On("Update", _a0)}
}

func (_c *MockUserRepository_Update_Call) Run(run func(_a0 *entity.User)) *MockUserRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.User))
	})
	return _c
}

func (_c *MockUserRepository_Update_Call) Return(_a0 error) *MockUserRepository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_Update_Call) RunAndReturn(run func(*entity.User) error) *MockUserRepository_Update_Call {
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
