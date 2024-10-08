// Code generated by mockery v2.46.2. DO NOT EDIT.

package mocks

import (
	configmodels "github.com/monitoror/monitoror/api/config/models"
	mock "github.com/stretchr/testify/mock"

	models "github.com/monitoror/monitoror/monitorables/jenkins/api/models"

	monitorormodels "github.com/monitoror/monitoror/models"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// Build provides a mock function with given fields: params
func (_m *Usecase) Build(params *models.BuildParams) (*monitorormodels.Tile, error) {
	ret := _m.Called(params)

	if len(ret) == 0 {
		panic("no return value specified for Build")
	}

	var r0 *monitorormodels.Tile
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.BuildParams) (*monitorormodels.Tile, error)); ok {
		return rf(params)
	}
	if rf, ok := ret.Get(0).(func(*models.BuildParams) *monitorormodels.Tile); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*monitorormodels.Tile)
		}
	}

	if rf, ok := ret.Get(1).(func(*models.BuildParams) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildGenerator provides a mock function with given fields: params
func (_m *Usecase) BuildGenerator(params interface{}) ([]configmodels.GeneratedTile, error) {
	ret := _m.Called(params)

	if len(ret) == 0 {
		panic("no return value specified for BuildGenerator")
	}

	var r0 []configmodels.GeneratedTile
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}) ([]configmodels.GeneratedTile, error)); ok {
		return rf(params)
	}
	if rf, ok := ret.Get(0).(func(interface{}) []configmodels.GeneratedTile); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]configmodels.GeneratedTile)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUsecase creates a new instance of Usecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *Usecase {
	mock := &Usecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
