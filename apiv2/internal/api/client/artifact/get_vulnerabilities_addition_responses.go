// Code generated by go-swagger; DO NOT EDIT.

package artifact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spotmaxtech/goharbor-client/v5/apiv2/model"
)

// GetVulnerabilitiesAdditionReader is a Reader for the GetVulnerabilitiesAddition structure.
type GetVulnerabilitiesAdditionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVulnerabilitiesAdditionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVulnerabilitiesAdditionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVulnerabilitiesAdditionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetVulnerabilitiesAdditionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetVulnerabilitiesAdditionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVulnerabilitiesAdditionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVulnerabilitiesAdditionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVulnerabilitiesAdditionOK creates a GetVulnerabilitiesAdditionOK with default headers values
func NewGetVulnerabilitiesAdditionOK() *GetVulnerabilitiesAdditionOK {
	return &GetVulnerabilitiesAdditionOK{}
}

/*GetVulnerabilitiesAdditionOK handles this case with default header values.

Success
*/
type GetVulnerabilitiesAdditionOK struct {
	/*The content type of the vulnerabilities addition
	 */
	ContentType string

	Payload string
}

func (o *GetVulnerabilitiesAdditionOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/additions/vulnerabilities][%d] getVulnerabilitiesAdditionOK  %+v", 200, o.Payload)
}

func (o *GetVulnerabilitiesAdditionOK) GetPayload() string {
	return o.Payload
}

func (o *GetVulnerabilitiesAdditionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Type
	o.ContentType = response.GetHeader("Content-Type")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVulnerabilitiesAdditionBadRequest creates a GetVulnerabilitiesAdditionBadRequest with default headers values
func NewGetVulnerabilitiesAdditionBadRequest() *GetVulnerabilitiesAdditionBadRequest {
	return &GetVulnerabilitiesAdditionBadRequest{}
}

/*GetVulnerabilitiesAdditionBadRequest handles this case with default header values.

Bad request
*/
type GetVulnerabilitiesAdditionBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetVulnerabilitiesAdditionBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/additions/vulnerabilities][%d] getVulnerabilitiesAdditionBadRequest  %+v", 400, o.Payload)
}

func (o *GetVulnerabilitiesAdditionBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetVulnerabilitiesAdditionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVulnerabilitiesAdditionUnauthorized creates a GetVulnerabilitiesAdditionUnauthorized with default headers values
func NewGetVulnerabilitiesAdditionUnauthorized() *GetVulnerabilitiesAdditionUnauthorized {
	return &GetVulnerabilitiesAdditionUnauthorized{}
}

/*GetVulnerabilitiesAdditionUnauthorized handles this case with default header values.

Unauthorized
*/
type GetVulnerabilitiesAdditionUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetVulnerabilitiesAdditionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/additions/vulnerabilities][%d] getVulnerabilitiesAdditionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetVulnerabilitiesAdditionUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetVulnerabilitiesAdditionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVulnerabilitiesAdditionForbidden creates a GetVulnerabilitiesAdditionForbidden with default headers values
func NewGetVulnerabilitiesAdditionForbidden() *GetVulnerabilitiesAdditionForbidden {
	return &GetVulnerabilitiesAdditionForbidden{}
}

/*GetVulnerabilitiesAdditionForbidden handles this case with default header values.

Forbidden
*/
type GetVulnerabilitiesAdditionForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetVulnerabilitiesAdditionForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/additions/vulnerabilities][%d] getVulnerabilitiesAdditionForbidden  %+v", 403, o.Payload)
}

func (o *GetVulnerabilitiesAdditionForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetVulnerabilitiesAdditionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVulnerabilitiesAdditionNotFound creates a GetVulnerabilitiesAdditionNotFound with default headers values
func NewGetVulnerabilitiesAdditionNotFound() *GetVulnerabilitiesAdditionNotFound {
	return &GetVulnerabilitiesAdditionNotFound{}
}

/*GetVulnerabilitiesAdditionNotFound handles this case with default header values.

Not found
*/
type GetVulnerabilitiesAdditionNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetVulnerabilitiesAdditionNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/additions/vulnerabilities][%d] getVulnerabilitiesAdditionNotFound  %+v", 404, o.Payload)
}

func (o *GetVulnerabilitiesAdditionNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetVulnerabilitiesAdditionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVulnerabilitiesAdditionInternalServerError creates a GetVulnerabilitiesAdditionInternalServerError with default headers values
func NewGetVulnerabilitiesAdditionInternalServerError() *GetVulnerabilitiesAdditionInternalServerError {
	return &GetVulnerabilitiesAdditionInternalServerError{}
}

/*GetVulnerabilitiesAdditionInternalServerError handles this case with default header values.

Internal server error
*/
type GetVulnerabilitiesAdditionInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetVulnerabilitiesAdditionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/additions/vulnerabilities][%d] getVulnerabilitiesAdditionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVulnerabilitiesAdditionInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetVulnerabilitiesAdditionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
