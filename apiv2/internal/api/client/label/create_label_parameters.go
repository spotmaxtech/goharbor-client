// Code generated by go-swagger; DO NOT EDIT.

package label

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/spotmaxtech/goharbor-client/v5/apiv2/model"
)

// NewCreateLabelParams creates a new CreateLabelParams object
// with the default values initialized.
func NewCreateLabelParams() *CreateLabelParams {
	var ()
	return &CreateLabelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateLabelParamsWithTimeout creates a new CreateLabelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateLabelParamsWithTimeout(timeout time.Duration) *CreateLabelParams {
	var ()
	return &CreateLabelParams{

		timeout: timeout,
	}
}

// NewCreateLabelParamsWithContext creates a new CreateLabelParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateLabelParamsWithContext(ctx context.Context) *CreateLabelParams {
	var ()
	return &CreateLabelParams{

		Context: ctx,
	}
}

// NewCreateLabelParamsWithHTTPClient creates a new CreateLabelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateLabelParamsWithHTTPClient(client *http.Client) *CreateLabelParams {
	var ()
	return &CreateLabelParams{
		HTTPClient: client,
	}
}

/*CreateLabelParams contains all the parameters to send to the API endpoint
for the create label operation typically these are written to a http.Request
*/
type CreateLabelParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*Label
	  The json object of label.

	*/
	Label *model.Label

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create label params
func (o *CreateLabelParams) WithTimeout(timeout time.Duration) *CreateLabelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create label params
func (o *CreateLabelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create label params
func (o *CreateLabelParams) WithContext(ctx context.Context) *CreateLabelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create label params
func (o *CreateLabelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create label params
func (o *CreateLabelParams) WithHTTPClient(client *http.Client) *CreateLabelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create label params
func (o *CreateLabelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the create label params
func (o *CreateLabelParams) WithXRequestID(xRequestID *string) *CreateLabelParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the create label params
func (o *CreateLabelParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithLabel adds the label to the create label params
func (o *CreateLabelParams) WithLabel(label *model.Label) *CreateLabelParams {
	o.SetLabel(label)
	return o
}

// SetLabel adds the label to the create label params
func (o *CreateLabelParams) SetLabel(label *model.Label) {
	o.Label = label
}

// WriteToRequest writes these params to a swagger request
func (o *CreateLabelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}

	}

	if o.Label != nil {
		if err := r.SetBodyParam(o.Label); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
