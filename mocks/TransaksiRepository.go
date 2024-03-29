// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	models "kredit-plus/models"

	mock "github.com/stretchr/testify/mock"
)

// TransaksiRepository is an autogenerated mock type for the TransaksiRepository type
type TransaksiRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: model
func (_m *TransaksiRepository) Create(model models.Transaksi) (models.Transaksi, error) {
	ret := _m.Called(model)

	var r0 models.Transaksi
	if rf, ok := ret.Get(0).(func(models.Transaksi) models.Transaksi); ok {
		r0 = rf(model)
	} else {
		r0 = ret.Get(0).(models.Transaksi)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Transaksi) error); ok {
		r1 = rf(model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: model
func (_m *TransaksiRepository) Delete(model models.Transaksi) (models.Transaksi, error) {
	ret := _m.Called(model)

	var r0 models.Transaksi
	if rf, ok := ret.Get(0).(func(models.Transaksi) models.Transaksi); ok {
		r0 = rf(model)
	} else {
		r0 = ret.Get(0).(models.Transaksi)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Transaksi) error); ok {
		r1 = rf(model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields:
func (_m *TransaksiRepository) FindAll() ([]models.Transaksi, error) {
	ret := _m.Called()

	var r0 []models.Transaksi
	if rf, ok := ret.Get(0).(func() []models.Transaksi); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Transaksi)
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
func (_m *TransaksiRepository) FindByID(ID int) (models.Transaksi, error) {
	ret := _m.Called(ID)

	var r0 models.Transaksi
	if rf, ok := ret.Get(0).(func(int) models.Transaksi); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(models.Transaksi)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: model
func (_m *TransaksiRepository) Update(model models.Transaksi) (models.Transaksi, error) {
	ret := _m.Called(model)

	var r0 models.Transaksi
	if rf, ok := ret.Get(0).(func(models.Transaksi) models.Transaksi); ok {
		r0 = rf(model)
	} else {
		r0 = ret.Get(0).(models.Transaksi)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Transaksi) error); ok {
		r1 = rf(model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTransaksiRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransaksiRepository creates a new instance of TransaksiRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransaksiRepository(t mockConstructorTestingTNewTransaksiRepository) *TransaksiRepository {
	mock := &TransaksiRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
