// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	entity "github.com/rahmatrdn/go-skeleton/entity"
)

// LogUsecase is an autogenerated mock type for the LogUsecase type
type LogUsecase struct {
	mock.Mock
}

// Log provides a mock function with given fields: status, message, funcName, err, logFields, processName
func (_m *LogUsecase) Log(status entity.LogType, message string, funcName string, err error, logFields map[string]string, processName string) {
	_m.Called(status, message, funcName, err, logFields, processName)
}

type mockConstructorTestingTNewLogUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewLogUsecase creates a new instance of LogUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLogUsecase(t mockConstructorTestingTNewLogUsecase) *LogUsecase {
	mock := &LogUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
