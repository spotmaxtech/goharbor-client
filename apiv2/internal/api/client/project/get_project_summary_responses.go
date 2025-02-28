// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spotmaxtech/goharbor-client/v5/apiv2/model"
)

// GetProjectSummaryReader is a Reader for the GetProjectSummary structure.
type GetProjectSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProjectSummaryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProjectSummaryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProjectSummaryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProjectSummaryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProjectSummaryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProjectSummaryOK creates a GetProjectSummaryOK with default headers values
func NewGetProjectSummaryOK() *GetProjectSummaryOK {
	return &GetProjectSummaryOK{}
}

/*GetProjectSummaryOK handles this case with default header values.

Get summary of the project successfully.
*/
type GetProjectSummaryOK struct {
	Payload *model.ProjectSummary
}

func (o *GetProjectSummaryOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/summary][%d] getProjectSummaryOK  %+v", 200, o.Payload)
}

func (o *GetProjectSummaryOK) GetPayload() *model.ProjectSummary {
	return o.Payload
}

func (o *GetProjectSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ProjectSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectSummaryBadRequest creates a GetProjectSummaryBadRequest with default headers values
func NewGetProjectSummaryBadRequest() *GetProjectSummaryBadRequest {
	return &GetProjectSummaryBadRequest{}
}

/*GetProjectSummaryBadRequest handles this case with default header values.

Bad request
*/
type GetProjectSummaryBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetProjectSummaryBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/summary][%d] getProjectSummaryBadRequest  %+v", 400, o.Payload)
}

func (o *GetProjectSummaryBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetProjectSummaryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectSummaryUnauthorized creates a GetProjectSummaryUnauthorized with default headers values
func NewGetProjectSummaryUnauthorized() *GetProjectSummaryUnauthorized {
	return &GetProjectSummaryUnauthorized{}
}

/*GetProjectSummaryUnauthorized handles this case with default header values.

Unauthorized
*/
type GetProjectSummaryUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetProjectSummaryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/summary][%d] getProjectSummaryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetProjectSummaryUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetProjectSummaryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectSummaryForbidden creates a GetProjectSummaryForbidden with default headers values
func NewGetProjectSummaryForbidden() *GetProjectSummaryForbidden {
	return &GetProjectSummaryForbidden{}
}

/*GetProjectSummaryForbidden handles this case with default header values.

Forbidden
*/
type GetProjectSummaryForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetProjectSummaryForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/summary][%d] getProjectSummaryForbidden  %+v", 403, o.Payload)
}

func (o *GetProjectSummaryForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetProjectSummaryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectSummaryNotFound creates a GetProjectSummaryNotFound with default headers values
func NewGetProjectSummaryNotFound() *GetProjectSummaryNotFound {
	return &GetProjectSummaryNotFound{}
}

/*GetProjectSummaryNotFound handles this case with default header values.

Not found
*/
type GetProjectSummaryNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetProjectSummaryNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/summary][%d] getProjectSummaryNotFound  %+v", 404, o.Payload)
}

func (o *GetProjectSummaryNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetProjectSummaryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectSummaryInternalServerError creates a GetProjectSummaryInternalServerError with default headers values
func NewGetProjectSummaryInternalServerError() *GetProjectSummaryInternalServerError {
	return &GetProjectSummaryInternalServerError{}
}

/*GetProjectSummaryInternalServerError handles this case with default header values.

Internal server error
*/
type GetProjectSummaryInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetProjectSummaryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/summary][%d] getProjectSummaryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetProjectSummaryInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetProjectSummaryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
