// Code generated by mockery v2.49.1. DO NOT EDIT.

package mocks

import (
	ent "github.com/choonhong/hotel-data-merge/ent"

	mock "github.com/stretchr/testify/mock"
)

// Cache is an autogenerated mock type for the Cache type
type Cache struct {
	mock.Mock
}

type Cache_Expecter struct {
	mock *mock.Mock
}

func (_m *Cache) EXPECT() *Cache_Expecter {
	return &Cache_Expecter{mock: &_m.Mock}
}

// Clear provides a mock function with given fields:
func (_m *Cache) Clear() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Clear")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Cache_Clear_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Clear'
type Cache_Clear_Call struct {
	*mock.Call
}

// Clear is a helper method to define mock.On call
func (_e *Cache_Expecter) Clear() *Cache_Clear_Call {
	return &Cache_Clear_Call{Call: _e.mock.On("Clear")}
}

func (_c *Cache_Clear_Call) Run(run func()) *Cache_Clear_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Cache_Clear_Call) Return(_a0 error) *Cache_Clear_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Cache_Clear_Call) RunAndReturn(run func() error) *Cache_Clear_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: v
func (_m *Cache) Get(v interface{}) ([]*ent.Hotel, error) {
	ret := _m.Called(v)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []*ent.Hotel
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}) ([]*ent.Hotel, error)); ok {
		return rf(v)
	}
	if rf, ok := ret.Get(0).(func(interface{}) []*ent.Hotel); ok {
		r0 = rf(v)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ent.Hotel)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(v)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Cache_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type Cache_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - v interface{}
func (_e *Cache_Expecter) Get(v interface{}) *Cache_Get_Call {
	return &Cache_Get_Call{Call: _e.mock.On("Get", v)}
}

func (_c *Cache_Get_Call) Run(run func(v interface{})) *Cache_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *Cache_Get_Call) Return(_a0 []*ent.Hotel, _a1 error) *Cache_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Cache_Get_Call) RunAndReturn(run func(interface{}) ([]*ent.Hotel, error)) *Cache_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: v, hotels
func (_m *Cache) Set(v interface{}, hotels []*ent.Hotel) error {
	ret := _m.Called(v, hotels)

	if len(ret) == 0 {
		panic("no return value specified for Set")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, []*ent.Hotel) error); ok {
		r0 = rf(v, hotels)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Cache_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type Cache_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - v interface{}
//   - hotels []*ent.Hotel
func (_e *Cache_Expecter) Set(v interface{}, hotels interface{}) *Cache_Set_Call {
	return &Cache_Set_Call{Call: _e.mock.On("Set", v, hotels)}
}

func (_c *Cache_Set_Call) Run(run func(v interface{}, hotels []*ent.Hotel)) *Cache_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}), args[1].([]*ent.Hotel))
	})
	return _c
}

func (_c *Cache_Set_Call) Return(_a0 error) *Cache_Set_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Cache_Set_Call) RunAndReturn(run func(interface{}, []*ent.Hotel) error) *Cache_Set_Call {
	_c.Call.Return(run)
	return _c
}

// NewCache creates a new instance of Cache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCache(t interface {
	mock.TestingT
	Cleanup(func())
}) *Cache {
	mock := &Cache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}