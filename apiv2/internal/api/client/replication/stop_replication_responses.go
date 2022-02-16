// Code generated by go-swagger; DO NOT EDIT.

package replication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spotmaxtech/goharbor-client/v5/apiv2/model"
)

// StopReplicationReader is a Reader for the StopReplication structure.
type StopReplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopReplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopReplicationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewStopReplicationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStopReplicationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopReplicationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStopReplicationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStopReplicationOK creates a StopReplicationOK with default headers values
func NewStopReplicationOK() *StopReplicationOK {
	return &StopReplicationOK{}
}

/*StopReplicationOK handles this case with default header values.

Success
*/
type StopReplicationOK struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *StopReplicationOK) Error() string {
	return fmt.Sprintf("[PUT /replication/executions/{id}][%d] stopReplicationOK ", 200)
}

func (o *StopReplicationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewStopReplicationUnauthorized creates a StopReplicationUnauthorized with default headers values
func NewStopReplicationUnauthorized() *StopReplicationUnauthorized {
	return &StopReplicationUnauthorized{}
}

/*StopReplicationUnauthorized handles this case with default header values.

Unauthorized
*/
type StopReplicationUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *StopReplicationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /replication/executions/{id}][%d] stopReplicationUnauthorized  %+v", 401, o.Payload)
}

func (o *StopReplicationUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *StopReplicationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopReplicationForbidden creates a StopReplicationForbidden with default headers values
func NewStopReplicationForbidden() *StopReplicationForbidden {
	return &StopReplicationForbidden{}
}

/*StopReplicationForbidden handles this case with default header values.

Forbidden
*/
type StopReplicationForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *StopReplicationForbidden) Error() string {
	return fmt.Sprintf("[PUT /replication/executions/{id}][%d] stopReplicationForbidden  %+v", 403, o.Payload)
}

func (o *StopReplicationForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *StopReplicationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopReplicationNotFound creates a StopReplicationNotFound with default headers values
func NewStopReplicationNotFound() *StopReplicationNotFound {
	return &StopReplicationNotFound{}
}

/*StopReplicationNotFound handles this case with default header values.

Not found
*/
type StopReplicationNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *StopReplicationNotFound) Error() string {
	return fmt.Sprintf("[PUT /replication/executions/{id}][%d] stopReplicationNotFound  %+v", 404, o.Payload)
}

func (o *StopReplicationNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *StopReplicationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopReplicationInternalServerError creates a StopReplicationInternalServerError with default headers values
func NewStopReplicationInternalServerError() *StopReplicationInternalServerError {
	return &StopReplicationInternalServerError{}
}

/*StopReplicationInternalServerError handles this case with default header values.

Internal server error
*/
type StopReplicationInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *StopReplicationInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /replication/executions/{id}][%d] stopReplicationInternalServerError  %+v", 500, o.Payload)
}

func (o *StopReplicationInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *StopReplicationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
