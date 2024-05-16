// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RetentionExecution retention execution
//
// swagger:model RetentionExecution
type RetentionExecution struct {

	// dry run
	DryRun bool `json:"dry_run,omitempty"`

	// end time
	EndTime string `json:"end_time,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// policy id
	PolicyID int64 `json:"policy_id,omitempty"`

	// start time
	StartTime string `json:"start_time,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// trigger
	Trigger string `json:"trigger,omitempty"`
}

// Validate validates this retention execution
func (m *RetentionExecution) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this retention execution based on context it is used
func (m *RetentionExecution) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RetentionExecution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RetentionExecution) UnmarshalBinary(b []byte) error {
	var res RetentionExecution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
