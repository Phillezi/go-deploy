// Code generated by go-swagger; DO NOT EDIT.

package auditlog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/kthcloud/go-deploy/pkg/imp/harbor/sdk/v2.0/models"
)

// ListAuditLogsReader is a Reader for the ListAuditLogs structure.
type ListAuditLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAuditLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAuditLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListAuditLogsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListAuditLogsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListAuditLogsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListAuditLogsOK creates a ListAuditLogsOK with default headers values
func NewListAuditLogsOK() *ListAuditLogsOK {
	return &ListAuditLogsOK{}
}

/*
ListAuditLogsOK describes a response with status code 200, with default header values.

Success
*/
type ListAuditLogsOK struct {

	/* Link refers to the previous page and next page
	 */
	Link string

	/* The total count of auditlogs
	 */
	XTotalCount int64

	Payload []*models.AuditLog
}

// IsSuccess returns true when this list audit logs o k response has a 2xx status code
func (o *ListAuditLogsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list audit logs o k response has a 3xx status code
func (o *ListAuditLogsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list audit logs o k response has a 4xx status code
func (o *ListAuditLogsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list audit logs o k response has a 5xx status code
func (o *ListAuditLogsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list audit logs o k response a status code equal to that given
func (o *ListAuditLogsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListAuditLogsOK) Error() string {
	return fmt.Sprintf("[GET /audit-logs][%d] listAuditLogsOK  %+v", 200, o.Payload)
}

func (o *ListAuditLogsOK) String() string {
	return fmt.Sprintf("[GET /audit-logs][%d] listAuditLogsOK  %+v", 200, o.Payload)
}

func (o *ListAuditLogsOK) GetPayload() []*models.AuditLog {
	return o.Payload
}

func (o *ListAuditLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// hydrates response header X-Total-Count
	hdrXTotalCount := response.GetHeader("X-Total-Count")

	if hdrXTotalCount != "" {
		valxTotalCount, err := swag.ConvertInt64(hdrXTotalCount)
		if err != nil {
			return errors.InvalidType("X-Total-Count", "header", "int64", hdrXTotalCount)
		}
		o.XTotalCount = valxTotalCount
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAuditLogsBadRequest creates a ListAuditLogsBadRequest with default headers values
func NewListAuditLogsBadRequest() *ListAuditLogsBadRequest {
	return &ListAuditLogsBadRequest{}
}

/*
ListAuditLogsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ListAuditLogsBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list audit logs bad request response has a 2xx status code
func (o *ListAuditLogsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list audit logs bad request response has a 3xx status code
func (o *ListAuditLogsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list audit logs bad request response has a 4xx status code
func (o *ListAuditLogsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list audit logs bad request response has a 5xx status code
func (o *ListAuditLogsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list audit logs bad request response a status code equal to that given
func (o *ListAuditLogsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ListAuditLogsBadRequest) Error() string {
	return fmt.Sprintf("[GET /audit-logs][%d] listAuditLogsBadRequest  %+v", 400, o.Payload)
}

func (o *ListAuditLogsBadRequest) String() string {
	return fmt.Sprintf("[GET /audit-logs][%d] listAuditLogsBadRequest  %+v", 400, o.Payload)
}

func (o *ListAuditLogsBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListAuditLogsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListAuditLogsUnauthorized creates a ListAuditLogsUnauthorized with default headers values
func NewListAuditLogsUnauthorized() *ListAuditLogsUnauthorized {
	return &ListAuditLogsUnauthorized{}
}

/*
ListAuditLogsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListAuditLogsUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list audit logs unauthorized response has a 2xx status code
func (o *ListAuditLogsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list audit logs unauthorized response has a 3xx status code
func (o *ListAuditLogsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list audit logs unauthorized response has a 4xx status code
func (o *ListAuditLogsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list audit logs unauthorized response has a 5xx status code
func (o *ListAuditLogsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list audit logs unauthorized response a status code equal to that given
func (o *ListAuditLogsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListAuditLogsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /audit-logs][%d] listAuditLogsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListAuditLogsUnauthorized) String() string {
	return fmt.Sprintf("[GET /audit-logs][%d] listAuditLogsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListAuditLogsUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListAuditLogsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListAuditLogsInternalServerError creates a ListAuditLogsInternalServerError with default headers values
func NewListAuditLogsInternalServerError() *ListAuditLogsInternalServerError {
	return &ListAuditLogsInternalServerError{}
}

/*
ListAuditLogsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ListAuditLogsInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list audit logs internal server error response has a 2xx status code
func (o *ListAuditLogsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list audit logs internal server error response has a 3xx status code
func (o *ListAuditLogsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list audit logs internal server error response has a 4xx status code
func (o *ListAuditLogsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list audit logs internal server error response has a 5xx status code
func (o *ListAuditLogsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list audit logs internal server error response a status code equal to that given
func (o *ListAuditLogsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ListAuditLogsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /audit-logs][%d] listAuditLogsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListAuditLogsInternalServerError) String() string {
	return fmt.Sprintf("[GET /audit-logs][%d] listAuditLogsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListAuditLogsInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListAuditLogsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
