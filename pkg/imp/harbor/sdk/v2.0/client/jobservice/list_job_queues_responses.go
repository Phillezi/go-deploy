// Code generated by go-swagger; DO NOT EDIT.

package jobservice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"go-deploy/pkg/imp/harbor/sdk/v2.0/models"
)

// ListJobQueuesReader is a Reader for the ListJobQueues structure.
type ListJobQueuesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListJobQueuesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListJobQueuesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListJobQueuesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListJobQueuesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListJobQueuesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListJobQueuesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListJobQueuesOK creates a ListJobQueuesOK with default headers values
func NewListJobQueuesOK() *ListJobQueuesOK {
	return &ListJobQueuesOK{}
}

/*
ListJobQueuesOK describes a response with status code 200, with default header values.

List job queue successfully.
*/
type ListJobQueuesOK struct {
	Payload []*models.JobQueue
}

// IsSuccess returns true when this list job queues o k response has a 2xx status code
func (o *ListJobQueuesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list job queues o k response has a 3xx status code
func (o *ListJobQueuesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list job queues o k response has a 4xx status code
func (o *ListJobQueuesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list job queues o k response has a 5xx status code
func (o *ListJobQueuesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list job queues o k response a status code equal to that given
func (o *ListJobQueuesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListJobQueuesOK) Error() string {
	return fmt.Sprintf("[GET /jobservice/queues][%d] listJobQueuesOK  %+v", 200, o.Payload)
}

func (o *ListJobQueuesOK) String() string {
	return fmt.Sprintf("[GET /jobservice/queues][%d] listJobQueuesOK  %+v", 200, o.Payload)
}

func (o *ListJobQueuesOK) GetPayload() []*models.JobQueue {
	return o.Payload
}

func (o *ListJobQueuesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListJobQueuesUnauthorized creates a ListJobQueuesUnauthorized with default headers values
func NewListJobQueuesUnauthorized() *ListJobQueuesUnauthorized {
	return &ListJobQueuesUnauthorized{}
}

/*
ListJobQueuesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListJobQueuesUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list job queues unauthorized response has a 2xx status code
func (o *ListJobQueuesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list job queues unauthorized response has a 3xx status code
func (o *ListJobQueuesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list job queues unauthorized response has a 4xx status code
func (o *ListJobQueuesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list job queues unauthorized response has a 5xx status code
func (o *ListJobQueuesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list job queues unauthorized response a status code equal to that given
func (o *ListJobQueuesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListJobQueuesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /jobservice/queues][%d] listJobQueuesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListJobQueuesUnauthorized) String() string {
	return fmt.Sprintf("[GET /jobservice/queues][%d] listJobQueuesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListJobQueuesUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListJobQueuesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListJobQueuesForbidden creates a ListJobQueuesForbidden with default headers values
func NewListJobQueuesForbidden() *ListJobQueuesForbidden {
	return &ListJobQueuesForbidden{}
}

/*
ListJobQueuesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListJobQueuesForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list job queues forbidden response has a 2xx status code
func (o *ListJobQueuesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list job queues forbidden response has a 3xx status code
func (o *ListJobQueuesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list job queues forbidden response has a 4xx status code
func (o *ListJobQueuesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list job queues forbidden response has a 5xx status code
func (o *ListJobQueuesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list job queues forbidden response a status code equal to that given
func (o *ListJobQueuesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListJobQueuesForbidden) Error() string {
	return fmt.Sprintf("[GET /jobservice/queues][%d] listJobQueuesForbidden  %+v", 403, o.Payload)
}

func (o *ListJobQueuesForbidden) String() string {
	return fmt.Sprintf("[GET /jobservice/queues][%d] listJobQueuesForbidden  %+v", 403, o.Payload)
}

func (o *ListJobQueuesForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListJobQueuesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListJobQueuesNotFound creates a ListJobQueuesNotFound with default headers values
func NewListJobQueuesNotFound() *ListJobQueuesNotFound {
	return &ListJobQueuesNotFound{}
}

/*
ListJobQueuesNotFound describes a response with status code 404, with default header values.

Not found
*/
type ListJobQueuesNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list job queues not found response has a 2xx status code
func (o *ListJobQueuesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list job queues not found response has a 3xx status code
func (o *ListJobQueuesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list job queues not found response has a 4xx status code
func (o *ListJobQueuesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list job queues not found response has a 5xx status code
func (o *ListJobQueuesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list job queues not found response a status code equal to that given
func (o *ListJobQueuesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ListJobQueuesNotFound) Error() string {
	return fmt.Sprintf("[GET /jobservice/queues][%d] listJobQueuesNotFound  %+v", 404, o.Payload)
}

func (o *ListJobQueuesNotFound) String() string {
	return fmt.Sprintf("[GET /jobservice/queues][%d] listJobQueuesNotFound  %+v", 404, o.Payload)
}

func (o *ListJobQueuesNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListJobQueuesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListJobQueuesInternalServerError creates a ListJobQueuesInternalServerError with default headers values
func NewListJobQueuesInternalServerError() *ListJobQueuesInternalServerError {
	return &ListJobQueuesInternalServerError{}
}

/*
ListJobQueuesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ListJobQueuesInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list job queues internal server error response has a 2xx status code
func (o *ListJobQueuesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list job queues internal server error response has a 3xx status code
func (o *ListJobQueuesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list job queues internal server error response has a 4xx status code
func (o *ListJobQueuesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list job queues internal server error response has a 5xx status code
func (o *ListJobQueuesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list job queues internal server error response a status code equal to that given
func (o *ListJobQueuesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ListJobQueuesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /jobservice/queues][%d] listJobQueuesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListJobQueuesInternalServerError) String() string {
	return fmt.Sprintf("[GET /jobservice/queues][%d] listJobQueuesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListJobQueuesInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListJobQueuesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
