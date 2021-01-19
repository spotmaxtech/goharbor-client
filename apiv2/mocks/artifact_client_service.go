// Code generated by mockery v2.5.1. DO NOT EDIT.

package mocks

import (
	artifact "github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client/artifact"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// MockArtifactClientService is an autogenerated mock type for the ClientService type
type MockArtifactClientService struct {
	mock.Mock
}

// AddLabel provides a mock function with given fields: params, authInfo
func (_m *MockArtifactClientService) AddLabel(params *artifact.AddLabelParams, authInfo runtime.ClientAuthInfoWriter) (*artifact.AddLabelOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *artifact.AddLabelOK
	if rf, ok := ret.Get(0).(func(*artifact.AddLabelParams, runtime.ClientAuthInfoWriter) *artifact.AddLabelOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifact.AddLabelOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*artifact.AddLabelParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CopyArtifact provides a mock function with given fields: params, authInfo
func (_m *MockArtifactClientService) CopyArtifact(params *artifact.CopyArtifactParams, authInfo runtime.ClientAuthInfoWriter) (*artifact.CopyArtifactCreated, error) {
	ret := _m.Called(params, authInfo)

	var r0 *artifact.CopyArtifactCreated
	if rf, ok := ret.Get(0).(func(*artifact.CopyArtifactParams, runtime.ClientAuthInfoWriter) *artifact.CopyArtifactCreated); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifact.CopyArtifactCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*artifact.CopyArtifactParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTag provides a mock function with given fields: params, authInfo
func (_m *MockArtifactClientService) CreateTag(params *artifact.CreateTagParams, authInfo runtime.ClientAuthInfoWriter) (*artifact.CreateTagCreated, error) {
	ret := _m.Called(params, authInfo)

	var r0 *artifact.CreateTagCreated
	if rf, ok := ret.Get(0).(func(*artifact.CreateTagParams, runtime.ClientAuthInfoWriter) *artifact.CreateTagCreated); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifact.CreateTagCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*artifact.CreateTagParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteArtifact provides a mock function with given fields: params, authInfo
func (_m *MockArtifactClientService) DeleteArtifact(params *artifact.DeleteArtifactParams, authInfo runtime.ClientAuthInfoWriter) (*artifact.DeleteArtifactOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *artifact.DeleteArtifactOK
	if rf, ok := ret.Get(0).(func(*artifact.DeleteArtifactParams, runtime.ClientAuthInfoWriter) *artifact.DeleteArtifactOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifact.DeleteArtifactOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*artifact.DeleteArtifactParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTag provides a mock function with given fields: params, authInfo
func (_m *MockArtifactClientService) DeleteTag(params *artifact.DeleteTagParams, authInfo runtime.ClientAuthInfoWriter) (*artifact.DeleteTagOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *artifact.DeleteTagOK
	if rf, ok := ret.Get(0).(func(*artifact.DeleteTagParams, runtime.ClientAuthInfoWriter) *artifact.DeleteTagOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifact.DeleteTagOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*artifact.DeleteTagParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddition provides a mock function with given fields: params, authInfo
func (_m *MockArtifactClientService) GetAddition(params *artifact.GetAdditionParams, authInfo runtime.ClientAuthInfoWriter) (*artifact.GetAdditionOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *artifact.GetAdditionOK
	if rf, ok := ret.Get(0).(func(*artifact.GetAdditionParams, runtime.ClientAuthInfoWriter) *artifact.GetAdditionOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifact.GetAdditionOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*artifact.GetAdditionParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetArtifact provides a mock function with given fields: params, authInfo
func (_m *MockArtifactClientService) GetArtifact(params *artifact.GetArtifactParams, authInfo runtime.ClientAuthInfoWriter) (*artifact.GetArtifactOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *artifact.GetArtifactOK
	if rf, ok := ret.Get(0).(func(*artifact.GetArtifactParams, runtime.ClientAuthInfoWriter) *artifact.GetArtifactOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifact.GetArtifactOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*artifact.GetArtifactParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListArtifacts provides a mock function with given fields: params, authInfo
func (_m *MockArtifactClientService) ListArtifacts(params *artifact.ListArtifactsParams, authInfo runtime.ClientAuthInfoWriter) (*artifact.ListArtifactsOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *artifact.ListArtifactsOK
	if rf, ok := ret.Get(0).(func(*artifact.ListArtifactsParams, runtime.ClientAuthInfoWriter) *artifact.ListArtifactsOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifact.ListArtifactsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*artifact.ListArtifactsParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTags provides a mock function with given fields: params, authInfo
func (_m *MockArtifactClientService) ListTags(params *artifact.ListTagsParams, authInfo runtime.ClientAuthInfoWriter) (*artifact.ListTagsOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *artifact.ListTagsOK
	if rf, ok := ret.Get(0).(func(*artifact.ListTagsParams, runtime.ClientAuthInfoWriter) *artifact.ListTagsOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifact.ListTagsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*artifact.ListTagsParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveLabel provides a mock function with given fields: params, authInfo
func (_m *MockArtifactClientService) RemoveLabel(params *artifact.RemoveLabelParams, authInfo runtime.ClientAuthInfoWriter) (*artifact.RemoveLabelOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *artifact.RemoveLabelOK
	if rf, ok := ret.Get(0).(func(*artifact.RemoveLabelParams, runtime.ClientAuthInfoWriter) *artifact.RemoveLabelOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifact.RemoveLabelOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*artifact.RemoveLabelParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockArtifactClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}
