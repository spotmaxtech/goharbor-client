// Code generated by go-swagger; DO NOT EDIT.

package robotv1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spotmaxtech/goharbor-client/v5/apiv2/model"
)

// CreateRobotV1Reader is a Reader for the CreateRobotV1 structure.
type CreateRobotV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRobotV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRobotV1Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRobotV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateRobotV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRobotV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateRobotV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateRobotV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRobotV1Created creates a CreateRobotV1Created with default headers values
func NewCreateRobotV1Created() *CreateRobotV1Created {
	return &CreateRobotV1Created{}
}

/*CreateRobotV1Created handles this case with default header values.

Created
*/
type CreateRobotV1Created struct {
	/*The location of the resource
	 */
	Location string
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.RobotCreated
}

func (o *CreateRobotV1Created) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/robots][%d] createRobotV1Created  %+v", 201, o.Payload)
}

func (o *CreateRobotV1Created) GetPayload() *model.RobotCreated {
	return o.Payload
}

func (o *CreateRobotV1Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateRobotV1BadRequest creates a CreateRobotV1BadRequest with default headers values
func NewCreateRobotV1BadRequest() *CreateRobotV1BadRequest {
	return &CreateRobotV1BadRequest{}
}

/*CreateRobotV1BadRequest handles this case with default header values.

Bad request
*/
type CreateRobotV1BadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateRobotV1BadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/robots][%d] createRobotV1BadRequest  %+v", 400, o.Payload)
}

func (o *CreateRobotV1BadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateRobotV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRobotV1Unauthorized creates a CreateRobotV1Unauthorized with default headers values
func NewCreateRobotV1Unauthorized() *CreateRobotV1Unauthorized {
	return &CreateRobotV1Unauthorized{}
}

/*CreateRobotV1Unauthorized handles this case with default header values.

Unauthorized
*/
type CreateRobotV1Unauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateRobotV1Unauthorized) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/robots][%d] createRobotV1Unauthorized  %+v", 401, o.Payload)
}

func (o *CreateRobotV1Unauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateRobotV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRobotV1Forbidden creates a CreateRobotV1Forbidden with default headers values
func NewCreateRobotV1Forbidden() *CreateRobotV1Forbidden {
	return &CreateRobotV1Forbidden{}
}

/*CreateRobotV1Forbidden handles this case with default header values.

Forbidden
*/
type CreateRobotV1Forbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateRobotV1Forbidden) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/robots][%d] createRobotV1Forbidden  %+v", 403, o.Payload)
}

func (o *CreateRobotV1Forbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateRobotV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRobotV1NotFound creates a CreateRobotV1NotFound with default headers values
func NewCreateRobotV1NotFound() *CreateRobotV1NotFound {
	return &CreateRobotV1NotFound{}
}

/*CreateRobotV1NotFound handles this case with default header values.

Not found
*/
type CreateRobotV1NotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateRobotV1NotFound) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/robots][%d] createRobotV1NotFound  %+v", 404, o.Payload)
}

func (o *CreateRobotV1NotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateRobotV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRobotV1InternalServerError creates a CreateRobotV1InternalServerError with default headers values
func NewCreateRobotV1InternalServerError() *CreateRobotV1InternalServerError {
	return &CreateRobotV1InternalServerError{}
}

/*CreateRobotV1InternalServerError handles this case with default header values.

Internal server error
*/
type CreateRobotV1InternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateRobotV1InternalServerError) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/robots][%d] createRobotV1InternalServerError  %+v", 500, o.Payload)
}

func (o *CreateRobotV1InternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateRobotV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
