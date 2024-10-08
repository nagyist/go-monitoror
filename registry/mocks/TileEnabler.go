// Code generated by mockery v2.46.2. DO NOT EDIT.

package mocks

import (
	models "github.com/monitoror/monitoror/models"
	mock "github.com/stretchr/testify/mock"

	params "github.com/monitoror/monitoror/internal/pkg/monitorable/params"
)

// TileEnabler is an autogenerated mock type for the TileEnabler type
type TileEnabler struct {
	mock.Mock
}

// Enable provides a mock function with given fields: variantName, paramsValidator, routePath
func (_m *TileEnabler) Enable(variantName models.VariantName, paramsValidator params.Validator, routePath string) {
	_m.Called(variantName, paramsValidator, routePath)
}

// NewTileEnabler creates a new instance of TileEnabler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTileEnabler(t interface {
	mock.TestingT
	Cleanup(func())
}) *TileEnabler {
	mock := &TileEnabler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
