// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	immutable "github.com/spotmaxtech/goharbor-client/v5/apiv2/internal/api/client/immutable"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// MockImmutableClientService is an autogenerated mock type for the ClientService type
type MockImmutableClientService struct {
	mock.Mock
}

// CreateImmuRule provides a mock function with given fields: params, authInfo
func (_m *MockImmutableClientService) CreateImmuRule(params *immutable.CreateImmuRuleParams, authInfo runtime.ClientAuthInfoWriter) (*immutable.CreateImmuRuleCreated, error) {
	ret := _m.Called(params, authInfo)

	var r0 *immutable.CreateImmuRuleCreated
	if rf, ok := ret.Get(0).(func(*immutable.CreateImmuRuleParams, runtime.ClientAuthInfoWriter) *immutable.CreateImmuRuleCreated); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*immutable.CreateImmuRuleCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*immutable.CreateImmuRuleParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteImmuRule provides a mock function with given fields: params, authInfo
func (_m *MockImmutableClientService) DeleteImmuRule(params *immutable.DeleteImmuRuleParams, authInfo runtime.ClientAuthInfoWriter) (*immutable.DeleteImmuRuleOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *immutable.DeleteImmuRuleOK
	if rf, ok := ret.Get(0).(func(*immutable.DeleteImmuRuleParams, runtime.ClientAuthInfoWriter) *immutable.DeleteImmuRuleOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*immutable.DeleteImmuRuleOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*immutable.DeleteImmuRuleParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListImmuRules provides a mock function with given fields: params, authInfo
func (_m *MockImmutableClientService) ListImmuRules(params *immutable.ListImmuRulesParams, authInfo runtime.ClientAuthInfoWriter) (*immutable.ListImmuRulesOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *immutable.ListImmuRulesOK
	if rf, ok := ret.Get(0).(func(*immutable.ListImmuRulesParams, runtime.ClientAuthInfoWriter) *immutable.ListImmuRulesOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*immutable.ListImmuRulesOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*immutable.ListImmuRulesParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockImmutableClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// UpdateImmuRule provides a mock function with given fields: params, authInfo
func (_m *MockImmutableClientService) UpdateImmuRule(params *immutable.UpdateImmuRuleParams, authInfo runtime.ClientAuthInfoWriter) (*immutable.UpdateImmuRuleOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *immutable.UpdateImmuRuleOK
	if rf, ok := ret.Get(0).(func(*immutable.UpdateImmuRuleParams, runtime.ClientAuthInfoWriter) *immutable.UpdateImmuRuleOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*immutable.UpdateImmuRuleOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*immutable.UpdateImmuRuleParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
