// Code generated by mockery v2.36.1. DO NOT EDIT.

package mocks

import (
	entities "github.com/capstone-kelompok-7/backend-disappear/module/entities"
	dto "github.com/capstone-kelompok-7/backend-disappear/module/feature/challenge/dto"

	mock "github.com/stretchr/testify/mock"
)

// ServiceChallengeInterface is an autogenerated mock type for the ServiceChallengeInterface type
type ServiceChallengeInterface struct {
	mock.Mock
}

// CalculatePaginationValues provides a mock function with given fields: page, totalItems, perPage
func (_m *ServiceChallengeInterface) CalculatePaginationValues(page int, totalItems int, perPage int) (int, int) {
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

// CreateChallenge provides a mock function with given fields: newData
func (_m *ServiceChallengeInterface) CreateChallenge(newData *entities.ChallengeModels) (*entities.ChallengeModels, error) {
	ret := _m.Called(newData)

	var r0 *entities.ChallengeModels
	var r1 error
	if rf, ok := ret.Get(0).(func(*entities.ChallengeModels) (*entities.ChallengeModels, error)); ok {
		return rf(newData)
	}
	if rf, ok := ret.Get(0).(func(*entities.ChallengeModels) *entities.ChallengeModels); ok {
		r0 = rf(newData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.ChallengeModels)
		}
	}

	if rf, ok := ret.Get(1).(func(*entities.ChallengeModels) error); ok {
		r1 = rf(newData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSubmitChallengeForm provides a mock function with given fields: form
func (_m *ServiceChallengeInterface) CreateSubmitChallengeForm(form *entities.ChallengeFormModels) (*entities.ChallengeFormModels, error) {
	ret := _m.Called(form)

	var r0 *entities.ChallengeFormModels
	var r1 error
	if rf, ok := ret.Get(0).(func(*entities.ChallengeFormModels) (*entities.ChallengeFormModels, error)); ok {
		return rf(form)
	}
	if rf, ok := ret.Get(0).(func(*entities.ChallengeFormModels) *entities.ChallengeFormModels); ok {
		r0 = rf(form)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.ChallengeFormModels)
		}
	}

	if rf, ok := ret.Get(1).(func(*entities.ChallengeFormModels) error); ok {
		r1 = rf(form)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteChallenge provides a mock function with given fields: id
func (_m *ServiceChallengeInterface) DeleteChallenge(id uint64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllChallenges provides a mock function with given fields: page, perPage
func (_m *ServiceChallengeInterface) GetAllChallenges(page int, perPage int) ([]*entities.ChallengeModels, int64, error) {
	ret := _m.Called(page, perPage)

	var r0 []*entities.ChallengeModels
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int) ([]*entities.ChallengeModels, int64, error)); ok {
		return rf(page, perPage)
	}
	if rf, ok := ret.Get(0).(func(int, int) []*entities.ChallengeModels); ok {
		r0 = rf(page, perPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.ChallengeModels)
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

// GetAllSubmitChallengeForm provides a mock function with given fields: page, perPage
func (_m *ServiceChallengeInterface) GetAllSubmitChallengeForm(page int, perPage int) ([]*entities.ChallengeFormModels, int64, error) {
	ret := _m.Called(page, perPage)

	var r0 []*entities.ChallengeFormModels
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int) ([]*entities.ChallengeFormModels, int64, error)); ok {
		return rf(page, perPage)
	}
	if rf, ok := ret.Get(0).(func(int, int) []*entities.ChallengeFormModels); ok {
		r0 = rf(page, perPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.ChallengeFormModels)
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

// GetChallengeById provides a mock function with given fields: id
func (_m *ServiceChallengeInterface) GetChallengeById(id uint64) (*entities.ChallengeModels, error) {
	ret := _m.Called(id)

	var r0 *entities.ChallengeModels
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (*entities.ChallengeModels, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint64) *entities.ChallengeModels); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.ChallengeModels)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChallengeByStatus provides a mock function with given fields: page, perPage, status
func (_m *ServiceChallengeInterface) GetChallengeByStatus(page int, perPage int, status string) ([]*entities.ChallengeModels, int64, error) {
	ret := _m.Called(page, perPage, status)

	var r0 []*entities.ChallengeModels
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int, string) ([]*entities.ChallengeModels, int64, error)); ok {
		return rf(page, perPage, status)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []*entities.ChallengeModels); ok {
		r0 = rf(page, perPage, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.ChallengeModels)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) int64); ok {
		r1 = rf(page, perPage, status)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(int, int, string) error); ok {
		r2 = rf(page, perPage, status)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetChallengeByTitle provides a mock function with given fields: page, perPage, title
func (_m *ServiceChallengeInterface) GetChallengeByTitle(page int, perPage int, title string) ([]*entities.ChallengeModels, int64, error) {
	ret := _m.Called(page, perPage, title)

	var r0 []*entities.ChallengeModels
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int, string) ([]*entities.ChallengeModels, int64, error)); ok {
		return rf(page, perPage, title)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []*entities.ChallengeModels); ok {
		r0 = rf(page, perPage, title)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.ChallengeModels)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) int64); ok {
		r1 = rf(page, perPage, title)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(int, int, string) error); ok {
		r2 = rf(page, perPage, title)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetChallengesBySearchAndStatus provides a mock function with given fields: page, perPage, search, status
func (_m *ServiceChallengeInterface) GetChallengesBySearchAndStatus(page int, perPage int, search string, status string) ([]*entities.ChallengeModels, int64, error) {
	ret := _m.Called(page, perPage, search, status)

	var r0 []*entities.ChallengeModels
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int, string, string) ([]*entities.ChallengeModels, int64, error)); ok {
		return rf(page, perPage, search, status)
	}
	if rf, ok := ret.Get(0).(func(int, int, string, string) []*entities.ChallengeModels); ok {
		r0 = rf(page, perPage, search, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.ChallengeModels)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string, string) int64); ok {
		r1 = rf(page, perPage, search, status)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(int, int, string, string) error); ok {
		r2 = rf(page, perPage, search, status)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetNextPage provides a mock function with given fields: currentPage, totalPages
func (_m *ServiceChallengeInterface) GetNextPage(currentPage int, totalPages int) int {
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
func (_m *ServiceChallengeInterface) GetPrevPage(currentPage int) int {
	ret := _m.Called(currentPage)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(currentPage)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetSubmitChallengeFormByDateRange provides a mock function with given fields: page, perPage, filterType
func (_m *ServiceChallengeInterface) GetSubmitChallengeFormByDateRange(page int, perPage int, filterType string) ([]*entities.ChallengeFormModels, int64, error) {
	ret := _m.Called(page, perPage, filterType)

	var r0 []*entities.ChallengeFormModels
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int, string) ([]*entities.ChallengeFormModels, int64, error)); ok {
		return rf(page, perPage, filterType)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []*entities.ChallengeFormModels); ok {
		r0 = rf(page, perPage, filterType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.ChallengeFormModels)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) int64); ok {
		r1 = rf(page, perPage, filterType)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(int, int, string) error); ok {
		r2 = rf(page, perPage, filterType)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetSubmitChallengeFormById provides a mock function with given fields: id
func (_m *ServiceChallengeInterface) GetSubmitChallengeFormById(id uint64) (*entities.ChallengeFormModels, error) {
	ret := _m.Called(id)

	var r0 *entities.ChallengeFormModels
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (*entities.ChallengeFormModels, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint64) *entities.ChallengeFormModels); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.ChallengeFormModels)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSubmitChallengeFormByStatus provides a mock function with given fields: page, perPage, status
func (_m *ServiceChallengeInterface) GetSubmitChallengeFormByStatus(page int, perPage int, status string) ([]*entities.ChallengeFormModels, int64, error) {
	ret := _m.Called(page, perPage, status)

	var r0 []*entities.ChallengeFormModels
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int, string) ([]*entities.ChallengeFormModels, int64, error)); ok {
		return rf(page, perPage, status)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []*entities.ChallengeFormModels); ok {
		r0 = rf(page, perPage, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.ChallengeFormModels)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) int64); ok {
		r1 = rf(page, perPage, status)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(int, int, string) error); ok {
		r2 = rf(page, perPage, status)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetSubmitChallengeFormByStatusAndDate provides a mock function with given fields: page, perPage, filterStatus, filterType
func (_m *ServiceChallengeInterface) GetSubmitChallengeFormByStatusAndDate(page int, perPage int, filterStatus string, filterType string) ([]*entities.ChallengeFormModels, int64, error) {
	ret := _m.Called(page, perPage, filterStatus, filterType)

	var r0 []*entities.ChallengeFormModels
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int, string, string) ([]*entities.ChallengeFormModels, int64, error)); ok {
		return rf(page, perPage, filterStatus, filterType)
	}
	if rf, ok := ret.Get(0).(func(int, int, string, string) []*entities.ChallengeFormModels); ok {
		r0 = rf(page, perPage, filterStatus, filterType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.ChallengeFormModels)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string, string) int64); ok {
		r1 = rf(page, perPage, filterStatus, filterType)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(int, int, string, string) error); ok {
		r2 = rf(page, perPage, filterStatus, filterType)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateChallenge provides a mock function with given fields: id, updatedChallenge
func (_m *ServiceChallengeInterface) UpdateChallenge(id uint64, updatedChallenge *entities.ChallengeModels) (*entities.ChallengeModels, error) {
	ret := _m.Called(id, updatedChallenge)

	var r0 *entities.ChallengeModels
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64, *entities.ChallengeModels) (*entities.ChallengeModels, error)); ok {
		return rf(id, updatedChallenge)
	}
	if rf, ok := ret.Get(0).(func(uint64, *entities.ChallengeModels) *entities.ChallengeModels); ok {
		r0 = rf(id, updatedChallenge)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.ChallengeModels)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64, *entities.ChallengeModels) error); ok {
		r1 = rf(id, updatedChallenge)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSubmitChallengeForm provides a mock function with given fields: id, updatedStatus
func (_m *ServiceChallengeInterface) UpdateSubmitChallengeForm(id uint64, updatedStatus dto.UpdateChallengeFormStatusRequest) (*entities.ChallengeFormModels, error) {
	ret := _m.Called(id, updatedStatus)

	var r0 *entities.ChallengeFormModels
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64, dto.UpdateChallengeFormStatusRequest) (*entities.ChallengeFormModels, error)); ok {
		return rf(id, updatedStatus)
	}
	if rf, ok := ret.Get(0).(func(uint64, dto.UpdateChallengeFormStatusRequest) *entities.ChallengeFormModels); ok {
		r0 = rf(id, updatedStatus)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.ChallengeFormModels)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64, dto.UpdateChallengeFormStatusRequest) error); ok {
		r1 = rf(id, updatedStatus)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewServiceChallengeInterface creates a new instance of ServiceChallengeInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceChallengeInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServiceChallengeInterface {
	mock := &ServiceChallengeInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
