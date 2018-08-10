// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CheckpointDeleteOptions options of deleting a checkpoint from a container
// swagger:model CheckpointDeleteOptions

type CheckpointDeleteOptions struct {

	// checkpoint dir
	CheckpointDir string `json:"CheckpointDir,omitempty"`

	// checkpoint ID
	CheckpointID string `json:"CheckpointID,omitempty"`
}

/* polymorph CheckpointDeleteOptions CheckpointDir false */

/* polymorph CheckpointDeleteOptions CheckpointID false */

// Validate validates this checkpoint delete options
func (m *CheckpointDeleteOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CheckpointDeleteOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckpointDeleteOptions) UnmarshalBinary(b []byte) error {
	var res CheckpointDeleteOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}