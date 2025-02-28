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

// DeleteRetentionReader is a Reader for the DeleteRetention structure.
type DeleteRetentionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRetentionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRetentionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteRetentionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRetentionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRetentionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRetentionOK creates a DeleteRetentionOK with default headers values
func NewDeleteRetentionOK() *DeleteRetentionOK {
	return &DeleteRetentionOK{}
}

/*DeleteRetentionOK handles this case with default header values.

Update Retention Policy successfully.
*/
type DeleteRetentionOK struct {
}

func (o *DeleteRetentionOK) Error() string {
	return fmt.Sprintf("[DELETE /retentions/{id}][%d] deleteRetentionOK ", 200)
}

func (o *DeleteRetentionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRetentionUnauthorized creates a DeleteRetentionUnauthorized with default headers values
func NewDeleteRetentionUnauthorized() *DeleteRetentionUnauthorized {
	return &DeleteRetentionUnauthorized{}
}

/*DeleteRetentionUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteRetentionUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteRetentionUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /retentions/{id}][%d] deleteRetentionUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRetentionUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteRetentionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRetentionForbidden creates a DeleteRetentionForbidden with default headers values
func NewDeleteRetentionForbidden() *DeleteRetentionForbidden {
	return &DeleteRetentionForbidden{}
}

/*DeleteRetentionForbidden handles this case with default header values.

Forbidden
*/
type DeleteRetentionForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteRetentionForbidden) Error() string {
	return fmt.Sprintf("[DELETE /retentions/{id}][%d] deleteRetentionForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRetentionForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteRetentionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRetentionInternalServerError creates a DeleteRetentionInternalServerError with default headers values
func NewDeleteRetentionInternalServerError() *DeleteRetentionInternalServerError {
	return &DeleteRetentionInternalServerError{}
}

/*DeleteRetentionInternalServerError handles this case with default header values.

Internal server error
*/
type DeleteRetentionInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteRetentionInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /retentions/{id}][%d] deleteRetentionInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRetentionInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteRetentionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
