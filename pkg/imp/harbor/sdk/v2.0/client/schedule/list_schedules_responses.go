// Code generated by go-swagger; DO NOT EDIT.

package schedule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"go-deploy/pkg/imp/harbor/sdk/v2.0/models"
)

// ListSchedulesReader is a Reader for the ListSchedules structure.
type ListSchedulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSchedulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSchedulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListSchedulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListSchedulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListSchedulesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListSchedulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListSchedulesOK creates a ListSchedulesOK with default headers values
func NewListSchedulesOK() *ListSchedulesOK {
	return &ListSchedulesOK{}
}

/*
ListSchedulesOK describes a response with status code 200, with default header values.

list schedule successfully.
*/
type ListSchedulesOK struct {

	/* Link to previous page and next page
	 */
	Link string

	/* The total count of available items
	 */
	XTotalCount int64

	Payload []*models.ScheduleTask
}

// IsSuccess returns true when this list schedules o k response has a 2xx status code
func (o *ListSchedulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list schedules o k response has a 3xx status code
func (o *ListSchedulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list schedules o k response has a 4xx status code
func (o *ListSchedulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list schedules o k response has a 5xx status code
func (o *ListSchedulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list schedules o k response a status code equal to that given
func (o *ListSchedulesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListSchedulesOK) Error() string {
	return fmt.Sprintf("[GET /schedules][%d] listSchedulesOK  %+v", 200, o.Payload)
}

func (o *ListSchedulesOK) String() string {
	return fmt.Sprintf("[GET /schedules][%d] listSchedulesOK  %+v", 200, o.Payload)
}

func (o *ListSchedulesOK) GetPayload() []*models.ScheduleTask {
	return o.Payload
}

func (o *ListSchedulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListSchedulesUnauthorized creates a ListSchedulesUnauthorized with default headers values
func NewListSchedulesUnauthorized() *ListSchedulesUnauthorized {
	return &ListSchedulesUnauthorized{}
}

/*
ListSchedulesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListSchedulesUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list schedules unauthorized response has a 2xx status code
func (o *ListSchedulesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list schedules unauthorized response has a 3xx status code
func (o *ListSchedulesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list schedules unauthorized response has a 4xx status code
func (o *ListSchedulesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list schedules unauthorized response has a 5xx status code
func (o *ListSchedulesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list schedules unauthorized response a status code equal to that given
func (o *ListSchedulesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListSchedulesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /schedules][%d] listSchedulesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListSchedulesUnauthorized) String() string {
	return fmt.Sprintf("[GET /schedules][%d] listSchedulesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListSchedulesUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListSchedulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListSchedulesForbidden creates a ListSchedulesForbidden with default headers values
func NewListSchedulesForbidden() *ListSchedulesForbidden {
	return &ListSchedulesForbidden{}
}

/*
ListSchedulesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListSchedulesForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list schedules forbidden response has a 2xx status code
func (o *ListSchedulesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list schedules forbidden response has a 3xx status code
func (o *ListSchedulesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list schedules forbidden response has a 4xx status code
func (o *ListSchedulesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list schedules forbidden response has a 5xx status code
func (o *ListSchedulesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list schedules forbidden response a status code equal to that given
func (o *ListSchedulesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListSchedulesForbidden) Error() string {
	return fmt.Sprintf("[GET /schedules][%d] listSchedulesForbidden  %+v", 403, o.Payload)
}

func (o *ListSchedulesForbidden) String() string {
	return fmt.Sprintf("[GET /schedules][%d] listSchedulesForbidden  %+v", 403, o.Payload)
}

func (o *ListSchedulesForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListSchedulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListSchedulesNotFound creates a ListSchedulesNotFound with default headers values
func NewListSchedulesNotFound() *ListSchedulesNotFound {
	return &ListSchedulesNotFound{}
}

/*
ListSchedulesNotFound describes a response with status code 404, with default header values.

Not found
*/
type ListSchedulesNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list schedules not found response has a 2xx status code
func (o *ListSchedulesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list schedules not found response has a 3xx status code
func (o *ListSchedulesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list schedules not found response has a 4xx status code
func (o *ListSchedulesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list schedules not found response has a 5xx status code
func (o *ListSchedulesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list schedules not found response a status code equal to that given
func (o *ListSchedulesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ListSchedulesNotFound) Error() string {
	return fmt.Sprintf("[GET /schedules][%d] listSchedulesNotFound  %+v", 404, o.Payload)
}

func (o *ListSchedulesNotFound) String() string {
	return fmt.Sprintf("[GET /schedules][%d] listSchedulesNotFound  %+v", 404, o.Payload)
}

func (o *ListSchedulesNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListSchedulesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListSchedulesInternalServerError creates a ListSchedulesInternalServerError with default headers values
func NewListSchedulesInternalServerError() *ListSchedulesInternalServerError {
	return &ListSchedulesInternalServerError{}
}

/*
ListSchedulesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ListSchedulesInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list schedules internal server error response has a 2xx status code
func (o *ListSchedulesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list schedules internal server error response has a 3xx status code
func (o *ListSchedulesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list schedules internal server error response has a 4xx status code
func (o *ListSchedulesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list schedules internal server error response has a 5xx status code
func (o *ListSchedulesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list schedules internal server error response a status code equal to that given
func (o *ListSchedulesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ListSchedulesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /schedules][%d] listSchedulesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListSchedulesInternalServerError) String() string {
	return fmt.Sprintf("[GET /schedules][%d] listSchedulesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListSchedulesInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListSchedulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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