// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spotmaxtech/goharbor-client/v5/apiv2/model"
)

// CreateInstanceReader is a Reader for the CreateInstance structure.
type CreateInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateInstanceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateInstanceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateInstanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateInstanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateInstanceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateInstanceCreated creates a CreateInstanceCreated with default headers values
func NewCreateInstanceCreated() *CreateInstanceCreated {
	return &CreateInstanceCreated{}
}

/*CreateInstanceCreated handles this case with default header values.

Created
*/
type CreateInstanceCreated struct {
	/*The location of the resource
	 */
	Location string
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *CreateInstanceCreated) Error() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances][%d] createInstanceCreated ", 201)
}

func (o *CreateInstanceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewCreateInstanceBadRequest creates a CreateInstanceBadRequest with default headers values
func NewCreateInstanceBadRequest() *CreateInstanceBadRequest {
	return &CreateInstanceBadRequest{}
}

/*CreateInstanceBadRequest handles this case with default header values.

Bad request
*/
type CreateInstanceBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateInstanceBadRequest) Error() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances][%d] createInstanceBadRequest  %+v", 400, o.Payload)
}

func (o *CreateInstanceBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateInstanceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInstanceUnauthorized creates a CreateInstanceUnauthorized with default headers values
func NewCreateInstanceUnauthorized() *CreateInstanceUnauthorized {
	return &CreateInstanceUnauthorized{}
}

/*CreateInstanceUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateInstanceUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances][%d] createInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateInstanceUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInstanceForbidden creates a CreateInstanceForbidden with default headers values
func NewCreateInstanceForbidden() *CreateInstanceForbidden {
	return &CreateInstanceForbidden{}
}

/*CreateInstanceForbidden handles this case with default header values.

Forbidden
*/
type CreateInstanceForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateInstanceForbidden) Error() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances][%d] createInstanceForbidden  %+v", 403, o.Payload)
}

func (o *CreateInstanceForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateInstanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInstanceNotFound creates a CreateInstanceNotFound with default headers values
func NewCreateInstanceNotFound() *CreateInstanceNotFound {
	return &CreateInstanceNotFound{}
}

/*CreateInstanceNotFound handles this case with default header values.

Not found
*/
type CreateInstanceNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateInstanceNotFound) Error() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances][%d] createInstanceNotFound  %+v", 404, o.Payload)
}

func (o *CreateInstanceNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateInstanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInstanceConflict creates a CreateInstanceConflict with default headers values
func NewCreateInstanceConflict() *CreateInstanceConflict {
	return &CreateInstanceConflict{}
}

/*CreateInstanceConflict handles this case with default header values.

Conflict
*/
type CreateInstanceConflict struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateInstanceConflict) Error() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances][%d] createInstanceConflict  %+v", 409, o.Payload)
}

func (o *CreateInstanceConflict) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateInstanceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInstanceInternalServerError creates a CreateInstanceInternalServerError with default headers values
func NewCreateInstanceInternalServerError() *CreateInstanceInternalServerError {
	return &CreateInstanceInternalServerError{}
}

/*CreateInstanceInternalServerError handles this case with default header values.

Internal server error
*/
type CreateInstanceInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances][%d] createInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateInstanceInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
