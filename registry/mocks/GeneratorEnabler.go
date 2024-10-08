// Code generated by mockery v2.46.2. DO NOT EDIT.

package mocks

import (
	configmodels "github.com/monitoror/monitoror/api/config/models"
	mock "github.com/stretchr/testify/mock"

	models "github.com/monitoror/monitoror/models"

	params "github.com/monitoror/monitoror/internal/pkg/monitorable/params"
)

// GeneratorEnabler is an autogenerated mock type for the GeneratorEnabler type
type GeneratorEnabler struct {
	mock.Mock
}

// Enable provides a mock function with given fields: variantName, generatorParamsValidator, tileGeneratorFunction
func (_m *GeneratorEnabler) Enable(variantName models.VariantName, generatorParamsValidator params.Validator, tileGeneratorFunction configmodels.TileGeneratorFunction) {
	_m.Called(variantName, generatorParamsValidator, tileGeneratorFunction)
}

// NewGeneratorEnabler creates a new instance of GeneratorEnabler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGeneratorEnabler(t interface {
	mock.TestingT
	Cleanup(func())
}) *GeneratorEnabler {
	mock := &GeneratorEnabler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
