// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	cart "advancerentbook-api/features/cart"

	mock "github.com/stretchr/testify/mock"
)

// CartData is an autogenerated mock type for the CartData type
type CartData struct {
	mock.Mock
}

// AddCart provides a mock function with given fields: userID, bookID, newCart
func (_m *CartData) AddCart(userID uint, bookID uint, newCart cart.Core) (cart.Core, error) {
	ret := _m.Called(userID, bookID, newCart)

	var r0 cart.Core
	if rf, ok := ret.Get(0).(func(uint, uint, cart.Core) cart.Core); ok {
		r0 = rf(userID, bookID, newCart)
	} else {
		r0 = ret.Get(0).(cart.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint, cart.Core) error); ok {
		r1 = rf(userID, bookID, newCart)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCart provides a mock function with given fields: userID, cartID
func (_m *CartData) DeleteCart(userID uint, cartID uint) error {
	ret := _m.Called(userID, cartID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(userID, cartID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ShowCart provides a mock function with given fields: userID
func (_m *CartData) ShowCart(userID uint) ([]cart.Core, error) {
	ret := _m.Called(userID)

	var r0 []cart.Core
	if rf, ok := ret.Get(0).(func(uint) []cart.Core); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cart.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCartData interface {
	mock.TestingT
	Cleanup(func())
}

// NewCartData creates a new instance of CartData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCartData(t mockConstructorTestingTNewCartData) *CartData {
	mock := &CartData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
