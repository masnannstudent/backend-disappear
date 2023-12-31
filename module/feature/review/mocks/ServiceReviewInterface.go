// Code generated by mockery v2.36.1. DO NOT EDIT.

package mocks

import (
	entities "github.com/capstone-kelompok-7/backend-disappear/module/entities"
	mock "github.com/stretchr/testify/mock"
)

// ServiceReviewInterface is an autogenerated mock type for the ServiceReviewInterface type
type ServiceReviewInterface struct {
	mock.Mock
}

// CountAverageRating provides a mock function with given fields: productID
func (_m *ServiceReviewInterface) CountAverageRating(productID uint64) (float64, error) {
	ret := _m.Called(productID)

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (float64, error)); ok {
		return rf(productID)
	}
	if rf, ok := ret.Get(0).(func(uint64) float64); ok {
		r0 = rf(productID)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateReview provides a mock function with given fields: newData
func (_m *ServiceReviewInterface) CreateReview(newData *entities.ReviewModels) (*entities.ReviewModels, error) {
	ret := _m.Called(newData)

	var r0 *entities.ReviewModels
	var r1 error
	if rf, ok := ret.Get(0).(func(*entities.ReviewModels) (*entities.ReviewModels, error)); ok {
		return rf(newData)
	}
	if rf, ok := ret.Get(0).(func(*entities.ReviewModels) *entities.ReviewModels); ok {
		r0 = rf(newData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.ReviewModels)
		}
	}

	if rf, ok := ret.Get(1).(func(*entities.ReviewModels) error); ok {
		r1 = rf(newData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateReviewImages provides a mock function with given fields: reviewData
func (_m *ServiceReviewInterface) CreateReviewImages(reviewData *entities.ReviewPhotoModels) (*entities.ReviewPhotoModels, error) {
	ret := _m.Called(reviewData)

	var r0 *entities.ReviewPhotoModels
	var r1 error
	if rf, ok := ret.Get(0).(func(*entities.ReviewPhotoModels) (*entities.ReviewPhotoModels, error)); ok {
		return rf(reviewData)
	}
	if rf, ok := ret.Get(0).(func(*entities.ReviewPhotoModels) *entities.ReviewPhotoModels); ok {
		r0 = rf(reviewData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.ReviewPhotoModels)
		}
	}

	if rf, ok := ret.Get(1).(func(*entities.ReviewPhotoModels) error); ok {
		r1 = rf(reviewData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReviewById provides a mock function with given fields: reviewID
func (_m *ServiceReviewInterface) GetReviewById(reviewID uint64) (*entities.ReviewModels, error) {
	ret := _m.Called(reviewID)

	var r0 *entities.ReviewModels
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (*entities.ReviewModels, error)); ok {
		return rf(reviewID)
	}
	if rf, ok := ret.Get(0).(func(uint64) *entities.ReviewModels); ok {
		r0 = rf(reviewID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.ReviewModels)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(reviewID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReviewsProductByID provides a mock function with given fields: productID
func (_m *ServiceReviewInterface) GetReviewsProductByID(productID uint64) (*entities.ProductModels, error) {
	ret := _m.Called(productID)

	var r0 *entities.ProductModels
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (*entities.ProductModels, error)); ok {
		return rf(productID)
	}
	if rf, ok := ret.Get(0).(func(uint64) *entities.ProductModels); ok {
		r0 = rf(productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.ProductModels)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewServiceReviewInterface creates a new instance of ServiceReviewInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceReviewInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServiceReviewInterface {
	mock := &ServiceReviewInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
