// Code generated by go-swagger; DO NOT EDIT.

package replication

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

// ListReplicationPoliciesReader is a Reader for the ListReplicationPolicies structure.
type ListReplicationPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListReplicationPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListReplicationPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListReplicationPoliciesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListReplicationPoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListReplicationPoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListReplicationPoliciesOK creates a ListReplicationPoliciesOK with default headers values
func NewListReplicationPoliciesOK() *ListReplicationPoliciesOK {
	return &ListReplicationPoliciesOK{}
}

/*ListReplicationPoliciesOK handles this case with default header values.

Success
*/
type ListReplicationPoliciesOK struct {
	/*Link refers to the previous page and next page
	 */
	Link string
	/*The total count of the resources
	 */
	XTotalCount int64

	Payload []*model.ReplicationPolicy
}

func (o *ListReplicationPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /replication/policies][%d] listReplicationPoliciesOK  %+v", 200, o.Payload)
}

func (o *ListReplicationPoliciesOK) GetPayload() []*model.ReplicationPolicy {
	return o.Payload
}

func (o *ListReplicationPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListReplicationPoliciesUnauthorized creates a ListReplicationPoliciesUnauthorized with default headers values
func NewListReplicationPoliciesUnauthorized() *ListReplicationPoliciesUnauthorized {
	return &ListReplicationPoliciesUnauthorized{}
}

/*ListReplicationPoliciesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListReplicationPoliciesUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListReplicationPoliciesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /replication/policies][%d] listReplicationPoliciesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListReplicationPoliciesUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListReplicationPoliciesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListReplicationPoliciesForbidden creates a ListReplicationPoliciesForbidden with default headers values
func NewListReplicationPoliciesForbidden() *ListReplicationPoliciesForbidden {
	return &ListReplicationPoliciesForbidden{}
}

/*ListReplicationPoliciesForbidden handles this case with default header values.

Forbidden
*/
type ListReplicationPoliciesForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListReplicationPoliciesForbidden) Error() string {
	return fmt.Sprintf("[GET /replication/policies][%d] listReplicationPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *ListReplicationPoliciesForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListReplicationPoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListReplicationPoliciesInternalServerError creates a ListReplicationPoliciesInternalServerError with default headers values
func NewListReplicationPoliciesInternalServerError() *ListReplicationPoliciesInternalServerError {
	return &ListReplicationPoliciesInternalServerError{}
}

/*ListReplicationPoliciesInternalServerError handles this case with default header values.

Internal server error
*/
type ListReplicationPoliciesInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListReplicationPoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /replication/policies][%d] listReplicationPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListReplicationPoliciesInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListReplicationPoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
