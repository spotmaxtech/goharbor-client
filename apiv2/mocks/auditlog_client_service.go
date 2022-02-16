// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	auditlog "github.com/spotmaxtech/goharbor-client/v5/apiv2/internal/api/client/auditlog"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// MockAuditlogClientService is an autogenerated mock type for the ClientService type
type MockAuditlogClientService struct {
	mock.Mock
}

// ListAuditLogs provides a mock function with given fields: params, authInfo
func (_m *MockAuditlogClientService) ListAuditLogs(params *auditlog.ListAuditLogsParams, authInfo runtime.ClientAuthInfoWriter) (*auditlog.ListAuditLogsOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *auditlog.ListAuditLogsOK
	if rf, ok := ret.Get(0).(func(*auditlog.ListAuditLogsParams, runtime.ClientAuthInfoWriter) *auditlog.ListAuditLogsOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auditlog.ListAuditLogsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*auditlog.ListAuditLogsParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockAuditlogClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}
