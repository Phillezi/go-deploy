// Code generated by go-swagger; DO NOT EDIT.

package ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"go-deploy/pkg/imp/harbor/sdk/v2.0/models"
)

// SearchLdapUserReader is a Reader for the SearchLdapUser structure.
type SearchLdapUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchLdapUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchLdapUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchLdapUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSearchLdapUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchLdapUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchLdapUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchLdapUserOK creates a SearchLdapUserOK with default headers values
func NewSearchLdapUserOK() *SearchLdapUserOK {
	return &SearchLdapUserOK{}
}

/*
SearchLdapUserOK describes a response with status code 200, with default header values.

Search ldap users successfully.
*/
type SearchLdapUserOK struct {
	Payload []*models.LdapUser
}

// IsSuccess returns true when this search ldap user o k response has a 2xx status code
func (o *SearchLdapUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search ldap user o k response has a 3xx status code
func (o *SearchLdapUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search ldap user o k response has a 4xx status code
func (o *SearchLdapUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search ldap user o k response has a 5xx status code
func (o *SearchLdapUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search ldap user o k response a status code equal to that given
func (o *SearchLdapUserOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchLdapUserOK) Error() string {
	return fmt.Sprintf("[GET /ldap/users/search][%d] searchLdapUserOK  %+v", 200, o.Payload)
}

func (o *SearchLdapUserOK) String() string {
	return fmt.Sprintf("[GET /ldap/users/search][%d] searchLdapUserOK  %+v", 200, o.Payload)
}

func (o *SearchLdapUserOK) GetPayload() []*models.LdapUser {
	return o.Payload
}

func (o *SearchLdapUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchLdapUserBadRequest creates a SearchLdapUserBadRequest with default headers values
func NewSearchLdapUserBadRequest() *SearchLdapUserBadRequest {
	return &SearchLdapUserBadRequest{}
}

/*
SearchLdapUserBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type SearchLdapUserBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this search ldap user bad request response has a 2xx status code
func (o *SearchLdapUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search ldap user bad request response has a 3xx status code
func (o *SearchLdapUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search ldap user bad request response has a 4xx status code
func (o *SearchLdapUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this search ldap user bad request response has a 5xx status code
func (o *SearchLdapUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this search ldap user bad request response a status code equal to that given
func (o *SearchLdapUserBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SearchLdapUserBadRequest) Error() string {
	return fmt.Sprintf("[GET /ldap/users/search][%d] searchLdapUserBadRequest  %+v", 400, o.Payload)
}

func (o *SearchLdapUserBadRequest) String() string {
	return fmt.Sprintf("[GET /ldap/users/search][%d] searchLdapUserBadRequest  %+v", 400, o.Payload)
}

func (o *SearchLdapUserBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *SearchLdapUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSearchLdapUserUnauthorized creates a SearchLdapUserUnauthorized with default headers values
func NewSearchLdapUserUnauthorized() *SearchLdapUserUnauthorized {
	return &SearchLdapUserUnauthorized{}
}

/*
SearchLdapUserUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SearchLdapUserUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this search ldap user unauthorized response has a 2xx status code
func (o *SearchLdapUserUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search ldap user unauthorized response has a 3xx status code
func (o *SearchLdapUserUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search ldap user unauthorized response has a 4xx status code
func (o *SearchLdapUserUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this search ldap user unauthorized response has a 5xx status code
func (o *SearchLdapUserUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this search ldap user unauthorized response a status code equal to that given
func (o *SearchLdapUserUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SearchLdapUserUnauthorized) Error() string {
	return fmt.Sprintf("[GET /ldap/users/search][%d] searchLdapUserUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchLdapUserUnauthorized) String() string {
	return fmt.Sprintf("[GET /ldap/users/search][%d] searchLdapUserUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchLdapUserUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *SearchLdapUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSearchLdapUserForbidden creates a SearchLdapUserForbidden with default headers values
func NewSearchLdapUserForbidden() *SearchLdapUserForbidden {
	return &SearchLdapUserForbidden{}
}

/*
SearchLdapUserForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SearchLdapUserForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this search ldap user forbidden response has a 2xx status code
func (o *SearchLdapUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search ldap user forbidden response has a 3xx status code
func (o *SearchLdapUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search ldap user forbidden response has a 4xx status code
func (o *SearchLdapUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this search ldap user forbidden response has a 5xx status code
func (o *SearchLdapUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this search ldap user forbidden response a status code equal to that given
func (o *SearchLdapUserForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SearchLdapUserForbidden) Error() string {
	return fmt.Sprintf("[GET /ldap/users/search][%d] searchLdapUserForbidden  %+v", 403, o.Payload)
}

func (o *SearchLdapUserForbidden) String() string {
	return fmt.Sprintf("[GET /ldap/users/search][%d] searchLdapUserForbidden  %+v", 403, o.Payload)
}

func (o *SearchLdapUserForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *SearchLdapUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSearchLdapUserInternalServerError creates a SearchLdapUserInternalServerError with default headers values
func NewSearchLdapUserInternalServerError() *SearchLdapUserInternalServerError {
	return &SearchLdapUserInternalServerError{}
}

/*
SearchLdapUserInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SearchLdapUserInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this search ldap user internal server error response has a 2xx status code
func (o *SearchLdapUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search ldap user internal server error response has a 3xx status code
func (o *SearchLdapUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search ldap user internal server error response has a 4xx status code
func (o *SearchLdapUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this search ldap user internal server error response has a 5xx status code
func (o *SearchLdapUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this search ldap user internal server error response a status code equal to that given
func (o *SearchLdapUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SearchLdapUserInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ldap/users/search][%d] searchLdapUserInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchLdapUserInternalServerError) String() string {
	return fmt.Sprintf("[GET /ldap/users/search][%d] searchLdapUserInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchLdapUserInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *SearchLdapUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
