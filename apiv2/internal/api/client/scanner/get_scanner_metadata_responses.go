// Code generated by go-swagger; DO NOT EDIT.

package scanner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spotmaxtech/goharbor-client/v5/apiv2/model"
)

// GetScannerMetadataReader is a Reader for the GetScannerMetadata structure.
type GetScannerMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScannerMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScannerMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetScannerMetadataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScannerMetadataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScannerMetadataInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScannerMetadataOK creates a GetScannerMetadataOK with default headers values
func NewGetScannerMetadataOK() *GetScannerMetadataOK {
	return &GetScannerMetadataOK{}
}

/*GetScannerMetadataOK handles this case with default header values.

The metadata of the specified scanner adapter
*/
type GetScannerMetadataOK struct {
	Payload *model.ScannerAdapterMetadata
}

func (o *GetScannerMetadataOK) Error() string {
	return fmt.Sprintf("[GET /scanners/{registration_id}/metadata][%d] getScannerMetadataOK  %+v", 200, o.Payload)
}

func (o *GetScannerMetadataOK) GetPayload() *model.ScannerAdapterMetadata {
	return o.Payload
}

func (o *GetScannerMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ScannerAdapterMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScannerMetadataUnauthorized creates a GetScannerMetadataUnauthorized with default headers values
func NewGetScannerMetadataUnauthorized() *GetScannerMetadataUnauthorized {
	return &GetScannerMetadataUnauthorized{}
}

/*GetScannerMetadataUnauthorized handles this case with default header values.

Unauthorized
*/
type GetScannerMetadataUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetScannerMetadataUnauthorized) Error() string {
	return fmt.Sprintf("[GET /scanners/{registration_id}/metadata][%d] getScannerMetadataUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScannerMetadataUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetScannerMetadataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScannerMetadataForbidden creates a GetScannerMetadataForbidden with default headers values
func NewGetScannerMetadataForbidden() *GetScannerMetadataForbidden {
	return &GetScannerMetadataForbidden{}
}

/*GetScannerMetadataForbidden handles this case with default header values.

Forbidden
*/
type GetScannerMetadataForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetScannerMetadataForbidden) Error() string {
	return fmt.Sprintf("[GET /scanners/{registration_id}/metadata][%d] getScannerMetadataForbidden  %+v", 403, o.Payload)
}

func (o *GetScannerMetadataForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetScannerMetadataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScannerMetadataInternalServerError creates a GetScannerMetadataInternalServerError with default headers values
func NewGetScannerMetadataInternalServerError() *GetScannerMetadataInternalServerError {
	return &GetScannerMetadataInternalServerError{}
}

/*GetScannerMetadataInternalServerError handles this case with default header values.

Internal server error
*/
type GetScannerMetadataInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetScannerMetadataInternalServerError) Error() string {
	return fmt.Sprintf("[GET /scanners/{registration_id}/metadata][%d] getScannerMetadataInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScannerMetadataInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetScannerMetadataInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
