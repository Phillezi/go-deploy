// Code generated by go-swagger; DO NOT EDIT.

package label

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"go-deploy/pkg/imp/harbor/sdk/v2.0/models"
)

// CreateLabelReader is a Reader for the CreateLabel structure.
type CreateLabelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateLabelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateLabelCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateLabelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateLabelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateLabelConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateLabelUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateLabelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateLabelCreated creates a CreateLabelCreated with default headers values
func NewCreateLabelCreated() *CreateLabelCreated {
	return &CreateLabelCreated{}
}

/*
CreateLabelCreated describes a response with status code 201, with default header values.

Create successfully.
*/
type CreateLabelCreated struct {

	/* The URL of the created resource
	 */
	Location string
}

// IsSuccess returns true when this create label created response has a 2xx status code
func (o *CreateLabelCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create label created response has a 3xx status code
func (o *CreateLabelCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create label created response has a 4xx status code
func (o *CreateLabelCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create label created response has a 5xx status code
func (o *CreateLabelCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create label created response a status code equal to that given
func (o *CreateLabelCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateLabelCreated) Error() string {
	return fmt.Sprintf("[POST /labels][%d] createLabelCreated ", 201)
}

func (o *CreateLabelCreated) String() string {
	return fmt.Sprintf("[POST /labels][%d] createLabelCreated ", 201)
}

func (o *CreateLabelCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	return nil
}

// NewCreateLabelBadRequest creates a CreateLabelBadRequest with default headers values
func NewCreateLabelBadRequest() *CreateLabelBadRequest {
	return &CreateLabelBadRequest{}
}

/*
CreateLabelBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateLabelBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create label bad request response has a 2xx status code
func (o *CreateLabelBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create label bad request response has a 3xx status code
func (o *CreateLabelBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create label bad request response has a 4xx status code
func (o *CreateLabelBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create label bad request response has a 5xx status code
func (o *CreateLabelBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create label bad request response a status code equal to that given
func (o *CreateLabelBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateLabelBadRequest) Error() string {
	return fmt.Sprintf("[POST /labels][%d] createLabelBadRequest  %+v", 400, o.Payload)
}

func (o *CreateLabelBadRequest) String() string {
	return fmt.Sprintf("[POST /labels][%d] createLabelBadRequest  %+v", 400, o.Payload)
}

func (o *CreateLabelBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateLabelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateLabelUnauthorized creates a CreateLabelUnauthorized with default headers values
func NewCreateLabelUnauthorized() *CreateLabelUnauthorized {
	return &CreateLabelUnauthorized{}
}

/*
CreateLabelUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateLabelUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create label unauthorized response has a 2xx status code
func (o *CreateLabelUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create label unauthorized response has a 3xx status code
func (o *CreateLabelUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create label unauthorized response has a 4xx status code
func (o *CreateLabelUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create label unauthorized response has a 5xx status code
func (o *CreateLabelUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create label unauthorized response a status code equal to that given
func (o *CreateLabelUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateLabelUnauthorized) Error() string {
	return fmt.Sprintf("[POST /labels][%d] createLabelUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateLabelUnauthorized) String() string {
	return fmt.Sprintf("[POST /labels][%d] createLabelUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateLabelUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateLabelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateLabelConflict creates a CreateLabelConflict with default headers values
func NewCreateLabelConflict() *CreateLabelConflict {
	return &CreateLabelConflict{}
}

/*
CreateLabelConflict describes a response with status code 409, with default header values.

Conflict
*/
type CreateLabelConflict struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create label conflict response has a 2xx status code
func (o *CreateLabelConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create label conflict response has a 3xx status code
func (o *CreateLabelConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create label conflict response has a 4xx status code
func (o *CreateLabelConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create label conflict response has a 5xx status code
func (o *CreateLabelConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create label conflict response a status code equal to that given
func (o *CreateLabelConflict) IsCode(code int) bool {
	return code == 409
}

func (o *CreateLabelConflict) Error() string {
	return fmt.Sprintf("[POST /labels][%d] createLabelConflict  %+v", 409, o.Payload)
}

func (o *CreateLabelConflict) String() string {
	return fmt.Sprintf("[POST /labels][%d] createLabelConflict  %+v", 409, o.Payload)
}

func (o *CreateLabelConflict) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateLabelConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateLabelUnsupportedMediaType creates a CreateLabelUnsupportedMediaType with default headers values
func NewCreateLabelUnsupportedMediaType() *CreateLabelUnsupportedMediaType {
	return &CreateLabelUnsupportedMediaType{}
}

/*
CreateLabelUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported MediaType
*/
type CreateLabelUnsupportedMediaType struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create label unsupported media type response has a 2xx status code
func (o *CreateLabelUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create label unsupported media type response has a 3xx status code
func (o *CreateLabelUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create label unsupported media type response has a 4xx status code
func (o *CreateLabelUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this create label unsupported media type response has a 5xx status code
func (o *CreateLabelUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this create label unsupported media type response a status code equal to that given
func (o *CreateLabelUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *CreateLabelUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /labels][%d] createLabelUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CreateLabelUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /labels][%d] createLabelUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CreateLabelUnsupportedMediaType) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateLabelUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateLabelInternalServerError creates a CreateLabelInternalServerError with default headers values
func NewCreateLabelInternalServerError() *CreateLabelInternalServerError {
	return &CreateLabelInternalServerError{}
}

/*
CreateLabelInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type CreateLabelInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create label internal server error response has a 2xx status code
func (o *CreateLabelInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create label internal server error response has a 3xx status code
func (o *CreateLabelInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create label internal server error response has a 4xx status code
func (o *CreateLabelInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create label internal server error response has a 5xx status code
func (o *CreateLabelInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create label internal server error response a status code equal to that given
func (o *CreateLabelInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateLabelInternalServerError) Error() string {
	return fmt.Sprintf("[POST /labels][%d] createLabelInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateLabelInternalServerError) String() string {
	return fmt.Sprintf("[POST /labels][%d] createLabelInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateLabelInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateLabelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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