// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spotmaxtech/goharbor-client/v5/apiv2/model"
)

// CreateProjectMemberReader is a Reader for the CreateProjectMember structure.
type CreateProjectMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProjectMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateProjectMemberCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateProjectMemberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateProjectMemberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateProjectMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateProjectMemberConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateProjectMemberInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateProjectMemberCreated creates a CreateProjectMemberCreated with default headers values
func NewCreateProjectMemberCreated() *CreateProjectMemberCreated {
	return &CreateProjectMemberCreated{}
}

/*CreateProjectMemberCreated handles this case with default header values.

Project member created successfully.
*/
type CreateProjectMemberCreated struct {
	/*The URL of the created resource
	 */
	Location string
}

func (o *CreateProjectMemberCreated) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/members][%d] createProjectMemberCreated ", 201)
}

func (o *CreateProjectMemberCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	return nil
}

// NewCreateProjectMemberBadRequest creates a CreateProjectMemberBadRequest with default headers values
func NewCreateProjectMemberBadRequest() *CreateProjectMemberBadRequest {
	return &CreateProjectMemberBadRequest{}
}

/*CreateProjectMemberBadRequest handles this case with default header values.

Bad request
*/
type CreateProjectMemberBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateProjectMemberBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/members][%d] createProjectMemberBadRequest  %+v", 400, o.Payload)
}

func (o *CreateProjectMemberBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateProjectMemberBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectMemberUnauthorized creates a CreateProjectMemberUnauthorized with default headers values
func NewCreateProjectMemberUnauthorized() *CreateProjectMemberUnauthorized {
	return &CreateProjectMemberUnauthorized{}
}

/*CreateProjectMemberUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateProjectMemberUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateProjectMemberUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/members][%d] createProjectMemberUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateProjectMemberUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateProjectMemberUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectMemberForbidden creates a CreateProjectMemberForbidden with default headers values
func NewCreateProjectMemberForbidden() *CreateProjectMemberForbidden {
	return &CreateProjectMemberForbidden{}
}

/*CreateProjectMemberForbidden handles this case with default header values.

Forbidden
*/
type CreateProjectMemberForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateProjectMemberForbidden) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/members][%d] createProjectMemberForbidden  %+v", 403, o.Payload)
}

func (o *CreateProjectMemberForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateProjectMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectMemberConflict creates a CreateProjectMemberConflict with default headers values
func NewCreateProjectMemberConflict() *CreateProjectMemberConflict {
	return &CreateProjectMemberConflict{}
}

/*CreateProjectMemberConflict handles this case with default header values.

Conflict
*/
type CreateProjectMemberConflict struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateProjectMemberConflict) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/members][%d] createProjectMemberConflict  %+v", 409, o.Payload)
}

func (o *CreateProjectMemberConflict) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateProjectMemberConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectMemberInternalServerError creates a CreateProjectMemberInternalServerError with default headers values
func NewCreateProjectMemberInternalServerError() *CreateProjectMemberInternalServerError {
	return &CreateProjectMemberInternalServerError{}
}

/*CreateProjectMemberInternalServerError handles this case with default header values.

Internal server error
*/
type CreateProjectMemberInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateProjectMemberInternalServerError) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/members][%d] createProjectMemberInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateProjectMemberInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateProjectMemberInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
