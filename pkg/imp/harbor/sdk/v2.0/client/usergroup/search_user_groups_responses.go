// Code generated by go-swagger; DO NOT EDIT.

package usergroup

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

// SearchUserGroupsReader is a Reader for the SearchUserGroups structure.
type SearchUserGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchUserGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchUserGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSearchUserGroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchUserGroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchUserGroupsOK creates a SearchUserGroupsOK with default headers values
func NewSearchUserGroupsOK() *SearchUserGroupsOK {
	return &SearchUserGroupsOK{}
}

/*
SearchUserGroupsOK describes a response with status code 200, with default header values.

Search groups successfully.
*/
type SearchUserGroupsOK struct {

	/* Link to previous page and next page
	 */
	Link string

	/* The total count of available items
	 */
	XTotalCount int64

	Payload []*models.UserGroupSearchItem
}

// IsSuccess returns true when this search user groups o k response has a 2xx status code
func (o *SearchUserGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search user groups o k response has a 3xx status code
func (o *SearchUserGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search user groups o k response has a 4xx status code
func (o *SearchUserGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search user groups o k response has a 5xx status code
func (o *SearchUserGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search user groups o k response a status code equal to that given
func (o *SearchUserGroupsOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchUserGroupsOK) Error() string {
	return fmt.Sprintf("[GET /usergroups/search][%d] searchUserGroupsOK  %+v", 200, o.Payload)
}

func (o *SearchUserGroupsOK) String() string {
	return fmt.Sprintf("[GET /usergroups/search][%d] searchUserGroupsOK  %+v", 200, o.Payload)
}

func (o *SearchUserGroupsOK) GetPayload() []*models.UserGroupSearchItem {
	return o.Payload
}

func (o *SearchUserGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSearchUserGroupsUnauthorized creates a SearchUserGroupsUnauthorized with default headers values
func NewSearchUserGroupsUnauthorized() *SearchUserGroupsUnauthorized {
	return &SearchUserGroupsUnauthorized{}
}

/*
SearchUserGroupsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SearchUserGroupsUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this search user groups unauthorized response has a 2xx status code
func (o *SearchUserGroupsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search user groups unauthorized response has a 3xx status code
func (o *SearchUserGroupsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search user groups unauthorized response has a 4xx status code
func (o *SearchUserGroupsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this search user groups unauthorized response has a 5xx status code
func (o *SearchUserGroupsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this search user groups unauthorized response a status code equal to that given
func (o *SearchUserGroupsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SearchUserGroupsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /usergroups/search][%d] searchUserGroupsUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchUserGroupsUnauthorized) String() string {
	return fmt.Sprintf("[GET /usergroups/search][%d] searchUserGroupsUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchUserGroupsUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *SearchUserGroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSearchUserGroupsInternalServerError creates a SearchUserGroupsInternalServerError with default headers values
func NewSearchUserGroupsInternalServerError() *SearchUserGroupsInternalServerError {
	return &SearchUserGroupsInternalServerError{}
}

/*
SearchUserGroupsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SearchUserGroupsInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this search user groups internal server error response has a 2xx status code
func (o *SearchUserGroupsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search user groups internal server error response has a 3xx status code
func (o *SearchUserGroupsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search user groups internal server error response has a 4xx status code
func (o *SearchUserGroupsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this search user groups internal server error response has a 5xx status code
func (o *SearchUserGroupsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this search user groups internal server error response a status code equal to that given
func (o *SearchUserGroupsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SearchUserGroupsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /usergroups/search][%d] searchUserGroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchUserGroupsInternalServerError) String() string {
	return fmt.Sprintf("[GET /usergroups/search][%d] searchUserGroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchUserGroupsInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *SearchUserGroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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