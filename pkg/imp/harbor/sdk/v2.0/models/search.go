// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Search search
//
// swagger:model Search
type Search struct {

	// Search results of the projects that matched the filter keywords.
	Project []*Project `json:"project"`

	// Search results of the repositories that matched the filter keywords.
	Repository []*SearchRepository `json:"repository"`
}

// Validate validates this search
func (m *Search) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepository(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Search) validateProject(formats strfmt.Registry) error {
	if swag.IsZero(m.Project) { // not required
		return nil
	}

	for i := 0; i < len(m.Project); i++ {
		if swag.IsZero(m.Project[i]) { // not required
			continue
		}

		if m.Project[i] != nil {
			if err := m.Project[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("project" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("project" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Search) validateRepository(formats strfmt.Registry) error {
	if swag.IsZero(m.Repository) { // not required
		return nil
	}

	for i := 0; i < len(m.Repository); i++ {
		if swag.IsZero(m.Repository[i]) { // not required
			continue
		}

		if m.Repository[i] != nil {
			if err := m.Repository[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("repository" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("repository" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this search based on the context it is used
func (m *Search) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRepository(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Search) contextValidateProject(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Project); i++ {

		if m.Project[i] != nil {
			if err := m.Project[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("project" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("project" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Search) contextValidateRepository(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Repository); i++ {

		if m.Repository[i] != nil {
			if err := m.Repository[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("repository" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("repository" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Search) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Search) UnmarshalBinary(b []byte) error {
	var res Search
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
