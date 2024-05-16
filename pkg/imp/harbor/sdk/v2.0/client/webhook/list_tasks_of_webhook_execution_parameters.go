// Code generated by go-swagger; DO NOT EDIT.

package webhook

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
	"github.com/go-openapi/swag"
)

// NewListTasksOfWebhookExecutionParams creates a new ListTasksOfWebhookExecutionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListTasksOfWebhookExecutionParams() *ListTasksOfWebhookExecutionParams {
	return &ListTasksOfWebhookExecutionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListTasksOfWebhookExecutionParamsWithTimeout creates a new ListTasksOfWebhookExecutionParams object
// with the ability to set a timeout on a request.
func NewListTasksOfWebhookExecutionParamsWithTimeout(timeout time.Duration) *ListTasksOfWebhookExecutionParams {
	return &ListTasksOfWebhookExecutionParams{
		timeout: timeout,
	}
}

// NewListTasksOfWebhookExecutionParamsWithContext creates a new ListTasksOfWebhookExecutionParams object
// with the ability to set a context for a request.
func NewListTasksOfWebhookExecutionParamsWithContext(ctx context.Context) *ListTasksOfWebhookExecutionParams {
	return &ListTasksOfWebhookExecutionParams{
		Context: ctx,
	}
}

// NewListTasksOfWebhookExecutionParamsWithHTTPClient creates a new ListTasksOfWebhookExecutionParams object
// with the ability to set a custom HTTPClient for a request.
func NewListTasksOfWebhookExecutionParamsWithHTTPClient(client *http.Client) *ListTasksOfWebhookExecutionParams {
	return &ListTasksOfWebhookExecutionParams{
		HTTPClient: client,
	}
}

