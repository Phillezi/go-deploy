// Code generated by go-swagger; DO NOT EDIT.

package scan_all

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

// NewGetLatestScheduledScanAllMetricsParams creates a new GetLatestScheduledScanAllMetricsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLatestScheduledScanAllMetricsParams() *GetLatestScheduledScanAllMetricsParams {
	return &GetLatestScheduledScanAllMetricsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLatestScheduledScanAllMetricsParamsWithTimeout creates a new GetLatestScheduledScanAllMetricsParams object
// with the ability to set a timeout on a request.
func NewGetLatestScheduledScanAllMetricsParamsWithTimeout(timeout time.Duration) *GetLatestScheduledScanAllMetricsParams {
	return &GetLatestScheduledScanAllMetricsParams{
		timeout: timeout,
	}
}

// NewGetLatestScheduledScanAllMetricsParamsWithContext creates a new GetLatestScheduledScanAllMetricsParams object
// with the ability to set a context for a request.
func NewGetLatestScheduledScanAllMetricsParamsWithContext(ctx context.Context) *GetLatestScheduledScanAllMetricsParams {
	return &GetLatestScheduledScanAllMetricsParams{
		Context: ctx,
	}
}

// NewGetLatestScheduledScanAllMetricsParamsWithHTTPClient creates a new GetLatestScheduledScanAllMetricsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLatestScheduledScanAllMetricsParamsWithHTTPClient(client *http.Client) *GetLatestScheduledScanAllMetricsParams {
	return &GetLatestScheduledScanAllMetricsParams{
		HTTPClient: client,
	}
}

/*
GetLatestScheduledScanAllMetricsParams contains all the parameters to send to the API endpoint

	for the get latest scheduled scan all metrics operation.

	Typically these are written to a http.Request.
*/
type GetLatestScheduledScanAllMetricsParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get latest scheduled scan all metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLatestScheduledScanAllMetricsParams) WithDefaults() *GetLatestScheduledScanAllMetricsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get latest scheduled scan all metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLatestScheduledScanAllMetricsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get latest scheduled scan all metrics params
func (o *GetLatestScheduledScanAllMetricsParams) WithTimeout(timeout time.Duration) *GetLatestScheduledScanAllMetricsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get latest scheduled scan all metrics params
func (o *GetLatestScheduledScanAllMetricsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get latest scheduled scan all metrics params
func (o *GetLatestScheduledScanAllMetricsParams) WithContext(ctx context.Context) *GetLatestScheduledScanAllMetricsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get latest scheduled scan all metrics params
func (o *GetLatestScheduledScanAllMetricsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get latest scheduled scan all metrics params
func (o *GetLatestScheduledScanAllMetricsParams) WithHTTPClient(client *http.Client) *GetLatestScheduledScanAllMetricsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get latest scheduled scan all metrics params
func (o *GetLatestScheduledScanAllMetricsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get latest scheduled scan all metrics params
func (o *GetLatestScheduledScanAllMetricsParams) WithXRequestID(xRequestID *string) *GetLatestScheduledScanAllMetricsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get latest scheduled scan all metrics params
func (o *GetLatestScheduledScanAllMetricsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLatestScheduledScanAllMetricsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
