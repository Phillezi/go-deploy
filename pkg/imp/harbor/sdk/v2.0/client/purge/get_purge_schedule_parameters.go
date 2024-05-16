// Code generated by go-swagger; DO NOT EDIT.

package purge

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

// NewGetPurgeScheduleParams creates a new GetPurgeScheduleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPurgeScheduleParams() *GetPurgeScheduleParams {
	return &GetPurgeScheduleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPurgeScheduleParamsWithTimeout creates a new GetPurgeScheduleParams object
// with the ability to set a timeout on a request.
func NewGetPurgeScheduleParamsWithTimeout(timeout time.Duration) *GetPurgeScheduleParams {
	return &GetPurgeScheduleParams{
		timeout: timeout,
	}
}

// NewGetPurgeScheduleParamsWithContext creates a new GetPurgeScheduleParams object
// with the ability to set a context for a request.
func NewGetPurgeScheduleParamsWithContext(ctx context.Context) *GetPurgeScheduleParams {
	return &GetPurgeScheduleParams{
		Context: ctx,
	}
}

// NewGetPurgeScheduleParamsWithHTTPClient creates a new GetPurgeScheduleParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPurgeScheduleParamsWithHTTPClient(client *http.Client) *GetPurgeScheduleParams {
	return &GetPurgeScheduleParams{
		HTTPClient: client,
	}
}

/*
GetPurgeScheduleParams contains all the parameters to send to the API endpoint

	for the get purge schedule operation.

	Typically these are written to a http.Request.
*/
type GetPurgeScheduleParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get purge schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPurgeScheduleParams) WithDefaults() *GetPurgeScheduleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get purge schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPurgeScheduleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get purge schedule params
func (o *GetPurgeScheduleParams) WithTimeout(timeout time.Duration) *GetPurgeScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get purge schedule params
func (o *GetPurgeScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get purge schedule params
func (o *GetPurgeScheduleParams) WithContext(ctx context.Context) *GetPurgeScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get purge schedule params
func (o *GetPurgeScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get purge schedule params
func (o *GetPurgeScheduleParams) WithHTTPClient(client *http.Client) *GetPurgeScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get purge schedule params
func (o *GetPurgeScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get purge schedule params
func (o *GetPurgeScheduleParams) WithXRequestID(xRequestID *string) *GetPurgeScheduleParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get purge schedule params
func (o *GetPurgeScheduleParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPurgeScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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