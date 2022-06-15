// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	model "gitlab.com/jeremylo/microsvc/productsvc/domain/model"

	mock "github.com/stretchr/testify/mock"
)

// CustomerService is an autogenerated mock type for the CustomerService type
type CustomerService struct {
	mock.Mock
}

// DeleteCustomer provides a mock function with given fields: id
func (_m *CustomerService) DeleteCustomer(id int64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchCustomer provides a mock function with given fields:
func (_m *CustomerService) FetchCustomer() *[]model.Customer {
	ret := _m.Called()

	var r0 *[]model.Customer
	if rf, ok := ret.Get(0).(func() *[]model.Customer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]model.Customer)
		}
	}

	return r0
}

// SaveCustomer provides a mock function with given fields: customer
func (_m *CustomerService) SaveCustomer(customer *model.Customer) error {
	ret := _m.Called(customer)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Customer) error); ok {
		r0 = rf(customer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ShowCustomer provides a mock function with given fields: id
func (_m *CustomerService) ShowCustomer(id int) *model.Customer {
	ret := _m.Called(id)

	var r0 *model.Customer
	if rf, ok := ret.Get(0).(func(int) *model.Customer); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Customer)
		}
	}

	return r0
}

// UpdateCustomer provides a mock function with given fields: customer, id
func (_m *CustomerService) UpdateCustomer(customer *model.Customer, id int64) error {
	ret := _m.Called(customer, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Customer, int64) error); ok {
		r0 = rf(customer, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
