// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	models "kredit-plus/models"

	mock "github.com/stretchr/testify/mock"
)

// LimitRepository is an autogenerated mock type for the LimitRepository type
type LimitRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: limit
func (_m *LimitRepository) Create(limit models.LimitKonsumen) (models.LimitKonsumen, error) {
	ret := _m.Called(limit)

	var r0 models.LimitKonsumen
	if rf, ok := ret.Get(0).(func(models.LimitKonsumen) models.LimitKonsumen); ok {
		r0 = rf(limit)
	} else {
		r0 = ret.Get(0).(models.LimitKonsumen)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.LimitKonsumen) error); ok {
		r1 = rf(limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: limit
func (_m *LimitRepository) Delete(limit models.LimitKonsumen) (models.LimitKonsumen, error) {
	ret := _m.Called(limit)

	var r0 models.LimitKonsumen
	if rf, ok := ret.Get(0).(func(models.LimitKonsumen) models.LimitKonsumen); ok {
		r0 = rf(limit)
	} else {
		r0 = ret.Get(0).(models.LimitKonsumen)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.LimitKonsumen) error); ok {
		r1 = rf(limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields:
func (_m *LimitRepository) FindAll() ([]models.LimitKonsumen, error) {
	ret := _m.Called()

	var r0 []models.LimitKonsumen
	if rf, ok := ret.Get(0).(func() []models.LimitKonsumen); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.LimitKonsumen)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: ID
func (_m *LimitRepository) FindByID(ID int) (models.LimitKonsumen, error) {
	ret := _m.Called(ID)

	var r0 models.LimitKonsumen
	if rf, ok := ret.Get(0).(func(int) models.LimitKonsumen); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(models.LimitKonsumen)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: limit
func (_m *LimitRepository) Update(limit models.LimitKonsumen) (models.LimitKonsumen, error) {
	ret := _m.Called(limit)

	var r0 models.LimitKonsumen
	if rf, ok := ret.Get(0).(func(models.LimitKonsumen) models.LimitKonsumen); ok {
		r0 = rf(limit)
	} else {
		r0 = ret.Get(0).(models.LimitKonsumen)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.LimitKonsumen) error); ok {
		r1 = rf(limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewLimitRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewLimitRepository creates a new instance of LimitRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLimitRepository(t mockConstructorTestingTNewLimitRepository) *LimitRepository {
	mock := &LimitRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
