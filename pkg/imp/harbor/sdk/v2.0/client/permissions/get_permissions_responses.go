// Code generated by go-swagger; DO NOT EDIT.

package permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kthcloud/go-deploy/pkg/imp/harbor/sdk/v2.0/models"
)

// GetPermissionsReader is a Reader for the GetPermissions structure.
type GetPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPermissionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPermissionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPermissionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPermissionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPermissionsOK creates a GetPermissionsOK with default headers values
func NewGetPermissionsOK() *GetPermissionsOK {
	return &GetPermissionsOK{}
}

/*
GetPermissionsOK describes a response with status code 200, with default header values.

Get permissions successfully.
*/
type GetPermissionsOK struct {
	Payload *models.Permissions
}

// IsSuccess returns true when this get permissions o k response has a 2xx status code
func (o *GetPermissionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get permissions o k response has a 3xx status code
func (o *GetPermissionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get permissions o k response has a 4xx status code
func (o *GetPermissionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get permissions o k response has a 5xx status code
func (o *GetPermissionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get permissions o k response a status code equal to that given
func (o *GetPermissionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPermissionsOK) Error() string {
	return fmt.Sprintf("[GET /permissions][%d] getPermissionsOK  %+v", 200, o.Payload)
}

func (o *GetPermissionsOK) String() string {
	return fmt.Sprintf("[GET /permissions][%d] getPermissionsOK  %+v", 200, o.Payload)
}

func (o *GetPermissionsOK) GetPayload() *models.Permissions {
	return o.Payload
}

func (o *GetPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Permissions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPermissionsUnauthorized creates a GetPermissionsUnauthorized with default headers values
func NewGetPermissionsUnauthorized() *GetPermissionsUnauthorized {
	return &GetPermissionsUnauthorized{}
}

/*
GetPermissionsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetPermissionsUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get permissions unauthorized response has a 2xx status code
func (o *GetPermissionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get permissions unauthorized response has a 3xx status code
func (o *GetPermissionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get permissions unauthorized response has a 4xx status code
func (o *GetPermissionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get permissions unauthorized response has a 5xx status code
func (o *GetPermissionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get permissions unauthorized response a status code equal to that given
func (o *GetPermissionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetPermissionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /permissions][%d] getPermissionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPermissionsUnauthorized) String() string {
	return fmt.Sprintf("[GET /permissions][%d] getPermissionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPermissionsUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetPermissionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPermissionsForbidden creates a GetPermissionsForbidden with default headers values
func NewGetPermissionsForbidden() *GetPermissionsForbidden {
	return &GetPermissionsForbidden{}
}

/*
GetPermissionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetPermissionsForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get permissions forbidden response has a 2xx status code
func (o *GetPermissionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get permissions forbidden response has a 3xx status code
func (o *GetPermissionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get permissions forbidden response has a 4xx status code
func (o *GetPermissionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get permissions forbidden response has a 5xx status code
func (o *GetPermissionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get permissions forbidden response a status code equal to that given
func (o *GetPermissionsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetPermissionsForbidden) Error() string {
	return fmt.Sprintf("[GET /permissions][%d] getPermissionsForbidden  %+v", 403, o.Payload)
}

func (o *GetPermissionsForbidden) String() string {
	return fmt.Sprintf("[GET /permissions][%d] getPermissionsForbidden  %+v", 403, o.Payload)
}

func (o *GetPermissionsForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetPermissionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPermissionsNotFound creates a GetPermissionsNotFound with default headers values
func NewGetPermissionsNotFound() *GetPermissionsNotFound {
	return &GetPermissionsNotFound{}
}

/*
GetPermissionsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetPermissionsNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get permissions not found response has a 2xx status code
func (o *GetPermissionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get permissions not found response has a 3xx status code
func (o *GetPermissionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get permissions not found response has a 4xx status code
func (o *GetPermissionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get permissions not found response has a 5xx status code
func (o *GetPermissionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get permissions not found response a status code equal to that given
func (o *GetPermissionsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetPermissionsNotFound) Error() string {
	return fmt.Sprintf("[GET /permissions][%d] getPermissionsNotFound  %+v", 404, o.Payload)
}

func (o *GetPermissionsNotFound) String() string {
	return fmt.Sprintf("[GET /permissions][%d] getPermissionsNotFound  %+v", 404, o.Payload)
}

func (o *GetPermissionsNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetPermissionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPermissionsInternalServerError creates a GetPermissionsInternalServerError with default headers values
func NewGetPermissionsInternalServerError() *GetPermissionsInternalServerError {
	return &GetPermissionsInternalServerError{}
}

/*
GetPermissionsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetPermissionsInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get permissions internal server error response has a 2xx status code
func (o *GetPermissionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get permissions internal server error response has a 3xx status code
func (o *GetPermissionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get permissions internal server error response has a 4xx status code
func (o *GetPermissionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get permissions internal server error response has a 5xx status code
func (o *GetPermissionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get permissions internal server error response a status code equal to that given
func (o *GetPermissionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetPermissionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /permissions][%d] getPermissionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPermissionsInternalServerError) String() string {
	return fmt.Sprintf("[GET /permissions][%d] getPermissionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPermissionsInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetPermissionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
