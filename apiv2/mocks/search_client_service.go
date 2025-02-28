// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"

	search "github.com/spotmaxtech/goharbor-client/v5/apiv2/internal/api/client/search"
)

// MockSearchClientService is an autogenerated mock type for the ClientService type
type MockSearchClientService struct {
	mock.Mock
}

// Search provides a mock function with given fields: params, authInfo
func (_m *MockSearchClientService) Search(params *search.SearchParams, authInfo runtime.ClientAuthInfoWriter) (*search.SearchOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *search.SearchOK
	if rf, ok := ret.Get(0).(func(*search.SearchParams, runtime.ClientAuthInfoWriter) *search.SearchOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*search.SearchOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*search.SearchParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockSearchClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}
