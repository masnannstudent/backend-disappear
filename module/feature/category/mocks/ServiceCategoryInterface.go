// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	entities "github.com/capstone-kelompok-7/backend-disappear/module/entities"
	mock "github.com/stretchr/testify/mock"
)

// ServiceCategoryInterface is an autogenerated mock type for the ServiceCategoryInterface type
type ServiceCategoryInterface struct {
	mock.Mock
}

// CalculatePaginationValues provides a mock function with given fields: page, totalItems, perPage
func (_m *ServiceCategoryInterface) CalculatePaginationValues(page int, totalItems int, perPage int) (int, int) {
	ret := _m.Called(page, totalItems, perPage)

	var r0 int
	var r1 int
	if rf, ok := ret.Get(0).(func(int, int, int) (int, int)); ok {
		return rf(page, totalItems, perPage)
	}
	if rf, ok := ret.Get(0).(func(int, int, int) int); ok {
		r0 = rf(page, totalItems, perPage)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(int, int, int) int); ok {
		r1 = rf(page, totalItems, perPage)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// CreateCategory provides a mock function with given fields: categoryData
func (_m *ServiceCategoryInterface) CreateCategory(categoryData *entities.CategoryModels) (*entities.CategoryModels, error) {
	ret := _m.Called(categoryData)

	var r0 *entities.CategoryModels
	var r1 error
	if rf, ok := ret.Get(0).(func(*entities.CategoryModels) (*entities.CategoryModels, error)); ok {
		return rf(categoryData)
	}
	if rf, ok := ret.Get(0).(func(*entities.CategoryModels) *entities.CategoryModels); ok {
		r0 = rf(categoryData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.CategoryModels)
		}
	}

	if rf, ok := ret.Get(1).(func(*entities.CategoryModels) error); ok {
		r1 = rf(categoryData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCategoryById provides a mock function with given fields: categoryID
func (_m *ServiceCategoryInterface) DeleteCategoryById(categoryID uint64) error {
	ret := _m.Called(categoryID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(categoryID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllCategory provides a mock function with given fields: page, perPage
func (_m *ServiceCategoryInterface) GetAllCategory(page int, perPage int) ([]*entities.CategoryModels, int64, error) {
	ret := _m.Called(page, perPage)

	var r0 []*entities.CategoryModels
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int) ([]*entities.CategoryModels, int64, error)); ok {
		return rf(page, perPage)
	}
	if rf, ok := ret.Get(0).(func(int, int) []*entities.CategoryModels); ok {
		r0 = rf(page, perPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.CategoryModels)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) int64); ok {
		r1 = rf(page, perPage)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(page, perPage)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetCategoryById provides a mock function with given fields: categoryID
func (_m *ServiceCategoryInterface) GetCategoryById(categoryID uint64) (*entities.CategoryModels, error) {
	ret := _m.Called(categoryID)

	var r0 *entities.CategoryModels
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (*entities.CategoryModels, error)); ok {
		return rf(categoryID)
	}
	if rf, ok := ret.Get(0).(func(uint64) *entities.CategoryModels); ok {
		r0 = rf(categoryID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.CategoryModels)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(categoryID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCategoryByName provides a mock function with given fields: page, perPage, name
func (_m *ServiceCategoryInterface) GetCategoryByName(page int, perPage int, name string) ([]*entities.CategoryModels, int64, error) {
	ret := _m.Called(page, perPage, name)

	var r0 []*entities.CategoryModels
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int, string) ([]*entities.CategoryModels, int64, error)); ok {
		return rf(page, perPage, name)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []*entities.CategoryModels); ok {
		r0 = rf(page, perPage, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.CategoryModels)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) int64); ok {
		r1 = rf(page, perPage, name)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(int, int, string) error); ok {
		r2 = rf(page, perPage, name)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetNextPage provides a mock function with given fields: currentPage, totalPages
func (_m *ServiceCategoryInterface) GetNextPage(currentPage int, totalPages int) int {
	ret := _m.Called(currentPage, totalPages)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int) int); ok {
		r0 = rf(currentPage, totalPages)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetPrevPage provides a mock function with given fields: currentPage
func (_m *ServiceCategoryInterface) GetPrevPage(currentPage int) int {
	ret := _m.Called(currentPage)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(currentPage)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// UpdateCategoryById provides a mock function with given fields: categoryID, updatedCategory
func (_m *ServiceCategoryInterface) UpdateCategoryById(categoryID uint64, updatedCategory *entities.CategoryModels) error {
	ret := _m.Called(categoryID, updatedCategory)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64, *entities.CategoryModels) error); ok {
		r0 = rf(categoryID, updatedCategory)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewServiceCategoryInterface creates a new instance of ServiceCategoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceCategoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServiceCategoryInterface {
	mock := &ServiceCategoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}