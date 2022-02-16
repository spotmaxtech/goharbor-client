// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spotmaxtech/goharbor-client/v5/apiv2/model"
)

// UpdateUserPasswordReader is a Reader for the UpdateUserPassword structure.
type UpdateUserPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateUserPasswordBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateUserPasswordUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateUserPasswordForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateUserPasswordInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateUserPasswordOK creates a UpdateUserPasswordOK with default headers values
func NewUpdateUserPasswordOK() *UpdateUserPasswordOK {
	return &UpdateUserPasswordOK{}
}

/*UpdateUserPasswordOK handles this case with default header values.

Success
*/
type UpdateUserPasswordOK struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *UpdateUserPasswordOK) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/password][%d] updateUserPasswordOK ", 200)
}

func (o *UpdateUserPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewUpdateUserPasswordBadRequest creates a UpdateUserPasswordBadRequest with default headers values
func NewUpdateUserPasswordBadRequest() *UpdateUserPasswordBadRequest {
	return &UpdateUserPasswordBadRequest{}
}

/*UpdateUserPasswordBadRequest handles this case with default header values.

Invalid user ID; Password does not meet requirement
*/
type UpdateUserPasswordBadRequest struct {
}

func (o *UpdateUserPasswordBadRequest) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/password][%d] updateUserPasswordBadRequest ", 400)
}

func (o *UpdateUserPasswordBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserPasswordUnauthorized creates a UpdateUserPasswordUnauthorized with default headers values
func NewUpdateUserPasswordUnauthorized() *UpdateUserPasswordUnauthorized {
	return &UpdateUserPasswordUnauthorized{}
}

/*UpdateUserPasswordUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateUserPasswordUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateUserPasswordUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/password][%d] updateUserPasswordUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateUserPasswordUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateUserPasswordUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserPasswordForbidden creates a UpdateUserPasswordForbidden with default headers values
func NewUpdateUserPasswordForbidden() *UpdateUserPasswordForbidden {
	return &UpdateUserPasswordForbidden{}
}

/*UpdateUserPasswordForbidden handles this case with default header values.

The caller does not have permission to update the password of the user with given ID, or the old password in request body is not correct.
*/
type UpdateUserPasswordForbidden struct {
}

func (o *UpdateUserPasswordForbidden) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/password][%d] updateUserPasswordForbidden ", 403)
}

func (o *UpdateUserPasswordForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserPasswordInternalServerError creates a UpdateUserPasswordInternalServerError with default headers values
func NewUpdateUserPasswordInternalServerError() *UpdateUserPasswordInternalServerError {
	return &UpdateUserPasswordInternalServerError{}
}

/*UpdateUserPasswordInternalServerError handles this case with default header values.

Internal server error
*/
type UpdateUserPasswordInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateUserPasswordInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/password][%d] updateUserPasswordInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateUserPasswordInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateUserPasswordInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
