// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	entities "github.com/capstone-kelompok-7/backend-disappear/module/entities"
	dto "github.com/capstone-kelompok-7/backend-disappear/module/feature/order/dto"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// RepositoryOrderInterface is an autogenerated mock type for the RepositoryOrderInterface type
type RepositoryOrderInterface struct {
	mock.Mock
}

// AcceptOrder provides a mock function with given fields: orderID, orderStatus
func (_m *RepositoryOrderInterface) AcceptOrder(orderID string, orderStatus string) error {
	ret := _m.Called(orderID, orderStatus)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(orderID, orderStatus)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckTransaction provides a mock function with given fields: orderID
func (_m *RepositoryOrderInterface) CheckTransaction(orderID string) (dto.Status, error) {
	ret := _m.Called(orderID)

	var r0 dto.Status
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (dto.Status, error)); ok {
		return rf(orderID)
	}
	if rf, ok := ret.Get(0).(func(string) dto.Status); ok {
		r0 = rf(orderID)
	} else {
		r0 = ret.Get(0).(dto.Status)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(orderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConfirmPayment provides a mock function with given fields: orderID, orderStatus, paymentStatus
func (_m *RepositoryOrderInterface) ConfirmPayment(orderID string, orderStatus string, paymentStatus string) error {
	ret := _m.Called(orderID, orderStatus, paymentStatus)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(orderID, orderStatus, paymentStatus)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateOrder provides a mock function with given fields: newOrder
func (_m *RepositoryOrderInterface) CreateOrder(newOrder *entities.OrderModels) (*entities.OrderModels, error) {
	ret := _m.Called(newOrder)

	var r0 *entities.OrderModels
	var r1 error
	if rf, ok := ret.Get(0).(func(*entities.OrderModels) (*entities.OrderModels, error)); ok {
		return rf(newOrder)
	}
	if rf, ok := ret.Get(0).(func(*entities.OrderModels) *entities.OrderModels); ok {
		r0 = rf(newOrder)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.OrderModels)
		}
	}

	if rf, ok := ret.Get(1).(func(*entities.OrderModels) error); ok {
		r1 = rf(newOrder)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields: page, perPage
func (_m *RepositoryOrderInterface) FindAll(page int, perPage int) ([]*entities.OrderModels, error) {
	ret := _m.Called(page, perPage)

	var r0 []*entities.OrderModels
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) ([]*entities.OrderModels, error)); ok {
		return rf(page, perPage)
	}
	if rf, ok := ret.Get(0).(func(int, int) []*entities.OrderModels); ok {
		r0 = rf(page, perPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.OrderModels)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(page, perPage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByName provides a mock function with given fields: page, perPage, name
func (_m *RepositoryOrderInterface) FindByName(page int, perPage int, name string) ([]*entities.OrderModels, error) {
	ret := _m.Called(page, perPage, name)

	var r0 []*entities.OrderModels
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, string) ([]*entities.OrderModels, error)); ok {
		return rf(page, perPage, name)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []*entities.OrderModels); ok {
		r0 = rf(page, perPage, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.OrderModels)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) error); ok {
		r1 = rf(page, perPage, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllOrdersByUserID provides a mock function with given fields: userID
func (_m *RepositoryOrderInterface) GetAllOrdersByUserID(userID uint64) ([]*entities.OrderModels, error) {
	ret := _m.Called(userID)

	var r0 []*entities.OrderModels
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) ([]*entities.OrderModels, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint64) []*entities.OrderModels); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.OrderModels)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllOrdersWithFilter provides a mock function with given fields: userID, orderStatus
func (_m *RepositoryOrderInterface) GetAllOrdersWithFilter(userID uint64, orderStatus string) ([]*entities.OrderModels, error) {
	ret := _m.Called(userID, orderStatus)

	var r0 []*entities.OrderModels
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64, string) ([]*entities.OrderModels, error)); ok {
		return rf(userID, orderStatus)
	}
	if rf, ok := ret.Get(0).(func(uint64, string) []*entities.OrderModels); ok {
		r0 = rf(userID, orderStatus)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.OrderModels)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64, string) error); ok {
		r1 = rf(userID, orderStatus)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderByDateRange provides a mock function with given fields: startDate, endDate, offset, limit
func (_m *RepositoryOrderInterface) GetOrderByDateRange(startDate time.Time, endDate time.Time, offset int, limit int) ([]*entities.OrderModels, error) {
	ret := _m.Called(startDate, endDate, offset, limit)

	var r0 []*entities.OrderModels
	var r1 error
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, int, int) ([]*entities.OrderModels, error)); ok {
		return rf(startDate, endDate, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, int, int) []*entities.OrderModels); ok {
		r0 = rf(startDate, endDate, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.OrderModels)
		}
	}

	if rf, ok := ret.Get(1).(func(time.Time, time.Time, int, int) error); ok {
		r1 = rf(startDate, endDate, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderByDateRangeAndPaymentStatus provides a mock function with given fields: startDate, endDate, status, offset, limit
func (_m *RepositoryOrderInterface) GetOrderByDateRangeAndPaymentStatus(startDate time.Time, endDate time.Time, status string, offset int, limit int) ([]*entities.OrderModels, error) {
	ret := _m.Called(startDate, endDate, status, offset, limit)

	var r0 []*entities.OrderModels
	var r1 error
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, string, int, int) ([]*entities.OrderModels, error)); ok {
		return rf(startDate, endDate, status, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, string, int, int) []*entities.OrderModels); ok {
		r0 = rf(startDate, endDate, status, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.OrderModels)
		}
	}

	if rf, ok := ret.Get(1).(func(time.Time, time.Time, string, int, int) error); ok {
		r1 = rf(startDate, endDate, status, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderByDateRangeAndPaymentStatusAndSearch provides a mock function with given fields: startDate, endDate, status, search, offset, limit
func (_m *RepositoryOrderInterface) GetOrderByDateRangeAndPaymentStatusAndSearch(startDate time.Time, endDate time.Time, status string, search string, offset int, limit int) ([]*entities.OrderModels, error) {
	ret := _m.Called(startDate, endDate, status, search, offset, limit)

	var r0 []*entities.OrderModels
	var r1 error
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, string, string, int, int) ([]*entities.OrderModels, error)); ok {
		return rf(startDate, endDate, status, search, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, string, string, int, int) []*entities.OrderModels); ok {
		r0 = rf(startDate, endDate, status, search, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.OrderModels)
		}
	}

	if rf, ok := ret.Get(1).(func(time.Time, time.Time, string, string, int, int) error); ok {
		r1 = rf(startDate, endDate, status, search, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderByDateRangeAndStatus provides a mock function with given fields: startDate, endDate, status, offset, limit
func (_m *RepositoryOrderInterface) GetOrderByDateRangeAndStatus(startDate time.Time, endDate time.Time, status string, offset int, limit int) ([]*entities.OrderModels, error) {
	ret := _m.Called(startDate, endDate, status, offset, limit)

	var r0 []*entities.OrderModels
	var r1 error
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, string, int, int) ([]*entities.OrderModels, error)); ok {
		return rf(startDate, endDate, status, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, string, int, int) []*entities.OrderModels); ok {
		r0 = rf(startDate, endDate, status, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.OrderModels)
		}
	}

	if rf, ok := ret.Get(1).(func(time.Time, time.Time, string, int, int) error); ok {
		r1 = rf(startDate, endDate, status, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderByDateRangeAndStatusAndSearch provides a mock function with given fields: startDate, endDate, status, search, offset, limit
func (_m *RepositoryOrderInterface) GetOrderByDateRangeAndStatusAndSearch(startDate time.Time, endDate time.Time, status string, search string, offset int, limit int) ([]*entities.OrderModels, error) {
	ret := _m.Called(startDate, endDate, status, search, offset, limit)

	var r0 []*entities.OrderModels
	var r1 error
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, string, string, int, int) ([]*entities.OrderModels, error)); ok {
		return rf(startDate, endDate, status, search, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, string, string, int, int) []*entities.OrderModels); ok {
		r0 = rf(startDate, endDate, status, search, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.OrderModels)
		}
	}

	if rf, ok := ret.Get(1).(func(time.Time, time.Time, string, string, int, int) error); ok {
		r1 = rf(startDate, endDate, status, search, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderById provides a mock function with given fields: orderID
func (_m *RepositoryOrderInterface) GetOrderById(orderID string) (*entities.OrderModels, error) {
	ret := _m.Called(orderID)

	var r0 *entities.OrderModels
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entities.OrderModels, error)); ok {
		return rf(orderID)
	}
	if rf, ok := ret.Get(0).(func(string) *entities.OrderModels); ok {
		r0 = rf(orderID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.OrderModels)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(orderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderByOrderStatus provides a mock function with given fields: orderStatus, offset, limit
func (_m *RepositoryOrderInterface) GetOrderByOrderStatus(orderStatus string, offset int, limit int) ([]*entities.OrderModels, error) {
	ret := _m.Called(orderStatus, offset, limit)

	var r0 []*entities.OrderModels
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int, int) ([]*entities.OrderModels, error)); ok {
		return rf(orderStatus, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(string, int, int) []*entities.OrderModels); ok {
		r0 = rf(orderStatus, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.OrderModels)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(orderStatus, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderByPaymentStatus provides a mock function with given fields: orderStatus, offset, limit
func (_m *RepositoryOrderInterface) GetOrderByPaymentStatus(orderStatus string, offset int, limit int) ([]*entities.OrderModels, error) {
	ret := _m.Called(orderStatus, offset, limit)

	var r0 []*entities.OrderModels
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int, int) ([]*entities.OrderModels, error)); ok {
		return rf(orderStatus, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(string, int, int) []*entities.OrderModels); ok {
		r0 = rf(orderStatus, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.OrderModels)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(orderStatus, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderBySearchAndPaymentStatus provides a mock function with given fields: status, search, offset, limit
func (_m *RepositoryOrderInterface) GetOrderBySearchAndPaymentStatus(status string, search string, offset int, limit int) ([]*entities.OrderModels, error) {
	ret := _m.Called(status, search, offset, limit)

	var r0 []*entities.OrderModels
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, int, int) ([]*entities.OrderModels, error)); ok {
		return rf(status, search, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(string, string, int, int) []*entities.OrderModels); ok {
		r0 = rf(status, search, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.OrderModels)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, int, int) error); ok {
		r1 = rf(status, search, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderCountByByOrderStatus provides a mock function with given fields: orderStatus
func (_m *RepositoryOrderInterface) GetOrderCountByByOrderStatus(orderStatus string) (int64, error) {
	ret := _m.Called(orderStatus)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int64, error)); ok {
		return rf(orderStatus)
	}
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(orderStatus)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(orderStatus)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderCountByByPaymentStatus provides a mock function with given fields: orderStatus
func (_m *RepositoryOrderInterface) GetOrderCountByByPaymentStatus(orderStatus string) (int64, error) {
	ret := _m.Called(orderStatus)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int64, error)); ok {
		return rf(orderStatus)
	}
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(orderStatus)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(orderStatus)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderCountByDateRange provides a mock function with given fields: startDate, endDate
func (_m *RepositoryOrderInterface) GetOrderCountByDateRange(startDate time.Time, endDate time.Time) (int64, error) {
	ret := _m.Called(startDate, endDate)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(time.Time, time.Time) (int64, error)); ok {
		return rf(startDate, endDate)
	}
	if rf, ok := ret.Get(0).(func(time.Time, time.Time) int64); ok {
		r0 = rf(startDate, endDate)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(time.Time, time.Time) error); ok {
		r1 = rf(startDate, endDate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderCountByDateRangeAndPaymentStatus provides a mock function with given fields: startDate, endDate, status
func (_m *RepositoryOrderInterface) GetOrderCountByDateRangeAndPaymentStatus(startDate time.Time, endDate time.Time, status string) (int64, error) {
	ret := _m.Called(startDate, endDate, status)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, string) (int64, error)); ok {
		return rf(startDate, endDate, status)
	}
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, string) int64); ok {
		r0 = rf(startDate, endDate, status)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(time.Time, time.Time, string) error); ok {
		r1 = rf(startDate, endDate, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderCountByDateRangeAndPaymentStatusAndSearch provides a mock function with given fields: startDate, endDate, status, search
func (_m *RepositoryOrderInterface) GetOrderCountByDateRangeAndPaymentStatusAndSearch(startDate time.Time, endDate time.Time, status string, search string) (int64, error) {
	ret := _m.Called(startDate, endDate, status, search)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, string, string) (int64, error)); ok {
		return rf(startDate, endDate, status, search)
	}
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, string, string) int64); ok {
		r0 = rf(startDate, endDate, status, search)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(time.Time, time.Time, string, string) error); ok {
		r1 = rf(startDate, endDate, status, search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderCountByDateRangeAndStatus provides a mock function with given fields: startDate, endDate, status
func (_m *RepositoryOrderInterface) GetOrderCountByDateRangeAndStatus(startDate time.Time, endDate time.Time, status string) (int64, error) {
	ret := _m.Called(startDate, endDate, status)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, string) (int64, error)); ok {
		return rf(startDate, endDate, status)
	}
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, string) int64); ok {
		r0 = rf(startDate, endDate, status)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(time.Time, time.Time, string) error); ok {
		r1 = rf(startDate, endDate, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderCountByDateRangeAndStatusAndSearch provides a mock function with given fields: startDate, endDate, status, search
func (_m *RepositoryOrderInterface) GetOrderCountByDateRangeAndStatusAndSearch(startDate time.Time, endDate time.Time, status string, search string) (int64, error) {
	ret := _m.Called(startDate, endDate, status, search)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, string, string) (int64, error)); ok {
		return rf(startDate, endDate, status, search)
	}
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, string, string) int64); ok {
		r0 = rf(startDate, endDate, status, search)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(time.Time, time.Time, string, string) error); ok {
		r1 = rf(startDate, endDate, status, search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderCountBySearchAndDateRange provides a mock function with given fields: startDate, endDate, search
func (_m *RepositoryOrderInterface) GetOrderCountBySearchAndDateRange(startDate time.Time, endDate time.Time, search string) (int64, error) {
	ret := _m.Called(startDate, endDate, search)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, string) (int64, error)); ok {
		return rf(startDate, endDate, search)
	}
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, string) int64); ok {
		r0 = rf(startDate, endDate, search)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(time.Time, time.Time, string) error); ok {
		r1 = rf(startDate, endDate, search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderCountBySearchAndPaymentStatus provides a mock function with given fields: status, search
func (_m *RepositoryOrderInterface) GetOrderCountBySearchAndPaymentStatus(status string, search string) (int64, error) {
	ret := _m.Called(status, search)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (int64, error)); ok {
		return rf(status, search)
	}
	if rf, ok := ret.Get(0).(func(string, string) int64); ok {
		r0 = rf(status, search)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(status, search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrdersBySearchAndDateRange provides a mock function with given fields: startDate, endDate, search, offset, limit
func (_m *RepositoryOrderInterface) GetOrdersBySearchAndDateRange(startDate time.Time, endDate time.Time, search string, offset int, limit int) ([]*entities.OrderModels, error) {
	ret := _m.Called(startDate, endDate, search, offset, limit)

	var r0 []*entities.OrderModels
	var r1 error
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, string, int, int) ([]*entities.OrderModels, error)); ok {
		return rf(startDate, endDate, search, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, string, int, int) []*entities.OrderModels); ok {
		r0 = rf(startDate, endDate, search, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.OrderModels)
		}
	}

	if rf, ok := ret.Get(1).(func(time.Time, time.Time, string, int, int) error); ok {
		r1 = rf(startDate, endDate, search, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrdersBySearchAndStatus provides a mock function with given fields: status, search, offset, limit
func (_m *RepositoryOrderInterface) GetOrdersBySearchAndStatus(status string, search string, offset int, limit int) ([]*entities.OrderModels, error) {
	ret := _m.Called(status, search, offset, limit)

	var r0 []*entities.OrderModels
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, int, int) ([]*entities.OrderModels, error)); ok {
		return rf(status, search, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(string, string, int, int) []*entities.OrderModels); ok {
		r0 = rf(status, search, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.OrderModels)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, int, int) error); ok {
		r1 = rf(status, search, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrdersCountBySearchAndStatus provides a mock function with given fields: status, search
func (_m *RepositoryOrderInterface) GetOrdersCountBySearchAndStatus(status string, search string) (int64, error) {
	ret := _m.Called(status, search)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (int64, error)); ok {
		return rf(status, search)
	}
	if rf, ok := ret.Get(0).(func(string, string) int64); ok {
		r0 = rf(status, search)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(status, search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalCustomerCountByName provides a mock function with given fields: name
func (_m *RepositoryOrderInterface) GetTotalCustomerCountByName(name string) (int64, error) {
	ret := _m.Called(name)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int64, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalOrderCount provides a mock function with given fields:
func (_m *RepositoryOrderInterface) GetTotalOrderCount() (int64, error) {
	ret := _m.Called()

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func() (int64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessGatewayPayment provides a mock function with given fields: totalAmountPaid, orderID, paymentMethod, name, email
func (_m *RepositoryOrderInterface) ProcessGatewayPayment(totalAmountPaid uint64, orderID string, paymentMethod string, name string, email string) (interface{}, error) {
	ret := _m.Called(totalAmountPaid, orderID, paymentMethod, name, email)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64, string, string, string, string) (interface{}, error)); ok {
		return rf(totalAmountPaid, orderID, paymentMethod, name, email)
	}
	if rf, ok := ret.Get(0).(func(uint64, string, string, string, string) interface{}); ok {
		r0 = rf(totalAmountPaid, orderID, paymentMethod, name, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(uint64, string, string, string, string) error); ok {
		r1 = rf(totalAmountPaid, orderID, paymentMethod, name, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Tracking provides a mock function with given fields: courier, awb
func (_m *RepositoryOrderInterface) Tracking(courier string, awb string) (map[string]interface{}, error) {
	ret := _m.Called(courier, awb)

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (map[string]interface{}, error)); ok {
		return rf(courier, awb)
	}
	if rf, ok := ret.Get(0).(func(string, string) map[string]interface{}); ok {
		r0 = rf(courier, awb)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(courier, awb)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOrderStatus provides a mock function with given fields: req
func (_m *RepositoryOrderInterface) UpdateOrderStatus(req *dto.UpdateOrderStatus) error {
	ret := _m.Called(req)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dto.UpdateOrderStatus) error); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRepositoryOrderInterface creates a new instance of RepositoryOrderInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepositoryOrderInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *RepositoryOrderInterface {
	mock := &RepositoryOrderInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
