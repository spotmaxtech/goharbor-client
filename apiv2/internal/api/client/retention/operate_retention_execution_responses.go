// Code generated by go-swagger; DO NOT EDIT.

package retention

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/spotmaxtech/goharbor-client/v5/apiv2/model"
)

// OperateRetentionExecutionReader is a Reader for the OperateRetentionExecution structure.
type OperateRetentionExecutionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OperateRetentionExecutionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOperateRetentionExecutionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewOperateRetentionExecutionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOperateRetentionExecutionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOperateRetentionExecutionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOperateRetentionExecutionOK creates a OperateRetentionExecutionOK with default headers values
func NewOperateRetentionExecutionOK() *OperateRetentionExecutionOK {
	return &OperateRetentionExecutionOK{}
}

/*OperateRetentionExecutionOK handles this case with default header values.

Stop a Retention job successfully.
*/
type OperateRetentionExecutionOK struct {
}

func (o *OperateRetentionExecutionOK) Error() string {
	return fmt.Sprintf("[PATCH /retentions/{id}/executions/{eid}][%d] operateRetentionExecutionOK ", 200)
}

func (o *OperateRetentionExecutionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOperateRetentionExecutionUnauthorized creates a OperateRetentionExecutionUnauthorized with default headers values
func NewOperateRetentionExecutionUnauthorized() *OperateRetentionExecutionUnauthorized {
	return &OperateRetentionExecutionUnauthorized{}
}

/*OperateRetentionExecutionUnauthorized handles this case with default header values.

Unauthorized
*/
type OperateRetentionExecutionUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *OperateRetentionExecutionUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /retentions/{id}/executions/{eid}][%d] operateRetentionExecutionUnauthorized  %+v", 401, o.Payload)
}

func (o *OperateRetentionExecutionUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *OperateRetentionExecutionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOperateRetentionExecutionForbidden creates a OperateRetentionExecutionForbidden with default headers values
func NewOperateRetentionExecutionForbidden() *OperateRetentionExecutionForbidden {
	return &OperateRetentionExecutionForbidden{}
}

/*OperateRetentionExecutionForbidden handles this case with default header values.

Forbidden
*/
type OperateRetentionExecutionForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *OperateRetentionExecutionForbidden) Error() string {
	return fmt.Sprintf("[PATCH /retentions/{id}/executions/{eid}][%d] operateRetentionExecutionForbidden  %+v", 403, o.Payload)
}

func (o *OperateRetentionExecutionForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *OperateRetentionExecutionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOperateRetentionExecutionInternalServerError creates a OperateRetentionExecutionInternalServerError with default headers values
func NewOperateRetentionExecutionInternalServerError() *OperateRetentionExecutionInternalServerError {
	return &OperateRetentionExecutionInternalServerError{}
}

/*OperateRetentionExecutionInternalServerError handles this case with default header values.

Internal server error
*/
type OperateRetentionExecutionInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *OperateRetentionExecutionInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /retentions/{id}/executions/{eid}][%d] operateRetentionExecutionInternalServerError  %+v", 500, o.Payload)
}

func (o *OperateRetentionExecutionInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *OperateRetentionExecutionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*OperateRetentionExecutionBody operate retention execution body
swagger:model OperateRetentionExecutionBody
*/
type OperateRetentionExecutionBody struct {

	// action
	Action string `json:"action,omitempty"`
}

// Validate validates this operate retention execution body
func (o *OperateRetentionExecutionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *OperateRetentionExecutionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *OperateRetentionExecutionBody) UnmarshalBinary(b []byte) error {
	var res OperateRetentionExecutionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
