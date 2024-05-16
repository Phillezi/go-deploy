// Code generated by go-swagger; DO NOT EDIT.

package permissions

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

// NewGetPermissionsParams creates a new GetPermissionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPermissionsParams() *GetPermissionsParams {
	return &GetPermissionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPermissionsParamsWithTimeout creates a new GetPermissionsParams object
// with the ability to set a timeout on a request.
func NewGetPermissionsParamsWithTimeout(timeout time.Duration) *GetPermissionsParams {
	return &GetPermissionsParams{
		timeout: timeout,
	}
}

// NewGetPermissionsParamsWithContext creates a new GetPermissionsParams object
// with the ability to set a context for a request.
func NewGetPermissionsParamsWithContext(ctx context.Context) *GetPermissionsParams {
	return &GetPermissionsParams{
		Context: ctx,
	}
}

// NewGetPermissionsParamsWithHTTPClient creates a new GetPermissionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPermissionsParamsWithHTTPClient(client *http.Client) *GetPermissionsParams {
	return &GetPermissionsParams{
		HTTPClient: client,
	}
}

/*
GetPermissionsParams contains all the parameters to send to the API endpoint

	for the get permissions operation.

	Typically these are written to a http.Request.
*/
type GetPermissionsParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get permissions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPermissionsParams) WithDefaults() *GetPermissionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get permissions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPermissionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get permissions params
func (o *GetPermissionsParams) WithTimeout(timeout time.Duration) *GetPermissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get permissions params
func (o *GetPermissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get permissions params
func (o *GetPermissionsParams) WithContext(ctx context.Context) *GetPermissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get permissions params
func (o *GetPermissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get permissions params
func (o *GetPermissionsParams) WithHTTPClient(client *http.Client) *GetPermissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get permissions params
func (o *GetPermissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get permissions params
func (o *GetPermissionsParams) WithXRequestID(xRequestID *string) *GetPermissionsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get permissions params
func (o *GetPermissionsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPermissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}