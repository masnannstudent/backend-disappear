// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// HandlerVoucherInterface is an autogenerated mock type for the HandlerVoucherInterface type
type HandlerVoucherInterface struct {
	mock.Mock
}

// ClaimVoucher provides a mock function with given fields:
func (_m *HandlerVoucherInterface) ClaimVoucher() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// CreateVoucher provides a mock function with given fields:
func (_m *HandlerVoucherInterface) CreateVoucher() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// DeleteVoucherById provides a mock function with given fields:
func (_m *HandlerVoucherInterface) DeleteVoucherById() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// GetAllVouchers provides a mock function with given fields:
func (_m *HandlerVoucherInterface) GetAllVouchers() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// GetAllVouchersToClaims provides a mock function with given fields:
func (_m *HandlerVoucherInterface) GetAllVouchersToClaims() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// GetVoucherById provides a mock function with given fields:
func (_m *HandlerVoucherInterface) GetVoucherById() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// GetVoucherUser provides a mock function with given fields:
func (_m *HandlerVoucherInterface) GetVoucherUser() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// UpdateVouchers provides a mock function with given fields:
func (_m *HandlerVoucherInterface) UpdateVouchers() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// NewHandlerVoucherInterface creates a new instance of HandlerVoucherInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHandlerVoucherInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *HandlerVoucherInterface {
	mock := &HandlerVoucherInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}