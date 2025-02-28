// Code generated by go-swagger; DO NOT EDIT.

package ldap

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

// NewPingLdapParams creates a new PingLdapParams object
// with the default values initialized.
func NewPingLdapParams() *PingLdapParams {
	var ()
	return &PingLdapParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPingLdapParamsWithTimeout creates a new PingLdapParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPingLdapParamsWithTimeout(timeout time.Duration) *PingLdapParams {
	var ()
	return &PingLdapParams{

		timeout: timeout,
	}
}

// NewPingLdapParamsWithContext creates a new PingLdapParams object
// with the default values initialized, and the ability to set a context for a request
func NewPingLdapParamsWithContext(ctx context.Context) *PingLdapParams {
	var ()
	return &PingLdapParams{

		Context: ctx,
	}
}

// NewPingLdapParamsWithHTTPClient creates a new PingLdapParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPingLdapParamsWithHTTPClient(client *http.Client) *PingLdapParams {
	var ()
	return &PingLdapParams{
		HTTPClient: client,
	}
}

/*PingLdapParams contains all the parameters to send to the API endpoint
for the ping ldap operation typically these are written to a http.Request
*/
type PingLdapParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*Ldapconf
	  ldap configuration. support input ldap service configuration. If it is a empty request, will load current configuration from the system.

	*/
	Ldapconf *model.LdapConf

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ping ldap params
func (o *PingLdapParams) WithTimeout(timeout time.Duration) *PingLdapParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ping ldap params
func (o *PingLdapParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ping ldap params
func (o *PingLdapParams) WithContext(ctx context.Context) *PingLdapParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ping ldap params
func (o *PingLdapParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ping ldap params
func (o *PingLdapParams) WithHTTPClient(client *http.Client) *PingLdapParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ping ldap params
func (o *PingLdapParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the ping ldap params
func (o *PingLdapParams) WithXRequestID(xRequestID *string) *PingLdapParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the ping ldap params
func (o *PingLdapParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithLdapconf adds the ldapconf to the ping ldap params
func (o *PingLdapParams) WithLdapconf(ldapconf *model.LdapConf) *PingLdapParams {
	o.SetLdapconf(ldapconf)
	return o
}

// SetLdapconf adds the ldapconf to the ping ldap params
func (o *PingLdapParams) SetLdapconf(ldapconf *model.LdapConf) {
	o.Ldapconf = ldapconf
}

// WriteToRequest writes these params to a swagger request
func (o *PingLdapParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Ldapconf != nil {
		if err := r.SetBodyParam(o.Ldapconf); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
