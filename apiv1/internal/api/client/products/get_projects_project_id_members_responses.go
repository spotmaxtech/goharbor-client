// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spotmaxtech/goharbor-client/v5/apiv1/model"
)

// GetProjectsProjectIDMembersReader is a Reader for the GetProjectsProjectIDMembers structure.
type GetProjectsProjectIDMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectsProjectIDMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectsProjectIDMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProjectsProjectIDMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProjectsProjectIDMembersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProjectsProjectIDMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProjectsProjectIDMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProjectsProjectIDMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProjectsProjectIDMembersOK creates a GetProjectsProjectIDMembersOK with default headers values
func NewGetProjectsProjectIDMembersOK() *GetProjectsProjectIDMembersOK {
	return &GetProjectsProjectIDMembersOK{}
}

/*GetProjectsProjectIDMembersOK handles this case with default header values.

Get project members successfully.
*/
type GetProjectsProjectIDMembersOK struct {
	Payload []*model.ProjectMemberEntity
}

func (o *GetProjectsProjectIDMembersOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/members][%d] getProjectsProjectIdMembersOK  %+v", 200, o.Payload)
}

func (o *GetProjectsProjectIDMembersOK) GetPayload() []*model.ProjectMemberEntity {
	return o.Payload
}

func (o *GetProjectsProjectIDMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsProjectIDMembersBadRequest creates a GetProjectsProjectIDMembersBadRequest with default headers values
func NewGetProjectsProjectIDMembersBadRequest() *GetProjectsProjectIDMembersBadRequest {
	return &GetProjectsProjectIDMembersBadRequest{}
}

/*GetProjectsProjectIDMembersBadRequest handles this case with default header values.

The project id is invalid.
*/
type GetProjectsProjectIDMembersBadRequest struct {
}

func (o *GetProjectsProjectIDMembersBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/members][%d] getProjectsProjectIdMembersBadRequest ", 400)
}

func (o *GetProjectsProjectIDMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDMembersUnauthorized creates a GetProjectsProjectIDMembersUnauthorized with default headers values
func NewGetProjectsProjectIDMembersUnauthorized() *GetProjectsProjectIDMembersUnauthorized {
	return &GetProjectsProjectIDMembersUnauthorized{}
}

/*GetProjectsProjectIDMembersUnauthorized handles this case with default header values.

User need to log in first.
*/
type GetProjectsProjectIDMembersUnauthorized struct {
}

func (o *GetProjectsProjectIDMembersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/members][%d] getProjectsProjectIdMembersUnauthorized ", 401)
}

func (o *GetProjectsProjectIDMembersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDMembersForbidden creates a GetProjectsProjectIDMembersForbidden with default headers values
func NewGetProjectsProjectIDMembersForbidden() *GetProjectsProjectIDMembersForbidden {
	return &GetProjectsProjectIDMembersForbidden{}
}

/*GetProjectsProjectIDMembersForbidden handles this case with default header values.

User in session does not have permission to the project.
*/
type GetProjectsProjectIDMembersForbidden struct {
}

func (o *GetProjectsProjectIDMembersForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/members][%d] getProjectsProjectIdMembersForbidden ", 403)
}

func (o *GetProjectsProjectIDMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDMembersNotFound creates a GetProjectsProjectIDMembersNotFound with default headers values
func NewGetProjectsProjectIDMembersNotFound() *GetProjectsProjectIDMembersNotFound {
	return &GetProjectsProjectIDMembersNotFound{}
}

/*GetProjectsProjectIDMembersNotFound handles this case with default header values.

Project ID does not exist.
*/
type GetProjectsProjectIDMembersNotFound struct {
}

func (o *GetProjectsProjectIDMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/members][%d] getProjectsProjectIdMembersNotFound ", 404)
}

func (o *GetProjectsProjectIDMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDMembersInternalServerError creates a GetProjectsProjectIDMembersInternalServerError with default headers values
func NewGetProjectsProjectIDMembersInternalServerError() *GetProjectsProjectIDMembersInternalServerError {
	return &GetProjectsProjectIDMembersInternalServerError{}
}

/*GetProjectsProjectIDMembersInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type GetProjectsProjectIDMembersInternalServerError struct {
}

func (o *GetProjectsProjectIDMembersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/members][%d] getProjectsProjectIdMembersInternalServerError ", 500)
}

func (o *GetProjectsProjectIDMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
