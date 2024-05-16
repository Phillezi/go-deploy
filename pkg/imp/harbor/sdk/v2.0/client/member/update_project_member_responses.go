// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"go-deploy/pkg/imp/harbor/sdk/v2.0/models"
)

// UpdateProjectMemberReader is a Reader for the UpdateProjectMember structure.
type UpdateProjectMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProjectMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProjectMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateProjectMemberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateProjectMemberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateProjectMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateProjectMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateProjectMemberInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateProjectMemberOK creates a UpdateProjectMemberOK with default headers values
func NewUpdateProjectMemberOK() *UpdateProjectMemberOK {
	return &UpdateProjectMemberOK{}
}

/*
UpdateProjectMemberOK describes a response with status code 200, with default header values.

Success
*/
type UpdateProjectMemberOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this update project member o k response has a 2xx status code
func (o *UpdateProjectMemberOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update project member o k response has a 3xx status code
func (o *UpdateProjectMemberOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project member o k response has a 4xx status code
func (o *UpdateProjectMemberOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update project member o k response has a 5xx status code
func (o *UpdateProjectMemberOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update project member o k response a status code equal to that given
func (o *UpdateProjectMemberOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateProjectMemberOK) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/members/{mid}][%d] updateProjectMemberOK ", 200)
}

func (o *UpdateProjectMemberOK) String() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/members/{mid}][%d] updateProjectMemberOK ", 200)
}

func (o *UpdateProjectMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewUpdateProjectMemberBadRequest creates a UpdateProjectMemberBadRequest with default headers values
func NewUpdateProjectMemberBadRequest() *UpdateProjectMemberBadRequest {
	return &UpdateProjectMemberBadRequest{}
}

/*
UpdateProjectMemberBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateProjectMemberBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update project member bad request response has a 2xx status code
func (o *UpdateProjectMemberBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project member bad request response has a 3xx status code
func (o *UpdateProjectMemberBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project member bad request response has a 4xx status code
func (o *UpdateProjectMemberBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update project member bad request response has a 5xx status code
func (o *UpdateProjectMemberBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update project member bad request response a status code equal to that given
func (o *UpdateProjectMemberBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateProjectMemberBadRequest) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/members/{mid}][%d] updateProjectMemberBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateProjectMemberBadRequest) String() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/members/{mid}][%d] updateProjectMemberBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateProjectMemberBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateProjectMemberBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateProjectMemberUnauthorized creates a UpdateProjectMemberUnauthorized with default headers values
func NewUpdateProjectMemberUnauthorized() *UpdateProjectMemberUnauthorized {
	return &UpdateProjectMemberUnauthorized{}
}

/*
UpdateProjectMemberUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateProjectMemberUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update project member unauthorized response has a 2xx status code
func (o *UpdateProjectMemberUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project member unauthorized response has a 3xx status code
func (o *UpdateProjectMemberUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project member unauthorized response has a 4xx status code
func (o *UpdateProjectMemberUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update project member unauthorized response has a 5xx status code
func (o *UpdateProjectMemberUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update project member unauthorized response a status code equal to that given
func (o *UpdateProjectMemberUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateProjectMemberUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/members/{mid}][%d] updateProjectMemberUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateProjectMemberUnauthorized) String() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/members/{mid}][%d] updateProjectMemberUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateProjectMemberUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateProjectMemberUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateProjectMemberForbidden creates a UpdateProjectMemberForbidden with default headers values
func NewUpdateProjectMemberForbidden() *UpdateProjectMemberForbidden {
	return &UpdateProjectMemberForbidden{}
}

/*
UpdateProjectMemberForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateProjectMemberForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update project member forbidden response has a 2xx status code
func (o *UpdateProjectMemberForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project member forbidden response has a 3xx status code
func (o *UpdateProjectMemberForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project member forbidden response has a 4xx status code
func (o *UpdateProjectMemberForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update project member forbidden response has a 5xx status code
func (o *UpdateProjectMemberForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update project member forbidden response a status code equal to that given
func (o *UpdateProjectMemberForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateProjectMemberForbidden) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/members/{mid}][%d] updateProjectMemberForbidden  %+v", 403, o.Payload)
}

func (o *UpdateProjectMemberForbidden) String() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/members/{mid}][%d] updateProjectMemberForbidden  %+v", 403, o.Payload)
}

func (o *UpdateProjectMemberForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateProjectMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateProjectMemberNotFound creates a UpdateProjectMemberNotFound with default headers values
func NewUpdateProjectMemberNotFound() *UpdateProjectMemberNotFound {
	return &UpdateProjectMemberNotFound{}
}

/*
UpdateProjectMemberNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateProjectMemberNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update project member not found response has a 2xx status code
func (o *UpdateProjectMemberNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project member not found response has a 3xx status code
func (o *UpdateProjectMemberNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project member not found response has a 4xx status code
func (o *UpdateProjectMemberNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update project member not found response has a 5xx status code
func (o *UpdateProjectMemberNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update project member not found response a status code equal to that given
func (o *UpdateProjectMemberNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateProjectMemberNotFound) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/members/{mid}][%d] updateProjectMemberNotFound  %+v", 404, o.Payload)
}

func (o *UpdateProjectMemberNotFound) String() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/members/{mid}][%d] updateProjectMemberNotFound  %+v", 404, o.Payload)
}

func (o *UpdateProjectMemberNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateProjectMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateProjectMemberInternalServerError creates a UpdateProjectMemberInternalServerError with default headers values
func NewUpdateProjectMemberInternalServerError() *UpdateProjectMemberInternalServerError {
	return &UpdateProjectMemberInternalServerError{}
}

/*
UpdateProjectMemberInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type UpdateProjectMemberInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update project member internal server error response has a 2xx status code
func (o *UpdateProjectMemberInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project member internal server error response has a 3xx status code
func (o *UpdateProjectMemberInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project member internal server error response has a 4xx status code
func (o *UpdateProjectMemberInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update project member internal server error response has a 5xx status code
func (o *UpdateProjectMemberInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update project member internal server error response a status code equal to that given
func (o *UpdateProjectMemberInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateProjectMemberInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/members/{mid}][%d] updateProjectMemberInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateProjectMemberInternalServerError) String() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/members/{mid}][%d] updateProjectMemberInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateProjectMemberInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateProjectMemberInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
