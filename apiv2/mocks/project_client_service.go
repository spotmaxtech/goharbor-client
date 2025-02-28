// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	runtime "github.com/go-openapi/runtime"
	project "github.com/spotmaxtech/goharbor-client/v5/apiv2/internal/api/client/project"
	mock "github.com/stretchr/testify/mock"
)

// MockProjectClientService is an autogenerated mock type for the ClientService type
type MockProjectClientService struct {
	mock.Mock
}

// CreateProject provides a mock function with given fields: params, authInfo
func (_m *MockProjectClientService) CreateProject(params *project.CreateProjectParams, authInfo runtime.ClientAuthInfoWriter) (*project.CreateProjectCreated, error) {
	ret := _m.Called(params, authInfo)

	var r0 *project.CreateProjectCreated
	if rf, ok := ret.Get(0).(func(*project.CreateProjectParams, runtime.ClientAuthInfoWriter) *project.CreateProjectCreated); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*project.CreateProjectCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*project.CreateProjectParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteProject provides a mock function with given fields: params, authInfo
func (_m *MockProjectClientService) DeleteProject(params *project.DeleteProjectParams, authInfo runtime.ClientAuthInfoWriter) (*project.DeleteProjectOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *project.DeleteProjectOK
	if rf, ok := ret.Get(0).(func(*project.DeleteProjectParams, runtime.ClientAuthInfoWriter) *project.DeleteProjectOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*project.DeleteProjectOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*project.DeleteProjectParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLogs provides a mock function with given fields: params, authInfo
func (_m *MockProjectClientService) GetLogs(params *project.GetLogsParams, authInfo runtime.ClientAuthInfoWriter) (*project.GetLogsOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *project.GetLogsOK
	if rf, ok := ret.Get(0).(func(*project.GetLogsParams, runtime.ClientAuthInfoWriter) *project.GetLogsOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*project.GetLogsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*project.GetLogsParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProject provides a mock function with given fields: params, authInfo
func (_m *MockProjectClientService) GetProject(params *project.GetProjectParams, authInfo runtime.ClientAuthInfoWriter) (*project.GetProjectOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *project.GetProjectOK
	if rf, ok := ret.Get(0).(func(*project.GetProjectParams, runtime.ClientAuthInfoWriter) *project.GetProjectOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*project.GetProjectOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*project.GetProjectParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProjectDeletable provides a mock function with given fields: params, authInfo
func (_m *MockProjectClientService) GetProjectDeletable(params *project.GetProjectDeletableParams, authInfo runtime.ClientAuthInfoWriter) (*project.GetProjectDeletableOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *project.GetProjectDeletableOK
	if rf, ok := ret.Get(0).(func(*project.GetProjectDeletableParams, runtime.ClientAuthInfoWriter) *project.GetProjectDeletableOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*project.GetProjectDeletableOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*project.GetProjectDeletableParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProjectSummary provides a mock function with given fields: params, authInfo
func (_m *MockProjectClientService) GetProjectSummary(params *project.GetProjectSummaryParams, authInfo runtime.ClientAuthInfoWriter) (*project.GetProjectSummaryOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *project.GetProjectSummaryOK
	if rf, ok := ret.Get(0).(func(*project.GetProjectSummaryParams, runtime.ClientAuthInfoWriter) *project.GetProjectSummaryOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*project.GetProjectSummaryOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*project.GetProjectSummaryParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetScannerOfProject provides a mock function with given fields: params, authInfo
func (_m *MockProjectClientService) GetScannerOfProject(params *project.GetScannerOfProjectParams, authInfo runtime.ClientAuthInfoWriter) (*project.GetScannerOfProjectOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *project.GetScannerOfProjectOK
	if rf, ok := ret.Get(0).(func(*project.GetScannerOfProjectParams, runtime.ClientAuthInfoWriter) *project.GetScannerOfProjectOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*project.GetScannerOfProjectOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*project.GetScannerOfProjectParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HeadProject provides a mock function with given fields: params, authInfo
func (_m *MockProjectClientService) HeadProject(params *project.HeadProjectParams, authInfo runtime.ClientAuthInfoWriter) (*project.HeadProjectOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *project.HeadProjectOK
	if rf, ok := ret.Get(0).(func(*project.HeadProjectParams, runtime.ClientAuthInfoWriter) *project.HeadProjectOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*project.HeadProjectOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*project.HeadProjectParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProjects provides a mock function with given fields: params, authInfo
func (_m *MockProjectClientService) ListProjects(params *project.ListProjectsParams, authInfo runtime.ClientAuthInfoWriter) (*project.ListProjectsOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *project.ListProjectsOK
	if rf, ok := ret.Get(0).(func(*project.ListProjectsParams, runtime.ClientAuthInfoWriter) *project.ListProjectsOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*project.ListProjectsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*project.ListProjectsParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListScannerCandidatesOfProject provides a mock function with given fields: params, authInfo
func (_m *MockProjectClientService) ListScannerCandidatesOfProject(params *project.ListScannerCandidatesOfProjectParams, authInfo runtime.ClientAuthInfoWriter) (*project.ListScannerCandidatesOfProjectOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *project.ListScannerCandidatesOfProjectOK
	if rf, ok := ret.Get(0).(func(*project.ListScannerCandidatesOfProjectParams, runtime.ClientAuthInfoWriter) *project.ListScannerCandidatesOfProjectOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*project.ListScannerCandidatesOfProjectOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*project.ListScannerCandidatesOfProjectParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetScannerOfProject provides a mock function with given fields: params, authInfo
func (_m *MockProjectClientService) SetScannerOfProject(params *project.SetScannerOfProjectParams, authInfo runtime.ClientAuthInfoWriter) (*project.SetScannerOfProjectOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *project.SetScannerOfProjectOK
	if rf, ok := ret.Get(0).(func(*project.SetScannerOfProjectParams, runtime.ClientAuthInfoWriter) *project.SetScannerOfProjectOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*project.SetScannerOfProjectOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*project.SetScannerOfProjectParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockProjectClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// UpdateProject provides a mock function with given fields: params, authInfo
func (_m *MockProjectClientService) UpdateProject(params *project.UpdateProjectParams, authInfo runtime.ClientAuthInfoWriter) (*project.UpdateProjectOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *project.UpdateProjectOK
	if rf, ok := ret.Get(0).(func(*project.UpdateProjectParams, runtime.ClientAuthInfoWriter) *project.UpdateProjectOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*project.UpdateProjectOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*project.UpdateProjectParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
