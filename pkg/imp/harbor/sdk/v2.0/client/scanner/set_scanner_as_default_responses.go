// Code generated by go-swagger; DO NOT EDIT.

package scanner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"go-deploy/pkg/imp/harbor/sdk/v2.0/models"
)

// SetScannerAsDefaultReader is a Reader for the SetScannerAsDefault structure.
type SetScannerAsDefaultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetScannerAsDefaultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetScannerAsDefaultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSetScannerAsDefaultUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSetScannerAsDefaultForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSetScannerAsDefaultInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetScannerAsDefaultOK creates a SetScannerAsDefaultOK with default headers values
func NewSetScannerAsDefaultOK() *SetScannerAsDefaultOK {
	return &SetScannerAsDefaultOK{}
}

/*
SetScannerAsDefaultOK describes a response with status code 200, with default header values.

Successfully set the specified scanner registration as system default
*/
type SetScannerAsDefaultOK struct {
}

// IsSuccess returns true when this set scanner as default o k response has a 2xx status code
func (o *SetScannerAsDefaultOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set scanner as default o k response has a 3xx status code
func (o *SetScannerAsDefaultOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set scanner as default o k response has a 4xx status code
func (o *SetScannerAsDefaultOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this set scanner as default o k response has a 5xx status code
func (o *SetScannerAsDefaultOK) IsServerError() bool {
	return false
}

// IsCode returns true when this set scanner as default o k response a status code equal to that given
func (o *SetScannerAsDefaultOK) IsCode(code int) bool {
	return code == 200
}

func (o *SetScannerAsDefaultOK) Error() string {
	return fmt.Sprintf("[PATCH /scanners/{registration_id}][%d] setScannerAsDefaultOK ", 200)
}

func (o *SetScannerAsDefaultOK) String() string {
	return fmt.Sprintf("[PATCH /scanners/{registration_id}][%d] setScannerAsDefaultOK ", 200)
}

func (o *SetScannerAsDefaultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetScannerAsDefaultUnauthorized creates a SetScannerAsDefaultUnauthorized with default headers values
func NewSetScannerAsDefaultUnauthorized() *SetScannerAsDefaultUnauthorized {
	return &SetScannerAsDefaultUnauthorized{}
}

/*
SetScannerAsDefaultUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SetScannerAsDefaultUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this set scanner as default unauthorized response has a 2xx status code
func (o *SetScannerAsDefaultUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set scanner as default unauthorized response has a 3xx status code
func (o *SetScannerAsDefaultUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set scanner as default unauthorized response has a 4xx status code
func (o *SetScannerAsDefaultUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this set scanner as default unauthorized response has a 5xx status code
func (o *SetScannerAsDefaultUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this set scanner as default unauthorized response a status code equal to that given
func (o *SetScannerAsDefaultUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SetScannerAsDefaultUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /scanners/{registration_id}][%d] setScannerAsDefaultUnauthorized  %+v", 401, o.Payload)
}

func (o *SetScannerAsDefaultUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /scanners/{registration_id}][%d] setScannerAsDefaultUnauthorized  %+v", 401, o.Payload)
}

func (o *SetScannerAsDefaultUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *SetScannerAsDefaultUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSetScannerAsDefaultForbidden creates a SetScannerAsDefaultForbidden with default headers values
func NewSetScannerAsDefaultForbidden() *SetScannerAsDefaultForbidden {
	return &SetScannerAsDefaultForbidden{}
}

/*
SetScannerAsDefaultForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SetScannerAsDefaultForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this set scanner as default forbidden response has a 2xx status code
func (o *SetScannerAsDefaultForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set scanner as default forbidden response has a 3xx status code
func (o *SetScannerAsDefaultForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set scanner as default forbidden response has a 4xx status code
func (o *SetScannerAsDefaultForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this set scanner as default forbidden response has a 5xx status code
func (o *SetScannerAsDefaultForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this set scanner as default forbidden response a status code equal to that given
func (o *SetScannerAsDefaultForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SetScannerAsDefaultForbidden) Error() string {
	return fmt.Sprintf("[PATCH /scanners/{registration_id}][%d] setScannerAsDefaultForbidden  %+v", 403, o.Payload)
}

func (o *SetScannerAsDefaultForbidden) String() string {
	return fmt.Sprintf("[PATCH /scanners/{registration_id}][%d] setScannerAsDefaultForbidden  %+v", 403, o.Payload)
}

func (o *SetScannerAsDefaultForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *SetScannerAsDefaultForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSetScannerAsDefaultInternalServerError creates a SetScannerAsDefaultInternalServerError with default headers values
func NewSetScannerAsDefaultInternalServerError() *SetScannerAsDefaultInternalServerError {
	return &SetScannerAsDefaultInternalServerError{}
}

/*
SetScannerAsDefaultInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SetScannerAsDefaultInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this set scanner as default internal server error response has a 2xx status code
func (o *SetScannerAsDefaultInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set scanner as default internal server error response has a 3xx status code
func (o *SetScannerAsDefaultInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set scanner as default internal server error response has a 4xx status code
func (o *SetScannerAsDefaultInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this set scanner as default internal server error response has a 5xx status code
func (o *SetScannerAsDefaultInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this set scanner as default internal server error response a status code equal to that given
func (o *SetScannerAsDefaultInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SetScannerAsDefaultInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /scanners/{registration_id}][%d] setScannerAsDefaultInternalServerError  %+v", 500, o.Payload)
}

func (o *SetScannerAsDefaultInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /scanners/{registration_id}][%d] setScannerAsDefaultInternalServerError  %+v", 500, o.Payload)
}

func (o *SetScannerAsDefaultInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *SetScannerAsDefaultInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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