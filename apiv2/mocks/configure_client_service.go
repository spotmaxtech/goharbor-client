// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	configure "github.com/spotmaxtech/goharbor-client/v5/apiv2/internal/api/client/configure"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// MockConfigureClientService is an autogenerated mock type for the ClientService type
type MockConfigureClientService struct {
	mock.Mock
}

// GetConfigurations provides a mock function with given fields: params, authInfo
func (_m *MockConfigureClientService) GetConfigurations(params *configure.GetConfigurationsParams, authInfo runtime.ClientAuthInfoWriter) (*configure.GetConfigurationsOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *configure.GetConfigurationsOK
	if rf, ok := ret.Get(0).(func(*configure.GetConfigurationsParams, runtime.ClientAuthInfoWriter) *configure.GetConfigurationsOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*configure.GetConfigurationsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*configure.GetConfigurationsParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInternalconfig provides a mock function with given fields: params, authInfo
func (_m *MockConfigureClientService) GetInternalconfig(params *configure.GetInternalconfigParams, authInfo runtime.ClientAuthInfoWriter) (*configure.GetInternalconfigOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *configure.GetInternalconfigOK
	if rf, ok := ret.Get(0).(func(*configure.GetInternalconfigParams, runtime.ClientAuthInfoWriter) *configure.GetInternalconfigOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*configure.GetInternalconfigOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*configure.GetInternalconfigParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockConfigureClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// UpdateConfigurations provides a mock function with given fields: params, authInfo
func (_m *MockConfigureClientService) UpdateConfigurations(params *configure.UpdateConfigurationsParams, authInfo runtime.ClientAuthInfoWriter) (*configure.UpdateConfigurationsOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *configure.UpdateConfigurationsOK
	if rf, ok := ret.Get(0).(func(*configure.UpdateConfigurationsParams, runtime.ClientAuthInfoWriter) *configure.UpdateConfigurationsOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*configure.UpdateConfigurationsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*configure.UpdateConfigurationsParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
