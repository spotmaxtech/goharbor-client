// Code generated by go-swagger; DO NOT EDIT.

package auditlog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/spotmaxtech/goharbor-client/v5/apiv2/model"
)

// ListAuditLogsReader is a Reader for the ListAuditLogs structure.
type ListAuditLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAuditLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAuditLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListAuditLogsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListAuditLogsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListAuditLogsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListAuditLogsOK creates a ListAuditLogsOK with default headers values
func NewListAuditLogsOK() *ListAuditLogsOK {
	return &ListAuditLogsOK{}
}

/*ListAuditLogsOK handles this case with default header values.

Success
*/
type ListAuditLogsOK struct {
	/*Link refers to the previous page and next page
	 */
	Link string
	/*The total count of auditlogs
	 */
	XTotalCount int64

	Payload []*model.AuditLog
}

func (o *ListAuditLogsOK) Error() string {
	return fmt.Sprintf("[GET /audit-logs][%d] listAuditLogsOK  %+v", 200, o.Payload)
}

func (o *ListAuditLogsOK) GetPayload() []*model.AuditLog {
	return o.Payload
}

func (o *ListAuditLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Link
	o.Link = response.GetHeader("Link")

	// response header X-Total-Count
	xTotalCount, err := swag.ConvertInt64(response.GetHeader("X-Total-Count"))
	if err != nil {
		return errors.InvalidType("X-Total-Count", "header", "int64", response.GetHeader("X-Total-Count"))
	}
	o.XTotalCount = xTotalCount

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAuditLogsBadRequest creates a ListAuditLogsBadRequest with default headers values
func NewListAuditLogsBadRequest() *ListAuditLogsBadRequest {
	return &ListAuditLogsBadRequest{}
}

/*ListAuditLogsBadRequest handles this case with default header values.

Bad request
*/
type ListAuditLogsBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListAuditLogsBadRequest) Error() string {
	return fmt.Sprintf("[GET /audit-logs][%d] listAuditLogsBadRequest  %+v", 400, o.Payload)
}

func (o *ListAuditLogsBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListAuditLogsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAuditLogsUnauthorized creates a ListAuditLogsUnauthorized with default headers values
func NewListAuditLogsUnauthorized() *ListAuditLogsUnauthorized {
	return &ListAuditLogsUnauthorized{}
}

/*ListAuditLogsUnauthorized handles this case with default header values.

Unauthorized
*/
type ListAuditLogsUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListAuditLogsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /audit-logs][%d] listAuditLogsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListAuditLogsUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListAuditLogsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAuditLogsInternalServerError creates a ListAuditLogsInternalServerError with default headers values
func NewListAuditLogsInternalServerError() *ListAuditLogsInternalServerError {
	return &ListAuditLogsInternalServerError{}
}

/*ListAuditLogsInternalServerError handles this case with default header values.

Internal server error
*/
type ListAuditLogsInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListAuditLogsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /audit-logs][%d] listAuditLogsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListAuditLogsInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListAuditLogsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
