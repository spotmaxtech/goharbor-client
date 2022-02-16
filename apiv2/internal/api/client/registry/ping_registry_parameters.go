// Code generated by go-swagger; DO NOT EDIT.

package registry

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

// NewPingRegistryParams creates a new PingRegistryParams object
// with the default values initialized.
func NewPingRegistryParams() *PingRegistryParams {
	var ()
	return &PingRegistryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPingRegistryParamsWithTimeout creates a new PingRegistryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPingRegistryParamsWithTimeout(timeout time.Duration) *PingRegistryParams {
	var ()
	return &PingRegistryParams{

		timeout: timeout,
	}
}

// NewPingRegistryParamsWithContext creates a new PingRegistryParams object
// with the default values initialized, and the ability to set a context for a request
func NewPingRegistryParamsWithContext(ctx context.Context) *PingRegistryParams {
	var ()
	return &PingRegistryParams{

		Context: ctx,
	}
}

// NewPingRegistryParamsWithHTTPClient creates a new PingRegistryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPingRegistryParamsWithHTTPClient(client *http.Client) *PingRegistryParams {
	var ()
	return &PingRegistryParams{
		HTTPClient: client,
	}
}

/*PingRegistryParams contains all the parameters to send to the API endpoint
for the ping registry operation typically these are written to a http.Request
*/
type PingRegistryParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*Registry
	  The registry

	*/
	Registry *model.RegistryPing

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ping registry params
func (o *PingRegistryParams) WithTimeout(timeout time.Duration) *PingRegistryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ping registry params
func (o *PingRegistryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ping registry params
func (o *PingRegistryParams) WithContext(ctx context.Context) *PingRegistryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ping registry params
func (o *PingRegistryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ping registry params
func (o *PingRegistryParams) WithHTTPClient(client *http.Client) *PingRegistryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ping registry params
func (o *PingRegistryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the ping registry params
func (o *PingRegistryParams) WithXRequestID(xRequestID *string) *PingRegistryParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the ping registry params
func (o *PingRegistryParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithRegistry adds the registry to the ping registry params
func (o *PingRegistryParams) WithRegistry(registry *model.RegistryPing) *PingRegistryParams {
	o.SetRegistry(registry)
	return o
}

// SetRegistry adds the registry to the ping registry params
func (o *PingRegistryParams) SetRegistry(registry *model.RegistryPing) {
	o.Registry = registry
}

// WriteToRequest writes these params to a swagger request
func (o *PingRegistryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Registry != nil {
		if err := r.SetBodyParam(o.Registry); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
