// Code generated by go-swagger; DO NOT EDIT.

package robot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spotmaxtech/goharbor-client/v5/apiv2/model"
)

// CreateRobotReader is a Reader for the CreateRobot structure.
type CreateRobotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRobotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRobotCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRobotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateRobotUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRobotForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateRobotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateRobotInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRobotCreated creates a CreateRobotCreated with default headers values
func NewCreateRobotCreated() *CreateRobotCreated {
	return &CreateRobotCreated{}
}

/*CreateRobotCreated handles this case with default header values.

Created
*/
type CreateRobotCreated struct {
	/*The location of the resource
	 */
	Location string
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.RobotCreated
}

func (o *CreateRobotCreated) Error() string {
	return fmt.Sprintf("[POST /robots][%d] createRobotCreated  %+v", 201, o.Payload)
}

func (o *CreateRobotCreated) GetPayload() *model.RobotCreated {
	return o.Payload
}

func (o *CreateRobotCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.RobotCreated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRobotBadRequest creates a CreateRobotBadRequest with default headers values
func NewCreateRobotBadRequest() *CreateRobotBadRequest {
	return &CreateRobotBadRequest{}
}

/*CreateRobotBadRequest handles this case with default header values.

Bad request
*/
type CreateRobotBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateRobotBadRequest) Error() string {
	return fmt.Sprintf("[POST /robots][%d] createRobotBadRequest  %+v", 400, o.Payload)
}

func (o *CreateRobotBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateRobotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRobotUnauthorized creates a CreateRobotUnauthorized with default headers values
func NewCreateRobotUnauthorized() *CreateRobotUnauthorized {
	return &CreateRobotUnauthorized{}
}

/*CreateRobotUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateRobotUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateRobotUnauthorized) Error() string {
	return fmt.Sprintf("[POST /robots][%d] createRobotUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateRobotUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateRobotUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRobotForbidden creates a CreateRobotForbidden with default headers values
func NewCreateRobotForbidden() *CreateRobotForbidden {
	return &CreateRobotForbidden{}
}

/*CreateRobotForbidden handles this case with default header values.

Forbidden
*/
type CreateRobotForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateRobotForbidden) Error() string {
	return fmt.Sprintf("[POST /robots][%d] createRobotForbidden  %+v", 403, o.Payload)
}

func (o *CreateRobotForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateRobotForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRobotNotFound creates a CreateRobotNotFound with default headers values
func NewCreateRobotNotFound() *CreateRobotNotFound {
	return &CreateRobotNotFound{}
}

/*CreateRobotNotFound handles this case with default header values.

Not found
*/
type CreateRobotNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateRobotNotFound) Error() string {
	return fmt.Sprintf("[POST /robots][%d] createRobotNotFound  %+v", 404, o.Payload)
}

func (o *CreateRobotNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateRobotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRobotInternalServerError creates a CreateRobotInternalServerError with default headers values
func NewCreateRobotInternalServerError() *CreateRobotInternalServerError {
	return &CreateRobotInternalServerError{}
}

/*CreateRobotInternalServerError handles this case with default header values.

Internal server error
*/
type CreateRobotInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateRobotInternalServerError) Error() string {
	return fmt.Sprintf("[POST /robots][%d] createRobotInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateRobotInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateRobotInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