/*
ListTasksOfWebhookExecutionParams contains all the parameters to send to the API endpoint

	for the list tasks of webhook execution operation.

	Typically these are written to a http.Request.
*/
type ListTasksOfWebhookExecutionParams struct {

	/* XIsResourceName.

	   The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.
	*/
	XIsResourceName *bool

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* ExecutionID.

	   Execution ID
	*/
	ExecutionID int64

	/* Page.

	   The page number

	   Format: int64
	   Default: 1
	*/
	Page *int64

	/* PageSize.

	   The size of per page

	   Format: int64
	   Default: 10
	*/
	PageSize *int64

	/* ProjectNameOrID.

	   The name or id of the project
	*/
	ProjectNameOrID string

	/* Q.

	   Query string to query resources. Supported query patterns are "exact match(k=v)", "fuzzy match(k=~v)", "range(k=[min~max])", "list with union releationship(k={v1 v2 v3})" and "list with intersetion relationship(k=(v1 v2 v3))". The value of range and list can be string(enclosed by " or '), integer or time(in format "2020-04-09 02:36:00"). All of these query patterns should be put in the query string "q=xxx" and splitted by ",". e.g. q=k1=v1,k2=~v2,k3=[min~max]
	*/
	Q *string

	/* Sort.

	   Sort the resource list in ascending or descending order. e.g. sort by field1 in ascending order and field2 in descending order with "sort=field1,-field2"
	*/
	Sort *string

	/* WebhookPolicyID.

	   The ID of the webhook policy

	   Format: int64
	*/
	WebhookPolicyID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list tasks of webhook execution params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTasksOfWebhookExecutionParams) WithDefaults() *ListTasksOfWebhookExecutionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list tasks of webhook execution params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTasksOfWebhookExecutionParams) SetDefaults() {
	var (
		xIsResourceNameDefault = bool(false)

		pageDefault = int64(1)

		pageSizeDefault = int64(10)
	)

	val := ListTasksOfWebhookExecutionParams{
		XIsResourceName: &xIsResourceNameDefault,
		Page:            &pageDefault,
		PageSize:        &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list tasks of webhook execution params
func (o *ListTasksOfWebhookExecutionParams) WithTimeout(timeout time.Duration) *ListTasksOfWebhookExecutionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list tasks of webhook execution params
func (o *ListTasksOfWebhookExecutionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list tasks of webhook execution params
func (o *ListTasksOfWebhookExecutionParams) WithContext(ctx context.Context) *ListTasksOfWebhookExecutionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list tasks of webhook execution params
func (o *ListTasksOfWebhookExecutionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list tasks of webhook execution params
func (o *ListTasksOfWebhookExecutionParams) WithHTTPClient(client *http.Client) *ListTasksOfWebhookExecutionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list tasks of webhook execution params
func (o *ListTasksOfWebhookExecutionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXIsResourceName adds the xIsResourceName to the list tasks of webhook execution params
func (o *ListTasksOfWebhookExecutionParams) WithXIsResourceName(xIsResourceName *bool) *ListTasksOfWebhookExecutionParams {
	o.SetXIsResourceName(xIsResourceName)
	return o
}

// SetXIsResourceName adds the xIsResourceName to the list tasks of webhook execution params
func (o *ListTasksOfWebhookExecutionParams) SetXIsResourceName(xIsResourceName *bool) {
	o.XIsResourceName = xIsResourceName
}

// WithXRequestID adds the xRequestID to the list tasks of webhook execution params
func (o *ListTasksOfWebhookExecutionParams) WithXRequestID(xRequestID *string) *ListTasksOfWebhookExecutionParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the list tasks of webhook execution params
func (o *ListTasksOfWebhookExecutionParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithExecutionID adds the executionID to the list tasks of webhook execution params
func (o *ListTasksOfWebhookExecutionParams) WithExecutionID(executionID int64) *ListTasksOfWebhookExecutionParams {
	o.SetExecutionID(executionID)
	return o
}

// SetExecutionID adds the executionId to the list tasks of webhook execution params
func (o *ListTasksOfWebhookExecutionParams) SetExecutionID(executionID int64) {
	o.ExecutionID = executionID
}

// WithPage adds the page to the list tasks of webhook execution params
func (o *ListTasksOfWebhookExecutionParams) WithPage(page *int64) *ListTasksOfWebhookExecutionParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list tasks of webhook execution params
func (o *ListTasksOfWebhookExecutionParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the list tasks of webhook execution params
func (o *ListTasksOfWebhookExecutionParams) WithPageSize(pageSize *int64) *ListTasksOfWebhookExecutionParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list tasks of webhook execution params
func (o *ListTasksOfWebhookExecutionParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithProjectNameOrID adds the projectNameOrID to the list tasks of webhook execution params
func (o *ListTasksOfWebhookExecutionParams) WithProjectNameOrID(projectNameOrID string) *ListTasksOfWebhookExecutionParams {
	o.SetProjectNameOrID(projectNameOrID)
	return o
}

// SetProjectNameOrID adds the projectNameOrId to the list tasks of webhook execution params
func (o *ListTasksOfWebhookExecutionParams) SetProjectNameOrID(projectNameOrID string) {
	o.ProjectNameOrID = projectNameOrID
}

// WithQ adds the q to the list tasks of webhook execution params
func (o *ListTasksOfWebhookExecutionParams) WithQ(q *string) *ListTasksOfWebhookExecutionParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the list tasks of webhook execution params
func (o *ListTasksOfWebhookExecutionParams) SetQ(q *string) {
	o.Q = q
}

// WithSort adds the sort to the list tasks of webhook execution params
func (o *ListTasksOfWebhookExecutionParams) WithSort(sort *string) *ListTasksOfWebhookExecutionParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the list tasks of webhook execution params
func (o *ListTasksOfWebhookExecutionParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithWebhookPolicyID adds the webhookPolicyID to the list tasks of webhook execution params
func (o *ListTasksOfWebhookExecutionParams) WithWebhookPolicyID(webhookPolicyID int64) *ListTasksOfWebhookExecutionParams {
	o.SetWebhookPolicyID(webhookPolicyID)
	return o
}

// SetWebhookPolicyID adds the webhookPolicyId to the list tasks of webhook execution params
func (o *ListTasksOfWebhookExecutionParams) SetWebhookPolicyID(webhookPolicyID int64) {
	o.WebhookPolicyID = webhookPolicyID
}

// WriteToRequest writes these params to a swagger request
func (o *ListTasksOfWebhookExecutionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XIsResourceName != nil {

		// header param X-Is-Resource-Name
		if err := r.SetHeaderParam("X-Is-Resource-Name", swag.FormatBool(*o.XIsResourceName)); err != nil {
			return err
		}
	}

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	// path param execution_id
	if err := r.SetPathParam("execution_id", swag.FormatInt64(o.ExecutionID)); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}
	}

	// path param project_name_or_id
	if err := r.SetPathParam("project_name_or_id", o.ProjectNameOrID); err != nil {
		return err
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	// path param webhook_policy_id
	if err := r.SetPathParam("webhook_policy_id", swag.FormatInt64(o.WebhookPolicyID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
