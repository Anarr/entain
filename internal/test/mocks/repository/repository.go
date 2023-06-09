// Code generated by mockery v2.15.0. DO NOT EDIT.

package mockedrepository

import (
	model "github.com/Anarr/entain/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CancelRequest provides a mock function with given fields: m
func (_m *Repository) CancelRequest(m model.Request) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Request) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetLatestRequests provides a mock function with given fields: limit
func (_m *Repository) GetLatestRequests(limit int) ([]*model.Request, error) {
	ret := _m.Called(limit)

	var r0 []*model.Request
	if rf, ok := ret.Get(0).(func(int) []*model.Request); ok {
		r0 = rf(limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Request)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitDefaultUser provides a mock function with given fields:
func (_m *Repository) InitDefaultUser() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveRequest provides a mock function with given fields: m
func (_m *Repository) SaveRequest(m model.Request) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Request) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUserBalance provides a mock function with given fields: m
func (_m *Repository) UpdateUserBalance(m model.Request) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Request) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
