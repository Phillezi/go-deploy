// Code generated by go-swagger; DO NOT EDIT.

package oidc

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
)

// NewPingOIDCParams creates a new PingOIDCParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPingOIDCParams() *PingOIDCParams {
	return &PingOIDCParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPingOIDCParamsWithTimeout creates a new PingOIDCParams object
// with the ability to set a timeout on a request.
func NewPingOIDCParamsWithTimeout(timeout time.Duration) *PingOIDCParams {
	return &PingOIDCParams{
		timeout: timeout,
	}
}

// NewPingOIDCParamsWithContext creates a new PingOIDCParams object
// with the ability to set a context for a request.
func NewPingOIDCParamsWithContext(ctx context.Context) *PingOIDCParams {
	return &PingOIDCParams{
		Context: ctx,
	}
}

// NewPingOIDCParamsWithHTTPClient creates a new PingOIDCParams object
// with the ability to set a custom HTTPClient for a request.
func NewPingOIDCParamsWithHTTPClient(client *http.Client) *PingOIDCParams {
	return &PingOIDCParams{
		HTTPClient: client,
	}
}

/*
PingOIDCParams contains all the parameters to send to the API endpoint

	for the ping OIDC operation.

	Typically these are written to a http.Request.
*/
type PingOIDCParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* Endpoint.

	   Request body for OIDC endpoint to be tested.
	*/
	Endpoint PingOIDCBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ping OIDC params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PingOIDCParams) WithDefaults() *PingOIDCParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ping OIDC params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PingOIDCParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ping OIDC params
func (o *PingOIDCParams) WithTimeout(timeout time.Duration) *PingOIDCParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ping OIDC params
func (o *PingOIDCParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ping OIDC params
func (o *PingOIDCParams) WithContext(ctx context.Context) *PingOIDCParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ping OIDC params
func (o *PingOIDCParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ping OIDC params
func (o *PingOIDCParams) WithHTTPClient(client *http.Client) *PingOIDCParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ping OIDC params
func (o *PingOIDCParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the ping OIDC params
func (o *PingOIDCParams) WithXRequestID(xRequestID *string) *PingOIDCParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the ping OIDC params
func (o *PingOIDCParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithEndpoint adds the endpoint to the ping OIDC params
func (o *PingOIDCParams) WithEndpoint(endpoint PingOIDCBody) *PingOIDCParams {
	o.SetEndpoint(endpoint)
	return o
}

// SetEndpoint adds the endpoint to the ping OIDC params
func (o *PingOIDCParams) SetEndpoint(endpoint PingOIDCBody) {
	o.Endpoint = endpoint
}

// WriteToRequest writes these params to a swagger request
func (o *PingOIDCParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.Endpoint); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
