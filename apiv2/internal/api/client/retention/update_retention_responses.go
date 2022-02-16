// Code generated by go-swagger; DO NOT EDIT.

package retention

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spotmaxtech/goharbor-client/v5/apiv2/model"
)

// UpdateRetentionReader is a Reader for the UpdateRetention structure.
type UpdateRetentionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRetentionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRetentionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRetentionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRetentionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateRetentionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRetentionOK creates a UpdateRetentionOK with default headers values
func NewUpdateRetentionOK() *UpdateRetentionOK {
	return &UpdateRetentionOK{}
}

/*UpdateRetentionOK handles this case with default header values.

Update Retention Policy successfully.
*/
type UpdateRetentionOK struct {
}

func (o *UpdateRetentionOK) Error() string {
	return fmt.Sprintf("[PUT /retentions/{id}][%d] updateRetentionOK ", 200)
}

func (o *UpdateRetentionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRetentionUnauthorized creates a UpdateRetentionUnauthorized with default headers values
func NewUpdateRetentionUnauthorized() *UpdateRetentionUnauthorized {
	return &UpdateRetentionUnauthorized{}
}

/*UpdateRetentionUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateRetentionUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateRetentionUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /retentions/{id}][%d] updateRetentionUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateRetentionUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateRetentionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRetentionForbidden creates a UpdateRetentionForbidden with default headers values
func NewUpdateRetentionForbidden() *UpdateRetentionForbidden {
	return &UpdateRetentionForbidden{}
}

/*UpdateRetentionForbidden handles this case with default header values.

Forbidden
*/
type UpdateRetentionForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateRetentionForbidden) Error() string {
	return fmt.Sprintf("[PUT /retentions/{id}][%d] updateRetentionForbidden  %+v", 403, o.Payload)
}

func (o *UpdateRetentionForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateRetentionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRetentionInternalServerError creates a UpdateRetentionInternalServerError with default headers values
func NewUpdateRetentionInternalServerError() *UpdateRetentionInternalServerError {
	return &UpdateRetentionInternalServerError{}
}

/*UpdateRetentionInternalServerError handles this case with default header values.

Internal server error
*/
type UpdateRetentionInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateRetentionInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /retentions/{id}][%d] updateRetentionInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateRetentionInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateRetentionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
