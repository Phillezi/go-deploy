// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"go-deploy/pkg/imp/harbor/sdk/v2.0/models"
)

// GetLogsOfWebhookTaskReader is a Reader for the GetLogsOfWebhookTask structure.
type GetLogsOfWebhookTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogsOfWebhookTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogsOfWebhookTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLogsOfWebhookTaskBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLogsOfWebhookTaskUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLogsOfWebhookTaskForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLogsOfWebhookTaskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLogsOfWebhookTaskInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLogsOfWebhookTaskOK creates a GetLogsOfWebhookTaskOK with default headers values
func NewGetLogsOfWebhookTaskOK() *GetLogsOfWebhookTaskOK {
	return &GetLogsOfWebhookTaskOK{}
}

/*
GetLogsOfWebhookTaskOK describes a response with status code 200, with default header values.

Get log success
*/
type GetLogsOfWebhookTaskOK struct {

	/* Content type of response
	 */
	ContentType string

	Payload string
}

// IsSuccess returns true when this get logs of webhook task o k response has a 2xx status code
func (o *GetLogsOfWebhookTaskOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get logs of webhook task o k response has a 3xx status code
func (o *GetLogsOfWebhookTaskOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get logs of webhook task o k response has a 4xx status code
func (o *GetLogsOfWebhookTaskOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get logs of webhook task o k response has a 5xx status code
func (o *GetLogsOfWebhookTaskOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get logs of webhook task o k response a status code equal to that given
func (o *GetLogsOfWebhookTaskOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetLogsOfWebhookTaskOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks/{task_id}/log][%d] getLogsOfWebhookTaskOK  %+v", 200, o.Payload)
}

func (o *GetLogsOfWebhookTaskOK) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks/{task_id}/log][%d] getLogsOfWebhookTaskOK  %+v", 200, o.Payload)
}

func (o *GetLogsOfWebhookTaskOK) GetPayload() string {
	return o.Payload
}

func (o *GetLogsOfWebhookTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Content-Type
	hdrContentType := response.GetHeader("Content-Type")

	if hdrContentType != "" {
		o.ContentType = hdrContentType
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsOfWebhookTaskBadRequest creates a GetLogsOfWebhookTaskBadRequest with default headers values
func NewGetLogsOfWebhookTaskBadRequest() *GetLogsOfWebhookTaskBadRequest {
	return &GetLogsOfWebhookTaskBadRequest{}
}

/*
GetLogsOfWebhookTaskBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetLogsOfWebhookTaskBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get logs of webhook task bad request response has a 2xx status code
func (o *GetLogsOfWebhookTaskBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get logs of webhook task bad request response has a 3xx status code
func (o *GetLogsOfWebhookTaskBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get logs of webhook task bad request response has a 4xx status code
func (o *GetLogsOfWebhookTaskBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get logs of webhook task bad request response has a 5xx status code
func (o *GetLogsOfWebhookTaskBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get logs of webhook task bad request response a status code equal to that given
func (o *GetLogsOfWebhookTaskBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetLogsOfWebhookTaskBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks/{task_id}/log][%d] getLogsOfWebhookTaskBadRequest  %+v", 400, o.Payload)
}

func (o *GetLogsOfWebhookTaskBadRequest) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks/{task_id}/log][%d] getLogsOfWebhookTaskBadRequest  %+v", 400, o.Payload)
}

func (o *GetLogsOfWebhookTaskBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetLogsOfWebhookTaskBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsOfWebhookTaskUnauthorized creates a GetLogsOfWebhookTaskUnauthorized with default headers values
func NewGetLogsOfWebhookTaskUnauthorized() *GetLogsOfWebhookTaskUnauthorized {
	return &GetLogsOfWebhookTaskUnauthorized{}
}

/*
GetLogsOfWebhookTaskUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetLogsOfWebhookTaskUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get logs of webhook task unauthorized response has a 2xx status code
func (o *GetLogsOfWebhookTaskUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get logs of webhook task unauthorized response has a 3xx status code
func (o *GetLogsOfWebhookTaskUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get logs of webhook task unauthorized response has a 4xx status code
func (o *GetLogsOfWebhookTaskUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get logs of webhook task unauthorized response has a 5xx status code
func (o *GetLogsOfWebhookTaskUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get logs of webhook task unauthorized response a status code equal to that given
func (o *GetLogsOfWebhookTaskUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetLogsOfWebhookTaskUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks/{task_id}/log][%d] getLogsOfWebhookTaskUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLogsOfWebhookTaskUnauthorized) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks/{task_id}/log][%d] getLogsOfWebhookTaskUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLogsOfWebhookTaskUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetLogsOfWebhookTaskUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsOfWebhookTaskForbidden creates a GetLogsOfWebhookTaskForbidden with default headers values
func NewGetLogsOfWebhookTaskForbidden() *GetLogsOfWebhookTaskForbidden {
	return &GetLogsOfWebhookTaskForbidden{}
}

/*
GetLogsOfWebhookTaskForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetLogsOfWebhookTaskForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get logs of webhook task forbidden response has a 2xx status code
func (o *GetLogsOfWebhookTaskForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get logs of webhook task forbidden response has a 3xx status code
func (o *GetLogsOfWebhookTaskForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get logs of webhook task forbidden response has a 4xx status code
func (o *GetLogsOfWebhookTaskForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get logs of webhook task forbidden response has a 5xx status code
func (o *GetLogsOfWebhookTaskForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get logs of webhook task forbidden response a status code equal to that given
func (o *GetLogsOfWebhookTaskForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetLogsOfWebhookTaskForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks/{task_id}/log][%d] getLogsOfWebhookTaskForbidden  %+v", 403, o.Payload)
}

func (o *GetLogsOfWebhookTaskForbidden) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks/{task_id}/log][%d] getLogsOfWebhookTaskForbidden  %+v", 403, o.Payload)
}

func (o *GetLogsOfWebhookTaskForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetLogsOfWebhookTaskForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsOfWebhookTaskNotFound creates a GetLogsOfWebhookTaskNotFound with default headers values
func NewGetLogsOfWebhookTaskNotFound() *GetLogsOfWebhookTaskNotFound {
	return &GetLogsOfWebhookTaskNotFound{}
}

/*
GetLogsOfWebhookTaskNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetLogsOfWebhookTaskNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get logs of webhook task not found response has a 2xx status code
func (o *GetLogsOfWebhookTaskNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get logs of webhook task not found response has a 3xx status code
func (o *GetLogsOfWebhookTaskNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get logs of webhook task not found response has a 4xx status code
func (o *GetLogsOfWebhookTaskNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get logs of webhook task not found response has a 5xx status code
func (o *GetLogsOfWebhookTaskNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get logs of webhook task not found response a status code equal to that given
func (o *GetLogsOfWebhookTaskNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetLogsOfWebhookTaskNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks/{task_id}/log][%d] getLogsOfWebhookTaskNotFound  %+v", 404, o.Payload)
}

func (o *GetLogsOfWebhookTaskNotFound) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks/{task_id}/log][%d] getLogsOfWebhookTaskNotFound  %+v", 404, o.Payload)
}

func (o *GetLogsOfWebhookTaskNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetLogsOfWebhookTaskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsOfWebhookTaskInternalServerError creates a GetLogsOfWebhookTaskInternalServerError with default headers values
func NewGetLogsOfWebhookTaskInternalServerError() *GetLogsOfWebhookTaskInternalServerError {
	return &GetLogsOfWebhookTaskInternalServerError{}
}

/*
GetLogsOfWebhookTaskInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetLogsOfWebhookTaskInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get logs of webhook task internal server error response has a 2xx status code
func (o *GetLogsOfWebhookTaskInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get logs of webhook task internal server error response has a 3xx status code
func (o *GetLogsOfWebhookTaskInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get logs of webhook task internal server error response has a 4xx status code
func (o *GetLogsOfWebhookTaskInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get logs of webhook task internal server error response has a 5xx status code
func (o *GetLogsOfWebhookTaskInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get logs of webhook task internal server error response a status code equal to that given
func (o *GetLogsOfWebhookTaskInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetLogsOfWebhookTaskInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks/{task_id}/log][%d] getLogsOfWebhookTaskInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLogsOfWebhookTaskInternalServerError) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}/executions/{execution_id}/tasks/{task_id}/log][%d] getLogsOfWebhookTaskInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLogsOfWebhookTaskInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetLogsOfWebhookTaskInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
