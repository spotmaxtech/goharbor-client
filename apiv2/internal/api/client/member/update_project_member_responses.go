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

// UpdateProjectMemberReader is a Reader for the UpdateProjectMember structure.
type UpdateProjectMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProjectMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProjectMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateProjectMemberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateProjectMemberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateProjectMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateProjectMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateProjectMemberInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateProjectMemberOK creates a UpdateProjectMemberOK with default headers values
func NewUpdateProjectMemberOK() *UpdateProjectMemberOK {
	return &UpdateProjectMemberOK{}
}

/*UpdateProjectMemberOK handles this case with default header values.

Success
*/
type UpdateProjectMemberOK struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *UpdateProjectMemberOK) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/members/{mid}][%d] updateProjectMemberOK ", 200)
}

func (o *UpdateProjectMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewUpdateProjectMemberBadRequest creates a UpdateProjectMemberBadRequest with default headers values
func NewUpdateProjectMemberBadRequest() *UpdateProjectMemberBadRequest {
	return &UpdateProjectMemberBadRequest{}
}

/*UpdateProjectMemberBadRequest handles this case with default header values.

Bad request
*/
type UpdateProjectMemberBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateProjectMemberBadRequest) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/members/{mid}][%d] updateProjectMemberBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateProjectMemberBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateProjectMemberBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectMemberUnauthorized creates a UpdateProjectMemberUnauthorized with default headers values
func NewUpdateProjectMemberUnauthorized() *UpdateProjectMemberUnauthorized {
	return &UpdateProjectMemberUnauthorized{}
}

/*UpdateProjectMemberUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateProjectMemberUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateProjectMemberUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/members/{mid}][%d] updateProjectMemberUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateProjectMemberUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateProjectMemberUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectMemberForbidden creates a UpdateProjectMemberForbidden with default headers values
func NewUpdateProjectMemberForbidden() *UpdateProjectMemberForbidden {
	return &UpdateProjectMemberForbidden{}
}

/*UpdateProjectMemberForbidden handles this case with default header values.

Forbidden
*/
type UpdateProjectMemberForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateProjectMemberForbidden) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/members/{mid}][%d] updateProjectMemberForbidden  %+v", 403, o.Payload)
}

func (o *UpdateProjectMemberForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateProjectMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectMemberNotFound creates a UpdateProjectMemberNotFound with default headers values
func NewUpdateProjectMemberNotFound() *UpdateProjectMemberNotFound {
	return &UpdateProjectMemberNotFound{}
}

/*UpdateProjectMemberNotFound handles this case with default header values.

Not found
*/
type UpdateProjectMemberNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateProjectMemberNotFound) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/members/{mid}][%d] updateProjectMemberNotFound  %+v", 404, o.Payload)
}

func (o *UpdateProjectMemberNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateProjectMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectMemberInternalServerError creates a UpdateProjectMemberInternalServerError with default headers values
func NewUpdateProjectMemberInternalServerError() *UpdateProjectMemberInternalServerError {
	return &UpdateProjectMemberInternalServerError{}
}

/*UpdateProjectMemberInternalServerError handles this case with default header values.

Internal server error
*/
type UpdateProjectMemberInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateProjectMemberInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/members/{mid}][%d] updateProjectMemberInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateProjectMemberInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateProjectMemberInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
