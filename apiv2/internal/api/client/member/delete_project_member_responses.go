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

// DeleteProjectMemberReader is a Reader for the DeleteProjectMember structure.
type DeleteProjectMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProjectMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteProjectMemberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteProjectMemberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteProjectMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteProjectMemberInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteProjectMemberOK creates a DeleteProjectMemberOK with default headers values
func NewDeleteProjectMemberOK() *DeleteProjectMemberOK {
	return &DeleteProjectMemberOK{}
}

/*DeleteProjectMemberOK handles this case with default header values.

Success
*/
type DeleteProjectMemberOK struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *DeleteProjectMemberOK) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/members/{mid}][%d] deleteProjectMemberOK ", 200)
}

func (o *DeleteProjectMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewDeleteProjectMemberBadRequest creates a DeleteProjectMemberBadRequest with default headers values
func NewDeleteProjectMemberBadRequest() *DeleteProjectMemberBadRequest {
	return &DeleteProjectMemberBadRequest{}
}

/*DeleteProjectMemberBadRequest handles this case with default header values.

Bad request
*/
type DeleteProjectMemberBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteProjectMemberBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/members/{mid}][%d] deleteProjectMemberBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteProjectMemberBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteProjectMemberBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectMemberUnauthorized creates a DeleteProjectMemberUnauthorized with default headers values
func NewDeleteProjectMemberUnauthorized() *DeleteProjectMemberUnauthorized {
	return &DeleteProjectMemberUnauthorized{}
}

/*DeleteProjectMemberUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteProjectMemberUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteProjectMemberUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/members/{mid}][%d] deleteProjectMemberUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteProjectMemberUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteProjectMemberUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectMemberForbidden creates a DeleteProjectMemberForbidden with default headers values
func NewDeleteProjectMemberForbidden() *DeleteProjectMemberForbidden {
	return &DeleteProjectMemberForbidden{}
}

/*DeleteProjectMemberForbidden handles this case with default header values.

Forbidden
*/
type DeleteProjectMemberForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteProjectMemberForbidden) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/members/{mid}][%d] deleteProjectMemberForbidden  %+v", 403, o.Payload)
}

func (o *DeleteProjectMemberForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteProjectMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectMemberInternalServerError creates a DeleteProjectMemberInternalServerError with default headers values
func NewDeleteProjectMemberInternalServerError() *DeleteProjectMemberInternalServerError {
	return &DeleteProjectMemberInternalServerError{}
}

/*DeleteProjectMemberInternalServerError handles this case with default header values.

Internal server error
*/
type DeleteProjectMemberInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteProjectMemberInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/members/{mid}][%d] deleteProjectMemberInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteProjectMemberInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteProjectMemberInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
